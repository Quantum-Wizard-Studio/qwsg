// Package installation owns the canonical local installed-package
// classification boundary. It does not install, update, migrate, or remove
// artifacts.
package installation

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"quantumwizard.hu/qwsg/internal/update"
)

type State string

const (
	NoInstallation           State = "no_installation"
	VerifiedSupported        State = "verified_supported_installation"
	SupportedUpgradeSource   State = "supported_upgrade_source"
	LegacyInstallation       State = "legacy_installation"
	UnknownInstallation      State = "unknown_unverified_installation"
	InconsistentInstallation State = "inconsistent_incomplete_installation"
)

type Reason string

const (
	ReasonNoPackageArtifacts      Reason = "no_package_artifacts"
	ReasonPackageVerified         Reason = "package_verified"
	ReasonUpgradeRouteVerified    Reason = "upgrade_route_verified"
	ReasonLegacyBinaryOnly        Reason = "legacy_binary_only"
	ReasonBinaryOnlyUnverified    Reason = "binary_only_unverified"
	ReasonArtifactTypeUnsafe      Reason = "artifact_type_unsafe"
	ReasonPackageLayoutIncomplete Reason = "package_layout_incomplete"
	ReasonReleaseMetadataInvalid  Reason = "release_metadata_invalid"
	ReasonBinaryIdentityInvalid   Reason = "binary_identity_invalid"
	ReasonProvenanceMismatch      Reason = "package_provenance_mismatch"
	ReasonUnsupportedVersion      Reason = "unsupported_installed_version"
)

type Result struct {
	State       State  `json:"state"`
	Reason      Reason `json:"reason"`
	Version     string `json:"version,omitempty"`
	MigrationID string `json:"migration_id,omitempty"`
}

type VersionOutput func(context.Context, string) ([]byte, error)

type Options struct {
	Root             string
	CandidateVersion string
	RunVersion       VersionOutput
}

type releaseProvenance struct {
	Schema, Version, Commit, Built, Platform string
}

var packagePaths = []string{
	"usr/local/bin/qwsg",
	"usr/local/lib/systemd/user/qwsg-guardian.service",
	"usr/local/share/doc/qwsg/RELEASE.json",
	"usr/local/share/doc/qwsg/README.md",
	"usr/local/share/doc/qwsg/INSTALL.md",
	"usr/local/share/doc/qwsg/LICENSE",
	"usr/local/share/doc/qwsg/CHANGELOG.md",
	"usr/local/share/doc/qwsg/qwsg-config.json",
}

func Classify(options Options) Result {
	root := options.Root
	if root == "" {
		root = "/"
	}
	if !filepath.IsAbs(root) || filepath.Clean(root) != root {
		return Result{State: UnknownInstallation, Reason: ReasonArtifactTypeUnsafe}
	}
	rootInfo, err := os.Lstat(root)
	if err != nil || !rootInfo.IsDir() || rootInfo.Mode()&os.ModeSymlink != 0 {
		return Result{State: UnknownInstallation, Reason: ReasonArtifactTypeUnsafe}
	}
	run := options.RunVersion
	if run == nil {
		run = runVersion
	}

	present := 0
	for _, relative := range packagePaths {
		exists, safe, err := safeRegularPath(root, relative)
		if err != nil || !safe {
			return Result{State: InconsistentInstallation, Reason: ReasonArtifactTypeUnsafe}
		}
		if !exists {
			continue
		}
		present++
	}
	if present == 0 {
		return Result{State: NoInstallation, Reason: ReasonNoPackageArtifacts}
	}

	binary := filepath.Join(root, packagePaths[0])
	if present == 1 {
		identity, err := probeIdentity(run, binary)
		if err == nil && legacy(identity.Version) {
			return Result{State: LegacyInstallation, Reason: ReasonLegacyBinaryOnly, Version: identity.Version}
		}
		version := ""
		if err == nil {
			version = identity.Version
		}
		return Result{State: UnknownInstallation, Reason: ReasonBinaryOnlyUnverified, Version: version}
	}
	if present != len(packagePaths) {
		return Result{State: InconsistentInstallation, Reason: ReasonPackageLayoutIncomplete}
	}

	metadata, err := readProvenance(filepath.Join(root, packagePaths[2]))
	if err != nil {
		return Result{State: InconsistentInstallation, Reason: ReasonReleaseMetadataInvalid}
	}
	parsed, err := update.ParseVersion(metadata.Version)
	if err != nil || parsed.Major != 1 {
		return Result{State: UnknownInstallation, Reason: ReasonUnsupportedVersion, Version: metadata.Version}
	}
	identity, err := probeIdentity(run, binary)
	if err != nil {
		return Result{State: InconsistentInstallation, Reason: ReasonBinaryIdentityInvalid}
	}
	if identity != metadata {
		return Result{State: InconsistentInstallation, Reason: ReasonProvenanceMismatch, Version: metadata.Version}
	}
	if options.CandidateVersion != "" {
		candidate, candidateErr := update.ParseVersion(options.CandidateVersion)
		if candidateErr != nil {
			return Result{State: UnknownInstallation, Reason: ReasonUnsupportedVersion, Version: metadata.Version}
		}
		if update.Compare(candidate, parsed) > 0 {
			migration, migrationErr := update.PlanMigration(metadata.Version, options.CandidateVersion)
			if migrationErr == nil && migration.Validate() == nil {
				return Result{State: SupportedUpgradeSource, Reason: ReasonUpgradeRouteVerified, Version: metadata.Version, MigrationID: migration.ID}
			}
			return Result{State: UnknownInstallation, Reason: ReasonUnsupportedVersion, Version: metadata.Version}
		}
	}
	return Result{State: VerifiedSupported, Reason: ReasonPackageVerified, Version: metadata.Version}
}

func safeRegularPath(root, relative string) (exists, safe bool, err error) {
	current := root
	parts := strings.Split(filepath.Clean(relative), string(filepath.Separator))
	for index, part := range parts {
		current = filepath.Join(current, part)
		info, statErr := os.Lstat(current)
		if errors.Is(statErr, os.ErrNotExist) {
			return false, true, nil
		}
		if statErr != nil || info.Mode()&os.ModeSymlink != 0 {
			return false, false, statErr
		}
		if index == len(parts)-1 {
			return true, info.Mode().IsRegular(), nil
		}
		if !info.IsDir() {
			return false, false, nil
		}
	}
	return false, false, nil
}

func readProvenance(path string) (releaseProvenance, error) {
	file, err := os.Open(path)
	if err != nil {
		return releaseProvenance{}, err
	}
	defer file.Close()
	var value releaseProvenance
	decoder := json.NewDecoder(io.LimitReader(file, 16<<10))
	decoder.DisallowUnknownFields()
	if err = decoder.Decode(&value); err != nil {
		return releaseProvenance{}, err
	}
	var extra any
	if decoder.Decode(&extra) != io.EOF || value.Schema != "qwsg.release/1" || value.Platform != "linux-amd64" || len(value.Commit) != 40 || !lowerHex(value.Commit) {
		return releaseProvenance{}, fmt.Errorf("release provenance invalid")
	}
	if _, err = time.Parse(time.RFC3339, value.Built); err != nil {
		return releaseProvenance{}, fmt.Errorf("release build time invalid")
	}
	return value, nil
}

func probeIdentity(run VersionOutput, binary string) (releaseProvenance, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	output, err := run(ctx, binary)
	if err != nil || len(output) == 0 || len(output) > 4096 {
		return releaseProvenance{}, fmt.Errorf("binary identity unavailable")
	}
	lines := strings.Split(strings.TrimSuffix(string(output), "\n"), "\n")
	if len(lines) != 3 || !strings.HasPrefix(lines[0], "QWSG ") || !strings.HasPrefix(lines[1], "commit: ") || !strings.HasPrefix(lines[2], "built: ") {
		return releaseProvenance{}, fmt.Errorf("binary identity invalid")
	}
	value := releaseProvenance{Schema: "qwsg.release/1", Version: strings.TrimPrefix(lines[0], "QWSG "), Commit: strings.TrimPrefix(lines[1], "commit: "), Built: strings.TrimPrefix(lines[2], "built: "), Platform: "linux-amd64"}
	if _, err = update.ParseVersion(value.Version); err != nil {
		return releaseProvenance{}, err
	}
	return value, nil
}

func runVersion(ctx context.Context, binary string) ([]byte, error) {
	command := exec.CommandContext(ctx, binary, "version")
	var output bytes.Buffer
	command.Stdout = &output
	command.Stderr = io.Discard
	err := command.Run()
	return output.Bytes(), err
}

func legacy(version string) bool {
	parsed, err := update.ParseVersion(version)
	return err == nil && parsed.Major == 0
}

func lowerHex(value string) bool {
	if value == "" || strings.ToLower(value) != value {
		return false
	}
	for _, r := range value {
		if !(r >= '0' && r <= '9') && !(r >= 'a' && r <= 'f') {
			return false
		}
	}
	return true
}

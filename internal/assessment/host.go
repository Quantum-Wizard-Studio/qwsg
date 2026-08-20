package assessment

import (
	"context"
	"errors"
	"os"
	"os/user"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"syscall"
	"time"

	"quantumwizard.hu/qwsg/internal/runner"
)

type Host interface {
	ReadOSRelease() ([]byte, error)
	Architecture() string
	EffectiveUID() int
	Run(context.Context, string) (runner.Result, error)
}

type LocalHost struct{ Runner runner.Runner }

func (LocalHost) ReadOSRelease() ([]byte, error) {
	file, err := os.Open("/etc/os-release")
	if err != nil {
		return nil, err
	}
	defer file.Close()
	buffer := make([]byte, 64*1024+1)
	n, err := file.Read(buffer)
	if n > 64*1024 {
		return nil, runner.ErrOutputLimit
	}
	if err != nil && n == 0 {
		return nil, err
	}
	return buffer[:n], nil
}
func (LocalHost) Architecture() string { return runtime.GOARCH }
func (LocalHost) EffectiveUID() int    { return os.Geteuid() }
func (host LocalHost) Run(ctx context.Context, id string) (runner.Result, error) {
	return host.Runner.Run(ctx, id)
}

func DefaultRunner() runner.Runner {
	environment := map[string][]string{}
	if directory, evidence := trustedRuntimeDirectory(os.Geteuid()); evidence == "systemd_user_runtime_directory_valid" {
		for _, id := range []string{"systemd_user", "systemd_unit_installed", "systemd_service_enabled", "systemd_service_active"} {
			environment[id] = []string{"XDG_RUNTIME_DIR=" + directory}
		}
	}
	return fixedRunner{base: runner.Bounded{Allowed: map[string]string{"glibc_version": "/usr/bin/getconf", "systemd_version": "/usr/bin/systemctl", "systemd_user": "/usr/bin/systemctl", "systemd_unit_installed": "/usr/bin/systemctl", "systemd_service_enabled": "/usr/bin/systemctl", "systemd_service_active": "/usr/bin/systemctl"}, TrustedEnvironment: environment, Timeout: 2 * time.Second, MaxOutput: 64 << 10}}
}

type fixedRunner struct{ base runner.Runner }

func (value fixedRunner) Run(ctx context.Context, id string, _ ...string) (runner.Result, error) {
	arguments := map[string][]string{
		"glibc_version":           {"GNU_LIBC_VERSION"},
		"systemd_version":         {"--version"},
		"systemd_user":            {"--user", "is-system-running"},
		"systemd_unit_installed":  {"--user", "cat", "qwsg-guardian.service"},
		"systemd_service_enabled": {"--user", "is-enabled", "qwsg-guardian.service"},
		"systemd_service_active":  {"--user", "is-active", "qwsg-guardian.service"},
	}
	args, ok := arguments[id]
	if !ok {
		return runner.Result{}, errors.New("assessment probe is not allowlisted")
	}
	return value.base.Run(ctx, id, args...)
}

func AssessInstall(ctx context.Context, host Host, now time.Time) Report {
	platform, osFinding := detectPlatform(host)
	findings := []Finding{osFinding}
	architecture := host.Architecture()
	archClass, archEvidence := Satisfied, "architecture_supported"
	if architecture != "amd64" {
		archClass, archEvidence = Incompatible, "architecture_unsupported"
	}
	findings = append(findings, Finding{RequirementID: "platform.architecture", Classification: archClass, EvidenceToken: archEvidence, ObservedValue: architecture})
	userClass, userEvidence := Satisfied, "ordinary_user"
	if host.EffectiveUID() == 0 {
		userClass, userEvidence = Incompatible, "root_runtime_unsupported"
	}
	findings = append(findings, Finding{RequirementID: "runtime.non_root", Classification: userClass, EvidenceToken: userEvidence})
	findings = append(findings, glibcFinding(ctx, host))
	findings = append(findings, systemdVersion(ctx, host), systemdUserManagerFinding(ctx, host))
	findings = append(findings, filesystemFinding(host))
	SortFindings(findings)
	domain := Summarize(findings)
	report := Report{SchemaName: SchemaName, SchemaVersion: SchemaVersion, ModelVersion: ModelVersion, RegistryVersion: RegistryVersion, AssessedAt: now.UTC(), Phase: "install", Platform: platform, Findings: findings, Domains: []DomainSummary{{Domain: "environment_dependencies", State: domain}, {Domain: "installation", State: domain}}}
	AttachGuidance(&report)
	report.NextActions = nextInstall(domain, report.Findings)
	return report
}

func systemdUserManagerFinding(ctx context.Context, host Host) Finding {
	if local, ok := host.(interface{ UserManagerFinding(context.Context) Finding }); ok {
		return local.UserManagerFinding(ctx)
	}
	result, err := host.Run(ctx, "systemd_user")
	return classifyUserManager(result, err)
}

func (host LocalHost) UserManagerFinding(ctx context.Context) Finding {
	_, evidence := trustedRuntimeDirectory(host.EffectiveUID())
	switch evidence {
	case "systemd_user_runtime_directory_missing":
		return Finding{RequirementID: "systemd.user_manager", Classification: MissingRequired, EvidenceToken: evidence}
	case "systemd_user_runtime_directory_unsafe":
		return Finding{RequirementID: "systemd.user_manager", Classification: UnknownVerification, EvidenceToken: evidence}
	}
	result, err := host.Run(ctx, "systemd_user")
	return classifyUserManager(result, err)
}

func classifyUserManager(result runner.Result, err error) Finding {
	state := strings.TrimSpace(string(result.Stdout))
	if state == "running" || state == "degraded" {
		return Finding{RequirementID: "systemd.user_manager", Classification: Satisfied, EvidenceToken: "systemd_user_manager_available", ObservedValue: state}
	}
	for _, transient := range []string{"initializing", "starting", "maintenance", "stopping"} {
		if state == transient {
			return Finding{RequirementID: "systemd.user_manager", Classification: UnknownVerification, EvidenceToken: "systemd_user_manager_starting"}
		}
	}
	if errors.Is(err, context.DeadlineExceeded) {
		return Finding{RequirementID: "systemd.user_manager", Classification: UnknownVerification, EvidenceToken: "systemd_user_probe_timeout"}
	}
	if errors.Is(err, runner.ErrOutputLimit) {
		return Finding{RequirementID: "systemd.user_manager", Classification: UnknownVerification, EvidenceToken: "systemd_user_probe_output_limit"}
	}
	if err != nil && state == "" {
		return Finding{RequirementID: "systemd.user_manager", Classification: MissingRequired, EvidenceToken: "systemd_user_manager_unavailable"}
	}
	if err != nil {
		return Finding{RequirementID: "systemd.user_manager", Classification: UnknownVerification, EvidenceToken: "systemd_user_probe_failed"}
	}
	return Finding{RequirementID: "systemd.user_manager", Classification: UnknownVerification, EvidenceToken: "systemd_user_state_unrecognized"}
}

func trustedRuntimeDirectory(uid int) (string, string) {
	if uid < 0 {
		return "", "systemd_user_runtime_directory_unsafe"
	}
	directory := filepath.Join("/run/user", strconv.Itoa(uid))
	return validateRuntimeDirectory(directory, uid)
}

func validateRuntimeDirectory(directory string, uid int) (string, string) {
	info, err := os.Lstat(directory)
	if os.IsNotExist(err) {
		return "", "systemd_user_runtime_directory_missing"
	}
	if err != nil || info.Mode()&os.ModeSymlink != 0 || !info.IsDir() || info.Mode().Perm()&0077 != 0 {
		return "", "systemd_user_runtime_directory_unsafe"
	}
	stat, ok := info.Sys().(*syscall.Stat_t)
	if !ok || int(stat.Uid) != uid {
		return "", "systemd_user_runtime_directory_unsafe"
	}
	return directory, "systemd_user_runtime_directory_valid"
}

func filesystemFinding(host Host) Finding {
	if local, ok := host.(interface{ FilesystemFinding() Finding }); ok {
		return local.FilesystemFinding()
	}
	return Finding{RequirementID: "filesystem.local_semantics", Classification: UnknownVerification, EvidenceToken: "filesystem_type_unknown"}
}

func (LocalHost) FilesystemFinding() Finding {
	current, err := user.Current()
	if err != nil || current.HomeDir == "" || !filepath.IsAbs(current.HomeDir) {
		return Finding{RequirementID: "filesystem.local_semantics", Classification: UnknownVerification, EvidenceToken: "filesystem_path_unavailable"}
	}
	paths, ok := effectiveAssessmentPaths(current.HomeDir)
	if !ok {
		return Finding{RequirementID: "filesystem.local_semantics", Classification: UnknownVerification, EvidenceToken: "filesystem_path_unsafe"}
	}
	for _, path := range paths {
		ancestor, evidence := safeExistingAncestor(path, os.Geteuid())
		if evidence != "" {
			return Finding{RequirementID: "filesystem.local_semantics", Classification: UnknownVerification, EvidenceToken: evidence}
		}
		var stat syscall.Statfs_t
		if err := syscall.Statfs(ancestor, &stat); err != nil {
			return Finding{RequirementID: "filesystem.local_semantics", Classification: UnknownVerification, EvidenceToken: "filesystem_path_unavailable"}
		}
		classification, evidence := classifyFilesystemType(uint64(stat.Type))
		if classification != Satisfied {
			return Finding{RequirementID: "filesystem.local_semantics", Classification: classification, EvidenceToken: evidence}
		}
	}
	return Finding{RequirementID: "filesystem.local_semantics", Classification: Satisfied, EvidenceToken: "filesystem_local_semantics_supported"}
}

func classifyFilesystemType(value uint64) (Classification, string) {
	switch value {
	case 0xef53, 0x58465342, 0x9123683e: // ext*, XFS, Btrfs
		return Satisfied, "filesystem_local_semantics_supported"
	default:
		return UnknownVerification, "filesystem_remote_or_overlay"
	}
}

func safeExistingAncestor(path string, uid int) (string, string) {
	path = filepath.Clean(path)
	if !filepath.IsAbs(path) || pathHasSymlink(path) {
		return "", "filesystem_path_unsafe"
	}
	for {
		info, err := os.Lstat(path)
		if err == nil {
			if info.Mode()&os.ModeSymlink != 0 || !info.IsDir() {
				return "", "filesystem_path_unsafe"
			}
			stat, ok := info.Sys().(*syscall.Stat_t)
			if !ok || int(stat.Uid) != uid {
				return "", "filesystem_path_unsafe"
			}
			return path, ""
		}
		if !os.IsNotExist(err) {
			return "", "filesystem_path_unavailable"
		}
		parent := filepath.Dir(path)
		if parent == path {
			return "", "filesystem_path_unavailable"
		}
		path = parent
	}
}

func effectiveAssessmentPaths(fallbackHome string) ([]string, bool) {
	home := os.Getenv("HOME")
	if home == "" {
		home = fallbackHome
	}
	configBase := os.Getenv("XDG_CONFIG_HOME")
	if configBase == "" {
		configBase = filepath.Join(home, ".config")
	}
	state := os.Getenv("QWSG_STATE_DIR")
	if state == "" {
		stateBase := os.Getenv("XDG_STATE_HOME")
		if stateBase == "" {
			stateBase = filepath.Join(home, ".local", "state")
		}
		state = filepath.Join(stateBase, "qwsg")
	}
	config := filepath.Join(configBase, "qwsg")
	if !filepath.IsAbs(config) || !filepath.IsAbs(state) || filepath.Clean(config) != config || filepath.Clean(state) != state {
		return nil, false
	}
	return []string{config, state}, true
}

func pathHasSymlink(path string) bool {
	current := string(filepath.Separator)
	for _, component := range strings.Split(strings.TrimPrefix(filepath.Clean(path), string(filepath.Separator)), string(filepath.Separator)) {
		if component == "" {
			continue
		}
		current = filepath.Join(current, component)
		info, err := os.Lstat(current)
		if os.IsNotExist(err) {
			return false
		}
		if err != nil || info.Mode()&os.ModeSymlink != 0 {
			return true
		}
	}
	return false
}

func glibcFinding(ctx context.Context, host Host) Finding {
	result, err := host.Run(ctx, "glibc_version")
	if err != nil {
		return Finding{RequirementID: "runtime.glibc", Classification: UnknownVerification, EvidenceToken: "glibc_version_unavailable"}
	}
	fields := strings.Fields(string(result.Stdout))
	if len(fields) != 2 || fields[0] != "glibc" {
		return Finding{RequirementID: "runtime.glibc", Classification: UnknownVerification, EvidenceToken: "glibc_version_invalid"}
	}
	return Finding{RequirementID: "runtime.glibc", Classification: Satisfied, EvidenceToken: "glibc_compatible", ObservedValue: fields[1]}
}

func detectPlatform(host Host) (Platform, Finding) {
	platform := Platform{ID: "unknown", Architecture: host.Architecture()}
	document, err := host.ReadOSRelease()
	if err != nil || len(document) > 64*1024 {
		return platform, Finding{RequirementID: "platform.operating_system", Classification: UnknownVerification, EvidenceToken: "os_release_unavailable"}
	}
	values := map[string]string{}
	for _, line := range strings.Split(string(document), "\n") {
		parts := strings.SplitN(strings.TrimSpace(line), "=", 2)
		if len(parts) == 2 {
			values[parts[0]] = strings.Trim(strings.TrimSpace(parts[1]), "\"'")
		}
	}
	platform.Distribution, platform.Version = values["ID"], values["VERSION_ID"]
	if platform.Distribution == "ubuntu" && platform.Version == "24.04" && platform.Architecture == "amd64" {
		platform.ID, platform.Supported = "ubuntu-24.04-amd64", true
		return platform, Finding{RequirementID: "platform.operating_system", Classification: Satisfied, EvidenceToken: "ubuntu_24_04_supported", ObservedValue: "Ubuntu 24.04"}
	}
	if platform.Distribution == "" || platform.Version == "" {
		return platform, Finding{RequirementID: "platform.operating_system", Classification: UnknownVerification, EvidenceToken: "os_release_incomplete"}
	}
	return platform, Finding{RequirementID: "platform.operating_system", Classification: Incompatible, EvidenceToken: "operating_system_unsupported", ObservedValue: platform.Distribution + " " + platform.Version}
}

func systemdVersion(ctx context.Context, host Host) Finding {
	result, err := host.Run(ctx, "systemd_version")
	if err != nil {
		return Finding{RequirementID: "systemd.version", Classification: UnknownVerification, EvidenceToken: "systemd_version_unavailable"}
	}
	fields := strings.Fields(string(result.Stdout))
	if len(fields) < 2 || fields[0] != "systemd" {
		return Finding{RequirementID: "systemd.version", Classification: UnknownVerification, EvidenceToken: "systemd_version_invalid"}
	}
	version, err := strconv.Atoi(fields[1])
	if err != nil {
		return Finding{RequirementID: "systemd.version", Classification: UnknownVerification, EvidenceToken: "systemd_version_invalid"}
	}
	if version < 255 {
		return Finding{RequirementID: "systemd.version", Classification: Incompatible, EvidenceToken: "systemd_version_too_old", ObservedValue: fields[1]}
	}
	return Finding{RequirementID: "systemd.version", Classification: Satisfied, EvidenceToken: "systemd_version_supported", ObservedValue: fields[1]}
}

func commandFinding(ctx context.Context, host Host, requirementID, probeID string, required bool) Finding {
	result, err := host.Run(ctx, probeID)
	if err == nil || (probeID == "systemd_user" && result.ExitCode == 1 && strings.TrimSpace(string(result.Stdout)) == "degraded") {
		return Finding{RequirementID: requirementID, Classification: Satisfied, EvidenceToken: probeID + "_available"}
	}
	classification := MissingOptional
	if required {
		classification = MissingRequired
	}
	return Finding{RequirementID: requirementID, Classification: classification, EvidenceToken: probeID + "_unavailable"}
}

func nextInstall(state SummaryState, findings []Finding) []string {
	if state == Ready || state == Partial {
		for _, finding := range findings {
			if finding.Classification == UnknownVerification && finding.Guidance != nil {
				return []string{"review_filesystem_verification_guidance", "install_release_artifacts", "run_qwsg_setup", "run_qwsg_readiness"}
			}
		}
		return []string{"install_release_artifacts", "run_qwsg_setup", "run_qwsg_readiness"}
	}
	for _, finding := range findings {
		if finding.RequirementID == "systemd.user_manager" && finding.Classification != Satisfied {
			return []string{"review_systemd_user_manager_guidance", "rerun_qwsg_install_check"}
		}
	}
	return []string{"resolve_required_findings", "rerun_qwsg_install_check"}
}

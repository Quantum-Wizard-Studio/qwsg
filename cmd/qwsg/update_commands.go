package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"syscall"
	"time"

	updatecore "quantumwizard.hu/qwsg/internal/update"
)

type localUpdateRecord struct{ Schema, Installed, Previous, Backup, UpdatedAt string }

var installedQWSGBinary = "/usr/local/bin/qwsg"

func runUpdate(args []string, out, errout io.Writer) int {
	if len(args) > 0 && isHelp(args[0]) {
		writeUpdateHelp(out)
		return 0
	}
	if len(args) > 0 && args[0] == "privileged-apply" {
		return runPrivilegedApply(args[1:], errout)
	}
	if len(args) > 0 && args[0] == "privileged-rollback" {
		return runPrivilegedRollback(args[1:], errout)
	}
	if len(args) > 0 && args[0] == "privileged-discard" {
		return runPrivilegedDiscard(args[1:], errout)
	}
	if len(args) > 0 && args[0] == "check" {
		if len(args) != 1 {
			return usageError(errout, "update check does not accept options")
		}
		return runUpdateCheck(out, errout)
	}
	if len(args) > 0 && args[0] == "status" {
		if len(args) != 1 {
			return usageError(errout, "update status does not accept options")
		}
		return runUpdateStatus(out, errout)
	}
	if len(args) > 0 && args[0] == "rollback" {
		if len(args) != 1 {
			return usageError(errout, "update rollback does not accept options")
		}
		return runUpdateRollback(out, errout)
	}
	archive, target, err := parseUpdateArgs(args)
	if err != nil {
		return usageError(errout, "%v", err)
	}
	return executeUpdate(archive, target, out, errout)
}

func runUpdateCheck(out, errout io.Writer) int {
	installed, err := installedVersion()
	if err != nil {
		fmt.Fprintln(errout, "Update check failed: installed QWSG identity unavailable.")
		return 1
	}
	ctx, cancel := context.WithTimeout(context.Background(), 35*time.Second)
	defer cancel()
	release, relation, err := updatecore.Discover(ctx, updatecore.HTTPClient(), "", installed)
	if err != nil {
		fmt.Fprintln(errout, "Update check failed: release discovery unavailable.")
		return 1
	}
	fmt.Fprintf(out, "Installed: %s\nAvailable: %s\nRelation: %s\n", safeText(installed), safeText(release.Version), relation)
	return 0
}

func parseUpdateArgs(args []string) (archive, target string, err error) {
	for i := 0; i < len(args); i++ {
		if i+1 >= len(args) {
			return "", "", fmt.Errorf("update option requires a value")
		}
		switch args[i] {
		case "--archive":
			archive = args[i+1]
		case "--version":
			target = args[i+1]
		default:
			return "", "", fmt.Errorf("unknown update option: %s", safeText(args[i]))
		}
		i++
	}
	if (archive == "") != (target == "") {
		return "", "", fmt.Errorf("--archive and --version must be used together")
	}
	return
}

func executeUpdate(localArchive, target string, out, errout io.Writer) int {
	if os.Geteuid() == 0 {
		fmt.Fprintln(errout, "Update orchestration must run as the intended non-root QWSG user.")
		return 1
	}
	installed, err := installedVersion()
	if err != nil {
		fmt.Fprintln(errout, "Update failed: installed QWSG identity unavailable.")
		return 1
	}
	state, err := localStateRoot()
	if err != nil {
		fmt.Fprintln(errout, "Update failed: local state unavailable.")
		return 1
	}
	updateRoot := filepath.Join(state, "update")
	if err = ensureUpdateRoot(updateRoot); err != nil {
		fmt.Fprintln(errout, "Update failed: private update state unavailable.")
		return 1
	}
	previousRecord, _ := loadUpdateRecord(updateRoot)
	var staged updatecore.Staged
	if localArchive != "" {
		if updatecore.Classify(installed, target) != updatecore.Newer {
			fmt.Fprintln(errout, "Update refused: target is not a supported newer version.")
			return 1
		}
		staged, err = updatecore.StageLocal(localArchive, localArchive+".sha256", target, updateRoot)
	} else {
		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
		defer cancel()
		var release updatecore.Release
		var relation updatecore.Relation
		release, relation, err = updatecore.Discover(ctx, updatecore.HTTPClient(), "", installed)
		if err == nil && relation != updatecore.Newer {
			fmt.Fprintf(out, "QWSG is not updated: available release is %s (%s).\n", safeText(release.Version), relation)
			return 0
		}
		if err == nil {
			staged, err = updatecore.Acquire(ctx, updatecore.HTTPClient(), release, updateRoot)
		}
	}
	if err != nil {
		fmt.Fprintln(errout, "Update failed: candidate acquisition or integrity verification failed.")
		return 1
	}
	defer os.RemoveAll(staged.Root)
	pkg, err := updatecore.VerifyPackage(staged)
	if err != nil {
		fmt.Fprintln(errout, "Update failed: package verification failed.")
		return 1
	}
	migration, err := updatecore.PlanMigration(installed, pkg.Provenance.Version)
	if err != nil || migration.Validate() != nil {
		fmt.Fprintln(errout, "Update refused: no deterministic compatible migration path.")
		return 1
	}
	enabled := commandState("is-enabled")
	active := commandState("is-active")
	if active == "yes" {
		if err = runSystemctl("stop"); err != nil {
			fmt.Fprintln(errout, "Update failed: Guardian could not be stopped.")
			return 1
		}
	}
	uid := strconv.Itoa(os.Getuid())
	txid := time.Now().UTC().Format("20060102T150405.000000000Z")
	backup := filepath.Join("/var/lib/qwsg/rollback", uid, txid)
	err = runSudo("privileged-apply", "--archive", staged.Archive, "--sidecar", staged.Sidecar, "--version", pkg.Provenance.Version, "--sha256", staged.SHA256, "--backup", backup, "--from", installed)
	if err == nil {
		err = runSystemctl("daemon-reload")
	}
	if err == nil && enabled == "yes" {
		err = runSystemctl("enable")
	}
	if err == nil && active == "yes" {
		err = runSystemctl("start")
	}
	if err == nil {
		err = validateInstalledVersion(pkg.Provenance.Version)
	}
	if err != nil {
		_ = runSudo("privileged-rollback", "--backup", backup)
		_ = runSystemctl("daemon-reload")
		if enabled == "yes" {
			_ = runSystemctl("enable")
		}
		if active == "yes" {
			_ = runSystemctl("start")
		}
		fmt.Fprintln(errout, "Update failed after mutation; automatic package rollback was attempted.")
		return 1
	}
	record := localUpdateRecord{Schema: "qwsg.update-local/1", Installed: pkg.Provenance.Version, Previous: installed, Backup: backup, UpdatedAt: time.Now().UTC().Format(time.RFC3339Nano)}
	if err = saveUpdateRecord(updateRoot, record); err != nil {
		fmt.Fprintln(errout, "Update installed but local rollback metadata could not be recorded.")
		return 1
	}
	if previousRecord.Backup != "" && previousRecord.Backup != backup {
		_ = runSudo("privileged-discard", "--backup", previousRecord.Backup)
	}
	fmt.Fprintf(out, "QWSG updated safely: %s -> %s\nRollback available: qwsg update rollback\n", safeText(installed), safeText(pkg.Provenance.Version))
	return 0
}

func runPrivilegedApply(args []string, errout io.Writer) int {
	if os.Geteuid() != 0 {
		return usageError(errout, "privileged update helper requires root")
	}
	values, err := parsePairs(args, "--archive", "--sidecar", "--version", "--sha256", "--backup", "--from")
	if err != nil {
		return usageError(errout, "%v", err)
	}
	if !validBackup(values["--backup"]) {
		return usageError(errout, "unsafe rollback path")
	}
	if err = os.MkdirAll("/var/lib/qwsg", 0700); err != nil {
		return 1
	}
	rootInfo, rootErr := os.Lstat("/var/lib/qwsg")
	if rootErr != nil || !rootInfo.IsDir() || rootInfo.Mode()&os.ModeSymlink != 0 {
		return 1
	}
	stageParent, err := os.MkdirTemp("/var/lib/qwsg", ".verified-stage-")
	if err != nil {
		fmt.Fprintln(errout, "privileged update staging failed")
		return 1
	}
	defer os.RemoveAll(stageParent)
	if err = os.Chmod(stageParent, 0700); err != nil {
		return 1
	}
	staged, err := updatecore.StageLocal(values["--archive"], values["--sidecar"], values["--version"], stageParent)
	if err != nil || staged.SHA256 != values["--sha256"] {
		fmt.Fprintln(errout, "privileged candidate re-verification failed")
		return 1
	}
	pkg, err := updatecore.VerifyPackage(staged)
	if err != nil {
		fmt.Fprintln(errout, "privileged package re-verification failed")
		return 1
	}
	if _, err = updatecore.Apply(pkg.Root, "/", values["--backup"], values["--from"]); err != nil {
		fmt.Fprintln(errout, "privileged update transaction failed")
		return 1
	}
	return 0
}
func runPrivilegedRollback(args []string, errout io.Writer) int {
	if os.Geteuid() != 0 {
		return usageError(errout, "privileged rollback helper requires root")
	}
	values, err := parsePairs(args, "--backup")
	if err != nil || !validBackup(values["--backup"]) {
		return usageError(errout, "unsafe rollback request")
	}
	if err = updatecore.Rollback("/", values["--backup"]); err != nil {
		fmt.Fprintln(errout, "privileged rollback transaction failed")
		return 1
	}
	if err = os.RemoveAll(values["--backup"]); err != nil {
		fmt.Fprintln(errout, "privileged rollback cleanup failed")
		return 1
	}
	return 0
}

func runPrivilegedDiscard(args []string, errout io.Writer) int {
	if os.Geteuid() != 0 {
		return usageError(errout, "privileged discard helper requires root")
	}
	values, err := parsePairs(args, "--backup")
	if err != nil || !validBackup(values["--backup"]) {
		return usageError(errout, "unsafe discard request")
	}
	if _, err = updatecore.ReadTransaction(values["--backup"]); err != nil {
		fmt.Fprintln(errout, "privileged discard validation failed")
		return 1
	}
	if err = os.RemoveAll(values["--backup"]); err != nil {
		fmt.Fprintln(errout, "privileged discard failed")
		return 1
	}
	return 0
}

func runUpdateStatus(out, errout io.Writer) int {
	root, err := localStateRoot()
	if err != nil {
		return 1
	}
	record, err := loadUpdateRecord(filepath.Join(root, "update"))
	if os.IsNotExist(err) {
		fmt.Fprintln(out, "No native update transaction is recorded.")
		return 0
	}
	if err != nil {
		fmt.Fprintln(errout, "Update status unavailable: local metadata invalid.")
		return 1
	}
	fmt.Fprintf(out, "Installed by updater: %s\nPrevious: %s\nRollback: available\n", safeText(record.Installed), safeText(record.Previous))
	return 0
}
func runUpdateRollback(out, errout io.Writer) int {
	if os.Geteuid() == 0 {
		fmt.Fprintln(errout, "Rollback orchestration must run as the intended non-root QWSG user.")
		return 1
	}
	root, err := localStateRoot()
	if err != nil {
		return 1
	}
	record, err := loadUpdateRecord(filepath.Join(root, "update"))
	if err != nil {
		fmt.Fprintln(errout, "Rollback unavailable: local metadata missing or invalid.")
		return 1
	}
	active := commandState("is-active")
	enabled := commandState("is-enabled")
	if active == "yes" {
		if err = runSystemctl("stop"); err != nil {
			return 1
		}
	}
	if err = runSudo("privileged-rollback", "--backup", record.Backup); err == nil {
		err = runSystemctl("daemon-reload")
	}
	if err == nil && enabled == "yes" {
		err = runSystemctl("enable")
	}
	if err == nil && active == "yes" {
		err = runSystemctl("start")
	}
	if err != nil {
		fmt.Fprintln(errout, "Rollback failed; Guardian was not reported ready.")
		return 1
	}
	if err = os.Remove(filepath.Join(root, "update", "current.json")); err != nil {
		return 1
	}
	fmt.Fprintf(out, "QWSG rolled back safely: %s -> %s\n", safeText(record.Installed), safeText(record.Previous))
	return 0
}

func ensureUpdateRoot(path string) error {
	if err := os.Mkdir(path, 0700); err != nil && !os.IsExist(err) {
		return err
	}
	info, err := os.Lstat(path)
	if err != nil || !info.IsDir() || info.Mode()&os.ModeSymlink != 0 || info.Mode().Perm() != 0700 || int(info.Sys().(*syscall.Stat_t).Uid) != os.Getuid() {
		return fmt.Errorf("unsafe update root")
	}
	return nil
}
func saveUpdateRecord(root string, r localUpdateRecord) error {
	data, err := json.Marshal(r)
	if err != nil {
		return err
	}
	data = append(data, '\n')
	tmp, err := os.CreateTemp(root, ".current-")
	if err != nil {
		return err
	}
	name := tmp.Name()
	defer os.Remove(name)
	if err = tmp.Chmod(0600); err == nil {
		_, err = tmp.Write(data)
	}
	if err == nil {
		err = tmp.Sync()
	}
	if ce := tmp.Close(); err == nil {
		err = ce
	}
	if err != nil {
		return err
	}
	return os.Rename(name, filepath.Join(root, "current.json"))
}
func loadUpdateRecord(root string) (localUpdateRecord, error) {
	path := filepath.Join(root, "current.json")
	info, err := os.Lstat(path)
	if err != nil {
		return localUpdateRecord{}, err
	}
	if !info.Mode().IsRegular() || info.Mode().Perm() != 0600 {
		return localUpdateRecord{}, fmt.Errorf("unsafe update metadata")
	}
	data, err := os.ReadFile(path)
	if err != nil {
		return localUpdateRecord{}, err
	}
	var r localUpdateRecord
	if json.Unmarshal(data, &r) != nil || r.Schema != "qwsg.update-local/1" || !validBackup(r.Backup) {
		return localUpdateRecord{}, fmt.Errorf("invalid update metadata")
	}
	return r, nil
}
func commandState(action string) string {
	cmd := exec.Command("/usr/bin/systemctl", "--user", action, "qwsg-guardian.service")
	if cmd.Run() == nil {
		return "yes"
	}
	return "no"
}
func runSystemctl(action string) error {
	args := []string{"--user", action}
	if action != "daemon-reload" {
		args = append(args, "qwsg-guardian.service")
	}
	return exec.Command("/usr/bin/systemctl", args...).Run()
}

func validateInstalledVersion(want string) error {
	output, err := exec.Command(installedQWSGBinary, "version").Output()
	if err != nil {
		return err
	}
	first := strings.SplitN(string(output), "\n", 2)[0]
	if first != "QWSG "+want {
		return fmt.Errorf("installed version mismatch")
	}
	return exec.Command("/usr/local/bin/qwsg", "config", "validate").Run()
}

func installedVersion() (string, error) {
	output, err := exec.Command(installedQWSGBinary, "version").Output()
	if err != nil {
		return "", err
	}
	line := strings.SplitN(string(output), "\n", 2)[0]
	if !strings.HasPrefix(line, "QWSG ") {
		return "", fmt.Errorf("installed version output invalid")
	}
	value := strings.TrimPrefix(line, "QWSG ")
	if _, err = updatecore.ParseVersion(value); err != nil {
		return "", err
	}
	return value, nil
}
func runSudo(args ...string) error {
	executable, err := os.Executable()
	if err != nil {
		return err
	}
	executable, err = filepath.Abs(executable)
	if err != nil {
		return err
	}
	all := append([]string{"--", executable, "update"}, args...)
	return exec.Command("/usr/bin/sudo", all...).Run()
}
func parsePairs(args []string, names ...string) (map[string]string, error) {
	allowed := map[string]bool{}
	for _, n := range names {
		allowed[n] = true
	}
	result := map[string]string{}
	for i := 0; i < len(args); i += 2 {
		if i+1 >= len(args) || !allowed[args[i]] || result[args[i]] != "" {
			return nil, fmt.Errorf("invalid helper arguments")
		}
		result[args[i]] = args[i+1]
	}
	for _, n := range names {
		if result[n] == "" {
			return nil, fmt.Errorf("missing helper argument")
		}
	}
	return result, nil
}
func validBackup(path string) bool {
	return strings.HasPrefix(path, "/var/lib/qwsg/rollback/") && !strings.Contains(path, "..") && filepath.Clean(path) == path
}
func writeUpdateHelp(out io.Writer) {
	fmt.Fprintln(out, "Usage:\n  qwsg update check\n  qwsg update\n  qwsg update --archive FILE --version VERSION\n  qwsg update status\n  qwsg update rollback")
}

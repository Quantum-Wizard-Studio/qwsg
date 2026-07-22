package collector

import (
	"bufio"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"os"
	"runtime"
	"sort"
	"strconv"
	"strings"
	"syscall"
	"time"

	"quantumwizard.hu/qwsg/internal/inventory"
	"quantumwizard.hu/qwsg/internal/runner"
)

var ErrUnsupported = errors.New("unsupported platform")
var ErrInvalidEvidence = errors.New("invalid evidence")

type Func struct {
	Name         string
	Capability   string
	Dependencies []string
	Timeout      time.Duration
	Check        func(context.Context) Availability
	Run          func(context.Context) ([]inventory.Item, []string, error)
}

func (f Func) Descriptor() Descriptor {
	capability := f.Capability
	if capability == "" {
		capability = f.Name
	}
	timeout := f.Timeout
	if timeout <= 0 {
		timeout = 2 * time.Second
	}
	return Descriptor{Name: f.Name, Version: "1.0.0", ContractVersion: ContractVersion, InventoryCompatibility: inventory.SchemaVersion, Capability: Capability{ID: capability, Description: capability + " inventory"}, SupportedPlatforms: []string{"linux"}, PrivilegeClass: "ordinary_user", Timeout: timeout, MaxTimeout: 5 * time.Second, OutputLimitBytes: 2 << 20, Dependencies: append([]string(nil), f.Dependencies...), SensitivityClasses: []string{"operational", "host_identifying", "network_sensitive"}}
}
func (f Func) Availability(ctx context.Context) Availability {
	if f.Check != nil {
		return f.Check(ctx)
	}
	return Availability{Available: true}
}
func (f Func) Collect(ctx context.Context, _ Request) Result {
	start := time.Now().UTC()
	items, sources, err := f.Run(ctx)
	end := time.Now().UTC()
	d := f.Descriptor()
	c := inventory.Category{CategoryID: d.Capability.ID, ContractVersion: ContractVersion, Status: inventory.Available, ObservedAt: start, CompletedAt: end, FreshUntil: end.Add(5 * time.Minute), DurationMS: time.Since(start).Milliseconds(), CollectorID: f.Name, PrivilegeUsed: "ordinary_user", SourceSummary: sources, Items: items, Errors: []inventory.InventoryError{}, Redactions: []string{}}
	if err != nil {
		c.Items = []inventory.Item{}
		c.Status = classify(err)
		c.Errors = []inventory.InventoryError{{Code: "collection_failed", CategoryID: d.Capability.ID, Class: string(c.Status), MessageKey: "collector.collection_failed", OccurredAt: end}}
	}
	return Result{ExecutionTimeMS: c.DurationMS, Timestamp: start, HealthStatus: c.Status, Warnings: []Warning{}, Errors: c.Errors, CollectedData: c, Metadata: map[string]string{}}
}
func classify(err error) inventory.CategoryStatus {
	if errors.Is(err, ErrUnsupported) {
		return inventory.Unsupported
	}
	if errors.Is(err, context.DeadlineExceeded) {
		return inventory.Timeout
	}
	if errors.Is(err, context.Canceled) {
		return inventory.Cancelled
	}
	if errors.Is(err, os.ErrPermission) {
		return inventory.PermissionDenied
	}
	if errors.Is(err, os.ErrNotExist) {
		return inventory.Unavailable
	}
	return inventory.Error
}
func fact(v any, source string) inventory.Fact {
	return inventory.Fact{Value: v, Quality: "observed", Sensitivity: "operational", Provenance: inventory.Provenance{SourceType: "kernel_virtual_file", SourceLabel: source, ObservedAt: time.Now().UTC()}}
}
func factUnit(v any, unit, source string) inventory.Fact {
	f := fact(v, source)
	f.Unit = unit
	return f
}
func redactedFact(sensitivity, reason, sourceType, source string) inventory.Fact {
	return inventory.Fact{Quality: "redacted", Sensitivity: sensitivity, Reason: reason, Provenance: inventory.Provenance{SourceType: sourceType, SourceLabel: source, ObservedAt: time.Now().UTC()}}
}
func item(id, kind string, facts map[string]inventory.Fact) inventory.Item {
	return inventory.Item{ID: id, Kind: kind, Facts: facts}
}
func readBounded(path string, max int) (string, error) {
	f, e := os.Open(path)
	if e != nil {
		return "", e
	}
	defer f.Close()
	b := make([]byte, max+1)
	n, e := f.Read(b)
	if n > max {
		return "", runner.ErrOutputLimit
	}
	if e != nil && !errors.Is(e, os.ErrClosed) && n == 0 {
		return "", e
	}
	return string(b[:n]), nil
}

func Default(r runner.Runner) []Collector {
	return []Collector{
		Func{Name: "collector_capabilities", Run: func(context.Context) ([]inventory.Item, []string, error) {
			ids := []string{"host", "os", "kernel", "cpu", "memory", "storage", "filesystem", "network", "virtualization", "services", "components"}
			out := make([]inventory.Item, 0, len(ids))
			for _, id := range ids {
				out = append(out, item(id, "collector_capability", map[string]inventory.Fact{"available": fact(true, "registry")}))
			}
			return out, []string{"collector registry"}, nil
		}},
		Func{Name: "host", Run: collectHost},
		Func{Name: "os", Run: collectOS},
		Func{Name: "kernel", Run: collectKernel},
		Func{Name: "cpu", Run: collectCPU},
		Func{Name: "memory", Run: collectMemory},
		Func{Name: "storage", Run: collectStorage},
		Func{Name: "filesystem", Dependencies: []string{"storage"}, Run: collectFilesystem},
		Func{Name: "network", Run: collectNetwork},
		Func{Name: "virtualization", Dependencies: []string{"host"}, Run: collectVirtualization},
		Func{Name: "services", Run: collectServices(r)},
		Func{Name: "components", Run: collectComponents(r)},
	}
}

func DefaultRegistry(r runner.Runner) (*Registry, error) {
	registry := NewRegistry()
	for _, c := range Default(r) {
		if err := registry.Register(c); err != nil {
			return nil, err
		}
	}
	return registry, nil
}
func collectOS(ctx context.Context) ([]inventory.Item, []string, error) {
	if err := ctx.Err(); err != nil {
		return nil, nil, err
	}
	s, e := readBounded("/etc/os-release", 64<<10)
	if e != nil {
		return nil, nil, e
	}
	vals := parseKeyValue(s)
	if vals["ID"] == "" {
		return nil, nil, ErrInvalidEvidence
	}
	return []inventory.Item{item("operating-system", "operating_system", map[string]inventory.Fact{"distribution_id": fact(vals["ID"], "os-release"), "version_id": fact(vals["VERSION_ID"], "os-release"), "variant_id": fact(vals["VARIANT_ID"], "os-release"), "architecture": fact(runtime.GOARCH, "Go runtime")})}, []string{"os-release", "Go runtime"}, nil
}
func parseKeyValue(s string) map[string]string {
	vals := map[string]string{}
	for _, l := range strings.Split(s, "\n") {
		l = strings.TrimSpace(l)
		if l == "" || strings.HasPrefix(l, "#") {
			continue
		}
		p := strings.SplitN(l, "=", 2)
		if len(p) == 2 {
			vals[p[0]] = strings.Trim(strings.TrimSpace(p[1]), "\"'")
		}
	}
	return vals
}
func unameFields() (map[string]string, error) {
	var u syscall.Utsname
	if err := syscall.Uname(&u); err != nil {
		return nil, err
	}
	return map[string]string{"sysname": utsString(u.Sysname[:]), "release": utsString(u.Release[:]), "version": utsString(u.Version[:]), "machine": utsString(u.Machine[:])}, nil
}
func utsString(value []int8) string {
	b := make([]byte, 0, len(value))
	for _, c := range value {
		if c == 0 {
			break
		}
		b = append(b, byte(c))
	}
	return string(b)
}
func collectKernel(ctx context.Context) ([]inventory.Item, []string, error) {
	if err := ctx.Err(); err != nil {
		return nil, nil, err
	}
	fields, err := unameFields()
	if err != nil {
		return nil, nil, err
	}
	return []inventory.Item{item("running-kernel", "kernel", map[string]inventory.Fact{"name": fact(fields["sysname"], "uname"), "release": fact(fields["release"], "uname"), "version": fact(fields["version"], "uname"), "architecture": fact(fields["machine"], "uname")})}, []string{"kernel uname"}, nil
}
func collectHost(ctx context.Context) ([]inventory.Item, []string, error) {
	if err := ctx.Err(); err != nil {
		return nil, nil, err
	}
	h, e := os.Hostname()
	if e != nil {
		return nil, nil, e
	}
	identityEvidence := h
	identitySource := "hostname fallback"
	if machineID, err := readBounded("/etc/machine-id", 4<<10); err == nil && strings.TrimSpace(machineID) != "" {
		identityEvidence = strings.TrimSpace(machineID)
		identitySource = "machine-id"
	}
	sum := sha256.Sum256([]byte("qwsg:subject:v1:" + identityEvidence))
	id := hex.EncodeToString(sum[:16])
	return []inventory.Item{item("local-host", "host", map[string]inventory.Fact{"display_label": redactedFact("host_identifying", "hostname_hidden", "system_api", "hostname"), "instance_id": fact(id, identitySource), "architecture": fact(runtime.GOARCH, "Go runtime")})}, []string{"privacy-safe subject identity", "system hostname (redacted)"}, nil
}
func collectCPU(ctx context.Context) ([]inventory.Item, []string, error) {
	if err := ctx.Err(); err != nil {
		return nil, nil, err
	}
	data, err := readBounded("/proc/cpuinfo", 4<<20)
	if err != nil {
		return nil, nil, err
	}
	fields := parseCPUInfo(data)
	fields["logical_cpus"] = runtime.NumCPU()
	fields["architecture"] = runtime.GOARCH
	facts := make(map[string]inventory.Fact, len(fields))
	for name, value := range fields {
		facts[name] = fact(value, "proc cpuinfo")
	}
	return []inventory.Item{item("cpu-summary", "cpu", facts)}, []string{"proc cpuinfo", "Go runtime"}, nil
}
func parseCPUInfo(data string) map[string]any {
	values := map[string]any{}
	physical := map[string]bool{}
	for _, line := range strings.Split(data, "\n") {
		parts := strings.SplitN(line, ":", 2)
		if len(parts) != 2 {
			continue
		}
		key, value := strings.TrimSpace(parts[0]), strings.TrimSpace(parts[1])
		switch key {
		case "vendor_id":
			if _, ok := values["vendor_id"]; !ok {
				values["vendor_id"] = value
			}
		case "model name":
			if _, ok := values["model_name"]; !ok {
				values["model_name"] = value
			}
		case "physical id":
			physical[value] = true
		}
	}
	if len(physical) > 0 {
		values["physical_packages"] = len(physical)
	}
	return values
}
func collectMemory(ctx context.Context) ([]inventory.Item, []string, error) {
	if err := ctx.Err(); err != nil {
		return nil, nil, err
	}
	s, e := readBounded("/proc/meminfo", 1<<20)
	if e != nil {
		return nil, nil, e
	}
	v, err := parseMemInfo(s)
	if err != nil {
		return nil, nil, err
	}
	return []inventory.Item{item("memory-summary", "memory", map[string]inventory.Fact{"total_bytes": factUnit(v["MemTotal"], "bytes", "proc meminfo"), "available_bytes": factUnit(v["MemAvailable"], "bytes", "proc meminfo"), "swap_total_bytes": factUnit(v["SwapTotal"], "bytes", "proc meminfo"), "swap_free_bytes": factUnit(v["SwapFree"], "bytes", "proc meminfo")})}, []string{"proc meminfo"}, nil
}
func parseMemInfo(s string) (map[string]int64, error) {
	v := map[string]int64{}
	sc := bufio.NewScanner(strings.NewReader(s))
	for sc.Scan() {
		p := strings.Fields(sc.Text())
		if len(p) >= 2 {
			n, _ := strconv.ParseInt(p[1], 10, 64)
			v[strings.TrimSuffix(p[0], ":")] = n * 1024
		}
	}
	if err := sc.Err(); err != nil || v["MemTotal"] <= 0 {
		return nil, ErrInvalidEvidence
	}
	return v, nil
}
func collectStorage(ctx context.Context) ([]inventory.Item, []string, error) {
	entries, err := os.ReadDir("/sys/class/block")
	if err != nil {
		return nil, nil, err
	}
	out := []inventory.Item{}
	for _, entry := range entries {
		if err := ctx.Err(); err != nil {
			return nil, nil, err
		}
		name := entry.Name()
		sizeText, err := readBounded("/sys/class/block/"+name+"/size", 128)
		if err != nil {
			continue
		}
		sectors, err := strconv.ParseInt(strings.TrimSpace(sizeText), 10, 64)
		if err != nil || sectors < 0 {
			continue
		}
		kind := "block_device"
		if _, err := os.Stat("/sys/class/block/" + name + "/partition"); err == nil {
			kind = "partition"
		}
		id := privacyID("block-device", name)
		facts := map[string]inventory.Fact{"device_name": redactedFact("host_identifying", "device_name_hidden", "kernel_virtual_file", "sys block"), "size_bytes": factUnit(sectors*512, "bytes", "sys block size"), "device_type": fact(kind, "sys block")}
		if value, err := readBounded("/sys/class/block/"+name+"/removable", 16); err == nil {
			facts["removable"] = fact(strings.TrimSpace(value) == "1", "sys block removable")
		}
		out = append(out, item(id, kind, facts))
	}
	sort.Slice(out, func(i, j int) bool { return out[i].ID < out[j].ID })
	if len(out) == 0 {
		return nil, nil, ErrInvalidEvidence
	}
	return out, []string{"sysfs block device metadata; device names redacted"}, nil
}
func collectFilesystem(ctx context.Context) ([]inventory.Item, []string, error) {
	if err := ctx.Err(); err != nil {
		return nil, nil, err
	}
	s, e := readBounded("/proc/self/mounts", 2<<20)
	if e != nil {
		return nil, nil, e
	}
	var out []inventory.Item
	seen := map[string]bool{}
	sc := bufio.NewScanner(strings.NewReader(s))
	for sc.Scan() {
		if err := ctx.Err(); err != nil {
			return nil, nil, err
		}
		p := strings.Fields(sc.Text())
		if len(p) < 4 || seen[p[1]] {
			continue
		}
		seen[p[1]] = true
		var st syscall.Statfs_t
		if syscall.Statfs(p[1], &st) != nil {
			continue
		}
		id := privacyID("mount", p[1])
		out = append(out, item(id, "filesystem", map[string]inventory.Fact{"mount_path": redactedFact("host_identifying", "mount_path_hidden", "filesystem_metadata", "mount table"), "filesystem_type": fact(p[2], "mount table"), "total_bytes": factUnit(int64(st.Blocks)*int64(st.Bsize), "bytes", "statfs"), "available_bytes": factUnit(int64(st.Bavail)*int64(st.Bsize), "bytes", "statfs"), "read_only": fact(hasMountOption(p[3], "ro"), "mount table")}))
	}
	if err := sc.Err(); err != nil {
		return nil, nil, err
	}
	sort.Slice(out, func(i, j int) bool { return out[i].ID < out[j].ID })
	return out, []string{"process mount table", "statfs"}, nil
}
func hasMountOption(options, wanted string) bool {
	for _, option := range strings.Split(options, ",") {
		if option == wanted {
			return true
		}
	}
	return false
}
func privacyID(namespace, evidence string) string {
	id := sha256.Sum256([]byte("qwsg:" + namespace + ":v1:" + evidence))
	return hex.EncodeToString(id[:16])
}
func collectNetwork(ctx context.Context) ([]inventory.Item, []string, error) {
	if err := ctx.Err(); err != nil {
		return nil, nil, err
	}
	interfaces, e := os.ReadDir("/sys/class/net")
	if e != nil {
		return nil, nil, e
	}
	out := make([]inventory.Item, 0, len(interfaces))
	for _, in := range interfaces {
		if err := ctx.Err(); err != nil {
			return nil, nil, err
		}
		name := in.Name()
		mtu := int64(0)
		if value, err := readBounded("/sys/class/net/"+name+"/mtu", 64); err == nil {
			mtu, _ = strconv.ParseInt(strings.TrimSpace(value), 10, 64)
		}
		up := false
		if value, err := readBounded("/sys/class/net/"+name+"/operstate", 64); err == nil {
			up = strings.TrimSpace(value) == "up"
		}
		loopback := false
		if value, err := readBounded("/sys/class/net/"+name+"/type", 64); err == nil {
			loopback = strings.TrimSpace(value) == "772"
		}
		out = append(out, item(privacyID("network-interface", name), "interface", map[string]inventory.Fact{"name": redactedFact("network_sensitive", "interface_name_hidden", "kernel_virtual_file", "sys network"), "up": fact(up, "sys network operstate"), "loopback": fact(loopback, "sys network type"), "mtu_bytes": factUnit(mtu, "bytes", "sys network mtu"), "addresses": redactedFact("network_sensitive", "network_addresses_hidden", "kernel_interface", "network addresses")}))
	}
	sort.Slice(out, func(i, j int) bool { return out[i].ID < out[j].ID })
	if len(out) == 0 {
		return nil, nil, ErrInvalidEvidence
	}
	return out, []string{"sysfs network metadata; names, addresses, and hardware identifiers redacted"}, nil
}
func collectVirtualization(ctx context.Context) ([]inventory.Item, []string, error) {
	if err := ctx.Err(); err != nil {
		return nil, nil, err
	}
	kind, technology, sources := detectVirtualization()
	return []inventory.Item{item("virtualization-context", "virtualization_context", map[string]inventory.Fact{"kind": fact(kind, "virtualization evidence"), "technology": fact(technology, "virtualization evidence")})}, sources, nil
}
func detectVirtualization() (string, string, []string) {
	if _, err := os.Stat("/.dockerenv"); err == nil {
		return "container", "docker_compatible", []string{"container marker"}
	}
	if cgroup, err := readBounded("/proc/1/cgroup", 1<<20); err == nil {
		lower := strings.ToLower(cgroup)
		for _, token := range []string{"containerd", "kubepods", "docker", "lxc"} {
			if strings.Contains(lower, token) {
				return "container", token, []string{"process 1 cgroup"}
			}
		}
	}
	if product, err := readBounded("/sys/class/dmi/id/product_name", 4<<10); err == nil {
		lower := strings.ToLower(strings.TrimSpace(product))
		technologies := []struct {
			name   string
			tokens []string
		}{{"kvm", []string{"kvm", "qemu"}}, {"vmware", []string{"vmware"}}, {"hyper_v", []string{"virtual machine"}}, {"virtualbox", []string{"virtualbox"}}}
		for _, technology := range technologies {
			for _, token := range technology.tokens {
				if strings.Contains(lower, token) {
					return "virtual_machine", technology.name, []string{"DMI product metadata"}
				}
			}
		}
	}
	return "none_detected", "unknown", []string{"bounded container and DMI evidence"}
}
func collectServices(r runner.Runner) func(context.Context) ([]inventory.Item, []string, error) {
	return func(ctx context.Context) ([]inventory.Item, []string, error) {
		res, e := r.Run(ctx, "systemctl", "list-units", "--type=service", "--state=running", "--no-legend", "--plain")
		if e != nil {
			return nil, nil, e
		}
		lines := strings.Split(strings.TrimSpace(string(res.Stdout)), "\n")
		out := []inventory.Item{}
		for i, l := range lines {
			if strings.TrimSpace(l) == "" {
				continue
			}
			out = append(out, item(strconv.Itoa(i+1), "service", map[string]inventory.Fact{"service_identity": redactedFact("operational", "service_identity_hidden", "service_manager", "systemd unit metadata"), "active": fact(true, "systemd unit metadata")}))
		}
		return out, []string{"systemd running-unit metadata; names redacted"}, nil
	}
}
func collectComponents(r runner.Runner) func(context.Context) ([]inventory.Item, []string, error) {
	return func(ctx context.Context) ([]inventory.Item, []string, error) {
		ids := []string{"go"}
		out := []inventory.Item{}
		for _, id := range ids {
			res, e := r.Run(ctx, id, "version")
			if e != nil {
				continue
			}
			fields := strings.Fields(string(res.Stdout))
			version := "unknown"
			if len(fields) > 2 {
				version = fields[2]
			}
			out = append(out, item(id, "runtime_component", map[string]inventory.Fact{"version": fact(version, "allowlisted version command")}))
		}
		if len(out) == 0 {
			return nil, nil, os.ErrNotExist
		}
		return out, []string{"allowlisted version commands"}, nil
	}
}

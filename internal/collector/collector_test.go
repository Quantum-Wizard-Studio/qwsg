package collector

import (
	"context"
	"errors"
	"os"
	"quantumwizard.hu/qwsg/internal/inventory"
	"quantumwizard.hu/qwsg/internal/runner"
	"reflect"
	"sort"
	"testing"
)

type fakeRunner struct{}

func (fakeRunner) Run(_ context.Context, id string, _ ...string) (runner.Result, error) {
	if id == "systemctl" {
		return runner.Result{Stdout: []byte("secret.service loaded active running hidden\n")}, nil
	}
	if id == "go" {
		return runner.Result{Stdout: []byte("go version go1.test linux/amd64\n")}, nil
	}
	return runner.Result{}, errors.New("unexpected")
}
func TestEveryDefaultCollectorReturnsContract(t *testing.T) {
	cs := Default(fakeRunner{})
	want := []string{"collector_capabilities", "components", "cpu", "filesystem", "host", "kernel", "memory", "network", "os", "services", "storage", "virtualization"}
	got := make([]string, 0, len(cs))
	for _, c := range cs {
		d := c.Descriptor()
		r := c.Collect(context.Background(), Request{Capability: d.Capability.ID})
		got = append(got, d.Capability.ID)
		if r.CollectedData.CategoryID == "" || r.CollectedData.ContractVersion != "1.0" || r.CollectedData.PrivilegeUsed != "ordinary_user" {
			t.Fatalf("bad contract: %#v", r)
		}
	}
	sort.Strings(got)
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("collectors %v, want %v", got, want)
	}
}

func TestCanonicalCollectorDependencies(t *testing.T) {
	dependencies := map[string][]string{}
	for _, c := range Default(fakeRunner{}) {
		dependencies[c.Descriptor().Name] = c.Descriptor().Dependencies
	}
	if !reflect.DeepEqual(dependencies["filesystem"], []string{"storage"}) || !reflect.DeepEqual(dependencies["virtualization"], []string{"host"}) {
		t.Fatalf("unexpected dependencies: %#v", dependencies)
	}
}

func TestLinuxEvidenceParsers(t *testing.T) {
	osRelease := parseKeyValue("# comment\nID=debian\nVERSION_ID=\"12\"\nVARIANT_ID='server'\n")
	if osRelease["ID"] != "debian" || osRelease["VERSION_ID"] != "12" || osRelease["VARIANT_ID"] != "server" {
		t.Fatalf("os-release: %#v", osRelease)
	}
	cpu := parseCPUInfo("vendor_id : GenuineIntel\nmodel name : Test CPU\nphysical id : 0\n\nphysical id : 1\n")
	if cpu["vendor_id"] != "GenuineIntel" || cpu["model_name"] != "Test CPU" || cpu["physical_packages"] != 2 {
		t.Fatalf("cpuinfo: %#v", cpu)
	}
	memory, err := parseMemInfo("MemTotal: 1024 kB\nMemAvailable: 512 kB\nSwapTotal: 128 kB\nSwapFree: 64 kB\n")
	if err != nil || memory["MemTotal"] != 1048576 || memory["SwapFree"] != 65536 {
		t.Fatalf("meminfo: %#v, %v", memory, err)
	}
	if _, err := parseMemInfo("MemFree: 2 kB\n"); !errors.Is(err, ErrInvalidEvidence) {
		t.Fatalf("expected invalid evidence, got %v", err)
	}
	if !hasMountOption("rw,nosuid,nodev,ro", "ro") || hasMountOption("errors=remount-ro,rw", "ro") {
		t.Fatal("mount option parsing is not token exact")
	}
}

func TestPrivacyIDsAreStableAndNamespaced(t *testing.T) {
	a := privacyID("interface", "eth0")
	if a != privacyID("interface", "eth0") || a == privacyID("device", "eth0") || len(a) != 32 {
		t.Fatalf("invalid privacy ID behavior: %q", a)
	}
}
func TestPermissionDenied(t *testing.T) {
	f := Func{Name: "x", Run: func(context.Context) ([]inventory.Item, []string, error) { return nil, nil, os.ErrPermission }}
	if got := f.Collect(context.Background(), Request{}).HealthStatus; got != "permission_denied" {
		t.Fatalf("got %s", got)
	}
}

func TestUnsupportedAndParseError(t *testing.T) {
	cases := []struct {
		err  error
		want inventory.CategoryStatus
	}{{ErrUnsupported, inventory.Unsupported}, {ErrInvalidEvidence, inventory.Error}}
	for _, tc := range cases {
		f := Func{Name: "x", Run: func(context.Context) ([]inventory.Item, []string, error) { return nil, nil, tc.err }}
		if got := f.Collect(context.Background(), Request{}).HealthStatus; got != tc.want {
			t.Fatalf("got %s, want %s", got, tc.want)
		}
	}
}

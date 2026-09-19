package updatemutation

import (
	"context"
	"errors"
	"os"
	"os/exec"
	"path/filepath"
	"syscall"
	"testing"
	"time"
)

func TestLockProcessProbe(t *testing.T) {
	root := os.Getenv("QWSG_MUTATION_LOCK_PROBE")
	if root == "" {
		return
	}
	lease, err := Acquire(root)
	if os.Getenv("QWSG_MUTATION_EXPECT_BUSY") == "1" {
		if !errors.Is(err, ErrContended) {
			t.Fatalf("expected contention, got %v", err)
		}
		return
	}
	if err != nil {
		t.Fatal(err)
	}
	if err = lease.Close(); err != nil {
		t.Fatal(err)
	}
}

func TestProcessContentionAndRelease(t *testing.T) {
	root := t.TempDir()
	if err := os.Chmod(root, 0700); err != nil {
		t.Fatal(err)
	}
	lease, err := Acquire(root)
	if err != nil {
		t.Fatal(err)
	}
	defer lease.Close()
	probe := func(busy string) {
		t.Helper()
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		cmd := exec.CommandContext(ctx, os.Args[0], "-test.run=^TestLockProcessProbe$")
		cmd.Env = append(os.Environ(), "QWSG_MUTATION_LOCK_PROBE="+root, "QWSG_MUTATION_EXPECT_BUSY="+busy)
		if out, err := cmd.CombinedOutput(); err != nil {
			t.Fatalf("probe: %v %s", err, out)
		}
	}
	probe("1")
	if _, err = Acquire(root); !errors.Is(err, ErrContended) {
		t.Fatalf("in-process reentry: %v", err)
	}
	before, err := os.Stat(filepath.Join(root, "automatic.lock"))
	if err != nil {
		t.Fatal(err)
	}
	if err = lease.Close(); err != nil {
		t.Fatal(err)
	}
	probe("0")
	after, err := os.Stat(filepath.Join(root, "automatic.lock"))
	if err != nil || !os.SameFile(before, after) {
		t.Fatal("lock inode was replaced")
	}
	// Historical automatic callers use the same inode and must still exclude us.
	f, err := os.OpenFile(filepath.Join(root, "automatic.lock"), os.O_RDWR, 0600)
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	if err = syscall.Flock(int(f.Fd()), syscall.LOCK_EX|syscall.LOCK_NB); err != nil {
		t.Fatal(err)
	}
	probe("1")
}

func TestUnsafeLockRefused(t *testing.T) {
	for _, scenario := range []string{"directory mode", "directory symlink", "lock mode", "lock symlink", "lock directory"} {
		t.Run(scenario, func(t *testing.T) {
			root := t.TempDir()
			if err := os.Chmod(root, 0700); err != nil {
				t.Fatal(err)
			}
			path := filepath.Join(root, "automatic.lock")
			var err error
			switch scenario {
			case "directory mode":
				err = os.Chmod(root, 0755)
			case "directory symlink":
				link := filepath.Join(t.TempDir(), "link")
				err = os.Symlink(root, link)
				root = link
			case "lock mode":
				err = os.WriteFile(path, nil, 0644)
			case "lock symlink":
				err = os.Symlink(filepath.Join(root, "missing"), path)
			case "lock directory":
				err = os.Mkdir(path, 0700)
			}
			if err != nil {
				t.Fatal(err)
			}
			lease, err := Acquire(root)
			if err == nil {
				lease.Close()
				t.Fatal("unsafe lock accepted")
			}
			if errors.Is(err, ErrContended) {
				t.Fatal("unsafe lock confused with contention")
			}
		})
	}
}

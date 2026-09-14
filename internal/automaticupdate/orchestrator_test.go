package automaticupdate

import (
	"context"
	"os"
	"path/filepath"
	"testing"

	"quantumwizard.hu/qwsg/internal/productcapability"
	"quantumwizard.hu/qwsg/internal/updatepolicy"
)

type unusedHost struct{}

func (unusedHost) Preflight(context.Context, string, string) error { panic("must not enter host") }
func (unusedHost) Apply(context.Context, PackageInput) (ApplyResult, error) {
	panic("must not enter host")
}
func (unusedHost) Validate(context.Context, string) error       { panic("must not enter host") }
func (unusedHost) Rollback(context.Context) error               { panic("must not enter host") }
func (unusedHost) Commit(context.Context, string, string) error { panic("must not enter host") }

func TestIndependentFileLockAndUnsafeLockRefuse(t *testing.T) {
	caps, _ := productcapability.Resolve(&productcapability.Declaration{Schema: productcapability.Schema, Product: productcapability.Pro, Grants: []productcapability.Capability{productcapability.UpdateManual, productcapability.UpdateAutomatic}})
	for _, scenario := range []string{"held descriptor", "symlink", "unsafe mode"} {
		t.Run(scenario, func(t *testing.T) {
			root := t.TempDir()
			if err := os.Chmod(root, 0700); err != nil {
				t.Fatal(err)
			}
			switch scenario {
			case "held descriptor":
				f, err := acquire(root)
				if err != nil {
					t.Fatal(err)
				}
				defer f.Close()
			case "symlink":
				if err := os.Symlink(filepath.Join(root, "missing"), filepath.Join(root, "automatic.lock")); err != nil {
					t.Fatal(err)
				}
			case "unsafe mode":
				if err := os.Chmod(root, 0755); err != nil {
					t.Fatal(err)
				}
			}
			r, err := Run(context.Background(), Request{Capabilities: caps, Policy: updatepolicy.Request{Mode: updatepolicy.Automatic}, StageParent: root}, unusedHost{})
			if err == nil || r.FailureCategory != "transaction_conflict" || r.MutationStarted {
				t.Fatalf("lock bypassed: %+v %v", r, err)
			}
		})
	}
}

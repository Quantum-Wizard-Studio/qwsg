package main

import (
	"bytes"
	"context"
	"errors"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"quantumwizard.hu/qwsg/internal/configurationstore"
	"quantumwizard.hu/qwsg/internal/productcapability"
)

func proTestAuthority() (productcapability.Set, error) {
	return productcapability.Resolve(&productcapability.Declaration{Schema: productcapability.Schema, Product: productcapability.Pro, Grants: []productcapability.Capability{productcapability.UpdateManual, productcapability.UpdateAutomatic}})
}

func TestProPolicyConfigurationHasNoUpdateSideEffects(t *testing.T) {
	home := t.TempDir()
	t.Setenv("HOME", home)
	t.Setenv("XDG_CONFIG_HOME", filepath.Join(home, "config"))
	t.Setenv("QWSG_STATE_DIR", filepath.Join(home, "state"))
	previous, fetch := installationCapabilities, updateMetadataFetch
	installationCapabilities = proTestAuthority
	updateMetadataFetch = func(context.Context) ([]byte, error) {
		t.Fatal("policy operation fetched update metadata")
		return nil, nil
	}
	t.Cleanup(func() { installationCapabilities = previous; updateMetadataFetch = fetch })
	var out, errout bytes.Buffer
	for _, args := range [][]string{{"config", "set", "update.policy", "automatic"}, {"config", "validate"}, {"update", "status"}} {
		out.Reset()
		errout.Reset()
		if code := run(args, &out, &errout); code != 0 {
			t.Fatalf("%v: %d %s", args, code, errout.String())
		}
	}
	if !strings.Contains(out.String(), "Update policy: automatic") || !strings.Contains(out.String(), "permitted (policy only)") {
		t.Fatal(out.String())
	}
	if _, err := os.Stat(filepath.Join(home, "state", "update")); !os.IsNotExist(err) {
		t.Fatalf("policy created transaction state: %v", err)
	}
	path, _ := configurationstore.DefaultPath(os.Getenv)
	before, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	installationCapabilities = previous
	out.Reset()
	errout.Reset()
	if code := run([]string{"config", "validate"}, &out, &errout); code == 0 || !strings.Contains(errout.String(), "update.automatic") {
		t.Fatalf("lost entitlement accepted: %d %s", code, errout.String())
	}
	after, _ := os.ReadFile(path)
	if !bytes.Equal(before, after) {
		t.Fatal("validation rewrote config")
	}
}

func TestCommunityPolicyVisibilityAndRefusalPreservesConfig(t *testing.T) {
	home := t.TempDir()
	t.Setenv("HOME", home)
	t.Setenv("XDG_CONFIG_HOME", filepath.Join(home, "config"))
	t.Setenv("QWSG_STATE_DIR", filepath.Join(home, "state"))
	var out, errout bytes.Buffer
	if code := run([]string{"update", "status"}, &out, &errout); code != 0 || !strings.Contains(out.String(), "Update policy: manual") || !strings.Contains(out.String(), "unavailable (not entitled)") {
		t.Fatalf("%d %s %s", code, out.String(), errout.String())
	}
	for _, value := range []string{"automatic", "unknown"} {
		errout.Reset()
		if code := run([]string{"config", "set", "update.policy", value}, &out, &errout); code == 0 {
			t.Fatal("unauthorized policy accepted")
		}
	}
	path, _ := configurationstore.DefaultPath(os.Getenv)
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		t.Fatal("refused policy wrote config")
	}
}

func TestUnknownAuthorityRefusesBeforeUpdateWork(t *testing.T) {
	previous, fetch := installationCapabilities, updateMetadataFetch
	t.Cleanup(func() { installationCapabilities = previous; updateMetadataFetch = fetch })
	updateMetadataFetch = func(context.Context) ([]byte, error) {
		t.Fatal("invalid authority reached update engine")
		return nil, nil
	}
	for _, authority := range []func() (productcapability.Set, error){
		func() (productcapability.Set, error) { return productcapability.Set{}, nil },
		func() (productcapability.Set, error) {
			s, _ := proTestAuthority()
			return s, errors.New("authority unavailable")
		},
	} {
		installationCapabilities = authority
		var out, errout bytes.Buffer
		if code := executeUpdate("", "", &out, &errout); code == 0 || !strings.Contains(errout.String(), "manual update capability unavailable") {
			t.Fatalf("%d %s", code, errout.String())
		}
	}
}

package main

import (
	"bytes"
	"io"
	"path/filepath"
	"testing"

	"quantumwizard.hu/qwsg/internal/changenotification"
)

func TestConfigurationChangeUsesManagedNotificationHook(t *testing.T) {
	home := t.TempDir()
	t.Setenv("HOME", home)
	t.Setenv("XDG_CONFIG_HOME", filepath.Join(home, "config"))
	previous := managedChangeDelivery
	defer func() { managedChangeDelivery = previous }()
	var got changenotification.Event
	managedChangeDelivery = func(e changenotification.Event, _ io.Writer) changenotification.DeliveryResult {
		got = e
		return changenotification.DeliveryAccepted
	}
	var out, errout bytes.Buffer
	if code := run([]string{"setup", "--accept-defaults"}, &out, &errout); code != 0 {
		t.Fatalf("setup=%d %s", code, errout.String())
	}
	if code := run([]string{"config", "set", "notification.lifecycle.enabled", "true"}, &out, &errout); code != 0 {
		t.Fatalf("set=%d %s", code, errout.String())
	}
	if got.Type != changenotification.ConfigurationChanged || got.Result != changenotification.Success {
		t.Fatalf("event=%+v", got)
	}
	out.Reset()
	if code := run([]string{"config", "get", "notification.lifecycle.enabled"}, &out, &errout); code != 0 || out.String() != "true\n" {
		t.Fatalf("get=%d %q %s", code, out.String(), errout.String())
	}
}

package main

import (
	"testing"

	"quantumwizard.hu/qwsg/internal/configuration"
)

func TestUpdateNotificationRequiresExplicitNotifyPolicy(t *testing.T) {
	for name, effective := range map[string]configuration.Effective{
		"absent": {},
		"manual": {Values: configuration.Model{Extensions: []configuration.Extension{{ID: "installer.update-policy", Fields: map[string]string{"policy": "manual"}}}}},
		"notify": {Values: configuration.Model{Extensions: []configuration.Extension{{ID: "installer.update-policy", Fields: map[string]string{"policy": "notify"}}}}},
	} {
		want := name == "notify"
		if got := updateNotificationEnabled(effective); got != want {
			t.Fatalf("%s: got %t want %t", name, got, want)
		}
	}
}

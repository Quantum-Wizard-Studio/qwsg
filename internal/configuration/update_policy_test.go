package configuration

import (
	"quantumwizard.hu/qwsg/internal/updatepolicy"
	"testing"
)

func TestUpdatePolicyBackwardCompatibility(t *testing.T) {
	base := mustBuiltIn(t)
	bytes, err := MarshalSourceCanonical(base)
	if err != nil {
		t.Fatal(err)
	}
	old, err := DecodeSource(bytes)
	if err != nil {
		t.Fatal(err)
	}
	effective, err := Resolve([]Source{old})
	if err != nil {
		t.Fatal(err)
	}
	request, err := UpdatePolicy(effective.Values)
	if err != nil || request.Mode != updatepolicy.Manual || request.Notify {
		t.Fatalf("old default: %+v %v", request, err)
	}
	for _, value := range []string{"manual", "notify", "automatic"} {
		base.Identity = ""
		extensions := []Extension{{ID: updatepolicy.ExtensionID, Version: "1.0", Fields: map[string]string{"policy": value}}}
		base.Patch.Extensions = &extensions
		resolved, err := Resolve([]Source{base})
		if err != nil {
			t.Fatal(err)
		}
		request, err := UpdatePolicy(resolved.Values)
		if err != nil || request.Notify != (value == "notify") || (request.Mode == updatepolicy.Automatic) != (value == "automatic") {
			t.Fatalf("%s: %+v %v", value, request, err)
		}
	}
}

func TestMalformedUpdatePolicyRefusedByConfiguration(t *testing.T) {
	for _, tc := range []struct {
		name, version       string
		fields              map[string]string
		required, duplicate bool
	}{
		{"unknown", "1.0", map[string]string{"policy": "future"}, false, false},
		{"empty", "1.0", map[string]string{"policy": ""}, false, false},
		{"missing", "1.0", map[string]string{}, false, false},
		{"schema", "2.0", map[string]string{"policy": "automatic"}, false, false},
		{"extra", "1.0", map[string]string{"policy": "automatic", "entitled": "true"}, false, false},
		{"required", "1.0", map[string]string{"policy": "manual"}, true, false},
		{"duplicate", "1.0", map[string]string{"policy": "manual"}, false, true},
	} {
		t.Run(tc.name, func(t *testing.T) {
			base := mustBuiltIn(t)
			base.Identity = ""
			extensions := []Extension{{ID: updatepolicy.ExtensionID, Version: tc.version, Fields: tc.fields, Required: tc.required}}
			if tc.duplicate {
				extensions = append(extensions, extensions[0])
			}
			base.Patch.Extensions = &extensions
			if _, err := NormalizeSource(base); err == nil {
				t.Fatal("malformed extension accepted")
			}
			if _, err := UpdatePolicy(Model{Extensions: extensions}); err == nil {
				t.Fatal("raw malformed model accepted")
			}
		})
	}
}

package updatepolicy

import (
	"errors"
	"quantumwizard.hu/qwsg/internal/productcapability"
	"testing"
)

func TestPolicyEligibility(t *testing.T) {
	community, _ := productcapability.Resolve(nil)
	pro, _ := productcapability.Resolve(&productcapability.Declaration{Schema: productcapability.Schema, Product: productcapability.Pro, Grants: []productcapability.Capability{productcapability.UpdateManual, productcapability.UpdateAutomatic}})
	for _, tc := range []struct {
		name string
		mode Mode
		set  productcapability.Set
		want error
	}{
		{"community default", Manual, community, nil},
		{"community automatic", Automatic, community, ErrAutomaticUnavailable},
		{"pro automatic", Automatic, pro, nil},
		{"pro manual", Manual, pro, nil},
		{"unknown policy", "unknown", pro, ErrPolicy},
		{"empty policy", "", pro, ErrPolicy},
		{"unknown authority", Automatic, productcapability.Set{}, productcapability.ErrAuthority},
	} {
		t.Run(tc.name, func(t *testing.T) {
			got, err := Evaluate(Request{Mode: tc.mode}, tc.set)
			if !errors.Is(err, tc.want) {
				t.Fatalf("state=%+v err=%v", got, err)
			}
			if err == nil && (got.Effective != tc.mode || !got.ManualAllowed || got.AutomaticAllowed != tc.set.Has(productcapability.UpdateAutomatic)) {
				t.Fatalf("wrong state %+v", got)
			}
			if err != nil && got != (State{}) {
				t.Fatal("refusal returned usable state")
			}
		})
	}
	if _, err := Evaluate(Request{Mode: Automatic, Notify: true}, pro); err == nil {
		t.Fatal("conflicting intent accepted")
	}
}

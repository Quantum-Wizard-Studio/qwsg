package productcapability

import "testing"

func TestCommunityDefaultAndImmutablePro(t *testing.T) {
	c, err := Resolve(nil)
	if err != nil || c.Product() != Community || !c.Has(UpdateManual) || c.Has(UpdateAutomatic) || c.Has("unknown") {
		t.Fatalf("community: %+v %v", c, err)
	}
	d := Declaration{Schema: Schema, Product: Pro, Grants: []Capability{UpdateManual, UpdateAutomatic}}
	p, err := Resolve(&d)
	if err != nil || !p.Has(UpdateAutomatic) || !p.Has(UpdateManual) {
		t.Fatalf("pro: %+v %v", p, err)
	}
	d.Grants[0] = "unknown"
	if !p.Has(UpdateManual) {
		t.Fatal("authority aliased mutable declaration")
	}
	if (Set{}).Valid() || (Set{}).Has(UpdateManual) || (Set{}).Has(UpdateAutomatic) {
		t.Fatal("zero authority granted capability")
	}
}

func TestInvalidAuthorityFailsClosed(t *testing.T) {
	for _, d := range []Declaration{
		{},
		{Schema: Schema, Product: "unknown", Grants: []Capability{UpdateManual}},
		{Schema: "qwsg.product-capabilities/2", Product: Pro, Grants: []Capability{UpdateManual, UpdateAutomatic}},
		{Schema: Schema, Product: Pro},
		{Schema: Schema, Product: Pro, Grants: []Capability{UpdateAutomatic}},
		{Schema: Schema, Product: Pro, Grants: []Capability{UpdateManual, "unknown"}},
		{Schema: Schema, Product: Pro, Grants: []Capability{UpdateManual, UpdateManual}},
		{Schema: Schema, Product: Community, Grants: []Capability{UpdateManual, UpdateAutomatic}},
	} {
		got, err := Resolve(&d)
		if err == nil || got.Valid() || got.Has(UpdateManual) || got.Has(UpdateAutomatic) {
			t.Fatalf("unsafe declaration %+v: %+v %v", d, got, err)
		}
	}
}

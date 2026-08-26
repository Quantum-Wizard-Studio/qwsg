package update

import "testing"

func releaseFixture(version, host string) apiRelease {
	r := apiRelease{TagName: "v" + version}
	a := "qwsg-" + version + "-linux-amd64.tar.gz"
	r.Assets = []struct {
		Name string `json:"name"`
		Size int64  `json:"size"`
		URL  string `json:"browser_download_url"`
	}{{a, 20, host + "/Quantum_Wizard_Studio/qwsg/releases/download/v" + version + "/" + a}, {a + ".sha256", 96, host + "/Quantum_Wizard_Studio/qwsg/releases/download/v" + version + "/" + a + ".sha256"}}
	return r
}

func TestSelectNewestValidCanonicalRelease(t *testing.T) {
	records := []apiRelease{releaseFixture("1.1.0", "https://git.quantumwizard.hu"), releaseFixture("1.2.0", "https://git.quantumwizard.hu")}
	release, relation, err := selectRelease(records, "1.1.0")
	if err != nil || release.Version != "1.2.0" || relation != Newer {
		t.Fatalf("got %+v %s %v", release, relation, err)
	}
}
func TestRejectNonCanonicalAssetOrigin(t *testing.T) {
	if _, _, err := selectRelease([]apiRelease{releaseFixture("1.2.0", "https://example.invalid")}, "1.1.0"); err == nil {
		t.Fatal("noncanonical assets accepted")
	}
}
func TestIgnoreDraftRelease(t *testing.T) {
	draft := releaseFixture("1.3.0", "https://git.quantumwizard.hu")
	draft.Draft = true
	release, _, err := selectRelease([]apiRelease{draft, releaseFixture("1.2.0", "https://git.quantumwizard.hu")}, "1.1.0")
	if err != nil || release.Version != "1.2.0" {
		t.Fatalf("got %+v %v", release, err)
	}
}
func TestRejectAmbiguousReleaseIdentity(t *testing.T) {
	r := releaseFixture("1.2.0", "https://git.quantumwizard.hu")
	if _, _, err := selectRelease([]apiRelease{r, r}, "1.1.0"); err == nil {
		t.Fatal("ambiguous release accepted")
	}
}
func TestRejectUnexpectedReleaseAsset(t *testing.T) {
	r := releaseFixture("1.2.0", "https://git.quantumwizard.hu")
	r.Assets = append(r.Assets, r.Assets[0])
	if _, _, err := selectRelease([]apiRelease{r}, "1.1.0"); err == nil {
		t.Fatal("unexpected asset accepted")
	}
}

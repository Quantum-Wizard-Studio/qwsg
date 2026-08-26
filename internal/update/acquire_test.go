package update

import (
	"context"
	"io"
	"net/http"
	"os"
	"strings"
	"testing"
)

type roundTripFunc func(*http.Request) (*http.Response, error)

func (f roundTripFunc) RoundTrip(r *http.Request) (*http.Response, error) { return f(r) }
func TestAcquireVerifiesSidecarAndPrivateStaging(t *testing.T) {
	archive := []byte("archive")
	hash := bytesSHA(archive)
	name := "qwsg-1.2.0-linux-amd64.tar.gz"
	bodies := map[string]string{"https://git.quantumwizard.hu/a": string(archive), "https://git.quantumwizard.hu/s": hash + "  " + name + "\n"}
	client := &http.Client{Transport: roundTripFunc(func(r *http.Request) (*http.Response, error) {
		body := bodies[r.URL.String()]
		return &http.Response{StatusCode: 200, Body: io.NopCloser(strings.NewReader(body)), Header: make(http.Header)}, nil
	})}
	parent := t.TempDir()
	if err := os.Chmod(parent, 0700); err != nil {
		t.Fatal(err)
	}
	release := Release{Version: "1.2.0", Archive: Asset{Name: name, URL: "https://git.quantumwizard.hu/a", Size: int64(len(archive))}, Sidecar: Asset{Name: name + ".sha256", URL: "https://git.quantumwizard.hu/s", Size: int64(len(bodies["https://git.quantumwizard.hu/s"]))}}
	staged, err := Acquire(context.Background(), client, release, parent)
	if err != nil {
		t.Fatal(err)
	}
	if staged.SHA256 != hash {
		t.Fatal("hash mismatch")
	}
}
func TestAcquireRejectsSizeMismatch(t *testing.T) {
	client := &http.Client{Transport: roundTripFunc(func(r *http.Request) (*http.Response, error) {
		return &http.Response{StatusCode: 200, Body: io.NopCloser(strings.NewReader("x")), Header: make(http.Header)}, nil
	})}
	parent := t.TempDir()
	if err := os.Chmod(parent, 0700); err != nil {
		t.Fatal(err)
	}
	release := Release{Version: "1.2.0", Archive: Asset{Name: "a", URL: "https://git.quantumwizard.hu/a", Size: 2}, Sidecar: Asset{Name: "s", URL: "https://git.quantumwizard.hu/s", Size: 2}}
	if _, err := Acquire(context.Background(), client, release, parent); err == nil {
		t.Fatal("size mismatch accepted")
	}
}

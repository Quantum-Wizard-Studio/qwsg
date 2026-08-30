package releasediscovery

import (
	"context"
	"crypto/tls"
	"net/http"
	"net/http/cookiejar"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
	"time"
)

func TestStaticSourceBoundedAnonymousFetchAndValidators(t *testing.T) {
	index, _ := signedFixture(t)
	payload := mutateJSON(t, index, func(*Index) {})
	server := httptest.NewTLSServer(http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		if request.Method != http.MethodGet || request.URL.RawQuery != "" || request.Header.Get("Accept") != MediaType || request.Header.Get("User-Agent") != "QWSG-release-index/1" {
			t.Errorf("unexpected request: method=%s query=%q headers=%v", request.Method, request.URL.RawQuery, request.Header)
		}
		for _, forbidden := range []string{"Authorization", "Cookie", "X-Hostname", "X-Server-ID", "X-Email"} {
			if request.Header.Get(forbidden) != "" {
				t.Errorf("private header transmitted: %s", forbidden)
			}
		}
		response.Header().Set("Content-Type", MediaType+"; charset=utf-8")
		response.Header().Set("ETag", `"fixture"`)
		response.Write(payload)
	}))
	defer server.Close()
	client := server.Client()
	client.Jar, _ = cookiejar.New(nil)
	endpoint, _ := url.Parse(server.URL)
	client.Jar.SetCookies(endpoint, []*http.Cookie{{Name: "private", Value: "must-not-send"}})
	source, err := NewStaticHTTPSource(server.URL+"/releases/index.json", "community-static", client)
	if err != nil {
		t.Fatal(err)
	}
	result, err := source.Fetch(context.Background(), FetchRequest{Channel: "stable"})
	if err != nil || result.NotModified || string(result.Manifest) != string(payload) || !result.Evidence.TransportAuthenticated || result.Evidence.SourceID != "community-static" || result.Evidence.Validators.ETag != `"fixture"` {
		t.Fatalf("result=%+v err=%v", result, err)
	}
}

func TestStaticSourceNotModifiedIsBodyFree(t *testing.T) {
	server := httptest.NewTLSServer(http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		if request.Header.Get("If-None-Match") != `"old"` {
			t.Errorf("validator missing")
		}
		response.Header().Set("ETag", `"old"`)
		response.WriteHeader(http.StatusNotModified)
	}))
	defer server.Close()
	source, _ := NewStaticHTTPSource(server.URL+"/releases/index.json", "community-static", server.Client())
	result, err := source.Fetch(context.Background(), FetchRequest{Channel: "stable", Validators: Validators{ETag: `"old"`}})
	if err != nil || !result.NotModified || len(result.Manifest) != 0 {
		t.Fatalf("result=%+v err=%v", result, err)
	}
}

func TestStaticSourceFailureCategories(t *testing.T) {
	tests := []struct {
		name        string
		status      int
		contentType string
		body        string
		failure     Failure
	}{
		{"HTTP status", http.StatusServiceUnavailable, MediaType, "unavailable", SourceHTTPStatus},
		{"media type", http.StatusOK, "text/html", "<html>", SourceMediaType},
		{"too large", http.StatusOK, MediaType, strings.Repeat("x", MaxIndexBytes+1), SourceTooLarge},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			server := httptest.NewTLSServer(http.HandlerFunc(func(response http.ResponseWriter, _ *http.Request) {
				response.Header().Set("Content-Type", tc.contentType)
				response.WriteHeader(tc.status)
				response.Write([]byte(tc.body))
			}))
			defer server.Close()
			source, _ := NewStaticHTTPSource(server.URL+"/releases/index.json", "community-static", server.Client())
			_, err := source.Fetch(context.Background(), FetchRequest{Channel: "stable"})
			if FailureOf(err) != tc.failure {
				t.Fatalf("failure=%s err=%v", FailureOf(err), err)
			}
		})
	}
}

func TestStaticSourceAuthorityRedirectTimeoutAndCancellation(t *testing.T) {
	if _, err := NewStaticHTTPSource("http://example.invalid/index.json", "community-static", nil); FailureOf(err) != SourceAuthority {
		t.Fatalf("HTTP endpoint failure=%s", FailureOf(err))
	}
	if _, err := NewStaticHTTPSource("https://example.invalid/releases/../index.json", "community-static", nil); FailureOf(err) != SourceAuthority {
		t.Fatalf("dot-segment endpoint failure=%s", FailureOf(err))
	}
	insecureClient := &http.Client{Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}}
	if _, err := NewStaticHTTPSource("https://example.invalid/releases/index.json", "community-static", insecureClient); FailureOf(err) != SourceAuthority {
		t.Fatalf("insecure TLS client failure=%s", FailureOf(err))
	}
	redirectTarget := httptest.NewTLSServer(http.HandlerFunc(func(http.ResponseWriter, *http.Request) {}))
	defer redirectTarget.Close()
	redirectSource := httptest.NewTLSServer(http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		http.Redirect(response, request, redirectTarget.URL+"/foreign/index.json", http.StatusFound)
	}))
	defer redirectSource.Close()
	source, _ := NewStaticHTTPSource(redirectSource.URL+"/releases/index.json", "community-static", redirectSource.Client())
	if _, err := source.Fetch(context.Background(), FetchRequest{Channel: "stable"}); FailureOf(err) != SourceAuthority {
		t.Fatalf("redirect failure=%s err=%v", FailureOf(err), err)
	}

	slow := httptest.NewTLSServer(http.HandlerFunc(func(response http.ResponseWriter, _ *http.Request) {
		time.Sleep(100 * time.Millisecond)
		response.Header().Set("Content-Type", MediaType)
		response.Write([]byte(`{}`))
	}))
	defer slow.Close()
	client := slow.Client()
	client.Timeout = 10 * time.Millisecond
	slowSource, _ := NewStaticHTTPSource(slow.URL+"/releases/index.json", "community-static", client)
	if _, err := slowSource.Fetch(context.Background(), FetchRequest{Channel: "stable"}); FailureOf(err) != SourceTimeout {
		t.Fatalf("timeout failure=%s err=%v", FailureOf(err), err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	if _, err := slowSource.Fetch(ctx, FetchRequest{Channel: "stable"}); FailureOf(err) != SourceCanceled {
		t.Fatalf("cancellation failure=%s err=%v", FailureOf(err), err)
	}
}

func TestStaticSourceRejectsUnsafeValidators(t *testing.T) {
	server := httptest.NewTLSServer(http.HandlerFunc(func(http.ResponseWriter, *http.Request) {}))
	defer server.Close()
	source, _ := NewStaticHTTPSource(server.URL+"/index.json", "community-static", server.Client())
	if _, err := source.Fetch(context.Background(), FetchRequest{Channel: "stable", Validators: Validators{ETag: "ok\r\nAuthorization: secret"}}); FailureOf(err) != SourceAuthority {
		t.Fatalf("unsafe validator failure=%s", FailureOf(err))
	}
}

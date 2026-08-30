package releasediscovery

import (
	"context"
	"crypto/tls"
	"errors"
	"io"
	"mime"
	"net"
	"net/http"
	"net/url"
	"path"
	"strings"
	"time"
)

var errSourceAuthority = errors.New("release source authority refused")

const (
	defaultFetchTimeout = 35 * time.Second
	maxValidatorBytes   = 256
)

type Validators struct {
	ETag         string `json:"etag,omitempty"`
	LastModified string `json:"last_modified,omitempty"`
}

type FetchRequest struct {
	Channel    string
	Validators Validators
}

type SourceEvidence struct {
	SourceID               string     `json:"source_id"`
	TransportAuthenticated bool       `json:"transport_authenticated"`
	Validators             Validators `json:"validators,omitempty"`
}

type FetchResult struct {
	NotModified bool
	Manifest    []byte
	Evidence    SourceEvidence
}

type ReleaseSource interface {
	Fetch(context.Context, FetchRequest) (FetchResult, error)
}

type StaticHTTPSource struct {
	endpoint *url.URL
	sourceID string
	client   *http.Client
}

func NewStaticHTTPSource(endpoint, sourceID string, client *http.Client) (*StaticHTTPSource, error) {
	parsed, err := url.Parse(endpoint)
	if err != nil || parsed.Scheme != "https" || parsed.Host == "" || parsed.User != nil || parsed.RawQuery != "" || parsed.Fragment != "" || parsed.Opaque != "" || parsed.EscapedPath() != parsed.Path || path.Clean(parsed.Path) != parsed.Path || parsed.Path == "/" || !safeToken(sourceID, 64) {
		return nil, fail(SourceAuthority)
	}
	if client == nil {
		client = boundedHTTPClient()
	}
	transport, err := secureTransport(client.Transport, parsed.Hostname())
	if err != nil {
		return nil, fail(SourceAuthority)
	}
	clone := *client
	clone.Transport = transport
	clone.Jar = nil
	if clone.Timeout <= 0 || clone.Timeout > defaultFetchTimeout {
		clone.Timeout = defaultFetchTimeout
	}
	clone.CheckRedirect = func(request *http.Request, via []*http.Request) error {
		if len(via) > 3 || request.URL.Scheme != "https" || request.URL.Host != parsed.Host || request.URL.EscapedPath() != parsed.EscapedPath() {
			return errSourceAuthority
		}
		return nil
	}
	return &StaticHTTPSource{endpoint: parsed, sourceID: sourceID, client: &clone}, nil
}

func secureTransport(roundTripper http.RoundTripper, hostname string) (*http.Transport, error) {
	if roundTripper == nil {
		roundTripper = http.DefaultTransport
	}
	base, ok := roundTripper.(*http.Transport)
	if !ok || base.DialTLS != nil || base.DialTLSContext != nil {
		return nil, errSourceAuthority
	}
	transport := base.Clone()
	transport.Proxy = nil
	transport.ProxyConnectHeader = nil
	transport.ForceAttemptHTTP2 = true
	transport.MaxIdleConns = 4
	transport.MaxIdleConnsPerHost = 2
	transport.IdleConnTimeout = boundedDuration(transport.IdleConnTimeout, 15*time.Second)
	transport.TLSHandshakeTimeout = boundedDuration(transport.TLSHandshakeTimeout, 5*time.Second)
	transport.ResponseHeaderTimeout = boundedDuration(transport.ResponseHeaderTimeout, 5*time.Second)
	transport.ExpectContinueTimeout = boundedDuration(transport.ExpectContinueTimeout, time.Second)
	configuration := transport.TLSClientConfig
	if configuration == nil {
		configuration = &tls.Config{}
	} else {
		configuration = configuration.Clone()
	}
	if configuration.InsecureSkipVerify || len(configuration.Certificates) != 0 || configuration.GetClientCertificate != nil || (configuration.ServerName != "" && configuration.ServerName != hostname) {
		return nil, errSourceAuthority
	}
	if configuration.MinVersion < tls.VersionTLS12 {
		configuration.MinVersion = tls.VersionTLS12
	}
	transport.TLSClientConfig = configuration
	return transport, nil
}

func boundedDuration(value, maximum time.Duration) time.Duration {
	if value <= 0 || value > maximum {
		return maximum
	}
	return value
}

func boundedHTTPClient() *http.Client {
	dialer := &net.Dialer{Timeout: 5 * time.Second, KeepAlive: 15 * time.Second}
	return &http.Client{
		Timeout: defaultFetchTimeout,
		Transport: &http.Transport{
			Proxy:                 nil,
			DialContext:           dialer.DialContext,
			ForceAttemptHTTP2:     true,
			MaxIdleConns:          4,
			MaxIdleConnsPerHost:   2,
			IdleConnTimeout:       15 * time.Second,
			TLSHandshakeTimeout:   5 * time.Second,
			ResponseHeaderTimeout: 5 * time.Second,
			ExpectContinueTimeout: time.Second,
			TLSClientConfig:       &tls.Config{MinVersion: tls.VersionTLS12},
		},
	}
}

func (s *StaticHTTPSource) Fetch(ctx context.Context, request FetchRequest) (FetchResult, error) {
	if !validChannel(request.Channel) || !safeValidator(request.Validators.ETag) || !safeValidator(request.Validators.LastModified) {
		return FetchResult{}, fail(SourceAuthority)
	}
	ctx, cancel := context.WithTimeout(ctx, defaultFetchTimeout)
	defer cancel()
	httpRequest, err := http.NewRequestWithContext(ctx, http.MethodGet, s.endpoint.String(), nil)
	if err != nil {
		return FetchResult{}, fail(SourceAuthority)
	}
	httpRequest.Header.Set("Accept", MediaType)
	httpRequest.Header.Set("User-Agent", "QWSG-release-index/1")
	if request.Validators.ETag != "" {
		httpRequest.Header.Set("If-None-Match", request.Validators.ETag)
	}
	if request.Validators.LastModified != "" {
		httpRequest.Header.Set("If-Modified-Since", request.Validators.LastModified)
	}
	response, err := s.client.Do(httpRequest)
	if err != nil {
		if errors.Is(err, errSourceAuthority) {
			return FetchResult{}, fail(SourceAuthority)
		}
		if errors.Is(ctx.Err(), context.Canceled) {
			return FetchResult{}, fail(SourceCanceled)
		}
		if errors.Is(ctx.Err(), context.DeadlineExceeded) || timeoutError(err) {
			return FetchResult{}, fail(SourceTimeout)
		}
		return FetchResult{}, fail(SourceTransport)
	}
	defer response.Body.Close()
	evidence := SourceEvidence{SourceID: s.sourceID, TransportAuthenticated: true, Validators: responseValidators(response.Header)}
	if response.StatusCode == http.StatusNotModified {
		return FetchResult{NotModified: true, Evidence: evidence}, nil
	}
	if response.StatusCode != http.StatusOK {
		return FetchResult{}, fail(SourceHTTPStatus)
	}
	mediaType, parameters, mediaErr := mime.ParseMediaType(response.Header.Get("Content-Type"))
	if mediaErr != nil || mediaType != MediaType || !validMediaParameters(parameters) {
		return FetchResult{}, fail(SourceMediaType)
	}
	if response.ContentLength > MaxIndexBytes {
		return FetchResult{}, fail(SourceTooLarge)
	}
	payload, readErr := io.ReadAll(io.LimitReader(response.Body, MaxIndexBytes+1))
	if readErr != nil {
		return FetchResult{}, fail(SourceTransport)
	}
	if len(payload) > MaxIndexBytes {
		return FetchResult{}, fail(SourceTooLarge)
	}
	return FetchResult{Manifest: payload, Evidence: evidence}, nil
}

func safeValidator(value string) bool {
	if len(value) > maxValidatorBytes {
		return false
	}
	for _, r := range value {
		if r < 0x20 || r > 0x7e {
			return false
		}
	}
	return true
}

func responseValidators(header http.Header) Validators {
	result := Validators{}
	if value := header.Get("ETag"); safeValidator(value) {
		result.ETag = value
	}
	if value := header.Get("Last-Modified"); safeValidator(value) {
		result.LastModified = value
	}
	return result
}

func validMediaParameters(parameters map[string]string) bool {
	for name, value := range parameters {
		if strings.ToLower(name) != "charset" || strings.ToLower(value) != "utf-8" {
			return false
		}
	}
	return true
}

func timeoutError(err error) bool {
	value, ok := err.(net.Error)
	return ok && value.Timeout()
}

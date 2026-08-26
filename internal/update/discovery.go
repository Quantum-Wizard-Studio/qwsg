package update

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"time"
)

const (
	CanonicalAPI = "https://git.quantumwizard.hu/api/v1/repos/Quantum_Wizard_Studio/qwsg"
	MaxMetadata  = 1 << 20
)

type Asset struct {
	Name, URL string
	Size      int64
}
type Release struct {
	Version, Tag     string
	Archive, Sidecar Asset
}
type apiRelease struct {
	TagName    string `json:"tag_name"`
	Draft      bool   `json:"draft"`
	Prerelease bool   `json:"prerelease"`
	Assets     []struct {
		Name string `json:"name"`
		Size int64  `json:"size"`
		URL  string `json:"browser_download_url"`
	} `json:"assets"`
}

func Discover(ctx context.Context, client *http.Client, apiBase, installed string) (Release, Relation, error) {
	if apiBase == "" {
		apiBase = CanonicalAPI
	}
	endpoint := strings.TrimRight(apiBase, "/") + "/releases?limit=50"
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return Release{}, Invalid, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return Release{}, Invalid, fmt.Errorf("release discovery: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return Release{}, Invalid, fmt.Errorf("release discovery HTTP %d", resp.StatusCode)
	}
	payload, err := io.ReadAll(io.LimitReader(resp.Body, MaxMetadata+1))
	if err != nil || len(payload) > MaxMetadata {
		return Release{}, Invalid, fmt.Errorf("release metadata exceeds bound")
	}
	var records []apiRelease
	if err = json.Unmarshal(payload, &records); err != nil {
		return Release{}, Invalid, fmt.Errorf("release metadata: %w", err)
	}
	if len(records) > 50 {
		return Release{}, Invalid, fmt.Errorf("release metadata exceeds bound")
	}
	return selectRelease(records, installed)
}

func selectRelease(records []apiRelease, installed string) (Release, Relation, error) {
	type candidate struct {
		release Release
		parsed  Version
	}
	var candidates []candidate
	for _, record := range records {
		if record.Draft || !strings.HasPrefix(record.TagName, "v") {
			continue
		}
		version := strings.TrimPrefix(record.TagName, "v")
		parsed, parseErr := ParseVersion(version)
		if parseErr != nil || parsed.Major != 1 {
			continue
		}
		if record.Prerelease != (len(parsed.Prerelease) > 0) {
			continue
		}
		archiveName := "qwsg-" + version + "-linux-amd64.tar.gz"
		archive, sidecar, ok := exactAssets(record.Assets, archiveName)
		if !ok {
			continue
		}
		candidates = append(candidates, candidate{Release{Version: version, Tag: record.TagName, Archive: archive, Sidecar: sidecar}, parsed})
	}
	if len(candidates) == 0 {
		return Release{}, Invalid, fmt.Errorf("no valid canonical release")
	}
	sort.Slice(candidates, func(i, j int) bool { return Compare(candidates[i].parsed, candidates[j].parsed) > 0 })
	if len(candidates) > 1 && Compare(candidates[0].parsed, candidates[1].parsed) == 0 {
		return Release{}, Invalid, fmt.Errorf("ambiguous canonical release identity")
	}
	relation := Classify(installed, candidates[0].release.Version)
	return candidates[0].release, relation, nil
}

func exactAssets(items []struct {
	Name string `json:"name"`
	Size int64  `json:"size"`
	URL  string `json:"browser_download_url"`
}, archive string) (Asset, Asset, bool) {
	if len(items) != 2 {
		return Asset{}, Asset{}, false
	}
	want := map[string]*Asset{archive: {}, archive + ".sha256": {}}
	seen := map[string]bool{}
	for _, item := range items {
		target, ok := want[item.Name]
		if !ok {
			continue
		}
		if seen[item.Name] || item.Size <= 0 || item.Size > 128<<20 || !canonicalAssetURL(item.URL, item.Name) {
			return Asset{}, Asset{}, false
		}
		*target = Asset{Name: item.Name, URL: item.URL, Size: item.Size}
		seen[item.Name] = true
	}
	return *want[archive], *want[archive+".sha256"], seen[archive] && seen[archive+".sha256"]
}

func canonicalAssetURL(raw, name string) bool {
	u, err := url.Parse(raw)
	if err != nil {
		return false
	}
	return u.Scheme == "https" && u.Host == "git.quantumwizard.hu" && strings.HasPrefix(u.Path, "/Quantum_Wizard_Studio/qwsg/releases/download/v") && strings.HasSuffix(u.Path, "/"+name) && u.RawQuery == "" && u.Fragment == ""
}

func HTTPClient() *http.Client {
	return &http.Client{Timeout: 30 * time.Second, CheckRedirect: func(req *http.Request, via []*http.Request) error {
		if len(via) > 3 || req.URL.Scheme != "https" || req.URL.Host != "git.quantumwizard.hu" {
			return fmt.Errorf("unsafe release redirect")
		}
		return nil
	}}
}

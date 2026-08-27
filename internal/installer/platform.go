package installer

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

type Platform struct{ OS, Version, Architecture string }

func (p Platform) String() string { return fmt.Sprintf("%s %s %s", p.OS, p.Version, p.Architecture) }
func (p Platform) Supported() bool {
	return p.OS == "ubuntu" && p.Version == "24.04" && p.Architecture == "amd64"
}

func Detect(osRelease io.Reader, machine string) (Platform, error) {
	values := map[string]string{}
	s := bufio.NewScanner(io.LimitReader(osRelease, 64<<10))
	for s.Scan() {
		key, value, ok := strings.Cut(s.Text(), "=")
		if ok {
			values[key] = strings.Trim(strings.TrimSpace(value), `"`)
		}
	}
	if err := s.Err(); err != nil {
		return Platform{}, err
	}
	arch := machine
	if arch == "x86_64" {
		arch = "amd64"
	}
	if values["ID"] == "" || values["VERSION_ID"] == "" {
		return Platform{}, fmt.Errorf("platform identity unavailable")
	}
	return Platform{OS: strings.ToLower(values["ID"]), Version: values["VERSION_ID"], Architecture: arch}, nil
}

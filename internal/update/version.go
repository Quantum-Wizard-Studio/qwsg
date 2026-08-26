package update

import (
	"fmt"
	"strconv"
	"strings"
)

type Version struct {
	Major, Minor, Patch int
	Prerelease          []string
}

func ParseVersion(raw string) (Version, error) {
	if raw == "" || strings.TrimSpace(raw) != raw || strings.Contains(raw, "+") {
		return Version{}, fmt.Errorf("invalid release identity")
	}
	base, pre, hasPre := raw, "", false
	if i := strings.IndexByte(raw, '-'); i >= 0 {
		base, pre, hasPre = raw[:i], raw[i+1:], true
	}
	parts := strings.Split(base, ".")
	if len(parts) != 3 {
		return Version{}, fmt.Errorf("invalid release identity")
	}
	n := [3]int{}
	for i, part := range parts {
		if part == "" || (len(part) > 1 && part[0] == '0') {
			return Version{}, fmt.Errorf("invalid release identity")
		}
		value, err := strconv.Atoi(part)
		if err != nil || value < 0 {
			return Version{}, fmt.Errorf("invalid release identity")
		}
		n[i] = value
	}
	v := Version{Major: n[0], Minor: n[1], Patch: n[2]}
	if hasPre {
		if pre == "" {
			return Version{}, fmt.Errorf("invalid release identity")
		}
		v.Prerelease = strings.Split(pre, ".")
		for _, item := range v.Prerelease {
			if item == "" || !validIdentifier(item) || (numeric(item) && len(item) > 1 && item[0] == '0') {
				return Version{}, fmt.Errorf("invalid release identity")
			}
		}
	}
	return v, nil
}

func validIdentifier(s string) bool {
	for _, r := range s {
		if !(r >= '0' && r <= '9') && !(r >= 'A' && r <= 'Z') && !(r >= 'a' && r <= 'z') && r != '-' {
			return false
		}
	}
	return true
}

func numeric(s string) bool { _, err := strconv.Atoi(s); return err == nil }

func Compare(a, b Version) int {
	for _, pair := range [][2]int{{a.Major, b.Major}, {a.Minor, b.Minor}, {a.Patch, b.Patch}} {
		if pair[0] < pair[1] {
			return -1
		}
		if pair[0] > pair[1] {
			return 1
		}
	}
	if len(a.Prerelease) == 0 && len(b.Prerelease) == 0 {
		return 0
	}
	if len(a.Prerelease) == 0 {
		return 1
	}
	if len(b.Prerelease) == 0 {
		return -1
	}
	for i := 0; i < len(a.Prerelease) && i < len(b.Prerelease); i++ {
		x, y := a.Prerelease[i], b.Prerelease[i]
		if x == y {
			continue
		}
		xn, yn := numeric(x), numeric(y)
		if xn && yn {
			xi, _ := strconv.Atoi(x)
			yi, _ := strconv.Atoi(y)
			if xi < yi {
				return -1
			}
			return 1
		}
		if xn {
			return -1
		}
		if yn {
			return 1
		}
		if x < y {
			return -1
		}
		return 1
	}
	if len(a.Prerelease) < len(b.Prerelease) {
		return -1
	}
	if len(a.Prerelease) > len(b.Prerelease) {
		return 1
	}
	return 0
}

type Relation string

const (
	Newer       Relation = "newer"
	Equal       Relation = "equal"
	Older       Relation = "older"
	Unsupported Relation = "unsupported"
	Invalid     Relation = "invalid"
)

func Classify(installed, candidate string) Relation {
	a, err := ParseVersion(installed)
	if err != nil {
		return Invalid
	}
	b, err := ParseVersion(candidate)
	if err != nil {
		return Invalid
	}
	if a.Major != b.Major || b.Major != 1 {
		return Unsupported
	}
	switch Compare(b, a) {
	case 1:
		return Newer
	case 0:
		return Equal
	default:
		return Older
	}
}

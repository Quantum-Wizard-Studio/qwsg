package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestCLIHelpVersionAndInvalid(t *testing.T) {
	for _, tc := range []struct {
		args     []string
		code     int
		contains string
	}{{[]string{"help"}, 0, "Usage"}, {[]string{"version"}, 0, "0.0.1"}, {[]string{"unknown"}, 1, "unknown command"}} {
		var out, err bytes.Buffer
		code := run(tc.args, &out, &err)
		if code != tc.code || !strings.Contains(out.String()+err.String(), tc.contains) {
			t.Fatalf("%v: %d %q", tc.args, code, out.String()+err.String())
		}
	}
}

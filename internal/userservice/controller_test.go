package userservice

import (
	"context"
	"quantumwizard.hu/qwsg/internal/runner"
	"testing"
)

type recording struct{ calls [][]string }

func (r *recording) Run(_ context.Context, id string, args ...string) (runner.Result, error) {
	r.calls = append(r.calls, append([]string{id}, args...))
	return runner.Result{}, nil
}
func TestControllerOwnsExactOperations(t *testing.T) {
	r := &recording{}
	c := Controller{runner: r}
	if err := c.Activate(context.Background()); err != nil {
		t.Fatal(err)
	}
	if len(r.calls) != 2 || r.calls[0][1] != "--user" || r.calls[0][2] != "daemon-reload" || r.calls[1][2] != "enable" || r.calls[1][4] != Unit {
		t.Fatalf("calls=%q", r.calls)
	}
}

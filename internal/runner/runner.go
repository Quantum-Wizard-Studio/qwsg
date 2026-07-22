package runner

import (
	"bytes"
	"context"
	"errors"
	"os/exec"
	"syscall"
	"time"
)

var ErrOutputLimit = errors.New("command output limit exceeded")

type Result struct {
	Stdout, Stderr []byte
	ExitCode       int
}
type Runner interface {
	Run(context.Context, string, ...string) (Result, error)
}
type Bounded struct {
	Allowed   map[string]string
	Timeout   time.Duration
	MaxOutput int
}

func (b Bounded) Run(parent context.Context, id string, args ...string) (Result, error) {
	path, ok := b.Allowed[id]
	if !ok {
		return Result{}, errors.New("command is not allowlisted")
	}
	ctx, cancel := context.WithTimeout(parent, b.Timeout)
	defer cancel()
	cmd := exec.CommandContext(ctx, path, args...)
	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
	cmd.Cancel = func() error {
		if cmd.Process == nil {
			return nil
		}
		return syscall.Kill(-cmd.Process.Pid, syscall.SIGKILL)
	}
	cmd.Env = []string{"PATH=/usr/bin:/bin", "LANG=C", "LC_ALL=C"}
	var out, errout bytes.Buffer
	ow := &limitWriter{w: &out, n: b.MaxOutput}
	ew := &limitWriter{w: &errout, n: b.MaxOutput}
	cmd.Stdout = ow
	cmd.Stderr = ew
	err := cmd.Run()
	if ow.exceeded || ew.exceeded {
		return Result{}, ErrOutputLimit
	}
	r := Result{Stdout: out.Bytes(), Stderr: errout.Bytes()}
	if ee, ok := err.(*exec.ExitError); ok {
		r.ExitCode = ee.ExitCode()
	}
	if ctx.Err() != nil {
		return r, ctx.Err()
	}
	return r, err
}

type limitWriter struct {
	w        *bytes.Buffer
	n        int
	exceeded bool
}

func (w *limitWriter) Write(p []byte) (int, error) {
	if len(p) > w.n {
		if w.n > 0 {
			_, _ = w.w.Write(p[:w.n])
		}
		w.n = 0
		w.exceeded = true
		return len(p), nil
	}
	w.n -= len(p)
	return w.w.Write(p)
}

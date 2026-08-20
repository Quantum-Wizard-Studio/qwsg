// Package userservice owns the narrow, explicitly authorized QWSG user-unit
// activation boundary. It cannot select another executable, unit, or argv.
package userservice

import (
	"context"
	"fmt"
	"quantumwizard.hu/qwsg/internal/runner"
	"time"
)

const Unit = "qwsg-guardian.service"

type Controller struct{ runner runner.Runner }

func New() Controller {
	return Controller{runner: runner.Bounded{Allowed: map[string]string{"systemctl": "/usr/bin/systemctl"}, Timeout: 10 * time.Second, MaxOutput: 64 << 10}}
}

func (c Controller) Activate(ctx context.Context) error {
	if _, err := c.runner.Run(ctx, "systemctl", "--user", "daemon-reload"); err != nil {
		return fmt.Errorf("guardian daemon reload failed")
	}
	if _, err := c.runner.Run(ctx, "systemctl", "--user", "enable", "--now", Unit); err != nil {
		return fmt.Errorf("guardian activation failed")
	}
	return nil
}

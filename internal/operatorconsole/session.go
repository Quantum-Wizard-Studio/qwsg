package operatorconsole

import (
	"bufio"
	"context"
	"fmt"
	"io"
)

func Run(ctx context.Context, input io.Reader, output io.Writer, provider OverviewProvider, initial State) error {
	state := initial
	rendered := false
	reader := bufio.NewScanner(io.LimitReader(input, 1<<20))
	reader.Buffer(make([]byte, 128), MaxInput+1)
	for {
		if rendered && state.Capabilities.Interactive {
			if _, err := io.WriteString(output, "\x1b[H\x1b[2J"); err != nil {
				return fmt.Errorf("console output: %w", err)
			}
		}
		if _, err := io.WriteString(output, Render(state)); err != nil {
			return fmt.Errorf("console output: %w", err)
		}
		rendered = true
		if state.Quit {
			return nil
		}
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}
		if !reader.Scan() {
			if err := reader.Err(); err != nil {
				return fmt.Errorf("console input: %w", err)
			}
			return nil
		}
		action := ParseAction(reader.Text())
		if action == Refresh {
			if provider == nil {
				state = ApplyRefresh(state, state.Overview, fmt.Errorf("provider unavailable"))
			} else {
				overview, err := provider.Refresh(ctx)
				state = ApplyRefresh(state, overview, err)
			}
			continue
		}
		state = Transition(state, action)
	}
}

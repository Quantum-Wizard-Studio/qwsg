package runtimeservice

import (
	"context"
	"os"
	"os/signal"
	"syscall"
)

// SignalContextFactory isolates operating-system signal registration so tests
// can verify the adapter without delivering a real process signal.
type SignalContextFactory func(context.Context, ...os.Signal) (context.Context, context.CancelFunc)

func OSSignalContext(parent context.Context, signals ...os.Signal) (context.Context, context.CancelFunc) {
	return signal.NotifyContext(parent, signals...)
}

// RunWithSignals maps only SIGINT and SIGTERM to Service cancellation. It does
// not install, supervise, restart, or detach a process.
func RunWithSignals(ctx context.Context, service Service, input Input, factory SignalContextFactory) (Result, error) {
	if factory == nil {
		return Result{}, os.ErrInvalid
	}
	bounded, stop := factory(ctx, os.Interrupt, syscall.SIGTERM)
	defer stop()
	return service.Run(bounded, input)
}

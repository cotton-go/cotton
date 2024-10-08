package contton

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/cotton-go/cotton/internal/cotton"
	"github.com/cotton-go/cotton/internal/reflection"
	"github.com/cotton-go/cotton/runtime/codegen"
)

type Implements struct {
	logger *slog.Logger
	implementsImpl
}

type implementsImpl struct{}

func Run[T any](ctx context.Context, app func(context.Context, *T) error) error {
	var cancel context.CancelFunc
	ctx, cancel = signal.NotifyContext(ctx, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	// Read config from SERVICE_CONFIG env variable, if non-empty.
	opts := cotton.Options{}
	if filename := os.Getenv("SERVICE_CONFIG"); filename != "" {
		contents, err := os.ReadFile(filename)
		if err != nil {
			return fmt.Errorf("config file: %w", err)
		}
		opts.ConfigFilename = filename
		opts.Config = string(contents)
	}

	regs := codegen.Registered()
	wlet, err := cotton.NewSingleWeavelet(ctx, regs, opts)
	if err != nil {
		return err
	}

	main, err := wlet.GetImpl(reflection.Type[T]())
	if err != nil {
		return err
	}

	err = app(ctx, main.(*T))
	cancel()
	return err
}

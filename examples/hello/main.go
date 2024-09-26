package main

import (
	"context"
	"log/slog"

	"github.com/cotton-go/cotton"
)

type app struct {
	cotton.WithConfig[Options] `config:"app"`
}

//go:generate cottongen -type=app -name=app -output=app_gen.go
func main() {
	cotton.Run(context.Background(), func(ctx context.Context, app *app) error {
		slog.Info("hello world", "cfg", app.Config())
		return nil
	})
}

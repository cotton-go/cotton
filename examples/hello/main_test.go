package main

import (
	"context"
	"testing"

	"github.com/cotton-go/cotton"
)

func TestMain(t *testing.T) {
	cotton.Run(context.Background(), func(ctx context.Context, app *app) error {
		t.Log("hello world", "cfg", app.Config())
		return nil
	})
}

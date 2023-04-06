package slog

import (
	"context"
	"testing"
)

func TestForkSpan(t *testing.T) {
	ctx := context.Background()

	logger := In(ctx)
	if logger == nil {
		t.Error("Logger is nil")
		return
	}
}

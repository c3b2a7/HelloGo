package concurrent

import (
	"context"
	"testing"
	"time"
)

func TestCancelMultiGoroutineWithCancelContext(t *testing.T) {
	CancelMultiGoroutineWithCancelContext()
}

func TestCancelMultiGoroutineWithNamedCancelContext(t *testing.T) {
	CancelMultiGoroutineWithNamedCancelContext()
}

func TestCancelWithCancelContext(t *testing.T) {
	CancelWithCancelContext()
}

func TestCancelWithChannel(t *testing.T) {
	CancelWithChannel()
}

func Test_watch(t *testing.T) {
	timeout, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	watch(timeout, func(ctx context.Context) string {
		return "watcher"
	})
}

package debug

import (
	"runtime/debug"
	"testing"
)

func TestReadBuildInfo(t *testing.T) {
	_, ok := debug.ReadBuildInfo()
	if !ok {
		t.Errorf("Failed to read build info")
	}
}

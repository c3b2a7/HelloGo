package iterator

import (
	"slices"
	"testing"
)

func TestSlice(t *testing.T) {
	data := []byte("hello world")
	backward := slices.Collect(slices.Values(data))

	if !slices.Equal(backward, []byte("hello world")) {
		t.Error("unexpected")
	}
}

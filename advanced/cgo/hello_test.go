package cgo

import (
	"testing"
)

func TestHello(t *testing.T) {
	Hello("hello cgo!")
}

package basic

import (
	"errors"
	"testing"
)

func Bar() error {
	return errors.New("error from Bar")
}

func Foo() (err error) {
	if err := Bar(); err != nil { // err被屏蔽
		return
	}
	return
}

func TestError(t *testing.T) {
	if err := Foo(); err != nil {
		return
	}
	t.Error("expected an error")
}

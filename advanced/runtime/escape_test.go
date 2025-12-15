package runtime

import (
	"fmt"
	"testing"
	"unsafe"
)

type A struct {
	S *string
}

func (f *A) String() string {
	return *f.S
}

type ATrick struct {
	S unsafe.Pointer
}

func (f *ATrick) String() string {
	return *(*string)(f.S)
}

func NewA(s string) A {
	return A{S: &s}
}

func NewATrick(s string) ATrick {
	return ATrick{S: noescape(unsafe.Pointer(&s))}
}

func noescape(p unsafe.Pointer) unsafe.Pointer {
	x := uintptr(p)
	return unsafe.Pointer(x ^ 0)
}

// go build -gcflags=-m
func TestNoEscape(t *testing.T) {
	s := "hello"
	f1 := NewA(s)
	f2 := NewATrick(s)
	s1 := f1.String()
	s2 := f2.String()
	s3 := s1 + s2
	fmt.Println(s1, s2, s3)
}

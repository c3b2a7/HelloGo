package reflect

import (
	"fmt"
	"reflect"
	"testing"
	"unsafe"
)

type stringer string

func (s stringer) String() string {
	return string(s)
}

func TestSprint(t *testing.T) {
	fmt.Println(Sprint(stringer("aa")))
}

func TestSprintAny(t *testing.T) {
	fmt.Println(SprintAny(stringer("aa")))

	var a = 1.2
	fmt.Printf("%#016x\n", *(*uint64)(reflect.ValueOf(&a).UnsafePointer())) // "0x3ff0000000000000"
	fmt.Println(Float64bits(a))
	fmt.Println(reflect.ValueOf(&a).UnsafePointer() == unsafe.Pointer(&a))
}

func Float64bits(f float64) uint64 { return *(*uint64)(unsafe.Pointer(&f)) }

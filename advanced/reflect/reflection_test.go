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
	if ret := Sprint(stringer("aa")); ret != "aa" {
		t.Errorf("unexpected result, expected: %s, actual: %s", "aa", ret)
	}
}

func TestSprintAny(t *testing.T) {
	if ret := SprintAny(stringer("aa")); ret != `"aa"` {
		t.Errorf("unexpected result, expected: %s, actual: %s", `"aa"`, ret)
	}

	var a = 1.2
	fmt.Printf("%#016x\n", *(*uint64)(reflect.ValueOf(&a).UnsafePointer())) // "0x3ff0000000000000"
	fmt.Println(Float64bits(a))
	fmt.Println(reflect.ValueOf(&a).UnsafePointer() == unsafe.Pointer(&a))
}

func Float64bits(f float64) uint64 { return *(*uint64)(unsafe.Pointer(&f)) }

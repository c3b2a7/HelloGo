package basic

import (
	"reflect"
	"slices"
	"testing"
	"unsafe"
)

func String(b []byte) (s string) {
	if len(b) == 0 {
		return ""
	}
	pbytes := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	pstring := (*reflect.StringHeader)(unsafe.Pointer(&s))
	pstring.Data = pbytes.Data
	pstring.Len = pbytes.Len
	return
}

func Slice(s string) (b []byte) {
	pbytes := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	pstring := (*reflect.StringHeader)(unsafe.Pointer(&s))
	pbytes.Data = pstring.Data
	pbytes.Len = pstring.Len
	pbytes.Cap = pstring.Len
	return
}

// BytesToString is equivalent to String
func BytesToString(b []byte) string {
	if len(b) == 0 {
		return ""
	}
	// 从字节指针和长度构造字符串，避免拷贝
	return unsafe.String(&b[0], len(b))
}

// StringToBytes is equivalent to Slice
func StringToBytes(s string) []byte {
	if len(s) == 0 {
		return nil
	}
	// 从字符串数据指针和长度构造字节切片，避免拷贝
	return unsafe.Slice(unsafe.StringData(s), len(s))
}

func TestBytesToStringHack(t *testing.T) {
	helloBytes := []byte("hello")
	helloString := string(helloBytes) // copy
	if helloString != "hello" {
		t.Error("unexpected")
	}

	helloStringHacked := String(helloBytes) // share the same memory
	//helloStringHacked := BytesToString(helloBytes) // share the same memory
	if helloStringHacked != "hello" {
		t.Error("unexpected")
	}

	helloBytes[0] = 'H'

	if helloString != "hello" { // unchanged
		t.Error("unexpected")
	}
	if helloStringHacked != "Hello" {
		t.Error("unexpected")
	}
}

func foo(a []int) {
	a = append(a, 1, 2, 3, 4, 5, 6, 7, 8)
	a[0] = 200
}

func TestSliceExtended(t *testing.T) {
	a := make([]int, 0, 20)
	a = append(a, 1, 2)

	foo(a)
	if !slices.Equal(a, []int{200, 2}) {
		t.Error("unexpected")
	}

	assertRawAddrCC(t, a)
	if !slices.Equal(a, []int{200, 2}) {
		t.Error("unexpected")
	}

	assertHeaderCC(t, a)
	if !slices.Equal(a, []int{200, 2}) {
		t.Error("unexpected")
	}
}

func assertHeaderCC(t *testing.T, a []int) {
	sh := (*reflect.SliceHeader)(unsafe.Pointer(&a))
	sh.Len = 10
	if !slices.Equal(a, []int{200, 2, 1, 2, 3, 4, 5, 6, 7, 8}) {
		t.Error("unexpected")
	}
}

func assertRawAddrCC(t *testing.T, a []int) {
	aptr := unsafe.Pointer(&a)
	psize := unsafe.Sizeof(aptr)

	internalLengthPointer := (*int64)(unsafe.Pointer(uintptr(aptr) + psize))

	*internalLengthPointer = 10
	if !slices.Equal(a, []int{200, 2, 1, 2, 3, 4, 5, 6, 7, 8}) {
		t.Error("unexpected")
	}
}

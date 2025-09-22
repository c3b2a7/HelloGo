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

func TestString(t *testing.T) {
	helloBytes := []byte("hello")

	helloStringHacked := String(helloBytes) // share the same memory
	if helloStringHacked != "hello" {
		t.Error("unexpected")
	}

	helloBytes[0] = 'H'
	if helloStringHacked != "Hello" {
		t.Error("unexpected")
	}
}

func Slice(s string) (b []byte) {
	pbytes := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	pstring := (*reflect.StringHeader)(unsafe.Pointer(&s))
	pbytes.Data = pstring.Data
	pbytes.Len = pstring.Len
	pbytes.Cap = pstring.Len
	return
}

func TestSlice(t *testing.T) {
	helloString := "hello"
	helloStringBytes := Slice(helloString)

	// 获取 string 的底层指针
	stringHdr := (*reflect.StringHeader)(unsafe.Pointer(&helloString))
	// 获取 []byte 的底层指针
	sliceHdr := (*reflect.SliceHeader)(unsafe.Pointer(&helloStringBytes))

	if stringHdr.Data != sliceHdr.Data {
		t.Errorf("memory not shared: string data=%#x slice data=%#x",
			stringHdr.Data, sliceHdr.Data)
	}
}

// BytesToString is equivalent to String
func BytesToString(b []byte) string {
	if len(b) == 0 {
		return ""
	}
	// 从字节指针和长度构造字符串，避免拷贝
	return unsafe.String(&b[0], len(b))
}

func TestBytesToString(t *testing.T) {
	helloBytes := []byte("hello")

	helloStringHacked := BytesToString(helloBytes) // share the same memory of helloBytes
	if helloStringHacked != "hello" {
		t.Error("unexpected")
	}

	helloBytes[0] = 'H'               // modify hello bytes
	if helloStringHacked != "Hello" { // changed
		t.Error("unexpected")
	}
}

// StringToBytes is equivalent to Slice
func StringToBytes(s string) []byte {
	if len(s) == 0 {
		return nil
	}
	// 从字符串数据指针和长度构造字节切片，避免拷贝
	return unsafe.Slice(unsafe.StringData(s), len(s))
}

func TestStringToBytes(t *testing.T) {
	helloString := "hello"
	helloStringBytes := StringToBytes(helloString)

	// 获取 string 的底层指针
	stringHdr := (*reflect.StringHeader)(unsafe.Pointer(&helloString))
	// 获取 []byte 的底层指针
	sliceHdr := (*reflect.SliceHeader)(unsafe.Pointer(&helloStringBytes))

	if stringHdr.Data != sliceHdr.Data {
		t.Errorf("memory not shared: string data=%#x slice data=%#x",
			stringHdr.Data, sliceHdr.Data)
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

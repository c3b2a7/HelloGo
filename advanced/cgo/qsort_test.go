package cgo

import (
	"slices"
	"testing"
	"unsafe"
)

func TestQSort(t *testing.T) {
	values := []int{42, 9, 101, 95, 27, 25}
	QSort(unsafe.Pointer(&values[0]), len(values), int(unsafe.Sizeof(values[0])), func(a, b unsafe.Pointer) int {
		pa, pb := (*int)(a), (*int)(b)
		return *pa - *pb
	})
	if !slices.Equal(values, []int{9, 25, 27, 42, 95, 101}) {
		t.Errorf("unexpected result: %v", values)
	}
}

func TestQSort_Slice(t *testing.T) {
	values := []int{42, 9, 101, 95, 27, 25}
	QSort_Slice(values, func(i int, i2 int) bool {
		return values[i] > values[i2]
	})
	if !slices.Equal(values, []int{9, 25, 27, 42, 95, 101}) {
		t.Errorf("unexpected result: %v", values)
	}
}

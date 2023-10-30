package cgo

import (
	"fmt"
	"testing"
	"unsafe"
)

func TestQSort(t *testing.T) {
	values := []int{42, 9, 101, 95, 27, 25}
	QSort(unsafe.Pointer(&values[0]), len(values), int(unsafe.Sizeof(values[0])), func(a, b unsafe.Pointer) int {
		pa, pb := (*int)(a), (*int)(b)
		return *pa - *pb
	})
	fmt.Println(values)
	QSort_Slice(values, func(i int, i2 int) bool {
		return values[i] > values[i2]
	})
	fmt.Println(values)
}

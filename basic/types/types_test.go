package types

import (
	"fmt"
	"testing"
	"unsafe"
)

func TestRef(t *testing.T) {
	swap := func(a, b *int) {
		*a, *b = *b, *a
	}
	a, b := 1, 2
	println(a, b) // 1 2
	swap(&a, &b)
	println(a, b) // 2 1
}

func TestArray(t *testing.T) {
	array := []int{}
	add := func(a []int) {
		a = append(a, 0)
	}
	add(array)
	fmt.Printf("%v", array)
}

func TestSlice(t *testing.T) {
	s := []int{}
	func(ref *[]int) {
		*ref = append(*ref, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	}(&s)
	fmt.Printf("%v\n", s)
}

func TestSlice2(t *testing.T) {
	s := []int{}
	func(ref []int) {
		ref = append(ref, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	}(s)
	fmt.Printf("%v\n", s)
}

func TestSlice3(t *testing.T) {
	s := []int{1, 2, 3, 4, 5}

	ptr := &s

	temp := *ptr
	temp = temp[:len(temp)-1]
	*ptr = temp

	fmt.Printf("%v\n", s)
	fmt.Printf("%v\n", s[:len(s)-1])
}

func TestSliceReCap(t *testing.T) {
	// 切片的定义
	//type slice struct {
	//    array unsafe.Pointer // 指向底层数组
	//    len   int
	//    cap   int
	//}

	a := make([]int64, 1, 2)
	// 虽然追加元素后没有重新赋值，其实底层数组已经更新了。
	_ = append(a, 6)
	// 由于 slice 的 len 字段不是指针，才未能更新正确的长度导致报错。
	// fmt.Println(a[1]) // 越界！
	// 强制访问 a[1]
	baseAddr := unsafe.Pointer(&a[0])
	offset := unsafe.Sizeof(a[0])
	fmt.Println(*(*int64)(unsafe.Pointer(uintptr(baseAddr) + offset))) // 6
}

func TestCopySlice(t *testing.T) {
	s := []int{1, 2, 3, 4, 5}
	temp := append([]int(nil), s...)
	temp[0] = 0
	fmt.Printf("%v\n%v\n", s, temp)
}

func TestChipType_String(t *testing.T) {
	fmt.Printf("%s %d", CPU, CPU)
	fmt.Printf("%s %d", GPU, GPU)
}

package types

import (
	"fmt"
	"testing"
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

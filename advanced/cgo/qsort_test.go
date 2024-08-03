package cgo

import (
	"cmp"
	"slices"
	"sort"
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

func TestQSortSlice(t *testing.T) {
	values := []int{42, 9, 101, 95, 27, 25}

	//QSortSlice(values, cmp.Compare[int])
	QSortSlice(values, func(i int, j int) int {
		return cmp.Compare(i, j)
	})

	if !slices.Equal(values, []int{9, 25, 27, 42, 95, 101}) {
		t.Errorf("unexpected result: %v", values)
	}
}

func TestQSortAndStanderLibrarySort(t *testing.T) {
	qsortValue := []int{42, 9, 101, 95, 27, 25}
	slicesValue := slices.Clone(qsortValue)
	sortValue := slices.Clone(qsortValue)

	c := func(i int, j int) int {
		return cmp.Compare(i, j)
	}

	QSortSlice(qsortValue, c)
	slices.SortFunc(slicesValue, c)
	sort.Slice(sortValue, func(i, j int) bool {
		// [sort.Slice] using less function that compare elements with indexes i and j
		return c(sortValue[i], sortValue[j]) < 0
	})

	if !slices.Equal(qsortValue, slicesValue) {
		t.Fatalf("unexpected result! qsortValue: %v, slicesValue: %v", qsortValue, slicesValue)
	}
	if !slices.Equal(slicesValue, sortValue) {
		t.Fatalf("unexpected result! slicesValue: %v, sortValue: %v", slicesValue, sortValue)
	}
}

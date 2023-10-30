package cgo

/*
#include <stdlib.h>

typedef int (*qsort_cmp_func_t)(const void* a, const void* b);
extern int _cgo_qsort_compare(void* a, void* b);
*/
import "C"
import (
	"fmt"
	"reflect"
	"sync"
	"unsafe"
)

type CompareFunc func(a, b unsafe.Pointer) int

var qsort_compare_info struct {
	cmp CompareFunc
	sync.Mutex
}

//export _cgo_qsort_compare
func _cgo_qsort_compare(a, b unsafe.Pointer) C.int {
	return C.int(qsort_compare_info.cmp(a, b))
}

func QSort(base unsafe.Pointer, num, size int, cmp CompareFunc) {
	qsort_compare_info.Lock()
	defer qsort_compare_info.Unlock()
	qsort_compare_info.cmp = cmp
	C.qsort(base, C.size_t(num), C.size_t(size), C.qsort_cmp_func_t(C._cgo_qsort_compare))
}

func QSort_Slice(slice interface{}, less func(int, int) bool) error {
	sv := reflect.ValueOf(slice)
	if sv.Kind() != reflect.Slice {
		return fmt.Errorf("qsort_slice called with non-slice value of type %T", slice)
	}
	if sv.Len() == 0 {
		return nil
	}
	//base := unsafe.Pointer(sv.Index(0).Addr().Pointer())
	base := sv.Index(0).Addr().UnsafePointer()
	num, size := sv.Len(), int(sv.Type().Elem().Size())
	QSort(base, num, size, func(a, b unsafe.Pointer) int {
		i := int((uintptr(a) - uintptr(base)) / uintptr(size))
		j := int((uintptr(b) - uintptr(base)) / uintptr(size))
		switch {
		case less(i, j):
			return -1
		case less(j, i):
			return +1
		default:
			return 0
		}
	})
	return nil
}

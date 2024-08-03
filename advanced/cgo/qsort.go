package cgo

/*
#include <stdlib.h>

typedef int (*qsort_cmp_func_t)(const void* a, const void* b);
extern int _cgo_qsort_compare(void* a, void* b);
*/
import "C"
import (
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

func QSortSlice[S ~[]E, E any](slice S, cmp func(E, E) int) {
	sv := reflect.ValueOf(slice)
	if sv.Len() == 0 {
		return
	}
	//base := unsafe.Pointer(sv.Index(0).Addr().Pointer())
	base := sv.Index(0).Addr().UnsafePointer()
	num, size := sv.Len(), int(sv.Type().Elem().Size())
	QSort(base, num, size, func(a, b unsafe.Pointer) int {
		i := int((uintptr(a) - uintptr(base)) / uintptr(size))
		j := int((uintptr(b) - uintptr(base)) / uintptr(size))
		return cmp(slice[i], slice[j])
	})
}

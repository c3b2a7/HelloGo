package cgo

/*
#cgo CFLAGS: -I./c/include
#include "hello.h"
*/
import "C"

func Hello(s string) {
	C.say_hello(C.CString(s))
}

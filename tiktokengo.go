//go:build !windows

package tiktokengo

/*
#cgo linux LDFLAGS: ${SRCDIR}tiktokenrs/target/release/libtiktoken.a -ldl
#cgo darwin LDFLAGS: ${SRCDIR}tiktokenrs/target/release/libtiktoken.a -ldl -framework Security -framework CoreFoundation

#include <stdlib.h>
extern char *hello_to_my_name(const char*);
*/

import "C"
import "unsafe"

func HelloToMyName(name string) string {
	n := C.CString(name)
	result := C.hello_to_my_name(n)
	goString := C.GoString(result)
	C.free(unsafe.Pointer(n))
	C.free(unsafe.Pointer(result))
	return goString
}

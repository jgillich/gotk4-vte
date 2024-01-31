// Code generated by girgen. DO NOT EDIT.

package vte

import (
	"unsafe"
)

// #include <stdlib.h>
// #include <vte/vte.h>
import "C"

const TEST_FLAGS_ALL = 18446744073709551615
const TEST_FLAGS_NONE = 0

// GetUserShell gets the user's shell, or NULL. In the latter case, the system
// default (usually "/bin/sh") should be used.
//
// The function returns the following values:
//
//   - filename: newly allocated string with the user's shell, or NULL.
//
func GetUserShell() string {
	var _cret *C.char // in

	_cret = C.vte_get_user_shell()

	var _filename string // out

	_filename = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _filename
}

// Code generated by girgen. DO NOT EDIT.

package vte

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <vte/vte.h>
import "C"

//export _gotk4_vte2_SelectionFunc
func _gotk4_vte2_SelectionFunc(arg1 *C.VteTerminal, arg2 C.glong, arg3 C.glong, arg4 C.gpointer) (cret C.gboolean) {
	var fn SelectionFunc
	{
		v := gbox.Get(uintptr(arg4))
		if v == nil {
			panic(`callback not found`)
		}
		fn = v.(SelectionFunc)
	}

	var _terminal *Terminal // out
	var _column int32       // out
	var _row int32          // out

	_terminal = wrapTerminal(coreglib.Take(unsafe.Pointer(arg1)))
	_column = int32(arg2)
	_row = int32(arg3)

	ok := fn(_terminal, _column, _row)

	var _ bool

	if ok {
		cret = C.TRUE
	}

	return cret
}

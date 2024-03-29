// Code generated by girgen. DO NOT EDIT.

package vte

import (
	"fmt"
	"unsafe"

	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <vte/vte.h>
import "C"

// GType values.
var (
	GTypeAlign = coreglib.Type(C.vte_align_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeAlign, F: marshalAlign},
	})
}

// Align: enumeration type that can be used to specify how the terminal uses
// extra allocated space.
type Align C.gint

const (
	// AlignStart: align to left/top.
	AlignStart Align = 0
	// AlignCenter: align to centre.
	AlignCenter Align = 1
	// AlignEnd: align to right/bottom.
	AlignEnd Align = 3
)

func marshalAlign(p uintptr) (interface{}, error) {
	return Align(coreglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for Align.
func (a Align) String() string {
	switch a {
	case AlignStart:
		return "Start"
	case AlignCenter:
		return "Center"
	case AlignEnd:
		return "End"
	default:
		return fmt.Sprintf("Align(%d)", a)
	}
}

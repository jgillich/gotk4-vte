// Code generated by girgen. DO NOT EDIT.

package vte

import (
	"fmt"
	_ "runtime/cgo"
	"strings"
	"unsafe"

	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
)

// #cgo pkg-config: vte-2.91-gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdlib.h>
// #include <glib-object.h>
// #include <vte/vte.h>
import "C"

// GType values.
var (
	GTypePtyError = coreglib.Type(C.vte_pty_error_get_type())
	GTypePtyFlags = coreglib.Type(C.vte_pty_flags_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypePtyError, F: marshalPtyError},
		coreglib.TypeMarshaler{T: GTypePtyFlags, F: marshalPtyFlags},
	})
}

type PtyError C.gint

const (
	// PtyErrorPtyHelperFailed: obsolete. Deprecated: 0.42.
	PtyErrorPtyHelperFailed PtyError = iota
	// PtyErrorPty98Failed: failure when using PTY98 to allocate the PTY.
	PtyErrorPty98Failed
)

func marshalPtyError(p uintptr) (interface{}, error) {
	return PtyError(coreglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for PtyError.
func (p PtyError) String() string {
	switch p {
	case PtyErrorPtyHelperFailed:
		return "PtyHelperFailed"
	case PtyErrorPty98Failed:
		return "Pty98Failed"
	default:
		return fmt.Sprintf("PtyError(%d)", p)
	}
}

// The function returns the following values:
//
func RegexErrorQuark() glib.Quark {
	var _cret C.GQuark // in

	_cret = C.vte_regex_error_quark()

	var _quark glib.Quark // out

	_quark = uint32(_cret)
	type _ = glib.Quark
	type _ = uint32

	return _quark
}

type PtyFlags C.guint

const (
	// PtyNoLastlog: unused. Deprecated: 0.38.
	PtyNoLastlog PtyFlags = 0b1
	// PtyNoUtmp: unused. Deprecated: 0.38.
	PtyNoUtmp PtyFlags = 0b10
	// PtyNoWtmp: unused. Deprecated: 0.38.
	PtyNoWtmp PtyFlags = 0b100
	// PtyNoHelper: unused. Deprecated: 0.38.
	PtyNoHelper PtyFlags = 0b1000
	// PtyNoFallback: unused. Deprecated: 0.38.
	PtyNoFallback PtyFlags = 0b10000
	// PtyNoSession: do not start a new session for the child in
	// vte_pty_child_setup(). See man:setsid(2) for more information. Since:
	// 0.58.
	PtyNoSession PtyFlags = 0b100000
	// PtyNoCtty: do not set the PTY as the controlling TTY for the child in
	// vte_pty_child_setup(). See man:tty_ioctl(4) for more information. Since:
	// 0.58.
	PtyNoCtty PtyFlags = 0b1000000
	// PtyDefault: default flags.
	PtyDefault PtyFlags = 0b0
)

func marshalPtyFlags(p uintptr) (interface{}, error) {
	return PtyFlags(coreglib.ValueFromNative(unsafe.Pointer(p)).Flags()), nil
}

// String returns the names in string for PtyFlags.
func (p PtyFlags) String() string {
	if p == 0 {
		return "PtyFlags(0)"
	}

	var builder strings.Builder
	builder.Grow(92)

	for p != 0 {
		next := p & (p - 1)
		bit := p - next

		switch bit {
		case PtyNoLastlog:
			builder.WriteString("NoLastlog|")
		case PtyNoUtmp:
			builder.WriteString("NoUtmp|")
		case PtyNoWtmp:
			builder.WriteString("NoWtmp|")
		case PtyNoHelper:
			builder.WriteString("NoHelper|")
		case PtyNoFallback:
			builder.WriteString("NoFallback|")
		case PtyNoSession:
			builder.WriteString("NoSession|")
		case PtyNoCtty:
			builder.WriteString("NoCtty|")
		case PtyDefault:
			builder.WriteString("Default|")
		default:
			builder.WriteString(fmt.Sprintf("PtyFlags(0b%b)|", bit))
		}

		p = next
	}

	return strings.TrimSuffix(builder.String(), "|")
}

// Has returns true if p contains other.
func (p PtyFlags) Has(other PtyFlags) bool {
	return (p & other) == other
}

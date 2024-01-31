// Code generated by girgen. DO NOT EDIT.

package vte

// #include <stdlib.h>
// #include <vte/vte.h>
import "C"

// GetMajorVersion returns the major version of the VTE library at runtime.
// Contrast this with VTE_MAJOR_VERSION which represents the version of the VTE
// library that the code was compiled with.
//
// The function returns the following values:
//
//   - guint: major version.
//
func GetMajorVersion() uint {
	var _cret C.guint // in

	_cret = C.vte_get_major_version()

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// GetMicroVersion returns the micro version of the VTE library at runtime.
// Contrast this with VTE_MICRO_VERSION which represents the version of the VTE
// library that the code was compiled with.
//
// The function returns the following values:
//
//   - guint: micro version.
//
func GetMicroVersion() uint {
	var _cret C.guint // in

	_cret = C.vte_get_micro_version()

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// GetMinorVersion returns the minor version of the VTE library at runtime.
// Contrast this with VTE_MINOR_VERSION which represents the version of the VTE
// library that the code was compiled with.
//
// The function returns the following values:
//
//   - guint: minor version.
//
func GetMinorVersion() uint {
	var _cret C.guint // in

	_cret = C.vte_get_minor_version()

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// Code generated by girgen. DO NOT EDIT.

package vte

// #include <stdlib.h>
// #include <vte/vte.h>
import "C"

// CharAttributes: deprecated: since version 0.68.
//
// An instance of this type is always passed by reference.
type CharAttributes struct {
	*charAttributes
}

// charAttributes is the struct that's finalized.
type charAttributes struct {
	native *C.VteCharAttributes
}

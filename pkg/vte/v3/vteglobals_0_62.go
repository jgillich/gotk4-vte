// Code generated by girgen. DO NOT EDIT.

package vte

// #include <stdlib.h>
// #include <vte/vte.h>
import "C"

// GetFeatureFlags gets features VTE was compiled with.
//
// The function returns the following values:
//
//   - featureFlags flags from FeatureFlags.
//
func GetFeatureFlags() FeatureFlags {
	var _cret C.VteFeatureFlags // in

	_cret = C.vte_get_feature_flags()

	var _featureFlags FeatureFlags // out

	_featureFlags = FeatureFlags(_cret)

	return _featureFlags
}

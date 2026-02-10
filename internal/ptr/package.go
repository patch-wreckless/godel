// Portions of this file are derived from
// http://github.com/patch-wreckless/go-ptr
//
// Copyright (c) 2026, Patch Wreckless <https://github.com/patch-wreckless>
// Licensed under the BSD 3-Clause License.
// See LICENSE.go-ptr file for details.

package ptr

// To returns a pointer to the given value.
func To[T any](t T) *T {
	return &t
}

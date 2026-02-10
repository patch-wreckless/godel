// Portions of this file are derived from
// https://github.com/patch-wreckless/go-ptr
//
// Copyright (c) 2026, Patch Wreckless <https://github.com/patch-wreckless>
// Licensed under the BSD 3-Clause License.
// See LICENSE.go-ptr file for details.

package ptr

import (
	"testing"
)

func TestTo(t *testing.T) {

	t.Run("returns pointer to given value", func(t *testing.T) {
		expected := -37
		p := To(expected)
		actual := *p
		if actual != expected {
			t.Errorf("expected %d; got %d", expected, actual)
		}
	})
}

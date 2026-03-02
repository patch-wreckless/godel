package fields

import (
	"fmt"
	"testing"

	"github.com/patch-wreckless/godel/internal/assert"
	"github.com/patch-wreckless/godel/internal/ptr"
)

func TestIndexExpr(t *testing.T) {

	t.Run("#String/returns bracketed zero-indexed postition",
		func(t *testing.T) {
			expected := "[42]"
			underTest := IndexExpr(42)
			actual := underTest.String()
			assert.Equal(t, expected, actual)
		})

	t.Run("#Access", func(t *testing.T) {

		t.Run("target is not indexable/returns not ok",
			func(t *testing.T) {

				testCases := []struct {
					name  string
					value any
				}{
					{name: "scalar", value: "string"},
					{name: "struct", value: struct{ Foo int }{}},
					{name: "map", value: map[string]string{"key": "val"}},

					{name: "pointer to scalar", value: ptr.To("string")},
					{name: "pointer to struct", value: &struct{ Foo int }{}},
					{name: "pointer to map", value: &map[string]string{"key": "val"}},

					{name: "pointer to pointer to scalar", value: ptr.To(ptr.To("string"))},
				}

				for _, tc := range testCases {

					t.Run(tc.name, func(t *testing.T) {
						underTest := IndexExpr(0)
						_, ok := underTest.Access(tc.value)
						assert.NotOk(t, ok)
					})
				}
			})

		t.Run("target is nil slice/returns untyped nil ok", func(t *testing.T) {
			var target []string
			underTest := IndexExpr(len(target))
			val, ok := underTest.Access(target)
			assert.Ok(t, ok)
			assert.Nil(t, val)
		})

		t.Run("target is non-nil slice/index out of range/returns untyped nil ok",
			func(t *testing.T) {
				target := []string{"zero", "one", "two"}
				underTest := IndexExpr(len(target))
				val, ok := underTest.Access(target)
				assert.Ok(t, ok)
				assert.Nil(t, val)
			})

		t.Run("target is non-nil slice/index in range/returns item value",
			func(t *testing.T) {
				expected := "expected"
				target := []string{"the", expected, "value"}
				underTest := IndexExpr(1)
				val, ok := underTest.Access(target)
				assert.Ok(t, ok)
				actual, ok := val.(string)
				assert.Ok(t, ok)
				assert.Equal(t, expected, actual)
			})

		t.Run("target is array/index out of range/returns not ok",
			// For arrays the length is part of the type so attempting to access
			// an out of range index is comparable attempting to access a field
			// that's not defined on a struct.
			func(t *testing.T) {
				target := [3]string{"zero", "one", "two"}
				underTest := IndexExpr(len(target))
				_, ok := underTest.Access(target)
				assert.NotOk(t, ok)
			})

		t.Run("target is array/index in range/returns item value",
			func(t *testing.T) {
				expected := "expected"
				target := [3]string{"the", expected, "value"}
				underTest := IndexExpr(1)
				val, ok := underTest.Access(target)
				assert.Ok(t, ok)
				actual, ok := val.(string)
				assert.Ok(t, ok)
				assert.Equal(t, expected, actual)
			})

		t.Run("target is defined type with kind slice/index in range/returns item value",
			func(t *testing.T) {
				type arr[T any] []T
				expected := "expected"
				target := arr[string]{"the", expected, "value"}
				underTest := IndexExpr(1)
				val, ok := underTest.Access(target)
				assert.Ok(t, ok)
				actual, ok := val.(string)
				assert.Ok(t, ok)
				assert.Equal(t, expected, actual)
			})

		t.Run("target is defined type with kind array/index in range/returns item value",
			func(t *testing.T) {
				type arr[T any] [3]T
				expected := "expected"
				target := arr[string]{"the", expected, "value"}
				underTest := IndexExpr(1)
				val, ok := underTest.Access(target)
				assert.Ok(t, ok)
				actual, ok := val.(string)
				assert.Ok(t, ok)
				assert.Equal(t, expected, actual)
			})

		t.Run("item is nil pointer/returns typed nil",
			func(t *testing.T) {
				target := []*string{ptr.To("the"), nil, ptr.To("value")}
				underTest := IndexExpr(1)
				val, ok := underTest.Access(target)
				assert.Ok(t, ok)
				actual, ok := val.(*string)
				assert.Ok(t, ok)
				assert.Nil(t, actual)
			})

		t.Run("item is non-nil pointer/returns item value",
			func(t *testing.T) {
				expected := ptr.To("expected")
				target := []*string{ptr.To("the"), expected, ptr.To("value")}
				underTest := IndexExpr(1)
				val, ok := underTest.Access(target)
				assert.Ok(t, ok)
				actual, ok := val.(*string)
				assert.Ok(t, ok)
				assert.Equal(t, expected, actual)
			})

		t.Run("target is nil pointer to slice/returns untyped nil ok",
			func(t *testing.T) {
				var target *[]string
				underTest := IndexExpr(1)
				val, ok := underTest.Access(target)
				assert.Ok(t, ok)
				assert.Nil(t, val)
			})

		t.Run("target is non-nil pointer to nil slice/returns untyped nil ok",
			func(t *testing.T) {
				target := ptr.To(([]string)(nil))
				underTest := IndexExpr(1)
				val, ok := underTest.Access(target)
				assert.Ok(t, ok)
				assert.Nil(t, val)
			})

		t.Run("target is non-nil pointer to non-nil slice/returns item value",
			func(t *testing.T) {
				expected := "expected"
				target := ptr.To([]string{"the", expected, "value"})
				underTest := IndexExpr(1)
				val, ok := underTest.Access(target)
				assert.Ok(t, ok)
				actual, ok := val.(string)
				assert.Ok(t, ok)
				assert.Equal(t, expected, actual)
			})
	})
}

func BenchmarkIndexExpr(b *testing.B) {

	indexExpr := IndexExpr(42)

	b.Run("#String", func(b *testing.B) {
		for range b.N {
			_ = indexExpr.String()
		}
	})

	b.Run("#stringUsingFmtSprintf", func(b *testing.B) {
		for range b.N {
			_ = indexExpr.stringUsingFmtSprintf()
		}
	})
}

func (i IndexExpr) stringUsingFmtSprintf() string {
	return fmt.Sprintf("[%d]", i)
}

package fields

import (
	"fmt"
	"testing"

	"github.com/patch-wreckless/godel/internal/ptr"
)

func TestIndexExpr(t *testing.T) {

	t.Run("#String/returns bracketed zero-indexed postition",
		func(t *testing.T) {
			expected := "[42]"
			underTest := IndexExpr(42)
			actual := underTest.String()
			if actual != expected {
				t.Errorf("expected %q; got %q", expected, actual)
			}
		})

	t.Run("#Access", func(t *testing.T) {

		t.Run("inapplicable to target type/returns not ok",
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
						val, ok := underTest.Access(tc.value)
						if ok {
							t.Errorf("expected not ok; got (%+v, true)", val)
						}
					})
				}
			})

		t.Run("target is nil slice/returns untyped nil ok", func(t *testing.T) {
			var target []string
			underTest := IndexExpr(len(target))
			val, ok := underTest.Access(target)
			if !ok {
				t.Error("expected ok to be true; got false")
			}
			if val != nil {
				t.Errorf("expected nil; got %+v", val)
			}
		})

		t.Run("target is non-nil slice/index out of range/returns untyped nil ok",
			func(t *testing.T) {
				target := []string{"zero", "one", "two"}
				underTest := IndexExpr(len(target))
				val, ok := underTest.Access(target)
				if !ok {
					t.Error("expected ok to be true; got false")
				}
				if val != nil {
					t.Errorf("expected nil; got %+v", val)
				}
			})

		t.Run("target is non-nil slice/index in range/returns item value",
			func(t *testing.T) {
				expected := "expected"
				target := []string{"the", expected, "value"}
				underTest := IndexExpr(1)
				val, ok := underTest.Access(target)
				if !ok {
					t.Error("expected ok to be true; got false")
				}
				actual, ok := val.(string)
				if !ok {
					t.Fatalf("expected %T; got %T", expected, val)
				}
				if actual != expected {
					t.Errorf("expected %q; got %q", expected, actual)
				}
			})

		t.Run("target is non-nil slice/item is nil pointer/returns typed nil",
			func(t *testing.T) {
				var expected *string
				target := []*string{ptr.To("the"), expected, ptr.To("value")}
				underTest := IndexExpr(1)
				val, ok := underTest.Access(target)
				if !ok {
					t.Error("expected ok to be true; got false")
				}
				actual, ok := val.(*string)
				if !ok {
					t.Fatalf("expected %T; got %T", expected, val)
				}
				if actual != expected {
					t.Errorf("expected %p; got %p", expected, actual)
				}
			})

		t.Run("target is non-nil slice/item is non-nil pointer/returns item value",
			func(t *testing.T) {
				expected := ptr.To("expected")
				target := []*string{ptr.To("the"), expected, ptr.To("value")}
				underTest := IndexExpr(1)
				val, ok := underTest.Access(target)
				if !ok {
					t.Error("expected ok to be true; got false")
				}
				actual, ok := val.(*string)
				if !ok {
					t.Fatalf("expected %T; got %T", expected, val)
				}
				if actual != expected {
					t.Errorf("expected %p; got %p", expected, actual)
				}
			})

		t.Run("target is defined type with kind slice/returns item value",
			func(t *testing.T) {
				type arr[T any] []T
				expected := "expected"
				target := arr[string]{"the", expected, "value"}
				underTest := IndexExpr(1)
				val, ok := underTest.Access(target)
				if !ok {
					t.Error("expected ok to be true; got false")
				}
				actual, ok := val.(string)
				if !ok {
					t.Fatalf("expected %T; got %T", expected, val)
				}
				if actual != expected {
					t.Errorf("expected %q; got %q", expected, actual)
				}
			})

		t.Run("target is nil pointer to slice/returns untyped nil ok",
			func(t *testing.T) {
				var target *[]string
				underTest := IndexExpr(1)
				val, ok := underTest.Access(target)
				if !ok {
					t.Error("expected ok to be true; got false")
				}
				if val != nil {
					t.Errorf("expected nil; got %+v", val)
				}
			})

		t.Run("target is non-nil pointer to nil slice/returns untyped nil ok",
			func(t *testing.T) {
				target := ptr.To(([]string)(nil))
				underTest := IndexExpr(1)
				val, ok := underTest.Access(target)
				if !ok {
					t.Error("expected ok to be true; got false")
				}
				if val != nil {
					t.Errorf("expected nil; got %+v", val)
				}
			})

		t.Run("value is pointer to non-nil slice/returns item value",
			func(t *testing.T) {
				expected := "expected"
				target := ptr.To([]string{"the", expected, "value"})
				underTest := IndexExpr(1)
				val, ok := underTest.Access(target)
				if !ok {
					t.Error("expected ok to be true; got false")
				}
				actual, ok := val.(string)
				if !ok {
					t.Fatalf("expected %T; got %T", expected, val)
				}
				if actual != expected {
					t.Errorf("expected %q; got %q", expected, actual)
				}
			})

		t.Run("target is array/index out of range/returns not ok",
			// For arrays the length is part of the type so attempting to access
			// an out of range index is comparable attempting to access a field
			// that's not defined on a struct.
			func(t *testing.T) {
				target := [3]string{"zero", "one", "two"}
				underTest := IndexExpr(len(target))
				val, ok := underTest.Access(target)
				if ok {
					t.Errorf("expected not ok; got (%+v, true)", val)
				}
			})

		t.Run("target is array/index in range/returns item value",
			func(t *testing.T) {
				expected := "expected"
				target := [3]string{"the", expected, "value"}
				underTest := IndexExpr(1)
				val, ok := underTest.Access(target)
				if !ok {
					t.Error("expected ok to be true; got false")
				}
				actual, ok := val.(string)
				if !ok {
					t.Fatalf("expected %T; got %T", expected, val)
				}
				if actual != expected {
					t.Errorf("expected %q; got %q", expected, actual)
				}
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

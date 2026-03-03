package fields

import (
	"fmt"
	"testing"

	"github.com/patch-wreckless/godel/internal/assert"
	"github.com/patch-wreckless/godel/internal/ptr"
)

func TestAllItemsExpr(t *testing.T) {

	t.Run("#String/returns AllItemsToken", func(t *testing.T) {
		expected := AllItemsToken
		underTest := AllItemsExpr{}
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
						underTest := AllItemsExpr{}
						_, ok := underTest.Access(tc.value)
						assert.NotOk(t, ok)
					})
				}
			})

		t.Run("target is nil slice/returns empty slice",
			func(t *testing.T) {
				var target []string
				underTest := AllItemsExpr{}
				actual, ok := underTest.Access(target)
				assert.Ok(t, ok)
				assert.NotNil(t, actual)
				assert.LenEq(t, 0, actual, func(lec *assert.LenEqConf[any, []any]) {
					lec.WithExpr("actual")
				})
			})

		t.Run("target is empty slice/returns empty slice",
			func(t *testing.T) {
				target := []string{}
				underTest := AllItemsExpr{}
				actual, ok := underTest.Access(target)
				assert.Ok(t, ok)
				assert.NotNil(t, actual)
				assert.LenEq(t, 0, actual, func(lec *assert.LenEqConf[any, []any]) {
					lec.WithExpr("actual")
				})
			})

		t.Run("target is non-empty slice/returns all values",
			func(t *testing.T) {
				target := []string{"the", "expected", "values"}
				expected := func() []any {
					r := make([]any, 0, len(target))
					for _, v := range target {
						r = append(r, v)
					}
					return r
				}()
				underTest := AllItemsExpr{}
				actual, ok := underTest.Access(target)
				assert.Ok(t, ok)
				assert.LenEq(t, len(expected), actual, func(lec *assert.LenEqConf[any, []any]) {
					lec.WithExpr("actual")
				})

				minLen := min(len(expected), len(actual))
				for i := range minLen {
					assert.Equal(t, expected[i], actual[i], func(ec *assert.EqualConf[any]) {
						ec.WithExpr(fmt.Sprintf("[%d]", i))
					})
				}
			})

		t.Run("target is empty array/returns empty slice",
			func(t *testing.T) {
				target := [0]string{}
				underTest := AllItemsExpr{}
				actual, ok := underTest.Access(target)
				assert.Ok(t, ok)
				assert.NotNil(t, actual)
				assert.LenEq(t, 0, actual)
			})

		t.Run("target is non-empty array/returns all values",
			func(t *testing.T) {
				target := [3]string{"the", "expected", "values"}
				expected := func() []any {
					r := make([]any, 0, len(target))
					for _, v := range target {
						r = append(r, v)
					}
					return r
				}()
				underTest := AllItemsExpr{}
				actual, ok := underTest.Access(target)
				assert.Ok(t, ok)
				assert.LenEq(t, len(expected), actual, func(lec *assert.LenEqConf[any, []any]) {
					lec.WithExpr("actual")
				})

				minLen := min(len(expected), len(actual))
				for i := range minLen {
					assert.Equal(t, expected[i], actual[i], func(ec *assert.EqualConf[any]) {
						ec.WithExpr(fmt.Sprintf("[%d]", i))
					})
				}
			})

		t.Run("target is defined type with kind slice/returns all values",
			func(t *testing.T) {
				type arr[T any] []T
				target := arr[string]{"the", "expected", "values"}
				expected := func() []any {
					r := make([]any, 0, len(target))
					for _, v := range target {
						r = append(r, v)
					}
					return r
				}()
				underTest := AllItemsExpr{}
				actual, ok := underTest.Access(target)
				assert.Ok(t, ok)
				assert.LenEq(t, len(expected), actual, func(lec *assert.LenEqConf[any, []any]) {
					lec.WithExpr("actual")
				})

				minLen := min(len(expected), len(actual))
				for i := range minLen {
					assert.Equal(t, expected[i], actual[i], func(ec *assert.EqualConf[any]) {
						ec.WithExpr(fmt.Sprintf("[%d]", i))
					})
				}
			})

		t.Run("target is defined type with kind array/returns all values",
			func(t *testing.T) {
				type arr[T any] [3]T
				target := arr[string]{"the", "expected", "values"}
				expected := func() []any {
					r := make([]any, 0, len(target))
					for _, v := range target {
						r = append(r, v)
					}
					return r
				}()
				underTest := AllItemsExpr{}
				actual, ok := underTest.Access(target)
				assert.Ok(t, ok)
				assert.LenEq(t, len(expected), actual, func(lec *assert.LenEqConf[any, []any]) {
					lec.WithExpr("actual")
				})

				minLen := min(len(expected), len(actual))
				for i := range minLen {
					assert.Equal(t, expected[i], actual[i], func(ec *assert.EqualConf[any]) {
						ec.WithExpr(fmt.Sprintf("[%d]", i))
					})
				}
			})

		t.Run("target is nil pointer to slice/returns untyped nil ok",
			func(t *testing.T) {
				var target *[]string
				underTest := AllItemsExpr{}
				val, ok := underTest.Access(target)
				assert.Ok(t, ok)
				assert.Nil(t, val)
			})

		t.Run("target is non-nil pointer to nil slice/returns untyped nil ok",
			func(t *testing.T) {
				target := ptr.To(([]string)(nil))
				underTest := AllItemsExpr{}
				actual, ok := underTest.Access(target)
				assert.Ok(t, ok)
				assert.NotNil(t, actual)
				assert.LenEq(t, 0, actual, func(lec *assert.LenEqConf[any, []any]) {
					lec.WithExpr("actual")
				})
			})

		t.Run("target is non-nil pointer to non-nil slice/returns item value",
			func(t *testing.T) {
				target := ptr.To([]string{"the", "expected", "values"})
				expected := func() []any {
					r := make([]any, 0, len(*target))
					for _, v := range *target {
						r = append(r, v)
					}
					return r
				}()
				underTest := AllItemsExpr{}
				actual, ok := underTest.Access(target)
				assert.Ok(t, ok)
				assert.LenEq(t, len(expected), actual, func(lec *assert.LenEqConf[any, []any]) {
					lec.WithExpr("actual")
				})

				minLen := min(len(expected), len(actual))
				for i := range minLen {
					assert.Equal(t, expected[i], actual[i], func(ec *assert.EqualConf[any]) {
						ec.WithExpr(fmt.Sprintf("[%d]", i))
					})
				}
			})
	})
}

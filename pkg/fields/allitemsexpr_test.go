package fields

import (
	"testing"

	"github.com/patch-wreckless/godel/internal/ptr"
)

func TestAllItemsExpr(t *testing.T) {

	t.Run("#String/returns AllItemsToken", func(t *testing.T) {
		expected := AllItemsToken
		underTest := AllItemsExpr{}
		actual := underTest.String()
		if actual != expected {
			t.Errorf("expected %q; got %q", expected, actual)
		}
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
						val, ok := underTest.Access(tc.value)
						if ok {
							t.Errorf("expected not ok; got (%+v, true)", val)
						}
					})
				}
			})

		t.Run("target is nil slice/returns empty slice",
			func(t *testing.T) {
				expected := []any{}
				var target []string
				underTest := AllItemsExpr{}
				actual, ok := underTest.Access(target)
				if !ok {
					t.Error("expected ok to be true; got false")
				}
				if actual == nil {
					t.Errorf("expected %v; got nil", expected)
				}
				if len(expected) != 0 {
					t.Errorf("expected empty slice; got %v", actual)
				}
			})

		t.Run("target is empty slice/returns empty slice",
			func(t *testing.T) {
				expected := []any{}
				target := []string{}
				underTest := AllItemsExpr{}
				actual, ok := underTest.Access(target)
				if !ok {
					t.Error("expected ok to be true; got false")
				}
				if actual == nil {
					t.Errorf("expected %v; got nil", expected)
				}
				if len(expected) != 0 {
					t.Errorf("expected empty slice; got %v", actual)
				}
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
				if !ok {
					t.Error("expected ok to be true; got false")
				}

				minLen := min(len(expected), len(actual))
				for i := range minLen {
					exp := expected[i]
					act := actual[i]
					if actual[i] != expected[i] {
						t.Errorf("expected [%d] to be %q; got %q", i, exp, act)
					}
				}
				if len(expected) > minLen {
					t.Errorf("missing expected items %v", expected[minLen:])
				}
				if len(actual) > minLen {
					t.Errorf("got unexpected items %v", actual[minLen:])
				}
			})

		t.Run("target is empty array/returns empty slice",
			func(t *testing.T) {
				expected := []any{}
				target := [0]string{}
				underTest := AllItemsExpr{}
				actual, ok := underTest.Access(target)
				if !ok {
					t.Error("expected ok to be true; got false")
				}
				if actual == nil {
					t.Errorf("expected %v; got nil", expected)
				}
				if len(expected) != 0 {
					t.Errorf("expected empty slice; got %v", actual)
				}
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
				if !ok {
					t.Error("expected ok to be true; got false")
				}

				minLen := min(len(expected), len(actual))
				for i := range minLen {
					exp := expected[i]
					act := actual[i]
					if actual[i] != expected[i] {
						t.Errorf("expected [%d] to be %q; got %q", i, exp, act)
					}
				}
				if len(expected) > minLen {
					t.Errorf("missing expected items %v", expected[minLen:])
				}
				if len(actual) > minLen {
					t.Errorf("got unexpected items %v", actual[minLen:])
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
				if !ok {
					t.Error("expected ok to be true; got false")
				}

				minLen := min(len(expected), len(actual))
				for i := range minLen {
					exp := expected[i]
					act := actual[i]
					if actual[i] != expected[i] {
						t.Errorf("expected [%d] to be %q; got %q", i, exp, act)
					}
				}
				if len(expected) > minLen {
					t.Errorf("missing expected items %v", expected[minLen:])
				}
				if len(actual) > minLen {
					t.Errorf("got unexpected items %v", actual[minLen:])
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
				if !ok {
					t.Error("expected ok to be true; got false")
				}

				minLen := min(len(expected), len(actual))
				for i := range minLen {
					exp := expected[i]
					act := actual[i]
					if actual[i] != expected[i] {
						t.Errorf("expected [%d] to be %q; got %q", i, exp, act)
					}
				}
				if len(expected) > minLen {
					t.Errorf("missing expected items %v", expected[minLen:])
				}
				if len(actual) > minLen {
					t.Errorf("got unexpected items %v", actual[minLen:])
				}
			})

		t.Run("target is nil pointer to slice/returns untyped nil ok",
			func(t *testing.T) {
				var target *[]string
				underTest := AllItemsExpr{}
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
				expected := []any{}
				target := ptr.To(([]string)(nil))
				underTest := AllItemsExpr{}
				actual, ok := underTest.Access(target)
				if !ok {
					t.Error("expected ok to be true; got false")
				}
				if actual == nil {
					t.Errorf("expected %v; got nil", expected)
				}
				if len(expected) != 0 {
					t.Errorf("expected empty slice; got %v", actual)
				}
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
				if !ok {
					t.Error("expected ok to be true; got false")
				}

				minLen := min(len(expected), len(actual))
				for i := range minLen {
					exp := expected[i]
					act := actual[i]
					if actual[i] != expected[i] {
						t.Errorf("expected [%d] to be %q; got %q", i, exp, act)
					}
				}
				if len(expected) > minLen {
					t.Errorf("missing expected items %v", expected[minLen:])
				}
				if len(actual) > minLen {
					t.Errorf("got unexpected items %v", actual[minLen:])
				}
			})
	})
}

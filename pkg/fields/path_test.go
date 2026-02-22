package fields

import (
	"testing"
)

func TestPath(t *testing.T) {

	t.Run("#Segments", func(t *testing.T) {

		t.Run("Path initialized with nil/returns nil", func(t *testing.T) {
			underTest := NewPath(nil)
			if segments := underTest.Segments(); segments != nil {
				t.Fatalf("expected nil; got %+v", segments)
			}
		})

		t.Run("Path initialized with empty slice/returns nil",
			func(t *testing.T) {
				underTest := NewPath([]PathSegment{})
				if segments := underTest.Segments(); segments != nil {
					t.Fatalf("expected nil; got %+v", segments)
				}
			})

		t.Run("Path initialized with segments/returns segments",
			func(t *testing.T) {
				expected := []PathSegment{
					NewPathSegment(MustFieldName("Foo")),
					NewPathSegment(IndexExpr(0)),
					NewPathSegment(MustFieldName("Bar")),
					NewPathSegment(AllItemsExpr{}),
					NewPathSegment(MustFieldName("Baz")),
				}
				underTest := NewPath(expected)
				actual := underTest.Segments()

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

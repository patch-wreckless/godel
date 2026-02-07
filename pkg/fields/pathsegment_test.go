package fields

import (
	"testing"
)

func TestPathSegment(t *testing.T) {

	var (
		fieldName    = MustFieldName("foo")
		indexExpr    = IndexExpr(47)
		allItemsExpr = AllItemsExpr{}
	)

	t.Run("#Value", func(t *testing.T) {

		t.Run("underlying value is FieldName/returns FieldName",
			func(t *testing.T) {
				expected := fieldName
				underTest := NewPathSegment(expected)
				val := underTest.Value()
				actual, ok := val.(FieldName)
				if !ok {
					t.Fatalf("expected %T; got %T", expected, val)
				}
				if actual != expected {
					t.Fatalf("expected %+v; got %+v", expected, val)
				}
			})

		t.Run("underlying value is IndexExpr/returns IndexExpr",
			func(t *testing.T) {
				expected := indexExpr
				underTest := NewPathSegment(expected)
				val := underTest.Value()
				actual, ok := val.(IndexExpr)
				if !ok {
					t.Fatalf("expected %T; got %T", expected, val)
				}
				if actual != expected {
					t.Fatalf("expected %+v; got %+v", expected, val)
				}
			})

		t.Run("underlying value is AllItemsExpr/returns AllItemsExpr",
			func(t *testing.T) {
				expected := allItemsExpr
				underTest := NewPathSegment(expected)
				val := underTest.Value()
				actual, ok := val.(AllItemsExpr)
				if !ok {
					t.Fatalf("expected %T; got %T", expected, val)
				}
				if actual != expected {
					t.Fatalf("expected %+v; got %+v", expected, val)
				}
			})
	})

	t.Run("#String", func(t *testing.T) {

		t.Run("underlying value is FieldName/returns FieldName.String()",
			func(t *testing.T) {
				expected := fieldName.String()
				underTest := NewPathSegment(fieldName)
				actual := underTest.String()
				if actual != expected {
					t.Fatalf("expected %q; got %q", expected, actual)
				}
			})

		t.Run("underlying value is IndexExpr/returns IndexExpr.String()",
			func(t *testing.T) {
				expected := indexExpr.String()
				underTest := NewPathSegment(indexExpr)
				actual := underTest.String()
				if actual != expected {
					t.Fatalf("expected %q; got %q", expected, actual)
				}
			})

		t.Run("underlying value is AllItemsExpr/returns AllItemsExpr.String()",
			func(t *testing.T) {
				expected := allItemsExpr.String()
				underTest := NewPathSegment(allItemsExpr)
				actual := underTest.String()
				if actual != expected {
					t.Fatalf("expected %q; got %q", expected, actual)
				}
			})
	})
}

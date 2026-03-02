package fields

import (
	"testing"

	"github.com/patch-wreckless/godel/internal/assert"
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
				assert.Ok(t, ok)
				assert.Equal(t, expected, actual)
			})

		t.Run("underlying value is IndexExpr/returns IndexExpr",
			func(t *testing.T) {
				expected := indexExpr
				underTest := NewPathSegment(expected)
				val := underTest.Value()
				actual, ok := val.(IndexExpr)
				assert.Ok(t, ok)
				assert.Equal(t, expected, actual)
			})

		t.Run("underlying value is AllItemsExpr/returns AllItemsExpr",
			func(t *testing.T) {
				expected := allItemsExpr
				underTest := NewPathSegment(expected)
				val := underTest.Value()
				actual, ok := val.(AllItemsExpr)
				assert.Ok(t, ok)
				assert.Equal(t, expected, actual)
			})
	})

	t.Run("#String", func(t *testing.T) {

		t.Run("underlying value is FieldName/returns FieldName.String()",
			func(t *testing.T) {
				expected := fieldName.String()
				underTest := NewPathSegment(fieldName)
				actual := underTest.String()
				assert.Equal(t, expected, actual)
			})

		t.Run("underlying value is IndexExpr/returns IndexExpr.String()",
			func(t *testing.T) {
				expected := indexExpr.String()
				underTest := NewPathSegment(indexExpr)
				actual := underTest.String()
				assert.Equal(t, expected, actual)
			})

		t.Run("underlying value is AllItemsExpr/returns AllItemsExpr.String()",
			func(t *testing.T) {
				expected := allItemsExpr.String()
				underTest := NewPathSegment(allItemsExpr)
				actual := underTest.String()
				assert.Equal(t, expected, actual)
			})
	})
}

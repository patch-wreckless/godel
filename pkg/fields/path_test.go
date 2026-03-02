package fields

import (
	"fmt"
	"testing"

	"github.com/patch-wreckless/godel/internal/assert"
)

func TestPath(t *testing.T) {

	t.Run("#Segments", func(t *testing.T) {

		t.Run("Path initialized with nil/returns nil", func(t *testing.T) {
			underTest := NewPath(nil)
			actual := underTest.Segments()
			assert.Nil(t, actual)
		})

		t.Run("Path initialized with empty slice/returns nil",
			func(t *testing.T) {
				underTest := NewPath([]PathSegment{})
				actual := underTest.Segments()
				assert.Nil(t, actual)
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

				assert.Equal(t, len(expected), len(actual), func(ec *assert.EqualConf[int]) {
					ec.WithExpr("len()")
				})
				minLen := min(len(expected), len(actual))
				for i := range minLen {
					assert.Equal(t, expected[i], actual[i], func(ec *assert.EqualConf[PathSegment]) {
						ec.WithExpr(fmt.Sprintf("[%d]", i))
					})
				}
			})
	})
}

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
	})
}

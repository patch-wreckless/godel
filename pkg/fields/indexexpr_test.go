package fields

import (
	"fmt"
	"testing"
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

package fields

import (
	"testing"
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
}

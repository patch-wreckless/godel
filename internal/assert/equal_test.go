package assert

import (
	"testing"
)

func TestEqual(t *testing.T) {

	t.Run("values are equal/does not report error", func(t *testing.T) {
		mockT := newMockTestingT(t)
		Equal(mockT, 5, 5)
	})

	t.Run("values are not equal", func(t *testing.T) {

		t.Run("reports error", func(t *testing.T) {
			mockT := newMockTestingT(t)
			mockT.Expect().Errorf("expected %v; got %v", 5, 4)
			Equal(mockT, 5, 4)
		})

		t.Run("expression configured/reports error with expression", func(t *testing.T) {
			mockT := newMockTestingT(t)
			mockT.Expect().Errorf("expected %s to be %v; got %v", "value", 5, 4)
			Equal(mockT, 5, 4, func(ec *EqualConf[int]) {
				ec.WithExpr("value")
			})
		})
	})
}

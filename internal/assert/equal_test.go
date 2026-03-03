package assert

import (
	"testing"
)

func TestEqual(t *testing.T) {

	t.Run("comparable values", func(t *testing.T) {

		t.Run("are equal/does not report error", func(t *testing.T) {
			mockT := newMockTestingT(t)
			Equal(mockT, 5, 5)
		})

		t.Run("are not equal/reports error", func(t *testing.T) {
			mockT := newMockTestingT(t)
			mockT.Expect().Errorf("expected %v; got %v", 5, 4)
			Equal(mockT, 5, 4)
		})
	})

	t.Run("interface values", func(t *testing.T) {

		t.Run("are equal/does not report error", func(t *testing.T) {
			mockT := newMockTestingT(t)
			Equal(mockT, "six", "six")
		})

		t.Run("are not equal/reports error", func(t *testing.T) {
			mockT := newMockTestingT(t)
			mockT.Expect().Errorf("expected %v; got %v", "six", "half of a dozen")
			Equal(mockT, "six", "half of a dozen")
		})
	})

	t.Run("reporting error", func(t *testing.T) {

		t.Run("with expression configured/error includes expression", func(t *testing.T) {
			mockT := newMockTestingT(t)
			mockT.Expect().Errorf("expected %s to be %v; got %v", "value", 5, 4)
			Equal(mockT, 5, 4, func(ec *EqualConf[int]) {
				ec.WithExpr("value")
			})
		})
	})
}

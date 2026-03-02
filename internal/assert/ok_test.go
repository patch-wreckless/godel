package assert

import "testing"

func TestOk(t *testing.T) {

	t.Run("value is true/does not report error", func(t *testing.T) {
		mockT := newMockTestingT(t)
		Ok(mockT, true)
	})

	t.Run("value is false/reports error", func(t *testing.T) {
		mockT := newMockTestingT(t)
		mockT.Expect().Errorf("expected %s to be %v; got %v", "ok", true, false)
		Ok(mockT, false)
	})
}

func TestNotOk(t *testing.T) {

	t.Run("value is false/does not report error", func(t *testing.T) {
		mockT := newMockTestingT(t)
		NotOk(mockT, false)
	})

	t.Run("value is true/reports error", func(t *testing.T) {
		mockT := newMockTestingT(t)
		mockT.Expect().Errorf("expected %s to be %v; got %v", "ok", false, true)
		NotOk(mockT, true)
	})
}

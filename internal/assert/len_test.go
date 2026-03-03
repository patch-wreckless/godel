package assert

import "testing"

func TestLenSliceEq(t *testing.T) {

	t.Run("expect negative/reports error", func(t *testing.T) {
		mockT := newMockTestingT(t)
		mockT.Expect().Errorf("expected %s to be %v; got %v", "len()", -2, 3)
		LenEq(mockT, -2, []int{1, 2, 3})
	})

	t.Run("nil slice", func(t *testing.T) {

		t.Run("expect zero/does not report error", func(t *testing.T) {
			mockT := newMockTestingT(t)
			var s []int
			LenEq(mockT, 0, s)
		})

		t.Run("expect non-zero/reports error", func(t *testing.T) {
			mockT := newMockTestingT(t)
			mockT.Expect().Errorf("expected %s to be %v; got %v", "len()", 3, 0)
			var s []int
			LenEq(mockT, 3, s)
		})
	})

	t.Run("non-nil empty slice", func(t *testing.T) {

		t.Run("expect zero/does not report error", func(t *testing.T) {
			mockT := newMockTestingT(t)
			LenEq(mockT, 0, []int{})
		})

		t.Run("expect non-zero/reports error", func(t *testing.T) {
			mockT := newMockTestingT(t)
			mockT.Expect().Errorf("expected %s to be %v; got %v", "len()", 3, 0)
			LenEq(mockT, 3, []int{})
		})
	})

	t.Run("non-empty slice", func(t *testing.T) {

		t.Run("expect matches len/does not report error", func(t *testing.T) {
			mockT := newMockTestingT(t)
			LenEq(mockT, 3, []int{1, 2, 3})
		})

		t.Run("expect does not match len/reports error", func(t *testing.T) {
			mockT := newMockTestingT(t)
			mockT.Expect().Errorf("expected %s to be %v; got %v", "len()", 2, 3)
			LenEq(mockT, 2, []int{1, 2, 3})
		})
	})
}

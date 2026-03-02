package assert

import (
	"testing"
	"unsafe"
)

func TestNil(t *testing.T) {

	t.Run("value is nil/does not report error", func(t *testing.T) {

		for _, tc := range nilTestCases {

			t.Run(tc.name, func(t *testing.T) {
				mockT := newMockTestingT(t)
				Nil(mockT, tc.val)
			})
		}
	})

	t.Run("value is not nil/reports error", func(t *testing.T) {
		mockT := newMockTestingT(t)
		mockT.Expect().Errorf("expected nil; got %v", 5)
		Nil(mockT, 5)
	})
}

func TestNotNil(t *testing.T) {

	t.Run("value is not nil/does not report error", func(t *testing.T) {
		mockT := newMockTestingT(t)
		NotNil(mockT, 5)
	})

	t.Run("value is nil/reports error", func(t *testing.T) {

		for _, tc := range nilTestCases {

			t.Run(tc.name, func(t *testing.T) {
				mockT := newMockTestingT(t)
				mockT.Expect().Error("expected non-nil")
				NotNil(mockT, tc.val)
			})
		}
	})
}

type nilTestCase struct {
	name string
	val  any
}

var nilTestCases = []nilTestCase{
	{
		name: "nil",
		val:  nil,
	},
	{
		name: "nil chan",
		val: func() any {
			var c chan int
			return c
		}(),
	},
	{
		name: "nil func",
		val: func() any {
			var fn func()
			return fn
		}(),
	},
	{
		name: "nil interface",
		val: func() any {
			var i interface{}
			return i
		}(),
	},
	{
		name: "nil map",
		val: func() any {
			var m map[int]int
			return m
		}(),
	},
	{
		name: "nil pointer",
		val: func() any {
			var p *int
			return p
		}(),
	},
	{
		name: "nil slice",
		val: func() any {
			var s []int
			return s
		}(),
	},
	{
		name: "nil uintptr",
		val: func() any {
			var p uintptr
			return p
		}(),
	},
	{
		name: "nil slice",
		val: func() any {
			var p unsafe.Pointer
			return p
		}(),
	},
}

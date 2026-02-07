package godel

import (
	"errors"
	"testing"
)

func TestMax(t *testing.T) {

	t.Run("#Check", func(t *testing.T) {

		t.Run("non-numeric value/returns Inapplicable", func(t *testing.T) {
			underTest := Min{Value: 10}
			expected := Inapplicable{
				Constraint: underTest,
				Value:      "non-numeric",
			}
			violations := underTest.Check(expected.Value)
			if n := len(violations); n != 1 {
				t.Fatalf("expected 1 violation; got %d", n)
			}
			var actual Inapplicable
			if err := violations[0].Error; !errors.As(err, &actual) {
				t.Fatalf(
					"expected violations[1].Error to be %T; got %T",
					expected,
					err)
			}
			if actual != expected {
				t.Fatalf(
					"expected violations[1].Error to be %v; got %v",
					expected,
					actual)
			}
		})

		const (
			maxVal = 10
			below  = 5
			above  = 15
		)

		type definedType int

		testCases := []struct {
			name        string
			value       any
			expectError bool
		}{
			{"int8 below min/no violation", int8(below), false},
			{"uint8 below min/no violation", uint8(below), false},
			{"int16 below min/no violation", int16(below), false},
			{"uint16 below min/no violation", uint16(below), false},
			{"int32 below min/no violation", int32(below), false},
			{"uint32 below min/no violation", uint32(below), false},
			{"int64 below min/no violation", int64(below), false},
			{"uint64 below min/no violation", uint64(below), false},
			{"int below min/no violation", int(below), false},
			{"uint below min/no violation", uint(below), false},
			{"float32 below min/no violation", float32(below), false},
			{"float64 below min/no violation", float64(below), false},
			{"defined type value below min/no violation", definedType(below), false},

			{"int8 equal to min/no violation", int8(maxVal), false},
			{"uint8 equal to min/no violation", uint8(maxVal), false},
			{"int16 equal to min/no violation", int16(maxVal), false},
			{"uint16 equal to min/no violation", uint16(maxVal), false},
			{"int32 equal to min/no violation", int32(maxVal), false},
			{"uint32 equal to min/no violation", uint32(maxVal), false},
			{"int64 equal to min/no violation", int64(maxVal), false},
			{"uint64 equal to min/no violation", uint64(maxVal), false},
			{"int equal to min/no violation", int(maxVal), false},
			{"uint equal to min/no violation", uint(maxVal), false},
			{"float32 equal to min/no violation", float32(float64(maxVal)), false},
			{"float64 equal to min/no violation", float64(maxVal), false},
			{"defined type value equal to min/no violation", definedType(maxVal), false},

			{"int8 above max/returns violation", int8(above), true},
			{"uint8 above max/returns violation", uint8(above), true},
			{"int16 above max/returns violation", int16(above), true},
			{"uint16 above max/returns violation", uint16(above), true},
			{"int32 above max/returns violation", int32(above), true},
			{"uint32 above max/returns violation", uint32(above), true},
			{"int64 above max/returns violation", int64(above), true},
			{"uint64 above max/returns violation", uint64(above), true},
			{"int above max/returns violation", int(above), true},
			{"uint above max/returns violation", uint(above), true},
			{"float32 above max/returns violation", float32(float64(above)), true},
			{"float64 above max/returns violation", float64(above), true},
			{"defined type value above max/returns violation", definedType(above), true},
		}

		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {
				underTest := Max{Value: maxVal}
				violations := underTest.Check(tc.value)
				if tc.expectError && len(violations) == 0 {
					t.Fatalf("expected violation; got none")
				}
				if !tc.expectError && len(violations) > 0 {
					t.Fatalf("expected no violations; got %+v", violations)
				}
			})
		}
	})
}

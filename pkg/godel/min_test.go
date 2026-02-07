package godel

import (
	"errors"
	"testing"
)

func TestMin(t *testing.T) {

	t.Run("#Check", func(t *testing.T) {

		t.Run("non-numeric value/returns Inapplicable", func(t *testing.T) {
			underTest := Max{Value: 10}
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
			minVal = 10
			below  = 5
			above  = 15
		)

		type definedType int

		testCases := []struct {
			name        string
			value       any
			expectError bool
		}{
			{"int8 above min/no violation", int8(above), false},
			{"uint8 above min/no violation", uint8(above), false},
			{"int16 above min/no violation", int16(above), false},
			{"uint16 above min/no violation", uint16(above), false},
			{"int32 above min/no violation", int32(above), false},
			{"uint32 above min/no violation", uint32(above), false},
			{"int64 above min/no violation", int64(above), false},
			{"uint64 above min/no violation", uint64(above), false},
			{"int above min/no violation", int(above), false},
			{"uint above min/no violation", uint(above), false},
			{"float32 above min/no violation", float32(float64(above)), false},
			{"float64 above min/no violation", float64(above), false},
			{"defined type value above min/no violation", definedType(above), false},

			{"int8 equal to min/no violation", int8(minVal), false},
			{"uint8 equal to min/no violation", uint8(minVal), false},
			{"int16 equal to min/no violation", int16(minVal), false},
			{"uint16 equal to min/no violation", uint16(minVal), false},
			{"int32 equal to min/no violation", int32(minVal), false},
			{"uint32 equal to min/no violation", uint32(minVal), false},
			{"int64 equal to min/no violation", int64(minVal), false},
			{"uint64 equal to min/no violation", uint64(minVal), false},
			{"int equal to min/no violation", int(minVal), false},
			{"uint equal to min/no violation", uint(minVal), false},
			{"float32 equal to min/no violation", float32(float64(minVal)), false},
			{"float64 equal to min/no violation", float64(minVal), false},
			{"defined type value equal to min/no violation", definedType(minVal), false},

			{"int8 below min/returns violation", int8(below), true},
			{"uint8 below min/returns violation", uint8(below), true},
			{"int16 below min/returns violation", int16(below), true},
			{"uint16 below min/returns violation", uint16(below), true},
			{"int32 below min/returns violation", int32(below), true},
			{"uint32 below min/returns violation", uint32(below), true},
			{"int64 below min/returns violation", int64(below), true},
			{"uint64 below min/returns violation", uint64(below), true},
			{"int below min/returns violation", int(below), true},
			{"uint below min/returns violation", uint(below), true},
			{"float32 below min/returns violation", float32(float64(below)), true},
			{"float64 below min/returns violation", float64(below), true},
			{"defined type value below min/returns violation", definedType(below), true},
		}

		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {
				underTest := Min{Value: minVal}
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

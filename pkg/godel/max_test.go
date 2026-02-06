package godel

import (
	"testing"
)

func TestMax(t *testing.T) {

	t.Run("#Check", func(t *testing.T) {

		t.Run("non-numeric value/panics", func(t *testing.T) {
			defer func() {
				if r := recover(); r == nil {
					t.Fatalf("expected panic; got none")
				}
			}()

			underTest := Max{Value: 10}
			underTest.Check("non-numeric")
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
			maxValue    int64
			expectError bool
		}{
			{"int8 below max/no violation", int8(below), maxVal, false},
			{"uint8 below max/no violation", uint8(below), maxVal, false},
			{"int16 below max/no violation", int16(below), maxVal, false},
			{"uint16 below max/no violation", uint16(below), maxVal, false},
			{"int32 below max/no violation", int32(below), maxVal, false},
			{"uint32 below max/no violation", uint32(below), maxVal, false},
			{"int64 below max/no violation", int64(below), maxVal, false},
			{"uint64 below max/no violation", uint64(below), maxVal, false},
			{"int below max/no violation", int(below), maxVal, false},
			{"uint below max/no violation", uint(below), maxVal, false},
			{"float32 below max/no violation", float32(below), maxVal, false},
			{"float64 below max/no violation", float64(below), maxVal, false},
			{"defined type value below max/no violation", definedType(below), maxVal, false},

			{"int8 equal to max/no violation", int8(maxVal), maxVal, false},
			{"uint8 equal to max/no violation", uint8(maxVal), maxVal, false},
			{"int16 equal to max/no violation", int16(maxVal), maxVal, false},
			{"uint16 equal to max/no violation", uint16(maxVal), maxVal, false},
			{"int32 equal to max/no violation", int32(maxVal), maxVal, false},
			{"uint32 equal to max/no violation", uint32(maxVal), maxVal, false},
			{"int64 equal to max/no violation", int64(maxVal), maxVal, false},
			{"uint64 equal to max/no violation", uint64(maxVal), maxVal, false},
			{"int equal to max/no violation", int(maxVal), maxVal, false},
			{"uint equal to max/no violation", uint(maxVal), maxVal, false},
			{"float32 equal to max/no violation", float32(float64(maxVal)), maxVal, false},
			{"float64 equal to max/no violation", float64(maxVal), maxVal, false},
			{"defined type value equal to max/no violation", definedType(maxVal), maxVal, false},

			{"int8 above max/returns violation", int8(above), maxVal, true},
			{"uint8 above max/returns violation", uint8(above), maxVal, true},
			{"int16 above max/returns violation", int16(above), maxVal, true},
			{"uint16 above max/returns violation", uint16(above), maxVal, true},
			{"int32 above max/returns violation", int32(above), maxVal, true},
			{"uint32 above max/returns violation", uint32(above), maxVal, true},
			{"int64 above max/returns violation", int64(above), maxVal, true},
			{"uint64 above max/returns violation", uint64(above), maxVal, true},
			{"int above max/returns violation", int(above), maxVal, true},
			{"uint above max/returns violation", uint(above), maxVal, true},
			{"float32 above max/returns violation", float32(float64(above)), maxVal, true},
			{"float64 above max/returns violation", float64(above), maxVal, true},
			{"defined type value above max/returns violation", definedType(above), maxVal, true},
		}

		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {
				underTest := Max{Value: tc.maxValue}
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

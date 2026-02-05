package godel

import (
	"reflect"
	"testing"
)

func TestMax(t *testing.T) {

	t.Run("#Path/returns .TargetPath", func(t *testing.T) {
		expected := []FieldName{"foo", "bar", "baz"}
		underTest := Max{TargetPath: expected}
		actual := underTest.Path()
		if !reflect.DeepEqual(expected, actual) {
			t.Fatalf("expected %+v; got %+v", expected, actual)
		}
	})

	t.Run("#Check", func(t *testing.T) {

		type definedType int

		testCases := []struct {
			name        string
			value       any
			maxValue    int64
			expectError bool
		}{
			{"int8 below max/no violation", int8(5), 10, false},
			{"int8 equal to max/no violation", int8(10), 10, false},
			{"int8 above max/returns violation", int8(15), 10, true},
			{"int16 below max/no violation", int16(5), 10, false},
			{"int16 equal to max/no violation", int16(10), 10, false},
			{"int16 above max/returns violation", int16(15), 10, true},
			{"int32 below max/no violation", int32(5), 10, false},
			{"int32 equal to max/no violation", int32(10), 10, false},
			{"int32 above max/returns violation", int32(15), 10, true},
			{"int64 below max/no violation", int64(5), 10, false},
			{"int64 equal to max/no violation", int64(10), 10, false},
			{"int64 above max/returns violation", int64(15), 10, true},
			{"int below max/no violation", int(5), 10, false},
			{"int equal to max/no violation", int(10), 10, false},
			{"int above max/returns violation", int(15), 10, true},
			{"uint8 below max/no violation", uint8(5), 10, false},
			{"uint8 equal to max/no violation", uint8(10), 10, false},
			{"uint8 above max/returns violation", uint8(15), 10, true},
			{"uint16 below max/no violation", uint16(5), 10, false},
			{"uint16 equal to max/no violation", uint16(10), 10, false},
			{"uint16 above max/returns violation", uint16(15), 10, true},
			{"uint32 below max/no violation", uint32(5), 10, false},
			{"uint32 equal to max/no violation", uint32(10), 10, false},
			{"uint32 above max/returns violation", uint32(15), 10, true},
			{"uint64 below max/no violation", uint64(5), 10, false},
			{"uint64 equal to max/no violation", uint64(10), 10, false},
			{"uint64 above max/returns violation", uint64(15), 10, true},
			{"uint below max/no violation", uint(5), 10, false},
			{"uint equal to max/no violation", uint(10), 10, false},
			{"uint above max/returns violation", uint(15), 10, true},
			{"float32 below max/no violation", float32(5.5), 10, false},
			{"float32 equal to max/no violation", float32(10.0), 10, false},
			{"float32 above max/returns violation", float32(15.2), 10, true},
			{"float64 below max/no violation", float64(5.5), 10, false},
			{"float64 equal to max/no violation", float64(10.0), 10, false},
			{"float64 above max/returns violation", float64(15.2), 10, true},
			{"defined type value below max/no violation", definedType(5), 10, false},
			{"defined type value equal to max/no violation", definedType(10), 10, false},
			{"defined type value above max/returns violation", definedType(15), 10, true},
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

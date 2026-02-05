package godel

import (
	"fmt"
	"reflect"
	"slices"
)

var _ Constraint = Max{}

// Max is a [Constraint] requiring a value to be less than or equal to a given value.
type Max struct {

	// TargetPath determines where the [Max] should be applied to the evaluated values.
	TargetPath []FieldName

	// Value is the maximum value that will be considered valid.
	Value int64
}

func (m Max) Path() []FieldName {
	return slices.Clone(m.TargetPath)
}

func (m Max) Check(val any) Violations {
	rv := reflect.ValueOf(val)
	kind := rv.Kind()

	valid := func() bool {
		switch kind {
		case reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uint:
			return rv.Uint() <= uint64(m.Value)
		case reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Int:
			return rv.Int() <= m.Value
		case reflect.Float32, reflect.Float64:
			return rv.Float() <= float64(m.Value)
		default:
			panic(fmt.Sprintf("Max.Check received non-numeric value: %v", val))
		}
	}()

	if !valid {
		return Violations{
			{
				Description: fmt.Sprintf("value %v exceeds max value %d", val, m.Value),
			},
		}
	}

	return nil
}

package godel

import (
	"fmt"
	"reflect"
)

var _ Constraint = Max{}

// Min is a [Constraint] requiring a value to be greater than or equal to a given value.
type Min struct {

	// Value is the minimum value that will be considered valid.
	Value int64
}

func (m Min) Check(val any) Violations {
	rv := reflect.ValueOf(val)
	kind := rv.Kind()

	valid := func() bool {
		switch kind {
		case reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uint:
			return rv.Uint() >= uint64(m.Value)
		case reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Int:
			return rv.Int() >= m.Value
		case reflect.Float32, reflect.Float64:
			return rv.Float() >= float64(m.Value)
		default:
			panic(fmt.Sprintf("Min.Check received non-numeric value: %v", val))
		}
	}()

	if !valid {
		return Violations{
			{
				Description: fmt.Sprintf("value %v is less than min value %d", val, m.Value),
			},
		}
	}

	return nil
}

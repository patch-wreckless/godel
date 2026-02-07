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

	lessThanViolation := func() Violations {
		return Violations{
			{
				Error: fmt.Errorf(
					"value %v is less than min value %d",
					val,
					m.Value),
			},
		}
	}

	switch kind {
	case reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uint:
		if rv.Uint() < uint64(m.Value) {
			return lessThanViolation()
		}
	case reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Int:
		if rv.Int() < m.Value {
			return lessThanViolation()
		}
	case reflect.Float32, reflect.Float64:
		if rv.Float() < float64(m.Value) {
			return lessThanViolation()
		}
	default:
		return Violations{{Error: Inapplicable{Constraint: m, Value: val}}}
	}

	return nil
}

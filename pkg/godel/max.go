package godel

import (
	"fmt"
	"reflect"
)

var _ Constraint = Max{}

// Max is a [Constraint] requiring a value to be less than or equal to a given value.
type Max struct {

	// Value is the maximum value that will be considered valid.
	Value int64
}

func (m Max) Check(val any) Violations {
	rv := reflect.ValueOf(val)
	kind := rv.Kind()

	greaterThanViolation := func() Violations {
		return Violations{
			{
				Error: fmt.Errorf(
					"value %v is greater than max value %d",
					val,
					m.Value),
			},
		}
	}

	switch kind {
	case reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uint:
		if rv.Uint() > uint64(m.Value) {
			return greaterThanViolation()
		}
	case reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Int:
		if rv.Int() > m.Value {
			return greaterThanViolation()
		}
	case reflect.Float32, reflect.Float64:
		if rv.Float() > float64(m.Value) {
			return greaterThanViolation()
		}
	default:
		return Violations{{Error: Inapplicable{Constraint: m, Value: val}}}
	}

	return nil
}

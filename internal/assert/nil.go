package assert

import (
	"reflect"
	"slices"
)

// Nil reports an error when the given value is not nil.
func Nil(t TestingT, actual any) {
	if t, ok := t.(interface{ Helper() }); ok {
		t.Helper()
	}

	if isNil(actual) {
		return
	}

	t.Errorf("expected nil; got %v", actual)
}

// NotNil reports an error when the given value is nil.
func NotNil(t TestingT, actual any) {
	if t, ok := t.(interface{ Helper() }); ok {
		t.Helper()
	}

	if !isNil(actual) {
		return
	}

	t.Error("expected non-nil")
}

func isNil(actual any) bool {

	if actual == nil {
		return true
	}

	val := reflect.ValueOf(actual)
	kind := val.Kind()

	if slices.Contains(nonNillableKinds, kind) {
		return false
	}

	if kind == reflect.Uintptr {
		return val.IsZero()
	}

	return val.IsNil()
}

var nonNillableKinds []reflect.Kind = []reflect.Kind{
	reflect.Bool,
	reflect.Int,
	reflect.Int8,
	reflect.Int16,
	reflect.Int32,
	reflect.Int64,
	reflect.Uint,
	reflect.Uint8,
	reflect.Uint16,
	reflect.Uint32,
	reflect.Uint64,
	reflect.Float32,
	reflect.Float64,
	reflect.Complex64,
	reflect.Complex128,
	reflect.Array,
	reflect.String,
	reflect.Struct,
}

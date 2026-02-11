package fields

import (
	"math"
	"reflect"
	"strconv"
)

// An IndexExpr is a [PathSegment] referencing an item in an ordered list.
type IndexExpr uint64

// String formats the [IndexExpr] into the bracket-notation used for indexing.
func (i IndexExpr) String() string {
	return "[" + strconv.FormatUint(uint64(i), 10) + "]"
}

// Access applies the [IndexExpr] to the target value as an index access
// expression and returns the accessed value and true if it was applicable,
// otherwise false.
func (i IndexExpr) Access(target any) (any, bool) {

	valType := reflect.TypeOf(target)
	for valType.Kind() == reflect.Pointer {
		valType = valType.Elem()
	}

	switch valType.Kind() {
	case reflect.Slice:
		// pass
	case reflect.Array:
		// Array lengths are part of their type, so an index expression that
		// would go out of bounds is inapplicable.
		if i >= IndexExpr(valType.Len()) {
			return nil, false
		}
	default:
		return nil, false
	}

	val := reflect.ValueOf(target)
	for val.Kind() == reflect.Pointer {
		val = val.Elem()
	}
	if val.Kind() == reflect.Invalid {
		// The field access was valid for the value type, but there is no value
		// to get a field from.
		return nil, true
	}

	// Slice lengths are not part of their type, so an index expression that
	// would go out out of bounds is applicable, but finds nothing.
	if val.Kind() == reflect.Slice && i >= IndexExpr(val.Len()) {
		return nil, true
	}

	// IndexExpr is uint64 so it may be out of range of int.
	if i > math.MaxInt {
		return nil, true
	}

	return val.Index(int(i)).Interface(), true
}

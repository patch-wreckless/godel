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

// Access evaluates i as an index access expression on target. If the index
// access is applicable to target then Access returns the item at the index and
// true, otherwise it returns nil and false.
//
// An index access is considered applicable if target’s type isa slice, an array
// where the index is in range, or a type defined as an applicable type.
// Applicability is defined recursively over pointer types: if T is applicable,
// then *T, **T, etc. are also applicable.
//
// If the index access is applicable but evaluation cannot proceed because
// some level of pointer indirection is nil, Access returns nil and true.
//
// If the index is applicable and every level of pointer indirection contains a
// non-nil value, Access returns the item value and true. If the item value
// itself is a nil pointer, the returned interface value is non-nil and holds a
// typed nil pointer.
func (i IndexExpr) Access(target any) (any, bool) {

	valType := reflect.TypeOf(target)
	for valType.Kind() == reflect.Pointer {
		valType = valType.Elem()
	}

	switch valType.Kind() {
	case reflect.Slice:
		// applicable
	case reflect.Array:
		// Array lengths are part of their type, so an index expression that
		// would go out of bounds is inapplicable.
		if i >= IndexExpr(valType.Len()) {
			return nil, false
		}
		// applicable
	default:
		return nil, false
	}

	val := reflect.ValueOf(target)
	for val.Kind() == reflect.Pointer {
		val = val.Elem()
	}
	if val.Kind() == reflect.Invalid {
		// The index access was valid for the value type, but there is no value
		// to get an item from.
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

package fields

import (
	"fmt"
	"reflect"
)

// AllItemsToken is a special index expression referring to every item in an
// ordered list.
const AllItemsToken string = "[*]"

// AllItemsExpr is a [PathSegment] representing the [AllItemsToken].
type AllItemsExpr struct{}

// String returns the [AllItemsToken].
func (AllItemsExpr) String() string {
	return AllItemsToken
}

// Access accesses the items at every index of target. If the all-items access
// is applicable to target then Access returns the item values in index order
// and true, otherwise it returns nil and false.
//
// An all-items access is considered applicable if an [IndexExpr] would be
// applicable. For arrays, an all-items access is never out of range.
//
// If the all-items access is applicable but evaluation cannot proceed because
// some level of pointer indirection is nil, Access returns nil and true. Note
// that a nil slice does not fall into this category because it's not a pointer,
// it contains the same set of items as a non-nil empty slice.
//
// If the all-items is applicable and every level of pointer indirection
// contains a non-nil value, Access returns a non-nil slice of any where each
// item is equivalent to what an [IndexExpr] on the corresponding index would
// have returned.
func (AllItemsExpr) Access(target any) ([]any, bool) {

	valType := reflect.TypeOf(target)
	for valType.Kind() == reflect.Pointer {
		valType = valType.Elem()
	}

	switch valType.Kind() {
	case reflect.Slice, reflect.Array:
		// applicable
	default:
		return nil, false
	}

	val := reflect.ValueOf(target)
	for val.Kind() == reflect.Pointer {
		val = val.Elem()
	}
	if val.Kind() == reflect.Invalid {
		// The all-items access was valid for the value type, but there is no
		// value to get items from.
		return nil, true
	}

	results := []any{}
	for i := range val.Len() {
		result, ok := IndexExpr(i).Access(val.Interface())
		if !ok {
			panic(fmt.Sprintf(
				"AllItemsExpr: IndexExpr(%d).Access failed on %T with length %d",
				i,
				val.Interface(),
				val.Len()))
		}
		results = append(results, result)
	}

	return results, true
}

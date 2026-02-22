package fields

import (
	"fmt"
	"reflect"
	"regexp"
)

// InvalidFieldName is the [error] returned when a [FieldName] with an invalid
// value is initialized or used.
type InvalidFieldName struct {

	// The Token containing the invalid field name.
	Token string
}

// Error implements [error].
func (i InvalidFieldName) Error() string {
	return fmt.Sprintf("invalid field name %q", i.Token)
}

// A FieldName is a [PathSegment] referencing a named field in structured data.
type FieldName struct {
	name  string
	valid bool
}

// NewFieldName initializes a [FieldName] with the given value, or returns an
// [error] if the given value is invalid.
func NewFieldName(name string) (FieldName, error) {
	if !validFieldName(name) {
		return FieldName{}, InvalidFieldName{Token: name}
	}
	return FieldName{
		name:  name,
		valid: true,
	}, nil
}

// MustFieldName initializes a [FieldName] with the given value, or panics if
// the given value is invalid.
func MustFieldName(name string) FieldName {
	fieldName, err := NewFieldName(name)
	if err != nil {
		panic(err.Error())
	}
	return fieldName
}

// String formats the [FieldName] in the dot-notation used for field access.
func (f FieldName) String() string {
	return "." + string(f.name)
}

// Access evaluates f as a field access expression on target. If the field
// access is applicable to target then Access returns the field value and
// true, otherwise it returns nil and false.
//
// A field access is considered applicable if target’s type is a struct with an
// exported field named f. Applicability is defined recursively over pointer
// types: if T is applicable, then *T, **T, etc. are also applicable.
//
// If the field access is applicable but evaluation cannot proceed because
// some level of pointer indirection is nil, Access returns nil and true.
//
// If the field is applicable and every level of pointer indirection contains a
// non-nil value, Access returns the field value and true. If the field value
// itself is a nil pointer, the returned interface value is non-nil and holds a
// typed nil pointer.
func (f FieldName) Access(target any) (any, bool) {

	valType := reflect.TypeOf(target)
	for valType.Kind() == reflect.Pointer {
		valType = valType.Elem()
	}

	if valType.Kind() != reflect.Struct {
		return nil, false
	}

	field, ok := valType.FieldByName(f.name)
	if !ok {
		return nil, false
	}

	if !field.IsExported() {
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

	return val.Field(field.Index[0]).Interface(), true
}

var validFieldNamePattern = regexp.MustCompile(`^[a-zA-Z_][a-zA-Z0-9_]*$`)

func validFieldName(name string) bool {
	return validFieldNamePattern.MatchString(name)
}

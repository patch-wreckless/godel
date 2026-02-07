package fields

import (
	"fmt"
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

var validFieldNamePattern = regexp.MustCompile(`^[a-zA-Z_][a-zA-Z0-9_]*$`)

func validFieldName(name string) bool {
	return validFieldNamePattern.MatchString(name)
}

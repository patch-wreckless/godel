package fields

import (
	"fmt"
	"regexp"
	"strconv"
)

var validIdentifierPattern = regexp.MustCompile(`^[a-zA-Z_][a-zA-Z0-9_]*$`)

// An Identifier represents a single component of a field access path, which
// can be either a named field or an array index.
type Identifier struct {

	// For named field identifiers, Name holds the field name. For array
	// indices, Name holds a string representation of the index (e.g., "[0]").
	Name string

	// For array indices, Index holds the numeric index value. For named field
	// identifiers, Index is nil.
	Index *int
}

// AsIndex returns the array index value and true for array index identifiers,
// or 0 and false for named field identifiers.
func (i Identifier) AsIndex() (int, bool) {
	if i.Index == nil {
		return 0, false
	}
	return *i.Index, true
}

// ParseIdentifier parses a string into an Identifier. It supports both named
// field identifiers (e.g., "foo") and array index identifiers (e.g., "[0]"),
// and returns an error if the input string is not a valid identifier.
func ParseIdentifier(s string) (Identifier, error) {
	if len(s) == 0 {
		return Identifier{}, fmt.Errorf("invalid empty identifier")
	}

	if s[0] == '[' && s[len(s)-1] == ']' {
		indexStr := s[1 : len(s)-1]
		index, err := strconv.Atoi(indexStr)
		if err != nil {
			return Identifier{}, fmt.Errorf("invalid array index %q", indexStr)
		}
		if index < 0 {
			return Identifier{}, fmt.Errorf("invalid negative array index: %d", index)
		}
		return Identifier{Name: s, Index: &index}, nil
	}

	if !validIdentifierPattern.MatchString(s) {
		return Identifier{}, fmt.Errorf("invalid identifier %q", s)
	}

	return Identifier{Name: s, Index: nil}, nil
}

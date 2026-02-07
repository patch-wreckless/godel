package fields

import (
	"fmt"
	"regexp"
	"strconv"
)

// Ensure Identifier is comparable.
var _ = Identifier{} == Identifier{}

// AllItemsToken is the string representation for the [AllItems] [Identifier].
const AllItemsToken = "[*]"

var validIdentifierPattern = regexp.MustCompile(`^[a-zA-Z_][a-zA-Z0-9_]*$`)

// An Identifier represents a single component of a field access path.
type Identifier struct {
	val any
}

// An IdentifierType represents a type of [Identifier].
type IdentifierType int

const (

	// IdentifierTypeFieldName is the [IdentifierType] for [FieldName].
	IdentifierTypeFieldName IdentifierType = iota + 1

	// IdentifierTypeIndex is the [IdentifierType] for [Index].
	IdentifierTypeIndex

	// [IdentifierTypeAllItems] is the [IdentifierType] for the [AllItems].
	IdentifierTypeAllItems IdentifierType = 4
)

// A FieldName is an [Identifier] for a named fields in some structured data.
type FieldName string

// An Index is an [Identifier] for an item in an order collection.
type Index uint64

// AllItems is a special [Identifier] representing every item in a collection.
type AllItems struct{}

// Type returns the [IdentifierType] of the [Identifier].
func (i Identifier) Type() IdentifierType {
	switch i.val.(type) {
	case FieldName:
		return IdentifierTypeFieldName
	case Index:
		return IdentifierTypeIndex
	case AllItems:
		return IdentifierTypeAllItems
	default:
		panic(fmt.Sprintf("unsupported identifier type %T", i.val))
	}
}

// FieldName returns the [FieldName] value of the [Identifier], or false if it
// is not a [FieldName].
func (i Identifier) FieldName() (FieldName, bool) {
	fieldName, ok := i.val.(FieldName)
	return fieldName, ok
}

// Index returns the [Index] value of the [Identifier], or false if it is not
// an [Index].
func (i Identifier) Index() (Index, bool) {
	index, ok := i.val.(Index)
	return index, ok
}

// AllItems returns the [AllItems] value of the [Identifier], or false if it is
// not [AllItems].
func (i Identifier) AllItems() (AllItems, bool) {
	allItems, ok := i.val.(AllItems)
	return allItems, ok
}

// NewIdentifier returns the given identifier value as an [Identifier].
func NewIdentifier[T FieldName | Index | AllItems](val T) Identifier {
	return Identifier{val: val}
}

// ParseIdentifier parses a string into an Identifier. It supports both named
// field identifiers (e.g., "foo") and array index identifiers (e.g., "[0]"),
// and returns an error if the input string is not a valid identifier.
func ParseIdentifier(str string) (Identifier, error) {
	if len(str) == 0 {
		return Identifier{}, fmt.Errorf("invalid empty identifier")
	}

	if str == AllItemsToken {
		return NewIdentifier(AllItems{}), nil
	}

	if str[0] == '[' && str[len(str)-1] == ']' {
		indexStr := str[1 : len(str)-1]
		index, err := strconv.Atoi(indexStr)
		if err != nil {
			return Identifier{}, fmt.Errorf("invalid array index %q", indexStr)
		}
		if index < 0 {
			return Identifier{}, fmt.Errorf("invalid negative array index: %d", index)
		}
		return NewIdentifier(Index(index)), nil
	}

	if !validIdentifierPattern.MatchString(str) {
		return Identifier{}, fmt.Errorf("invalid identifier %q", str)
	}

	return NewIdentifier(FieldName(str)), nil
}

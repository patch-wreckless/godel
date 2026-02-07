package fields

// An PathSegment represents a component of a field access path.
type PathSegment struct {
	val PathSegmentValue
}

// PathSegmentValue is the interface a value must implement to be the
// underlying value for a [PathSegment].
type PathSegmentValue interface {
	String() string
}

// PathSegmentValueTypes constrains the types a [PathSegment] can hold to a
// known set of types to approximate the guarantees of an enum type and
// exhaustive matching.
//
// This interface also helps ensure that the types in the list remain
// comparable by preventing them from being used if they're not.
type PathSegmentValueTypes interface {
	FieldName | IndexExpr | AllItemsExpr
	comparable

	PathSegmentValue
}

// NewPathSegment wraps the given value as a [PathSegment].
func NewPathSegment[T PathSegmentValueTypes](val T) PathSegment {
	return PathSegment{val: val}
}

// Value returns the underlying [PathSegmentValueTypes] from the [PathSegment].
func (p PathSegment) Value() any {
	return p.val
}

// String returns the underlying [PathSegmentValue] formated as a component of
// a field access path.
func (p PathSegment) String() string {
	return p.val.String()
}

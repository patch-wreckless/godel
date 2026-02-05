package godel

// A Constraint is a criteria that a value needs to meet in order to be valid.
type Constraint interface {

	// Path returns a series of [FieldName] values representing where the [Constraint] will
	// be applied.
	Path() []FieldName

	// Check evaluates the given
	Check(val any) Violations
}

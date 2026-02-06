package godel

// A Constraint is a criteria that a value needs to meet in order to be valid.
type Constraint interface {

	// Check evaluates the given
	Check(val any) Violations
}

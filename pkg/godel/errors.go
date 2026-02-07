package godel

import (
	"fmt"
)

// Inapplicable is the [error] returned when a [Constraint] is evaluated on a
// value it's not applicable to.
type Inapplicable struct {

	// The [Constraint] was not applicable to the value.
	Constraint Constraint

	// The value the [Constraint] was not applicable to.
	Value any
}

// Error implements [error].
func (err Inapplicable) Error() string {
	return fmt.Sprintf(
		"constraint '%v' is inapplicable to value '%v'",
		err.Constraint,
		err.Value)
}

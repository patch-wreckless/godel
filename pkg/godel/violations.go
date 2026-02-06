package godel

import (
	"fmt"
)

// A Violation represents an instance of a value not satisfying a [Constraint].
type Violation struct {

	// Description explains why the value did not satisfy the [Constraint].
	Description string
}

// Violations are a collection of [Violation] values.
type Violations []Violation

// Err returns a ViolationsError if the target [Violations] has one or more values,
// or nil if it's empty.
func (v Violations) Err() error {
	if len(v) == 0 {
		return nil
	}
	return ViolationsError(v)
}

// A ViolationsError is a [Violations] value that implements [error].
//
// Using a separate type prevents a function from returning a nil [Violations] as an
// [error] and inadvertendly creating a non-nil interface value with a nil underlying
// value (see https://go.dev/tour/methods/12).
type ViolationsError Violations

// Error returns a string representation of a [ViolationsError] to implement [error].
func (err ViolationsError) Error() string {
	return fmt.Sprintf("%d violation(s)", len(err))
}

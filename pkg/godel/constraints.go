package godel

var _ Constraint = Constraints{}

// Constraints is a group of [Constraint] values that can be evaluated together.
type Constraints []Constraint

// Check implements [Constraint] by evaluating all contained constraints and
// aggregating any violations.
func (c Constraints) Check(val any) Violations {

	var allViolations Violations

	for _, constraint := range c {
		violations := constraint.Check(val)
		if violations != nil {
			allViolations = append(allViolations, violations...)
		}
	}

	return allViolations
}

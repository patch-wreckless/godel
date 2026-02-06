package godel

type mockConstraint struct {
	violations Violations
}

func (m mockConstraint) Check(_ any) Violations {
	return m.violations
}

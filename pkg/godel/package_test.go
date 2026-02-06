package godel

import (
	"slices"
	"testing"
)

type mockConstraint struct {
	violations Violations
}

func (m mockConstraint) Check(_ any) Violations {
	return m.violations
}

func compareViolations(t *testing.T, expected, actual Violations) {
	if len(expected) != len(actual) {
		t.Errorf("expected %d violations; got %d: %v", len(expected), len(actual), actual)
	}
	for _, e := range expected {
		if !slices.Contains(actual, e) {
			t.Errorf("missing expected violation: %v", e)
		}
	}
	for _, e := range actual {
		if !slices.Contains(expected, e) {
			t.Errorf("unexpected violation: %v", e)
		}
	}
}

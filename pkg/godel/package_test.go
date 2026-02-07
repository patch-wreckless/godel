package godel

import (
	"errors"
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

	if expected, actual := len(expected), len(actual); actual != expected {
		t.Errorf("expected %d violation(s); got %d", expected, actual)
	}

	for _, expected := range expected {
		contains := slices.ContainsFunc(actual, func(actual Violation) bool {
			return violationEq(expected, actual)
		})
		if !contains {
			t.Errorf("missing expected violation: %v", expected)
		}
	}

	for _, actual := range actual {
		contains := slices.ContainsFunc(expected, func(expected Violation) bool {
			return violationEq(actual, expected)
		})
		if !contains {
			t.Errorf("got unexpected violation: %v", actual)
		}
	}
}

func violationEq(a, b Violation) bool {

	if !errors.Is(a.Error, b.Error) {
		return false
	}

	pathA, pathB := a.Path.Segments(), b.Path.Segments()
	if len(pathA) != len(pathB) {
		return false
	}
	for i, v := range pathA {
		if pathB[i] != v {
			return false
		}
	}

	return true
}

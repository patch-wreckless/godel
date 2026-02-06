package godel

import (
	"slices"
	"testing"
)

func TestConstraints(t *testing.T) {

	t.Run("#Check", func(t *testing.T) {

		t.Run("constraints is nil/returns nil", func(t *testing.T) {
			var underTest Constraints
			violations := underTest.Check(5)
			if violations != nil {
				t.Fatalf("expected no violations; got %v", violations)
			}
		})

		t.Run("constraints is empty/returns nil", func(t *testing.T) {
			underTest := Constraints{}
			violations := underTest.Check(5)
			if violations != nil {
				t.Fatalf("expected no violations; got %v", violations)
			}
		})

		compareViolations := func(t *testing.T, expected, actual Violations) {
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

		t.Run("single constraint returns no violation/returns nil", func(t *testing.T) {
			underTest := Constraints{
				mockConstraint{violations: nil},
			}
			violations := underTest.Check(5)
			if violations != nil {
				t.Fatalf("expected no violations; got %v", violations)
			}
		})

		t.Run("single constraint returns single violation/returns the violation", func(t *testing.T) {
			expected := Violations{
				Violation{Description: "constraint violated"},
			}
			underTest := Constraints{
				mockConstraint{violations: slices.Clone(expected)},
			}
			actual := underTest.Check(nil)
			compareViolations(t, expected, actual)
		})

		t.Run("single constraint returns multiple violation/returns all violations", func(t *testing.T) {
			expected := Violations{
				Violation{Description: "constraint violated"},
				Violation{Description: "another constraint violated"},
			}
			underTest := Constraints{
				mockConstraint{violations: slices.Clone(expected)},
			}
			actual := underTest.Check(nil)
			compareViolations(t, expected, actual)
		})

		t.Run("multiple constraints return no violation/returns nil", func(t *testing.T) {
			underTest := Constraints{
				mockConstraint{violations: nil},
				mockConstraint{violations: nil},
			}
			violations := underTest.Check(5)
			if violations != nil {
				t.Fatalf("expected no violations; got %v", violations)
			}
		})

		t.Run("multiple constraints return violations/returns all violations", func(t *testing.T) {
			expected := Violations{
				Violation{Description: "constraint violated"},
				Violation{Description: "another constraint violated"},
				Violation{Description: "yet another constraint violated"},
			}
			underTest := Constraints{
				mockConstraint{violations: slices.Clone(expected[:2])},
				mockConstraint{violations: slices.Clone(expected[2:])},
			}
			actual := underTest.Check(nil)
			compareViolations(t, expected, actual)
		})
	})
}

package godel

import (
	"errors"
	"testing"
)

func TestViolations(t *testing.T) {

	t.Run("#Err", func(t *testing.T) {

		t.Run("Violations is nil/returns nil", func(t *testing.T) {
			var underTest Violations
			if err := underTest.Err(); err != nil {
				t.Fatalf("expected nil; got %v", err)
			}
		})

		t.Run("Violations is empty/returns nil", func(t *testing.T) {
			var underTest Violations
			if err := underTest.Err(); err != nil {
				t.Fatalf("expected nil; got %v", err)
			}
		})

		t.Run("Violations is non-empty/returns ViolationsError", func(t *testing.T) {
			underTest := Violations{
				{
					Error: errors.New("some violation"),
				},
			}
			err := underTest.Err()
			if _, ok := err.(ViolationsError); !ok {
				t.Fatalf("expected %T; got %T", ViolationsError{}, err)
			}
		})
	})
}

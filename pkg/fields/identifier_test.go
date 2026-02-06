package fields

import (
	"testing"
)

func TestIdentifier(t *testing.T) {

	t.Run("#AsIndex", func(t *testing.T) {

		t.Run("Index is nil/returns _, false", func(t *testing.T) {
			identifier := Identifier{
				Index: nil,
			}

			if _, ok := identifier.AsIndex(); ok {
				t.Fatalf("expected ok to be false; got true")
			}
		})

		t.Run("Index is non-nil/returns value, true", func(t *testing.T) {
			expectedIndex := 5
			identifier := Identifier{
				Index: &expectedIndex,
			}

			actual, ok := identifier.AsIndex()
			if !ok {
				t.Fatalf("expected ok to be true; got false")
			}
			if actual != expectedIndex {
				t.Fatalf("expected index %d; got %d", expectedIndex, actual)
			}
		})
	})

	t.Run("#ParseIdentifier", func(t *testing.T) {

		t.Run("empty identifier/returns error", func(t *testing.T) {
			_, err := ParseIdentifier("")
			if err == nil {
				t.Fatalf("expected error; got nil")
			}
		})

		t.Run("field identifier with invalid characters/returns error", func(t *testing.T) {
			_, err := ParseIdentifier("foo-bar")
			if err == nil {
				t.Fatalf("expected error; got nil")
			}
		})

		t.Run("field identifier starts with digit/returns error", func(t *testing.T) {
			_, err := ParseIdentifier("1foo")
			if err == nil {
				t.Fatalf("expected error; got nil")
			}
		})

		t.Run("valid field identifier/returns Identifier", func(t *testing.T) {

			for _, tc := range []string{
				"foo",
				"fooBar",
				"FooBar",
				"foo_bar1",
				"_fooBar",
			} {

				t.Run(tc, func(t *testing.T) {
					expected := Identifier{
						Name:  tc,
						Index: nil,
					}
					actual, err := ParseIdentifier(tc)
					if err != nil {
						t.Fatalf("expected no error; got %v", err)
					}
					compareIdentifier(t, expected, actual)
				})
			}
		})

		t.Run("empty index expression/returns error", func(t *testing.T) {
			_, err := ParseIdentifier("[]")
			if err == nil {
				t.Fatalf("expected error; got nil")
			}
		})

		t.Run("non-numeric index expression/returns error", func(t *testing.T) {
			_, err := ParseIdentifier("[foo]")
			if err == nil {
				t.Fatalf("expected error; got nil")
			}
		})

		t.Run("negative numeric index expression/returns error", func(t *testing.T) {
			_, err := ParseIdentifier("[-1]")
			if err == nil {
				t.Fatalf("expected error; got nil")
			}
		})

		t.Run("valid numeric index expression/returns Identifier", func(t *testing.T) {
			expectedIndex := 3
			expected := Identifier{
				Name:  "[3]",
				Index: &expectedIndex,
			}
			actual, err := ParseIdentifier("[3]")
			if err != nil {
				t.Fatalf("expected no error; got %v", err)
			}
			compareIdentifier(t, expected, actual)
		})
	})
}

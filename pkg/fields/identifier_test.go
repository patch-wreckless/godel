package fields

import (
	"testing"
)

func TestIdentifier(t *testing.T) {

	t.Run("#ParseIdentifier", func(t *testing.T) {

		t.Run("empty identifier/returns error", func(t *testing.T) {
			_, err := ParseIdentifier("")
			if err == nil {
				t.Fatalf("expected error; got nil")
			}
		})

		t.Run("field name", func(t *testing.T) {

			t.Run("with invalid characters/returns error", func(t *testing.T) {
				_, err := ParseIdentifier("foo-bar")
				if err == nil {
					t.Fatalf("expected error; got nil")
				}
			})

			t.Run("with leading digit/returns error", func(t *testing.T) {
				_, err := ParseIdentifier("1foo")
				if err == nil {
					t.Fatalf("expected error; got nil")
				}
			})

			t.Run("is valid/returns identifier", func(t *testing.T) {

				testCases := []string{
					"foo",
					"fooBar",
					"FooBar",
					"foo_bar1",
					"_fooBar",
				}

				for _, tc := range testCases {
					t.Run(tc, func(t *testing.T) {
						_, err := ParseIdentifier(tc)
						if err != nil {
							t.Errorf("expected no error; got %v", err)
						}
					})
				}
			})
		})

		t.Run("index expression", func(t *testing.T) {

			t.Run("is empty/returns error", func(t *testing.T) {
				_, err := ParseIdentifier("[]")
				if err == nil {
					t.Fatalf("expected error; got nil")
				}
			})

			t.Run("is non-numeric/returns error", func(t *testing.T) {
				_, err := ParseIdentifier("[foo]")
				if err == nil {
					t.Fatalf("expected error; got nil")
				}
			})

			t.Run("is negative/returns error", func(t *testing.T) {
				_, err := ParseIdentifier("[-1]")
				if err == nil {
					t.Fatalf("expected error; got nil")
				}
			})

			t.Run("is valid/returns identifier", func(t *testing.T) {
				_, err := ParseIdentifier("[3]")
				if err != nil {
					t.Fatalf("expected no error; got %v", err)
				}
			})
		})

		t.Run("all items/returns identifier", func(t *testing.T) {
			_, err := ParseIdentifier("[*]")
			if err != nil {
				t.Fatalf("expected no error; got %v", err)
			}
		})
	})

	t.Run("#Type", func(t *testing.T) {

		testCases := []struct {
			name      string
			underTest Identifier
			expected  IdentifierType
		}{
			{
				name:      "field name",
				underTest: NewIdentifier(FieldName("foo")),
				expected:  IdentifierTypeFieldName,
			},
			{
				name:      "index",
				underTest: NewIdentifier(Index(47)),
				expected:  IdentifierTypeIndex,
			},
			{
				name:      "field name",
				underTest: NewIdentifier(AllItems{}),
				expected:  IdentifierTypeAllItems,
			},
		}

		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {
				actual := tc.underTest.Type()
				if actual != tc.expected {
					t.Fatalf(
						"expected IdentifierType(%d); got IdentifierType(%d)",
						tc.expected,
						actual)
				}
			})
		}
	})

	t.Run("#FieldName", func(t *testing.T) {

		testCases := []struct {
			name      string
			underTest Identifier
			expectOk  bool
			expected  FieldName
		}{
			{
				name:      "field name",
				underTest: NewIdentifier(FieldName("foo")),
				expectOk:  true,
				expected:  FieldName("foo"),
			},
			{
				name:      "index",
				underTest: NewIdentifier(Index(47)),
				expectOk:  false,
			},
			{
				name:      "field name",
				underTest: NewIdentifier(AllItems{}),
				expectOk:  false,
			},
		}

		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {
				actual, ok := tc.underTest.FieldName()
				if ok != tc.expectOk {
					t.Fatalf("expected ok to be %v; got %v", tc.expectOk, ok)
				}
				if actual != tc.expected {
					t.Fatalf("expected %v; got %v", tc.expected, actual)
				}
			})
		}
	})

	t.Run("#Index", func(t *testing.T) {

		testCases := []struct {
			name      string
			underTest Identifier
			expectOk  bool
			expected  Index
		}{
			{
				name:      "field name",
				underTest: NewIdentifier(FieldName("foo")),
				expectOk:  false,
			},
			{
				name:      "index",
				underTest: NewIdentifier(Index(47)),
				expectOk:  true,
				expected:  Index(47),
			},
			{
				name:      "field name",
				underTest: NewIdentifier(AllItems{}),
				expectOk:  false,
			},
		}

		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {
				actual, ok := tc.underTest.Index()
				if ok != tc.expectOk {
					t.Fatalf("expected ok to be %v; got %v", tc.expectOk, ok)
				}
				if actual != tc.expected {
					t.Fatalf("expected %v; got %v", tc.expected, actual)
				}
			})
		}
	})

	t.Run("#AllItems", func(t *testing.T) {

		testCases := []struct {
			name      string
			underTest Identifier
			expectOk  bool
		}{
			{
				name:      "field name",
				underTest: NewIdentifier(FieldName("foo")),
				expectOk:  false,
			},
			{
				name:      "index",
				underTest: NewIdentifier(Index(47)),
				expectOk:  false,
			},
			{
				name:      "field name",
				underTest: NewIdentifier(AllItems{}),
				expectOk:  true,
			},
		}

		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {
				_, ok := tc.underTest.AllItems()
				if ok != tc.expectOk {
					t.Fatalf("expected ok to be %v; got %v", tc.expectOk, ok)
				}
			})
		}
	})
}

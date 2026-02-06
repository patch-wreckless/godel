package fields

import (
	"testing"
)

func TestPath(t *testing.T) {

	t.Run("#Split", func(t *testing.T) {

		t.Run("path contains invalid identifier/panic", func(t *testing.T) {

			for _, tc := range []string{
				"foo-bar",
				"[-1]",
			} {

				defer func() {
					if r := recover(); r == nil {
						t.Fatalf("expected panic; got none")
					}
				}()
				underTest := Path{str: tc}
				_ = underTest.Split()
			}
		})

		t.Run("path contains empty identifier/panic", func(t *testing.T) {
			defer func() {
				if r := recover(); r == nil {
					t.Fatalf("expected panic; got none")
				}
			}()
			underTest := Path{str: "foo..bar"}
			_ = underTest.Split()
		})

		t.Run("path has leading empty identifier/panic", func(t *testing.T) {
			defer func() {
				if r := recover(); r == nil {
					t.Fatalf("expected panic; got none")
				}
			}()
			underTest := Path{str: ".foo"}
			_ = underTest.Split()
		})

		t.Run("path has trailing empty identifier/panic", func(t *testing.T) {
			defer func() {
				if r := recover(); r == nil {
					t.Fatalf("expected panic; got none")
				}
			}()
			underTest := Path{str: "foo."}
			_ = underTest.Split()
		})

		t.Run("empty path/returns empty", func(t *testing.T) {
			underTest := Path{str: ""}
			identifiers := underTest.Split()
			if len(identifiers) != 0 {
				t.Fatalf("expected empty identifiers; got %v", identifiers)
			}
		})

		t.Run(
			"path contains single field identifier/returns identifier",
			func(t *testing.T) {
				expected := []Identifier{{Name: "foo", Index: nil}}
				underTest := Path{str: expected[0].Name}
				actual := underTest.Split()
				compareIdentifiers(t, expected, actual)
			})

		t.Run(
			"path contains multiple field identifiers/returns identifiers in order",
			func(t *testing.T) {
				expected := []Identifier{
					{Name: "foo", Index: nil},
					{Name: "bar", Index: nil},
					{Name: "baz", Index: nil},
				}
				underTest := Path{str: "foo.bar.baz"}
				identifiers := underTest.Split()
				if len(identifiers) != len(expected) {
					t.Fatalf(
						"expected %d identifiers; got %d: %v",
						len(expected),
						len(identifiers),
						identifiers)
				}
				for i, exp := range expected {
					actual := identifiers[i]
					compareIdentifier(t, exp, actual)
				}
			})

		t.Run(
			"path contains single array index/returns identifier",
			func(t *testing.T) {
				expected := []Identifier{
					{
						Name: "[27]",
						Index: func() *int {
							x := 27
							return &x
						}(),
					},
				}
				underTest := Path{str: expected[0].Name}
				actual := underTest.Split()
				compareIdentifiers(t, expected, actual)
			})

		t.Run(
			"path contains multiple array indices/returns identifier in order",
			func(t *testing.T) {
				expected := []Identifier{
					{
						Name: "[27]",
						Index: func() *int {
							x := 27
							return &x
						}(),
					},
					{
						Name: "[42]",
						Index: func() *int {
							x := 42
							return &x
						}(),
					},
				}
				underTest := Path{str: "[27][42]"}
				actual := underTest.Split()
				compareIdentifiers(t, expected, actual)
			})

		t.Run(
			"path contains field then array index/returns identifier in order",
			func(t *testing.T) {
				expected := []Identifier{
					{Name: "foo"},
					{
						Name: "[42]",
						Index: func() *int {
							x := 42
							return &x
						}(),
					},
				}
				underTest := Path{str: "foo[42]"}
				actual := underTest.Split()
				compareIdentifiers(t, expected, actual)
			})

		t.Run(
			"path contains array index then field/returns identifier in order",
			func(t *testing.T) {
				expected := []Identifier{
					{
						Name: "[27]",
						Index: func() *int {
							x := 27
							return &x
						}(),
					},
					{Name: "bar"},
				}
				underTest := Path{str: "[27].bar"}
				actual := underTest.Split()
				compareIdentifiers(t, expected, actual)
			})
	})
}

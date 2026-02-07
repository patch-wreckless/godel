package fields

import (
	"testing"
)

func TestPath(t *testing.T) {

	t.Run("#Split", func(t *testing.T) {

		t.Run("path contains invalid identifier/panic", func(t *testing.T) {

			testCases := []string{
				"foo-bar",
				"[-1]",
			}

			for _, tc := range testCases {

				t.Run(tc, func(*testing.T) {
					defer func() {
						if r := recover(); r == nil {
							t.Fatalf("expected panic; got none")
						}
					}()
					underTest := Path{str: tc}
					_ = underTest.Split()
				})
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
			"path contains single field name/returns identifier",
			func(t *testing.T) {
				expected := []Identifier{{val: FieldName("foo")}}
				underTest := Path{str: "foo"}
				actual := underTest.Split()
				compareIdentifiers(t, expected, actual)
			})

		t.Run(
			"path contains multiple name identifiers/returns identifiers in order",
			func(t *testing.T) {
				expected := []Identifier{
					{val: FieldName("foo")},
					{val: FieldName("bar")},
					{val: FieldName("baz")},
				}
				underTest := Path{str: "foo.bar.baz"}
				actual := underTest.Split()
				compareIdentifiers(t, expected, actual)
			})

		t.Run(
			"path contains single array index/returns identifier",
			func(t *testing.T) {
				expected := []Identifier{{val: Index(27)}}
				underTest := Path{str: "[27]"}
				actual := underTest.Split()
				compareIdentifiers(t, expected, actual)
			})

		t.Run(
			"path contains multiple array indices/returns identifier in order",
			func(t *testing.T) {
				expected := []Identifier{
					{val: Index(27)},
					{val: Index(42)},
				}
				underTest := Path{str: "[27][42]"}
				actual := underTest.Split()
				compareIdentifiers(t, expected, actual)
			})

		t.Run(
			"path contains field then array index/returns identifier in order",
			func(t *testing.T) {
				expected := []Identifier{
					{val: FieldName("foo")},
					{val: Index(42)},
				}
				underTest := Path{str: "foo[42]"}
				actual := underTest.Split()
				compareIdentifiers(t, expected, actual)
			})

		t.Run(
			"path contains array index then field/returns identifier in order",
			func(t *testing.T) {
				expected := []Identifier{
					{val: Index(27)},
					{val: FieldName("bar")},
				}
				underTest := Path{str: "[27].bar"}
				actual := underTest.Split()
				compareIdentifiers(t, expected, actual)
			})
	})
}

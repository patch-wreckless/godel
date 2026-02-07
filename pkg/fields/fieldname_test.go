package fields

import (
	"errors"
	"fmt"
	"testing"
)

func TestInvalidFieldName(t *testing.T) {

	t.Run("errors.Is", func(t *testing.T) {

		t.Run("Tokens are equal/returns true", func(t *testing.T) {
			a := InvalidFieldName{Token: "foo"}
			b := InvalidFieldName{Token: "foo"}
			if !errors.Is(a, b) {
				t.Errorf("expected true; got false")
			}
		})

		t.Run("Tokens are not equal/returns false", func(t *testing.T) {
			a := InvalidFieldName{Token: "foo"}
			b := InvalidFieldName{Token: "bar"}
			if errors.Is(a, b) {
				t.Errorf("expected false; got true")
			}
		})
	})
}

func TestFieldName(t *testing.T) {

	invalidFieldNames := []string{
		"foo-bar",
		"1fooBar",
		"@fooBar",
	}

	validFieldNames := []string{
		"foo",
		"fooBar",
		"FooBar",
		"foo_bar1",
		"_fooBar",
	}

	t.Run("#NewFieldName", func(t *testing.T) {

		t.Run("invalid field name/returns InvalidFieldName",
			func(t *testing.T) {

				for _, tc := range invalidFieldNames {

					t.Run(tc, func(t *testing.T) {
						expected := InvalidFieldName{Token: tc}
						_, err := NewFieldName(tc)
						if err == nil {
							t.Fatalf("expected error; got nil")
						}
						var actual InvalidFieldName
						if !errors.As(err, &actual) {
							t.Fatalf("expected %T; got %T", expected, err)
						}
						if actual != expected {
							t.Fatalf("expected %+v; got %+v", expected, err)
						}
					})
				}
			})

		t.Run("valid field name/returns FieldName with no error",
			func(t *testing.T) {

				for _, tc := range validFieldNames {

					t.Run(tc, func(t *testing.T) {
						_, err := NewFieldName(tc)
						if err != nil {
							t.Fatalf("expected no error; got %q", err.Error())
						}
					})
				}
			})
	})

	t.Run("#MustFieldName", func(t *testing.T) {

		t.Run("invalid field name/panics",
			func(t *testing.T) {

				for _, tc := range invalidFieldNames {

					t.Run(tc, func(t *testing.T) {
						defer func() {
							if recover() == nil {
								t.Errorf("expected panic; got none")
							}
						}()
						_ = MustFieldName(tc)
					})
				}
			})

		t.Run("valid field name/returns FieldName with no panic",
			func(t *testing.T) {

				for _, tc := range validFieldNames {

					t.Run(tc, func(t *testing.T) {
						defer func() {
							if e := recover(); e != nil {
								t.Errorf("expected no panic; got %+v", e)
							}
						}()
						_ = MustFieldName(tc)
					})
				}
			})
	})

	t.Run("#String/returns dot-prefixed field name", func(t *testing.T) {
		expected := ".Foo"
		underTest := MustFieldName("Foo")
		actual := underTest.String()
		if actual != expected {
			t.Errorf("expected %q; got %q", expected, actual)
		}
	})
}

func BenchmarkFieldName(b *testing.B) {

	fieldName := MustFieldName("Foo")

	b.Run("#String", func(b *testing.B) {
		for range b.N {
			_ = fieldName.String()
		}
	})

	b.Run("#stringUsingFmtSprintf", func(b *testing.B) {
		for range b.N {
			_ = fieldName.stringUsingFmtSprintf()
		}
	})
}

func (i FieldName) stringUsingFmtSprintf() string {
	return fmt.Sprintf(".%s", i)
}

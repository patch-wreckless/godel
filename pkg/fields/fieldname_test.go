package fields

import (
	"errors"
	"fmt"
	"testing"

	"github.com/patch-wreckless/godel/internal/assert"
	"github.com/patch-wreckless/godel/internal/ptr"
)

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
						var actual InvalidFieldName
						if !errors.As(err, &actual) {
							t.Fatalf("expected %T; got %T", expected, err)
						}
						assert.Equal(t, expected, actual)
					})
				}
			})

		t.Run("valid field name/returns FieldName with no error",
			func(t *testing.T) {

				for _, tc := range validFieldNames {

					t.Run(tc, func(t *testing.T) {
						_, err := NewFieldName(tc)
						assert.Nil(t, err)
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
		assert.Equal(t, expected, actual)
	})

	t.Run("#Access", func(t *testing.T) {

		t.Run("target is not struct/returns not ok",
			func(t *testing.T) {

				testCases := []struct {
					name  string
					value any
				}{
					{name: "scalar", value: "string"},
					{name: "slice", value: []string{"slice"}},
					{name: "map", value: map[string]string{"key": "val"}},

					{name: "pointer to scalar", value: ptr.To("string")},
					{name: "pointer to slice", value: &[]string{"slice"}},
					{name: "pointer to map", value: &map[string]string{"key": "val"}},

					{name: "pointer to pointer to scalar", value: ptr.To(ptr.To("string"))},
				}

				for _, tc := range testCases {

					t.Run(tc.name, func(t *testing.T) {
						underTest := MustFieldName("Foo")
						_, ok := underTest.Access(tc.value)
						assert.NotOk(t, ok)
					})
				}
			})

		t.Run("target is struct with no matching field/returns not ok",
			func(t *testing.T) {
				target := struct{}{}
				underTest := MustFieldName("Foo")
				_, ok := underTest.Access(target)
				assert.NotOk(t, ok)
			})

		t.Run("target is struct with matching unexpported field/returns not ok",
			func(t *testing.T) {
				target := struct{ foo string }{
					foo: "unexpected",
				}
				underTest := MustFieldName("foo")
				_, ok := underTest.Access(target)
				assert.NotOk(t, ok)
			})

		t.Run("target is struct with matching exported field/returns field value",
			func(t *testing.T) {
				expected := "expected"
				target := struct{ Foo string }{
					Foo: expected,
				}
				underTest := MustFieldName("Foo")
				val, ok := underTest.Access(target)
				assert.Ok(t, ok)
				actual, ok := val.(string)
				assert.Ok(t, ok)
				assert.Equal(t, expected, actual)
			})

		t.Run("matching field is nil pointer/returns typed nil",
			func(t *testing.T) {
				target := struct{ Foo *string }{
					Foo: nil,
				}
				underTest := MustFieldName("Foo")
				val, ok := underTest.Access(target)
				assert.Ok(t, ok)
				actual, ok := val.(*string)
				assert.Ok(t, ok)
				assert.Nil(t, actual)
			})

		t.Run("matching field is non-nil pointer/returns pointer value",
			func(t *testing.T) {
				expected := ptr.To("expected")
				target := struct{ Foo *string }{
					Foo: expected,
				}
				underTest := MustFieldName("Foo")
				val, ok := underTest.Access(target)
				assert.Ok(t, ok)
				actual, ok := val.(*string)
				assert.Ok(t, ok)
				assert.Equal(t, expected, actual)
			})

		t.Run("target is nil pointer to struct/returns untyped nil ok",
			func(t *testing.T) {
				var target *struct{ Foo string }
				underTest := MustFieldName("Foo")
				val, ok := underTest.Access(target)
				assert.Ok(t, ok)
				assert.Nil(t, val)
			})

		t.Run("target is non-nil pointer to struct/returns field value",
			func(t *testing.T) {
				expected := "expected"
				target := &struct{ Foo string }{
					Foo: expected,
				}
				underTest := MustFieldName("Foo")
				val, ok := underTest.Access(target)
				assert.Ok(t, ok)
				actual, ok := val.(string)
				assert.Ok(t, ok)
				assert.Equal(t, expected, actual)
			})

		t.Run("target is nil pointer to pointer to struct/returns untyped nil ok",
			func(t *testing.T) {
				var target **struct{ Foo string }
				underTest := MustFieldName("Foo")
				val, ok := underTest.Access(target)
				assert.Ok(t, ok)
				assert.Nil(t, val)
			})

		t.Run("target is non-nil pointer to nil pointer to struct/returns untyped nil ok",
			func(t *testing.T) {
				target := ptr.To((*struct{ Foo string })(nil))
				underTest := MustFieldName("Foo")
				val, ok := underTest.Access(target)
				assert.Ok(t, ok)
				assert.Nil(t, val)
			})

		t.Run("target is non-nil pointer to non-nil pointer to struct/returns field value",
			func(t *testing.T) {
				expected := "expected"
				target := ptr.To(&struct{ Foo string }{
					Foo: expected,
				})
				underTest := MustFieldName("Foo")
				val, ok := underTest.Access(target)
				assert.Ok(t, ok)
				actual, ok := val.(string)
				assert.Ok(t, ok)
				assert.Equal(t, expected, actual)
			})
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

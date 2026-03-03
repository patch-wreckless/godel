package assert

import "fmt"

// LenEq reports an error to t when the the length of s is not expected.
func LenEq[T any, S ~[]T](
	t TestingT,
	expected int,
	s S,
	configure ...LenEqConfBuilder[T, S],
) {
	if t, ok := t.(interface{ Helper() }); ok {
		t.Helper()
	}
	conf := LenEqConf[T, S]{}
	for _, fn := range configure {
		fn(&conf)
	}

	if conf.expr != "" {
		Equal(t, expected, len(s), func(ec *EqualConf[int]) {
			ec.WithExpr(fmt.Sprintf("len(%s)", conf.expr))
		})
	} else {
		Equal(t, expected, len(s), func(ec *EqualConf[int]) {
			ec.WithExpr("len()")
		})
	}

}

// A LenEqConfBuilder configures an [LenEqConf].
type LenEqConfBuilder[T any, S ~[]T] func(*LenEqConf[T, S])

type LenEqConf[T any, S ~[]T] struct {
	expr string
}

func (c *LenEqConf[T, S]) WithExpr(expr string) *LenEqConf[T, S] {
	c.expr = expr
	return c
}

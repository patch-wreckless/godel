package assert

// Equal reports an error to t when the expected and actual values aren't equal.
func Equal[T comparable](
	t TestingT,
	expected, actual T,
	configure ...EqualConfBuilder[T],
) {
	if t, ok := t.(interface{ Helper() }); ok {
		t.Helper()
	}

	if actual == expected {
		return
	}

	conf := EqualConf[T]{}
	for _, fn := range configure {
		fn(&conf)
	}

	if conf.expr != "" {
		t.Errorf("expected %s to be %v; got %v", conf.expr, expected, actual)
	} else {
		t.Errorf("expected %v; got %v", expected, actual)
	}
}

// An EqualConfBuilder configures an [EqualConf].
type EqualConfBuilder[T comparable] func(*EqualConf[T])

type EqualConf[T comparable] struct {
	expr string
}

func (c *EqualConf[T]) WithExpr(expr string) *EqualConf[T] {
	c.expr = expr
	return c
}

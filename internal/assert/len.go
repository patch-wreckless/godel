package assert

// LenEq reports an error to t when the the length of s is not expected.
func LenEq[T any, S ~[]T](t TestingT, expected int, s S) {
	if t, ok := t.(interface{ Helper() }); ok {
		t.Helper()
	}

	Equal(t, expected, len(s), func(ec *EqualConf[int]) {
		ec.WithExpr("len()")
	})
}

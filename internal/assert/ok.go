package assert

// Ok reports an error if the given value is not true.
func Ok(t TestingT, ok bool) {
	if t, ok := t.(interface{ Helper() }); ok {
		t.Helper()
	}

	Equal(t, true, ok, func(ec *EqualConf[bool]) {
		ec.WithExpr("ok")
	})
}

// NotOk reports an error if the given value is true.
func NotOk(t TestingT, ok bool) {
	if t, ok := t.(interface{ Helper() }); ok {
		t.Helper()
	}

	Equal(t, false, ok, func(ec *EqualConf[bool]) {
		ec.WithExpr("ok")
	})
}

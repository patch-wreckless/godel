package assert

// TestingT is the subset of the [testing.T] interface the assert package
// depends on.
type TestingT interface {
	Error(args ...any)
	Errorf(format string, args ...any)
}

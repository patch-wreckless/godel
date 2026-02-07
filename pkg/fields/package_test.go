package fields

import (
	"testing"
)

func compareIdentifiers(t *testing.T, expected, actual []Identifier) {
	if expected, actual := len(expected), len(actual); actual != expected {
		t.Errorf("expected %d identifier(s); got %d", expected, actual)
	}

	for i := range min(len(expected), len(actual)) {
		if !identifierEq(expected[i], actual[i]) {
			t.Errorf(
				"expected Identifier[%d] to be %+v; got %+v",
				i,
				expected[i],
				actual[i])
		}
	}

	if len(expected) > len(actual) {
		for i := len(actual); i < len(expected); i++ {
			t.Errorf("missing expected identifier: %v", expected[i])
		}
	} else if len(actual) > len(expected) {
		for i := len(expected); i < len(actual); i++ {
			t.Errorf("unexpected identifier: %v", actual[i])
		}
	}
}

func identifierEq(a, b Identifier) bool {
	return a == b
}

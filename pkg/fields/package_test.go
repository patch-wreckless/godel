package fields

import (
	"testing"
)

func compareIdentifier(t *testing.T, expected, actual Identifier) {
	if !identiferEq(expected, actual) {
		t.Errorf("expected Identifier %+v; got %+v", expected, actual)
	}
}

func compareIdentifiers(t *testing.T, expected, actual []Identifier) {
	if len(expected) != len(actual) {
		t.Errorf("expected %d violations; got %d: %v", len(expected), len(actual), actual)
	}

	for i := range min(len(expected), len(actual)) {
		if !identiferEq(expected[i], actual[i]) {
			t.Errorf("expected Identifier[%d] to be %+v; got %+v", i, expected[i], actual[i])
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

func identiferEq(a, b Identifier) bool {
	if a.Name != b.Name {
		return false
	}

	nilOrVal := func(ptr *int) any {
		if ptr == nil {
			return nil
		}
		return *ptr
	}
	if nilOrVal(a.Index) != nilOrVal(b.Index) {
		return false
	}

	return true
}

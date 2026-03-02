package assert

import (
	"reflect"
	"testing"
)

func newMockTestingT(t *testing.T) *mockTestingT {
	m := &mockTestingT{
		t: t,
	}
	t.Cleanup(func() {
		for _, call := range m.expectedErrorCalls {
			t.Errorf("missing expected call to TestingT.Error: %v", call)
		}
		for _, call := range m.expectedErrorfCalls {
			t.Errorf("missing expected call to TestingT.Errorf: %v", call)
		}
	})
	return m
}

type mockTestingT struct {
	t *testing.T

	expectedErrorCalls []mockTestingTErrorCall

	expectedErrorfCalls []mockTestingTErrorfCall
}

type mockTestingTExpects struct {
	mock *mockTestingT
}

func (m *mockTestingT) Expect() mockTestingTExpects {
	return mockTestingTExpects{
		mock: m,
	}
}

type mockTestingTErrorCall struct {
	args []any
}

func (s mockTestingTExpects) Error(args ...any) {
	s.mock.expectedErrorCalls = append(
		s.mock.expectedErrorCalls,
		mockTestingTErrorCall{args: args})
}

func (m *mockTestingT) Error(args ...any) {
	actual := mockTestingTErrorCall{
		args: args,
	}
	if len(m.expectedErrorCalls) == 0 {
		m.t.Errorf("unexpected call to TestingT.Error: %v", actual)
		return
	}
	expected := m.expectedErrorCalls[0]
	m.expectedErrorCalls = m.expectedErrorCalls[1:]

	if !reflect.DeepEqual(expected, actual) {
		m.t.Errorf("TestingT.Error: expected %v; got %v", expected, actual)
	}
}

type mockTestingTErrorfCall struct {
	format string
	args   []any
}

func (s mockTestingTExpects) Errorf(format string, args ...any) {
	s.mock.expectedErrorfCalls = append(
		s.mock.expectedErrorfCalls,
		mockTestingTErrorfCall{format: format, args: args})
}

func (m *mockTestingT) Errorf(format string, args ...any) {
	actual := mockTestingTErrorfCall{
		format: format,
		args:   args,
	}
	if len(m.expectedErrorfCalls) == 0 {
		m.t.Errorf("unexpected call to TestingT.Errorf: %v", actual)
		return
	}
	expected := m.expectedErrorfCalls[0]
	m.expectedErrorfCalls = m.expectedErrorfCalls[1:]

	if !reflect.DeepEqual(expected, actual) {
		m.t.Errorf("TestingT.Errorf: expected %v; got %v", expected, actual)
	}
}

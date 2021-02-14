package tests

import "testing"

func Eq(t *testing.T, expected interface{}, actual interface{}) {
	if expected != actual {
		t.Errorf("expected \"%v\", but found \"%v\"", expected, actual)
	}
}

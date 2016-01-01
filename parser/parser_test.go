package parser

import (
	"testing"
)

func TestParseRType(t *testing.T) {
	testString := "add $t0, $t1, $t5"
	want := RType{"add", "$t0", "$t1", "$t5"}
	if got := parseRType(testString); got != want {
		t.Errorf("Not matching!!!\nInput: %s\nExpected: %+v\nGot: %+v\n", testString, want, got)
	}
}

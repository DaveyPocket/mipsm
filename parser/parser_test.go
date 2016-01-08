package parser

import (
	"testing"
)

//TODO Write more tests
func TestParse(t *testing.T) {
	testString := "add $t0, $t1, $t5"
	want := RType{"add", "$t0", "$t1", "$t5"}
	if got := Parse(testString); got != want {
		t.Errorf("Not matching!!!\nInput: %s\nExpected: %+v\nGot: %+v\n", testString, want, got)
	}
}

func TestParseRType(t *testing.T) {
	testString := "add $t0, $t1, $t5"
	want := RType{"add", "$t0", "$t1", "$t5"}
	if got := parseRType(testString); got != want {
		t.Errorf("Not matching!!!\nInput: %s\nExpected: %+v\nGot: %+v\n", testString, want, got)
	}
}

func TestParseIBranch(t *testing.T) {
	testString := "bne $s0, $s1, myRoutine"
	want := IType{"bne", "$s1", "$s0", "myRoutine"}
	if got := parseIBranch(testString); got != want {
		t.Errorf("Not matching!!!\nInput: %s\nExpected: %+v\nGot: %+v\n", testString, want, got)
	}
}

func TestParseIImmediate(t *testing.T) {
	testString := "addiu $t0, $t5, -50"
	want := IType{"addiu", "$t0", "$t5", "-50"}
	if got := parseIImmediate(testString); got != want {
		t.Errorf("Not matching!!!\nInput: %s\nExpected: %+v\nGot: %+v\n", testString, want, got)
	}
}

func TestParseIBaseOffset(t *testing.T) {
	testString := "lw $t0, -20($ra)"
	want := IType{"lw", "$t0", "$ra", "-20"}
	if got := parseIBaseOffset(testString); got != want {
		t.Errorf("Not matching!!!\nInput: %s\nExpected: %+v\nGot: %+v\n", testString, want, got)
	}

	testString = "sw $t0, 04($s2)"
	want = IType{"sw", "$s2", "$t0", "04"}
	if got := parseIBaseOffset(testString); got != want {
		t.Errorf("Not matching!!!\nInput: %s\nExpected: %+v\nGot: %+v\n", testString, want, got)
	}
}
func TestParseJType(t *testing.T) {
	testString := "j megaLoop"
	want := JType{"j", "megaLoop"}
	if got := parseJType(testString); got != want {
		t.Errorf("Not matching!!!\nInput: %s\nExpected: %+v\nGot: %+v\n", testString, want, got)
	}
}

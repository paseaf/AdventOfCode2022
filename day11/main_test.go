package main

import (
	"testing"
)

func Test_solvePuzzle1(t *testing.T) {
	expected := 10605
	actual := solvePuzzle1("./test_input.txt")
	if expected != actual {
		t.Errorf("incorrect result: expected %v, got %v", expected, actual)
	}
}
func Test_solvePuzzle2(t *testing.T) {
	expected := 2713310158
	actual := solvePuzzle2("./test_input.txt")
	if expected != actual {
		t.Errorf("incorrect result: expected %v, got %v", expected, actual)
	}
}

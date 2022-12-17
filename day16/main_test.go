package main

import (
	"testing"
)

func Test_solvePuzzle1(t *testing.T) {
	expected := 1651
	actual := solvePuzzle1("./test_input.txt")
	if expected != actual {
		t.Errorf("incorrect result: expected %v, got %v", expected, actual)
	}
}

package main

import (
	"testing"
)

func Test_solvePuzzle1(t *testing.T) {
	expected := 26
	actual := solvePuzzle1("./test_input.txt", 10)
	if expected != actual {
		t.Errorf("incorrect result: expected %v, got %v", expected, actual)
	}
}

func Test_solvePuzzle2(t *testing.T) {
	expected := 56000011
	actual := solvePuzzle2("./test_input.txt", 20)
	if expected != actual {
		t.Errorf("incorrect result: expected %v, got %v", expected, actual)
	}
}

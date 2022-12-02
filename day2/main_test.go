package main

import "testing"

func Test_calculateScore(t *testing.T) {
	expected := 15
	actual := calculateScore("./test_input.txt")
	if expected != actual {
		t.Errorf("incorrect result: expected %v, got %v", expected, actual)
	}
}

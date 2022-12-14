package main

import "testing"

func Test_findContaining(t *testing.T) {
	expected := 2
	actual := findContainingPairs("./test_input.txt")
	if expected != actual {
		t.Errorf("incorrect result: expected %v, got %v", expected, actual)
	}
}
func Test_findOverlapping(t *testing.T) {
	expected := 4
	actual := findOverlappingPairs("./test_input.txt")
	if expected != actual {
		t.Errorf("incorrect result: expected %v, got %v", expected, actual)
	}
}

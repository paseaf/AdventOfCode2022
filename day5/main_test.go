package main

import "testing"

func Test_findTopCreates(t *testing.T) {
	expected := "CMZ"
	actual := findTopCrates("./test_input.txt")
	if expected != actual {
		t.Errorf("incorrect result: expected %v, got %v", expected, actual)
	}
}

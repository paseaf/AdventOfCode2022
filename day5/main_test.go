package main

import "testing"

func Test_findTopCreates(t *testing.T) {
	expected := "CMZ"
	actual := findTopCrates("./test_input.txt")
	if expected != actual {
		t.Errorf("incorrect result: expected %v, got %v", expected, actual)
	}
}
func Test_findTopCreates2(t *testing.T) {
	expected := "MCD"
	actual := findTopCrates2("./test_input.txt")
	if expected != actual {
		t.Errorf("incorrect result: expected %v, got %v", expected, actual)
	}
}

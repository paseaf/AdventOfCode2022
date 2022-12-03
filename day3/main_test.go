package main

import "testing"

func Test_getPrioSum(t *testing.T) {
	expected := 157
	actual := getPrioSum1("./test_input.txt")
	if expected != actual {
		t.Errorf("incorrect result: expected %v, got %v", expected, actual)
	}
}
func Test_getPrioSum2(t *testing.T) {
	expected := 70
	actual := getPrioSum2("./test_input.txt")
	if expected != actual {
		t.Errorf("incorrect result: expected %v, got %v", expected, actual)
	}
}

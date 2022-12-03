package main

import "testing"

func Test_findMaxCalories(t *testing.T) {
	expected := 24000
	actual := findMaxCalories("./test_input.txt")
	if expected != actual {
		t.Errorf("incorrect result: expected %v, got %v", expected, actual)
	}
}
func Test_findTopThree(t *testing.T) {
	expected := 45000
	actual := findTopThree("./test_input.txt")
	if expected != actual {
		t.Errorf("incorrect result: expected %v, got %v", expected, actual)
	}
}

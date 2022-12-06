package main

import "testing"

func Test_findFirstMarker(t *testing.T) {
	expected := 10
	actual := findFirstMarker("./test_input.txt", 4)
	if expected != actual {
		t.Errorf("incorrect result: expected %v, got %v", expected, actual)
	}
}
func Test_findFirstMarker2(t *testing.T) {
	expected := 29
	actual := findFirstMarker("./test_input.txt", 14)
	if expected != actual {
		t.Errorf("incorrect result: expected %v, got %v", expected, actual)
	}
}

package main

import (
	"testing"
)

func Test_findTotalDirSize(t *testing.T) {
	expected := 95437
	actual := findTotalDirSize("./test_input.txt")
	if expected != actual {
		t.Errorf("incorrect result: expected %v, got %v", expected, actual)
	}
}

func Test_findMinDirSizeToFree(t *testing.T) {
	expected := 24933642
	actual := findMinDirSizeToFree("./test_input.txt")
	if expected != actual {
		t.Errorf("incorrect result: expected %v, got %v", expected, actual)
	}
}

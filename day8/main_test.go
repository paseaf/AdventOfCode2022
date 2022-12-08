package main

import (
	"testing"
)

func Test_findVisibleTrees(t *testing.T) {
	expected := 21
	actual := findVisibleTrees("./test_input.txt")
	if expected != actual {
		t.Errorf("incorrect result: expected %v, got %v", expected, actual)
	}
}

func Test_findHighestScenicScore(t *testing.T) {
	expected := 8
	actual := findHighestScenicScore("./test_input.txt")
	if expected != actual {
		t.Errorf("incorrect result: expected %v, got %v", expected, actual)
	}
}

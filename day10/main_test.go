package main

import (
	"strings"
	"testing"
)

func Test_solvePuzzle1(t *testing.T) {
	expected := 13140
	actual := solvePuzzle1("./test_input.txt")
	if expected != actual {
		t.Errorf("incorrect result: expected %v, got %v", expected, actual)
	}
}

func Test_solvePuzzle2(t *testing.T) {
	expectedStr := "##..##..##..##..##..##..##..##..##..##..\n###...###...###...###...###...###...###.\n####....####....####....####....####....\n#####.....#####.....#####.....#####.....\n######......######......######......####\n#######.......#######.......#######....."

	actual := solvePuzzle2("./test_input.txt")
	actualStr := ""
	for i, row := range actual {
		rowStr := strings.Join(row, "")
		actualStr += rowStr
		if i != len(actual)-1 {
			actualStr += "\n"
		}
	}

	if expectedStr != actualStr {
		t.Errorf("incorrect result: expected %v, got %v", expectedStr, actualStr)
	}
}

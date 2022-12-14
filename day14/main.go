package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func main() {
	result1 := solvePuzzle1(filepath.Join("./day14", "input.txt"))
	fmt.Printf("Result1 : %v\n", result1)
	result2 := solvePuzzle2(filepath.Join("./day14", "input.txt"))
	fmt.Printf("Result2 : %v\n", result2)
}

type Pos struct {
	row int
	col int
}

var ROWS = 200
var COLS = 1000
var SOURCE = Pos{0, 500}

func solvePuzzle1(path string) (result int) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Phase 1: prepare
	cave := make([][]bool, ROWS) // true: filled
	for i, _ := range cave {
		cave[i] = make([]bool, COLS)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " -> ")
		for i := 0; i < len(parts)-1; i++ {
			posA := parsePos(parts[i])
			posB := parsePos(parts[i+1])
			cave = mark(cave, posA, posB)
		}
	}

	// Phase 2: drop
	cnt := 0
	for true {
		goneForever := drop(&cave)
		if goneForever {
			return cnt
		}
		cnt++
	}
	return -1
}

func drop(cave *[][]bool) (goneForever bool) {
	cur := SOURCE
	lowerBound := 180
	for cur.row < lowerBound {
		if (*cave)[cur.row+1][cur.col] == false {
			cur.row++
		} else if (*cave)[cur.row+1][cur.col-1] == false {
			cur.row++
			cur.col--
		} else if (*cave)[cur.row+1][cur.col+1] == false {
			cur.row++
			cur.col++
		} else {
			(*cave)[cur.row][cur.col] = true
			return false
		}
	}
	return true
}

func parsePos(raw string) Pos {
	parts := strings.Split(raw, ",")
	col, _ := strconv.Atoi(parts[0])
	row, _ := strconv.Atoi(parts[1])
	return Pos{row, col}
}

func mark(cave [][]bool, start Pos, end Pos) [][]bool {
	if start.row == end.row {
		min, max := start.col, end.col
		if max < min {
			min, max = max, min
		}
		for col := min; col <= max; col++ {
			cave[start.row][col] = true
		}
	} else {
		min, max := start.row, end.row
		if max < min {
			min, max = max, min
		}
		for row := min; row <= max; row++ {
			cave[row][start.col] = true
		}
	}
	return cave
}

func solvePuzzle2(path string) (result int) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Phase 1: prepare
	cave := make([][]bool, ROWS) // true: filled
	for i, _ := range cave {
		cave[i] = make([]bool, COLS)
	}

	scanner := bufio.NewScanner(file)
	lowerBound := 0
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " -> ")
		for i := 0; i < len(parts)-1; i++ {
			posA := parsePos(parts[i])
			posB := parsePos(parts[i+1])
			cave = mark(cave, posA, posB)
			if posA.row > lowerBound {
				lowerBound = posA.row
			}
			if posB.row > lowerBound {
				lowerBound = posB.row
			}
		}
	}
	lowerBound++

	// Phase 2: drop
	cnt := 0
	for true {
		blocked := drop2(&cave, lowerBound)
		cnt++
		if blocked {
			return cnt
		}
	}
	return -1
}

func drop2(cave *[][]bool, lowerBound int) (blocked bool) {
	cur := SOURCE
	for true {
		if cur.row == lowerBound {
			(*cave)[cur.row][cur.col] = true
			return false
		}
		if (*cave)[cur.row+1][cur.col] == false {
			cur.row++
		} else if (*cave)[cur.row+1][cur.col-1] == false {
			cur.row++
			cur.col--
		} else if (*cave)[cur.row+1][cur.col+1] == false {
			cur.row++
			cur.col++
		} else {
			(*cave)[cur.row][cur.col] = true
			return (*cave)[SOURCE.row][SOURCE.col]
		}
	}
	return true
}

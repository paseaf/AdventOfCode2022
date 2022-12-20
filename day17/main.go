package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {
	result1 := solvePuzzle1(filepath.Join("./day17", "input.txt"))
	fmt.Printf("Result1 : %v\n", result1)
	result2 := solvePuzzle2(filepath.Join("./day17", "input.txt"))
	fmt.Printf("Result2 : %v\n", result2)
}

type Cave struct {
	content []uint8
}

type Rock struct {
	content []uint8
	bottom  int
}

func newCave() (cave Cave) {
	cave = Cave{}
	cave.content = make([]uint8, 0)
	return cave
}

func newRock(shape int, bottom int) (rock Rock) {
	rock = Rock{bottom: bottom}

	switch shape {
	case 0:
		// |7..####0|
		var line1 uint8 = 16 + 8 + 4 + 2
		rock.content = []uint8{line1}
	case 1:
		// +
		var line1 uint8 = 8
		var line2 uint8 = 16 + 8 + 4
		var line3 uint8 = 8
		rock.content = []uint8{line1, line2, line3}
	case 2:
		// ..# -- line 3
		// ..#
		// ### -- line1
		var line1 uint8 = 4 + 8 + 16
		var line2 uint8 = 4
		var line3 uint8 = 4
		rock.content = []uint8{line1, line2, line3}
	case 3:
		// |
		var line1 uint8 = 16
		rock.content = []uint8{line1, line1, line1, line1}
	case 4:
		// o
		var line1 uint8 = 16 + 8
		var line2 uint8 = 16 + 8
		rock.content = []uint8{line1, line2}
	}
	return rock
}
func solvePuzzle1(path string) (result int) {
	return simulate(path, 2022)
}
func solvePuzzle2(path string) (result int) {
	return simulate(path, 1000000000000)
}
func simulate(path string, rockCnt int) (result int) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Phase 1: prepare
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	dirs := scanner.Text()

	cave := newCave()
	dirId := 0
	slice := 100000000
	for i := 0; i < rockCnt; i++ {

		if i%slice == 0 {
			fmt.Printf("Current progress: %v/10000\n", i/slice)
		}
		// shiftAndDrop rock
		shapeId := i % 5
		// pad three rows
		cave.pad(3)
		// gen rock
		rock := newRock(shapeId, len(cave.content))
		cave.pad(len(rock.content))
		curDir := dirs[dirId]
		for {
			cave.shift(&rock, curDir)
			dirId = (dirId + 1) % len(dirs)
			curDir = dirs[dirId]
			hasDropped := cave.drop(&rock, curDir)
			if !hasDropped {
				break
			}
		}

		cave.persist(rock)
		cave.removeEmptyTop()
	}

	return len(cave.content)
}

// persist rock as part of the cave
func (cave *Cave) persist(rock Rock) {
	for i, row := range rock.content {
		caveRowId := i + rock.bottom
		cave.content[caveRowId] += row
	}
}

// remove top empty lines
func (cave *Cave) removeEmptyTop() {
	top := len(cave.content) - 1
	for top >= 0 {
		if cave.content[top] != 0 {
			break
		}
		top--
	}
	cave.content = cave.content[0 : top+1]
}

func (cave *Cave) unshift(rock *Rock, dir byte) {
	if dir == '<' {
		cave.shift(rock, '>')
	} else {
		cave.shift(rock, '<')
	}
}
func (cave *Cave) shift(rock *Rock, dir byte) {
	after := make([]uint8, len(rock.content))
	// shift left/right
	canShift := true
	if dir == '<' {
		for rowIdx, row := range rock.content {
			// Step 1: Move
			// check if hits wall
			if row>>6 == 1 {
				canShift = false
				break
			}
			// check if collides
			row = row << 1
			after[rowIdx] = row
			caveRowId := rock.bottom + rowIdx
			caveRow := cave.content[caveRowId]
			// collides with cave
			if collides(row, caveRow) {
				canShift = false
				break
			}
		}
		if canShift {
			rock.content = after
		}
	} else {
		// shift right
		canMove := true
		for rowIdx, row := range rock.content {
			// Step 1: Move
			// check if hits wall
			if row%2 == 1 {
				canMove = false
				break
			}
			// check if collides
			row = row >> 1
			after[rowIdx] = row
			caveRowId := rock.bottom + rowIdx
			caveRow := cave.content[caveRowId]
			// collides with cave
			if collides(row, caveRow) {
				canMove = false
				break
			}
		}
		if canMove {
			rock.content = after
		}
	}
}

func (cave *Cave) drop(rock *Rock, dir byte) (hasDropped bool) {
	hasDropped = true
	for rowIdx, row := range rock.content {
		caveRowBelowId := rock.bottom + rowIdx - 1
		if caveRowBelowId < 0 {
			hasDropped = false
			break
		}
		caveRowBelow := cave.content[caveRowBelowId]
		if collides(row, caveRowBelow) {
			hasDropped = false
			break
		}
	}
	if hasDropped {
		rock.bottom--
	}
	return hasDropped
}

func collides(row1 uint8, row2 uint8) bool {
	return (row1 | row2) != (row1 + row2)
}

func (cave *Cave) pad(rows int) {
	padding := make([]uint8, rows)
	cave.content = append(cave.content, padding...)
}

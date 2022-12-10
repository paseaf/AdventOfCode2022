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
	result1 := solvePuzzle1(filepath.Join("./day10", "input.txt"))
	fmt.Printf("Result1 : %v\n", result1)

	result2 := solvePuzzle2(filepath.Join("./day10", "input.txt"))
	fmt.Println("Result2")
	for _, row := range result2 {
		fmt.Println(row)
	}
}

func solvePuzzle1(path string) (result int) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	curCycle := 1
	x := 1

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if isTargetCycle(curCycle) {
			result += curCycle * x
		}
		line := scanner.Text()
		parts := strings.Split(line, " ")
		if len(parts) == 1 {
			// noop
			curCycle++
		} else {
			curCycle++
			if isTargetCycle(curCycle) {
				result += curCycle * x
			}
			curCycle++
			operand, _ := strconv.Atoi(parts[1])
			x += operand
		}
	}

	return result
}

func isTargetCycle(cycle int) bool {
	targetCycles := []int{20, 60, 100, 140, 180, 220}
	for _, targetCycle := range targetCycles {
		if cycle == targetCycle {
			return true
		}
	}
	return false
}

func solvePuzzle2(path string) (canvas [][]string) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	x := 1
	hasBuffer := false
	operand := 0
	scanner := bufio.NewScanner(file)
	for i := 0; i < 6; i++ {
		row := make([]string, 40)
		canvas = append(canvas, row)
		for j := 0; j < 40; j++ {
			// draw crt
			if x-1 <= j && j <= x+1 {
				canvas[i][j] = "#"
			} else {
				canvas[i][j] = "."
			}
			// read input
			if !hasBuffer {
				scanner.Scan()
				line := scanner.Text()
				parts := strings.Split(line, " ")
				if len(parts) == 1 {
					// noop
				} else {
					hasBuffer = true
					operand, _ = strconv.Atoi(parts[1])
				}
			} else {
				// update sprite
				x += operand
				hasBuffer = false
			}
		}
	}

	return canvas
}

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

func main() {
	result := findMaxCalories(filepath.Join("day1", "input.txt"))
	fmt.Printf("Result is %v\n", result)
}

func findMaxCalories(path string) (result int) {
	// 1. read input
	file, err := os.Open(path)
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	curResult := 0
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			result = max(result, curResult)
			curResult = 0
		} else {
			curNum, err := strconv.Atoi(line)
			check(err)
			curResult += curNum
		}
	}
	return result
}

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

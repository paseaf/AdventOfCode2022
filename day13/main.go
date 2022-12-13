package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {
	result1 := solvePuzzle1(filepath.Join("./day13", "input.txt"))
	fmt.Printf("Result1 : %v\n", result1)
	result2 := solvePuzzle2(filepath.Join("./day13", "input.txt"))
	fmt.Printf("Result2 : %v\n", result2)
}

func solvePuzzle1(path string) (result int) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	idx := 1
	for scanner.Scan() {
		left := parseLine(scanner.Text())
		scanner.Scan()
		right := parseLine(scanner.Text())
		scanner.Scan()

		isInOrder := compare(left, right) < 0
		if isInOrder {
			result += idx
		}
		idx++
	}

	return result
}

func sort(items []any) []any {
	for i := 0; i < len(items)-1; i++ {
		min := i
		for j := i + 1; j < len(items); j++ {
			if compare(items[j], items[min]) < 0 {
				min = j
			}
		}
		items[i], items[min] = items[min], items[i]
	}
	return items
}

func solvePuzzle2(path string) (result int) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	items := make([]any, 0)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			continue
		}
		item := parseLine(line)
		items = append(items, item)
	}
	items = append(items, parseLine("[[2]]"))
	items = append(items, parseLine("[[6]]"))

	items = sort(items)
	var a, b int
	for i, item := range items {
		outerList, okO := item.([]any)
		if !okO || len(outerList) != 1 {
			continue
		}
		innerList, okI := outerList[0].([]any)
		if okI && len(innerList) == 1 {
			num, ok := innerList[0].(float64)
			if ok {
				if num > 1.5 && num < 2.5 {
					a = i + 1
				} else if num > 5.5 && num < 6.5 {
					b = i + 1
				}
			}
		}
	}
	return a * b
}

func parseLine(line string) any {
	var result any
	if err := json.Unmarshal([]byte(line), &result); err != nil {
		log.Fatal(err)
	}
	return result
}

// -1: left < right 0: same 1: left > right
func compare(left any, right any) int {
	numL, okL := left.(float64)
	numR, okR := right.(float64)
	if okL && okR {
		return int(numL) - int(numR)
	}

	leftList := parseList(left)
	rightList := parseList(right)

	for i := range leftList {
		if i >= len(rightList) {
			return 1
		}
		if res := compare(leftList[i], rightList[i]); res != 0 {
			return res
		}
	}

	if len(leftList) == len(rightList) {
		return 0
	}

	return -1
}

func parseList(item any) (result []any) {
	switch item.(type) {
	case []any:
		result = item.([]any)
	case float64:
		result = []any{item}
	}
	return result
}

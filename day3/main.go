package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {
	result1 := getPrioSum1(filepath.Join("./day3", "input.txt"))
	fmt.Printf("Result1 : %v\n", result1)

	result2 := getPrioSum2(filepath.Join("./day3", "input.txt"))
	fmt.Printf("Result2 : %v\n", result2)
}

func getPrioSum1(path string) (result int) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			continue
		}
		lineResult := getSackPrioSum(line)
		result += lineResult
	}
	return result
}

func getSackPrioSum(line string) int {
	midPos := len(line) / 2
	// left half
	leftHas := map[byte]bool{}
	for i := 0; i < midPos; i++ {
		char := line[i]
		leftHas[char] = true
	}

	// right half
	resultHas := map[byte]bool{}
	for i := midPos; i < len(line); i++ {
		char := line[i]
		if leftHas[char] {
			resultHas[char] = true
		}
	}
	// sum
	sum := 0
	for char, exists := range resultHas {
		if exists {
			sum += toPrio(char)
		}
	}
	return sum
}

func getPrioSum2(path string) int {
	// 1. get line count
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	size := 3

	result := 0
	for true {
		curResult := findGroupPriority(scanner, size)
		if curResult == -1 {
			return result
		}
		result += curResult
	}

	return result
}

func findGroupPriority(scanner *bufio.Scanner, size int) int {
	charCnt := map[rune]int{}
	for i := 0; i < size-1; i++ {
		scanner.Scan()
		line := scanner.Text()
		charSet := map[rune]bool{}
		for _, char := range line {
			charSet[char] = true
		}

		for char, _ := range charSet {
			charCnt[char]++
		}
	}
	// last line
	scanner.Scan()
	line := scanner.Text()
	for _, char := range line {
		if charCnt[char] == size-1 {
			return toPrio(byte(char))
		}
	}
	return -1
}

func toPrio(char byte) (priority int) {
	ascii := int(char)
	if ascii <= int('Z') {
		return ascii - int('A') + 27
	} else {
		return ascii - int('a') + 1
	}
}

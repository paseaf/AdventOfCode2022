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
	result1 := findContainingPairs(filepath.Join("./day4", "input.txt"))
	fmt.Printf("Result1 : %v\n", result1)

}

func findContainingPairs(path string) (result int) {
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
		if isContaining(line) {
			result++
		}
	}
	return result
}

func isContaining(line string) bool {
	groups := strings.Split(line, ",")
	ranges := [][]int{}
	for _, group := range groups {
		bounds := strings.Split(group, "-")
		lower, err := strconv.Atoi(bounds[0])
		check(err)
		upper, err := strconv.Atoi(bounds[1])
		check(err)
		ranges = append(ranges, []int{lower, upper})
	}

	leftLarger := ranges[0][0] <= ranges[1][0] && ranges[0][1] >= ranges[1][1]
	rightLarger := ranges[1][0] <= ranges[0][0] && ranges[1][1] >= ranges[0][1]

	return leftLarger || rightLarger
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {
	score := calculateScore(filepath.Join("./day2", "input.txt"))
	fmt.Printf("Total score: %v", score)
}

func calculateScore(path string) (score int) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > 0 {
			score += winScore[line]
		}
	}
	return score
}

var winScore = map[string]int{
	"C X": 1 + 6,
	"A X": 1 + 3,
	"B X": 1,
	"A Y": 2 + 6,
	"B Y": 2 + 3,
	"C Y": 2,
	"B Z": 3 + 6,
	"C Z": 3 + 3,
	"A Z": 3,
}

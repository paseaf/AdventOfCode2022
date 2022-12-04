package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {
	score1 := calculateScore(filepath.Join("./day2", "input.txt"))
	fmt.Printf("Total score1: %v\n", score1)

	score2 := calculateScore2(filepath.Join("./day2", "input.txt"))
	fmt.Printf("Total score2: %v\n", score2)
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

func calculateScore2(path string) (score int) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > 0 {
			score += winScorePart2[line]
		}
	}
	return score
}

var winScorePart2 = map[string]int{
	// lose
	"A X": 3 + 0,
	"B X": 1 + 0,
	"C X": 2 + 0,
	// draw
	"A Y": 1 + 3,
	"B Y": 2 + 3,
	"C Y": 3 + 3,
	// win
	"B Z": 3 + 6,
	"C Z": 1 + 6,
	"A Z": 2 + 6,
}

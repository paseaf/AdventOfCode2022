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
	result1 := findTopCrates(filepath.Join("./day5", "input.txt"))
	fmt.Printf("Result1 : %v\n", result1)
}

func findTopCrates(path string) (result string) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	crates := parseCrates(scanner)
	fmt.Println(crates)
	moves := parseMoves(scanner)
	result = moveCrates(crates, moves)
	return result
}

func parseCrates(scanner *bufio.Scanner) (crates []string) {
	var lines []string
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		} else {
			lines = append(lines, line)
		}
	}
	size := len(strings.Split(lines[len(lines)-1], "   "))
	crates = make([]string, size+1)
	lines = lines[0 : len(lines)-1]

	for _, line := range lines {
		for i := 1; i <= size; i++ {
			charId := (i-1)*4 + 1
			if charId >= len(line) {
				break
			}
			if line[charId] != ' ' {
				crates[i] = string(line[charId]) + crates[i]
			}
		}
	}

	return crates
}

// moves: [from, to, cnt]
func parseMoves(scanner *bufio.Scanner) (moves [][]int) {
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		cnt, _ := strconv.Atoi(parts[1])
		from, _ := strconv.Atoi(parts[3])
		to, _ := strconv.Atoi(parts[5])
		moves = append(moves, []int{from, to, cnt})
	}
	return moves
}

func moveCrates(crates []string, moves [][]int) (top string) {
	for _, move := range moves {
		from := move[0]
		to := move[1]
		cnt := move[2]

		splitPos := len(crates[from]) - cnt
		toMove := crates[from][splitPos:]
		for i := len(toMove) - 1; i >= 0; i-- {
			crates[to] += string(toMove[i])
		}
		crates[from] = crates[from][0:splitPos]
	}
	for i := 1; i < len(crates); i++ {
		curTop := " "
		crate := crates[i]
		if len(crate) > 0 {
			curTop = string(crate[len(crate)-1])
		}
		top += curTop
	}
	return top
}

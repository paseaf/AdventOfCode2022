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
	result1 := solvePuzzle1(filepath.Join("./day9", "input.txt"))
	fmt.Printf("Result1 : %v\n", result1)
	result2 := solvePuzzle2(filepath.Join("./day9", "input.txt"))
	fmt.Printf("Result2 : %v\n", result2)
}

func solvePuzzle1(path string) (result int) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// visited: "x, y" -> bool
	visited := map[string]bool{"0,0": true}
	scanner := bufio.NewScanner(file)
	head := []int{0, 0}
	tail := []int{0, 0}
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		dir := parts[0]
		steps, _ := strconv.Atoi(parts[1])
		for i := 0; i < steps; i++ {
			head = moveHead(head, dir)
			tail = moveTail(tail, head)
			tailPos := strconv.Itoa(tail[0]) + "," + strconv.Itoa(tail[1])
			visited[tailPos] = true
		}
	}

	return len(visited)
}

func moveHead(head []int, dir string) []int {
	x := head[0]
	y := head[1]
	switch dir {
	case "U":
		y++
	case "D":
		y--
	case "L":
		x--
	case "R":
		x++
	}
	return []int{x, y}
}

func moveTail(tail []int, head []int) []int {
	hX := head[0]
	hY := head[1]
	tX := tail[0]
	tY := tail[1]

	x := tX
	y := tY
	distX := abs(hX - tX)
	distY := abs(hY - tY)
	if distX == 2 {
		if tX < hX {
			x++
		} else {
			x--
		}
		if distY < 2 {
			y = hY
		}
	}
	if distY == 2 {
		if tY < hY {
			y++
		} else {
			y--
		}
		if distX < 2 {
			x = hX
		}
	}

	return []int{x, y}
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func solvePuzzle2(path string) (result int) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// visited: "x, y" -> bool
	visited := map[string]bool{"0,0": true}
	scanner := bufio.NewScanner(file)
	head := []int{0, 0}
	var tails [][]int
	tailLen := 9
	for i := 0; i < tailLen; i++ {
		tails = append(tails, []int{0, 0})
	}

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		dir := parts[0]
		steps, _ := strconv.Atoi(parts[1])
		for i := 0; i < steps; i++ {
			head = moveHead(head, dir)
			tails = moveTails(tails, head)
			tail := tails[len(tails)-1]
			tailPos := strconv.Itoa(tail[0]) + "," + strconv.Itoa(tail[1])
			visited[tailPos] = true
		}
	}

	return len(visited)
}

func moveTails(tails [][]int, head []int) [][]int {
	var newTails [][]int
	for _, tail := range tails {
		newTail := moveTail(tail, head)
		newTails = append(newTails, newTail)
		head = newTail
	}

	return newTails
}

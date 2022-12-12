package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {
	result1 := solvePuzzle1(filepath.Join("./day12", "input.txt"))
	fmt.Printf("Result1 : %v\n", result1)
	result2 := solvePuzzle2(filepath.Join("./day12", "input.txt"))
	fmt.Printf("Result2 : %v\n", result2)
}

type Node [2]int

func solvePuzzle1(path string) (result int) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	graph, startNodes := parseGraph(file, 'S')
	result = 1 << 32
	for _, startNode := range startNodes {
		dist := findShortestDist(graph, startNode)
		if dist < result {
			result = dist
		}
	}
	return result
}

func solvePuzzle2(path string) (result int) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	graph, startNodes := parseGraph(file, 'a')

	result = 1 << 32
	for _, startNode := range startNodes {
		dist := findShortestDist(graph, startNode)
		if dist > 0 && dist < result {
			result = dist
		}
	}
	return result
}

func parseGraph(file *os.File, startChar rune) (graph [][]rune, startNodes []Node) {
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		row := make([]rune, 0)
		for i, char := range line {
			row = append(row, char)
			if char == 'S' || char == startChar {
				node := [2]int{len(graph), i}
				startNodes = append(startNodes, node)
			}
		}
		graph = append(graph, row)
	}
	return graph, startNodes
}

func findShortestDist(graph [][]rune, start Node) (dist int) {
	rows := len(graph)
	cols := len(graph[0])
	queue := [][2]int{start}
	visited := map[[2]int]bool{}
	graph[start[0]][start[1]] = 'a'
	for len(queue) > 0 {
		var nextSteps [][2]int
		for _, node := range queue {
			if visited[node] {
				continue
			}
			row := node[0]
			col := node[1]
			curVal := graph[row][col]

			rowUp := row - 1
			up := [2]int{rowUp, col}
			if rowUp >= 0 && graph[rowUp][col] <= curVal+1 {
				if curVal == 'z' && graph[rowUp][col] == 'E' {
					return dist + 1
				}
				nextSteps = append(nextSteps, up)
			}

			rowDown := row + 1
			down := [2]int{rowDown, col}
			if rowDown < rows && graph[rowDown][col] <= curVal+1 {
				if curVal == 'z' && graph[rowDown][col] == 'E' {
					return dist + 1
				}
				nextSteps = append(nextSteps, down)
			}

			colLeft := col - 1
			left := [2]int{row, colLeft}
			if colLeft >= 0 && graph[row][colLeft] <= curVal+1 {
				if curVal == 'z' && graph[row][colLeft] == 'E' {
					return dist + 1
				}
				nextSteps = append(nextSteps, left)
			}

			colRight := col + 1
			right := [2]int{row, colRight}
			if colRight < cols && graph[row][colRight] <= curVal+1 {
				if curVal == 'z' && graph[row][colRight] == 'E' {
					return dist + 1
				}
				nextSteps = append(nextSteps, right)
			}

			visited[[2]int{row, col}] = true
		}
		queue = nextSteps
		dist++
	}
	return -1
}

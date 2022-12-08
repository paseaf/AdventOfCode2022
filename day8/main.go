package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {
	result1 := findVisibleTrees(filepath.Join("./day8", "input.txt"))
	fmt.Printf("Result1 : %v\n", result1)
	result2 := findHighestScenicScore(filepath.Join("./day8", "input.txt"))
	fmt.Printf("Result2 : %v\n", result2)
}

func findVisibleTrees(path string) (result int) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	treeGrid := parseTreeGrid(file)
	visibleTrees := countVisibleTrees(treeGrid)

	return visibleTrees
}

func parseTreeGrid(file *os.File) (treeGrid [][]int) {
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		var row []int
		for _, char := range line {
			row = append(row, int(char-'0'))
		}
		treeGrid = append(treeGrid, row)
	}
	return treeGrid
}

func countVisibleTrees(treeGrid [][]int) int {
	m := len(treeGrid)
	n := len(treeGrid[0])
	count := 0
	for row := 0; row < m; row++ {
		for col := 0; col < n; col++ {
			if row == 0 || row == m-1 || col == 0 || col == n-1 {
				count++
			} else {
				if checkVisible(treeGrid, row, col) {
					count++
				}
			}
		}
	}
	return count
}

func checkVisible(treeGrid [][]int, curRow int, curCol int) bool {
	m := len(treeGrid)
	curHeight := treeGrid[curRow][curCol]
	n := len(treeGrid[0])

	upVisible := true
	for row := 0; row < curRow; row++ {
		if treeGrid[row][curCol] >= curHeight {
			upVisible = false
			break
		}
	}
	if upVisible {
		return true
	}

	downVisible := true
	for row := curRow + 1; row < m; row++ {
		if treeGrid[row][curCol] >= curHeight {
			downVisible = false
			break
		}
	}
	if downVisible {
		return true
	}

	leftVisible := true
	for col := 0; col < curCol; col++ {
		if treeGrid[curRow][col] >= curHeight {
			leftVisible = false
			break
		}
	}
	if leftVisible {
		return true
	}

	rightVisible := true
	for col := curCol + 1; col < n; col++ {
		if treeGrid[curRow][col] >= curHeight {
			rightVisible = false
			break
		}
	}
	if rightVisible {
		return true
	}

	return false
}

func findHighestScenicScore(path string) (score int) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	treeGrid := parseTreeGrid(file)
	score = parseHighestScenicScore(treeGrid)

	return score
}

func parseHighestScenicScore(treeGrid [][]int) (score int) {
	m := len(treeGrid)
	n := len(treeGrid[0])
	for row := 0; row < m; row++ {
		for col := 0; col < n; col++ {
			curScore := checkScore(treeGrid, row, col)
			if curScore > score {
				score = curScore
			}
		}
	}
	return score
}

func checkScore(treeGrid [][]int, curRow int, curCol int) int {
	m := len(treeGrid)
	curHeight := treeGrid[curRow][curCol]
	n := len(treeGrid[0])

	upScore := 0
	for row := curRow - 1; row >= 0; row-- {
		if treeGrid[row][curCol] >= curHeight {
			upScore++
			break
		}
		upScore++
	}

	downScore := 0
	for row := curRow + 1; row < m; row++ {
		if treeGrid[row][curCol] >= curHeight {
			downScore++
			break
		}
		downScore++
	}

	leftScore := 0
	for col := curCol - 1; col >= 0; col-- {
		if treeGrid[curRow][col] >= curHeight {
			leftScore++
			break
		}
		leftScore++
	}

	rightScore := 0
	for col := curCol + 1; col < n; col++ {
		if treeGrid[curRow][col] >= curHeight {
			rightScore++
			break
		}
		rightScore++
	}

	return upScore * downScore * leftScore * rightScore
}

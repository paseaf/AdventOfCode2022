package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

func main() {
	result1 := findMaxCalories(filepath.Join("day1", "input.txt"))
	fmt.Printf("Result1 is %v\n", result1)

	result2 := findTopThree(filepath.Join("day1", "input.txt"))
	fmt.Printf("Result2 is %v\n", result2)
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

func findTopThree(path string) (result int) {
	// 1. read input
	file, err := os.Open(path)
	check(err)
	defer file.Close()

	maxHeap := &IntHeap{}
	heap.Init(maxHeap)
	scanner := bufio.NewScanner(file)
	curResult := 0
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			heap.Push(maxHeap, curResult)
			curResult = 0
		} else {
			curNum, err := strconv.Atoi(line)
			check(err)
			curResult += curNum
		}
	}
	// get top 3
	for i := 0; i < 3; i++ {
		result += heap.Pop(maxHeap).(int)
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

// An IntHeap is a MAX-heap of ints.
// Modified from https://pkg.go.dev/container/heap#example-package-IntHeap
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {
	result1 := findFirstMarker(filepath.Join("./day6", "input.txt"), 4)
	fmt.Printf("Result1 : %v\n", result1)
	result2 := findFirstMarker(filepath.Join("./day6", "input.txt"), 14)
	fmt.Printf("Result2 : %v\n", result2)
}

// return the index of first char after marker
func findFirstMarker(path string, markerLen int) (result int) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	counter := map[byte]int{}
	var queue []byte
	pos := 0
	for char, err := reader.ReadByte(); err == nil; char, err = reader.ReadByte() {
		counter[char] += 1
		queue = append(queue, char)
		pos++

		if len(queue) > markerLen {
			first := queue[0]
			queue = queue[1:]
			counter[first] -= 1
			if counter[first] == 0 {
				delete(counter, first)
			}
		}

		if len(counter) == markerLen {
			return pos
		}
	}

	return -1
}

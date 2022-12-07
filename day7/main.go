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
	result1 := findTotalDirSize(filepath.Join("./day7", "input.txt"))
	fmt.Printf("Result1 : %v\n", result1)
	result2 := findMinDirSizeToFree(filepath.Join("./day7", "input.txt"))
	fmt.Printf("Result2 : %v\n", result2)
}

// return the total size of dirs with size <= 100000
func findTotalDirSize(path string) (result int) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fileSizes := parseFileSizes(file)
	dirSizes := collectDirSizes(fileSizes)

	// step 3: collect dirs
	for _, dirSize := range dirSizes {
		if dirSize <= 100000 {
			result += dirSize
		}
	}
	return result
}

func findMinDirSizeToFree(path string) int {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fileSizes := parseFileSizes(file)
	dirSizes := collectDirSizes(fileSizes)

	diskSize := 70000000
	used := dirSizes["/"]
	actualFree := diskSize - used
	expectedFree := 30000000
	minDirSizeToFree := diskSize
	// step 3: collect dirs
	for _, dirSize := range dirSizes {
		if actualFree+dirSize >= expectedFree {
			if dirSize < minDirSizeToFree {
				minDirSizeToFree = dirSize
			}
		}
	}
	return minDirSizeToFree
}

func parseFileSizes(file *os.File) map[string]int {
	scanner := bufio.NewScanner(file)
	fileSizes := map[string]int{}
	curPath := ""

	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "$ cd ") {
			parts := strings.Split(line, " ")
			path := parts[2]
			if path == ".." {
				curPath = parentPath(curPath)
			} else {
				if isAbsolutePath(path) {
					curPath = path
				} else {
					curPath += path
				}
				// mark dirs with an ending "/"
				if curPath[len(curPath)-1] != '/' {
					curPath += "/"
				}
			}
		} else if strings.HasPrefix(line, "$ ls") {
			continue
		} else {
			// handle sizes
			parts := strings.Split(line, " ")
			if parts[0] != "dir" {
				fileSize, _ := strconv.Atoi(parts[0])
				fileName := parts[1]
				filePath := curPath + fileName
				fileSizes[filePath] = fileSize
			}
		}
	}
	return fileSizes
}

func collectDirSizes(fileSizes map[string]int) map[string]int {
	dirSizes := map[string]int{}
	for path, size := range fileSizes {
		dirs := getDescendants(path)
		for _, dir := range dirs {
			dirSizes[dir] += size
		}
	}
	return dirSizes
}

func parentPath(dirPath string) string {
	parts := strings.Split(dirPath, "/")
	parts = append(parts[0:len(parts)-2], parts[len(parts)-1])
	return strings.Join(parts, "/")
}
func isAbsolutePath(path string) bool {
	return path[0] == '/'
}

func getDescendants(filePath string) (dirs []string) {
	for i := 0; i < len(filePath); i++ {
		if filePath[i] == '/' {
			dirs = append(dirs, filePath[0:i+1])
		}
	}
	return dirs
}

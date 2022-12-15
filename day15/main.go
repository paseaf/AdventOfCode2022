package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

func main() {
	result1 := solvePuzzle1(filepath.Join("./day15", "input.txt"), 2000000)
	fmt.Printf("Result1 : %v\n", result1)
	result2 := solvePuzzle2(filepath.Join("./day15", "input.txt"), 4000000)
	fmt.Printf("Result2 : %v\n", result2)
}

type Pos struct {
	x int
	y int
}

func solvePuzzle1(path string, y int) (result int) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Phase 1: prepare
	scanner := bufio.NewScanner(file)
	sensorBeamPairs := map[Pos]Pos{}
	beamSet := map[Pos]bool{}
	for scanner.Scan() {
		sensor, beam := parseLine(scanner.Text())
		sensorBeamPairs[sensor] = beam
		beamSet[beam] = true
	}

	// Phase 2:
	// Idea:
	//  for each sensor at (x, y)
	//  grow from (x, yTarget) to x++, x--
	//   until both sides out of bound
	//   count all non-beam positions and return
	cnt := 0
	for sensor, beam := range sensorBeamPairs {
		cnt += expand(sensor, beam, &beamSet, y)
	}
	return cnt
}

func parseLine(line string) (sensor Pos, beam Pos) {
	parts := strings.Split(line, " ")
	sensorX, _ := strconv.Atoi(parts[2][2 : len(parts[2])-1])
	sensorY, _ := strconv.Atoi(parts[3][2 : len(parts[3])-1])
	beamX, _ := strconv.Atoi(parts[8][2 : len(parts[8])-1])
	beamY, _ := strconv.Atoi(parts[9][2:len(parts[9])])
	return Pos{sensorX, sensorY}, Pos{beamX, beamY}
}

func expand(sensor Pos, beam Pos, beamSet *map[Pos]bool, y int) (cnt int) {
	maxDist := getDist(sensor, beam)
	// left
	curPos := Pos{sensor.x, y}
	for getDist(curPos, sensor) <= maxDist {
		if !(*beamSet)[curPos] {
			(*beamSet)[curPos] = true
			cnt++
		}
		curPos.x--
	}
	// right
	curPos = Pos{sensor.x, y}
	for getDist(curPos, sensor) <= maxDist {
		if !(*beamSet)[curPos] {
			(*beamSet)[curPos] = true
			cnt++
		}
		curPos.x++
	}
	return cnt
}

func getDist(a Pos, b Pos) int {
	xDist := abs(a.x - b.x)
	yDist := abs(a.y - b.y)
	return xDist + yDist
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
func solvePuzzle2(path string, boundary int) (result int) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Phase 1: prepare
	scanner := bufio.NewScanner(file)
	sensorRange := map[Pos]int{}
	beamSet := map[Pos]bool{}
	for scanner.Scan() {
		sensor, beam := parseLine(scanner.Text())
		sensorRange[sensor] = getDist(sensor, beam)
		beamSet[beam] = true
	}

	// Phase 2:
	// Idea 1: scan each row for y in [0, boundary]
	//   BAD TOO SLOW
	// Idea 2: get x range for each y, and merge ranges
	//   SUPER FAST!!!
	for y := 0; y <= boundary; y++ {
		if y%100000 == 0 {
			fmt.Printf("Progress:\t %v %%\n", y*100/boundary)
		}
		var ranges [][2]int
		for sensor, dist := range sensorRange {
			if left, right, ok := getRange(sensor, dist, y); ok {
				ranges = append(ranges, [2]int{left, right})
			}
		}
		// sort
		sort.Slice(ranges, func(i, j int) bool {
			return ranges[i][0] < ranges[j][0]
		})
		// merge
		merged := ranges[0]
		for i := 1; i < len(ranges); i++ {
			if ranges[i][0] > merged[1] {
				hiddenX := ranges[i][0] - 1
				return hiddenX*4000000 + y
			}
			if ranges[i][1] > merged[1] {
				merged[1] = ranges[i][1]
			}
		}
	}
	return -1
}

func getRange(sensor Pos, maxDist int, y int) (left int, right int, ok bool) {
	x := sensor.x
	rest := maxDist - abs(sensor.y-y)
	return x - rest, x + rest, rest >= 0
}

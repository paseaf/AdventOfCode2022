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
	result1 := solvePuzzle1(filepath.Join("./day16", "input.txt"))
	fmt.Printf("Result1 : %v\n", result1)
}

type State struct {
	rate         int
	release      int
	openedValves map[string]bool
	curRound     int
	curValve     string
	prevMove     string
}

type Valve struct {
	rate      int
	neighbors []string
}

func (state State) copy() (copyState State) {
	copyState = state
	copyOpenedValves := map[string]bool{}
	for k, v := range state.openedValves {
		copyOpenedValves[k] = v
	}
	copyState.openedValves = copyOpenedValves
	return copyState
}

func solvePuzzle1(path string) (result int) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Phase 1: prepare
	valves := map[string]Valve{}
	scanner := bufio.NewScanner(file)
	initValue := "AA"
	for scanner.Scan() {
		valve, rate, neighbors := parseLine(scanner.Text())
		valves[valve] = Valve{rate, neighbors}
	}

	// Phase 2: BFS
	// TO cut: valves with rate 0
	initState := State{
		rate:         0,
		release:      0,
		openedValves: map[string]bool{initValue: true},
		curRound:     0,
		curValve:     initValue,
	}
	curStates := []State{initState}
	for minute := 1; minute <= 30; minute++ {
		newStates := make([]State, 0)
		for _, state := range curStates {
			// if all opened. only increase release
			if inefficient(state) {
				continue
			}
			state.curRound++
			// all opened. no need to move or open
			if len(state.openedValves) == len(valves) {
				state := state.copy()
				state.release += state.rate
				if state.release >= result {
					result = state.release
					newStates = append(newStates, state)
				}
			} else {
				// case 1: open cur valve if possible
				if !state.openedValves[state.curValve] {
					state := state.copy()
					state.release += state.rate
					if state.release > result {
						result = state.release
					}
					state.rate += valves[state.curValve].rate
					state.prevMove = "open"
					state.openedValves[state.curValve] = true
					newStates = append(newStates, state)
				}
				// case 2: move to another valve and open it for free if rate == 0
				for _, neighbor := range valves[state.curValve].neighbors {
					state := state.copy()
					if neighbor == state.prevMove {
						continue
					}
					state.prevMove = state.curValve
					state.curValve = neighbor
					if valves[state.curValve].rate == 0 {
						state.openedValves[state.curValve] = true
					}
					state.release += state.rate
					if state.release > result {
						result = state.release
					}
					newStates = append(newStates, state)
				}
			}
			// open
		}
		fmt.Println(minute, result)
		curStates = newStates
	}

	return result
}

func inefficient(state State) bool {
	if state.curRound > 20 && state.rate < 30 {
		return true
	}
	if state.curRound > 24 && state.rate < 50 {
		return true
	}
	return false
}

func parseLine(line string) (valve string, rate int, neighbors []string) {
	parts := strings.Split(line, " ")
	valve = parts[1]
	rate, _ = strconv.Atoi(parts[4][5 : len(parts[4])-1])
	for i := 9; i < len(parts); i++ {
		part := parts[i]
		if part[len(part)-1] == ',' {
			part = part[0 : len(part)-1]
		}
		neighbors = append(neighbors, part)
	}

	return valve, rate, neighbors
}

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
	result1 := solvePuzzle1(filepath.Join("./day11", "input.txt"))
	fmt.Printf("Result1 : %v\n", result1)
	result2 := solvePuzzle2(filepath.Join("./day11", "input.txt"))
	fmt.Printf("Result2 : %v\n", result2)
}

type Monkey struct {
	items    []int
	operator string
	operand  int
	div      int
	trueId   int
	falseId  int
}

func solvePuzzle1(path string) (result int) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	monkeys := parseMonkeys(file)

	counter := playGame(monkeys, 20, 3)
	first, second := findWinners(counter)
	return first * second
}

func solvePuzzle2(path string) (result int) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	monkeys := parseMonkeys(file)

	counter := playGame(monkeys, 10000, 1)
	first, second := findWinners(counter)
	return first * second
}

func playGame(monkeys []Monkey, rounds int, dropRate int) (counter []int) {
	gcd := 1
	for _, monkey := range monkeys {
		gcd *= monkey.div
	}

	counter = make([]int, len(monkeys))
	for i := 0; i < rounds; i++ {
		for monkeyId, monkey := range monkeys {
			for _, old := range monkey.items {
				counter[monkeyId] += 1
				var newVal int
				switch monkey.operator {
				case "*":
					newVal = old * monkey.operand
				case "+":
					newVal = old + monkey.operand
				case "sqr":
					newVal = old * old
				default:
					fmt.Println("Unexpected operator", monkey.operator)
				}

				newVal /= dropRate
				newVal %= gcd
				if newVal%monkey.div == 0 {
					monkeys[monkey.trueId].items = append(monkeys[monkey.trueId].items, newVal)
				} else {
					monkeys[monkey.falseId].items = append(monkeys[monkey.falseId].items, newVal)
				}
			}
			monkeys[monkeyId].items = make([]int, 0)
		}
	}
	return counter
}

func findWinners(counter []int) (int, int) {
	first, second := -1, -1
	for _, num := range counter {
		if num > first {
			second = first
			first = num
		} else if num > second {
			second = num
		}
	}
	return first, second
}

func parseMonkeys(file *os.File) []Monkey {
	var monkeys []Monkey
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		monkey := makeMonkey(scanner)
		monkeys = append(monkeys, monkey)
		scanner.Scan() // skip blank line
	}
	return monkeys
}

func makeMonkey(scanner *bufio.Scanner) Monkey {
	scanner.Scan()
	items := parseItems(scanner.Text())

	scanner.Scan()
	operator, operand := parseOp(scanner.Text())

	scanner.Scan()
	div := parseDiv(scanner.Text())

	scanner.Scan()
	trueId := parseId(scanner.Text())
	scanner.Scan()
	falseId := parseId(scanner.Text())

	monkey := Monkey{
		items:    items,
		operator: operator,
		operand:  operand,
		div:      div,
		trueId:   trueId,
		falseId:  falseId,
	}
	return monkey
}

func parseItems(line string) (items []int) {
	itemStr := strings.Split(line, ": ")[1]
	strParts := strings.Split(itemStr, ", ")
	for _, part := range strParts {
		item, _ := strconv.Atoi(part)
		items = append(items, item)
	}

	return items
}

func parseOp(line string) (operator string, operand int) {
	part := strings.Split(line, "old ")[1]
	ops := strings.Split(part, " ")
	if ops[1] == "old" {
		operator = "sqr"
	} else {
		operator = ops[0]
		operand, _ = strconv.Atoi(ops[1])
	}

	return operator, operand
}

func parseDiv(line string) (div int) {
	part := strings.Split(line, "by ")[1]
	div, _ = strconv.Atoi(part)

	return div
}

func parseId(line string) (id int) {
	part := strings.Split(line, "monkey ")[1]
	id, _ = strconv.Atoi(part)

	return id
}

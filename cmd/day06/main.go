package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type operation struct {
	numbers   []int
	operation string
}

func parseInput(input string) []operation {
	lines := strings.Split(input, "\n")

	output := make([]operation, 0)

	for i := 0; i < len(lines); i++ {
		columns := strings.Fields(lines[i])

		for j, column := range columns {
			if i == 0 {
				parsed, _ := strconv.Atoi(column)
				op := operation{
					numbers: []int{parsed},
				}
				output = append(output, op)
			} else if i == len(lines)-1 {
				output[j].operation = column
			} else {
				parsed, _ := strconv.Atoi(column)
				output[j].numbers = append(output[j].numbers, parsed)
			}
		}
	}

	return output
}

func partOne(input []operation) int {
	total := 0
	for _, op := range input {
		current := 0
		if op.operation == "*" {
			current = 1
		}
		for _, num := range op.numbers {
			if op.operation == "+" {
				current += num
			}
			if op.operation == "*" {
				current *= num
			}
		}
		total += current
	}
	return total
}

func parseInputPartTwo(input string) []operation {
	lines := strings.Split(input, "\n")
	fields := strings.Fields(lines[0])

	rows := len(lines) - 1
	cols := len(lines[0])

	output := make([]operation, len(fields))
	i := 0

	for col := cols - 1; col >= 0; col-- {
		current := 0
		for row := 0; row < rows; row++ {
			ch := lines[row][col]
			if ch >= '0' && ch <= '9' {
				if current != 0 {
					current *= 10
				}
				digit := int(ch - '0')
				current += digit
			}
		}
		if current == 0 {
			i++
			continue
		}

		output[i].numbers = append(output[i].numbers, current)
	}

	operations := strings.Fields(lines[len(lines)-1])

	for i, op := range operations {
		output[len(output)-1-i].operation = op
	}

	return output
}

func partTwo(input []operation) int {
	total := 0
	for _, op := range input {
		current := 0
		if op.operation == "*" {
			current = 1
		}
		for _, num := range op.numbers {
			if op.operation == "+" {
				current += num
			}
			if op.operation == "*" {
				current *= num
			}
		}
		total += current
	}
	return total
}

func main() {
	d, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	path := fmt.Sprintf("%s/data/06.txt", d)

	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	b, err := io.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}

	contents := string(b)

	in := parseInput(contents)
	fmt.Printf("Part 1: %d\n", partOne(in))

	twoIn := parseInputPartTwo(contents)
	fmt.Printf("Part 2: %d\n", partTwo(twoIn))
}

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

/*
	The goal of this puzzle is to count the number of times a safe dial ends at 0.

	The numbers range from 0-99.
	Each line in the input represents a single move left or right.
*/

func turnDial(dialPos int, command string) int {
	moveLeft := false
	num := 0
	for _, ch := range command {
		if ch == 'L' {
			moveLeft = true
		}

		if ch >= '0' && ch <= '9' {
			val := int(ch - '0')
			num *= 10
			num += val
		}
	}

	// We are only counting times we land on zero, not go past it so going more than a full spin is not important.
	num = num % 100

	// Going left 25 lands on the same number as going right 75, so we will keep the return statement simpler.
	if moveLeft {
		num = 100 - num
	}

	return (dialPos + num) % 100
}

func main() {
	d, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	path := fmt.Sprintf("%s/cmd/day01/data/input.txt", d)

	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(f)

	dialPos := 50
	count := 0

	for scanner.Scan() {
		line := scanner.Text()

		dialPos = turnDial(dialPos, line)

		if dialPos == 0 {
			count++
		}
	}

	fmt.Printf("Times the dial hit zero: %d\n", count)
}

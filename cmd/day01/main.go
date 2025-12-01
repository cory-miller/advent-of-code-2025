package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// mod is a floor style modulo operation since Go uses truncated modulo.
//
// I use this because I want -25 % 100 to be 75, where -25 % 100 == -25 in Go.
func mod(a, b int) int {
	return (a%b + b) % b
}

func parseInput(command string) (bool, int) {
	left := false
	num := 0
	for _, ch := range command {
		if ch == 'L' {
			left = true
		}

		if ch >= '0' && ch <= '9' {
			val := int(ch - '0')
			num *= 10
			num += val
		}
	}

	return left, num
}

func turnDial(dialPos int, command string) int {
	moveLeft, num := parseInput(command)

	// We are only counting times we land on zero, not go past it so going more than a full spin is not important.
	num = num % 100

	// Going left 25 lands on the same number as going right 75, so we will keep the return statement simpler.
	if moveLeft {
		num = 100 - num
	}

	return (dialPos + num) % 100
}

func turnDialCountPasses(dialPos int, command string) (int /* position after spin */, int /* times passing zero */) {
	moveLeft, num := parseInput(command)

	// We will count the number of full rotations ahead of time.
	fullRot := num / 100
	num = num % 100

	if moveLeft {
		num *= -1
	}

	count := 0

	// We moved right and passed or stopped at 0.
	if dialPos+num >= 100 {
		count++
	}

	// We moved left and passed or stopped at 0.
	// This one needs to account for starting at 0. The right movement naturally accounts for that by checking for 100.
	if dialPos > 0 && dialPos+num <= 0 {
		count++
	}

	return mod(dialPos+num, 100), fullRot + count
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

	partTwoPos := 50
	partTwoCount := 0

	for scanner.Scan() {
		line := scanner.Text()

		dialPos = turnDial(dialPos, line)

		if dialPos == 0 {
			count++
		}

		nextTwoPos, addTwoCount := turnDialCountPasses(partTwoPos, line)
		partTwoPos = nextTwoPos
		partTwoCount += addTwoCount

		if dialPos != partTwoPos {
			log.Fatal("Part Two did not match")
		}
	}

	fmt.Printf("Part One: %d\n", count)
	fmt.Printf("Part Two: %d\n", partTwoCount)
}

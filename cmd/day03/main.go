package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func partOne(battery string) int {
	maxDig := 0
	maxDigIdx := 0

	for i := 0; i < len(battery)-1; i++ {
		if int(battery[i]-'0') > maxDig {
			maxDig = int(battery[i] - '0')
			maxDigIdx = i
		}

		if maxDig == 9 {
			break
		}
	}

	minDig := 0

	for i := len(battery) - 1; i > maxDigIdx; i-- {
		if int(battery[i]-'0') > minDig {
			minDig = int(battery[i] - '0')
		}

		if minDig == 9 {
			break
		}
	}

	return maxDig*10 + minDig
}

func maxIndex(in []rune, start, end int) int {
	m := 0
	index := 0

	for i := start; i <= end; i++ {
		ch := in[i]
		if int(ch-'0') > m {
			m = int(ch - '0')
			index = i
		}
		if m == 9 {
			break
		}
	}
	return index
}

func partTwo(battery string) int {
	// Suppose you have a string of length 15 and need to form a 12-digit number.
	// We can take the string, leaving the final 11 characters, and see what the biggest number the first part of the
	// range has. Repeat this until we form a 12-digit number.
	var chars []rune
	mi := 0
	for len(chars) < 12 {
		m := maxIndex([]rune(battery), mi, len(battery)-12+len(chars))
		chars = append(chars, rune(battery[m]))
		mi = m + 1
	}

	final, err := strconv.Atoi(string(chars))
	if err != nil {
		log.Fatal(err)
	}

	return final
}

func main() {
	d, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	path := fmt.Sprintf("%s/data/03.txt", d)

	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(f)

	partOneSum := 0
	partTwoSum := 0

	for scanner.Scan() {
		battery := scanner.Text()

		partOneSum += partOne(battery)
		partTwoSum += partTwo(battery)
	}

	fmt.Printf("Part One: %d\n", partOneSum)
	fmt.Printf("Part Two: %d\n", partTwoSum)
}

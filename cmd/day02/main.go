package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

// Pair is a simple data structure to track the ranges provided by the input
type Pair struct {
	start int
	end   int
}

func convertToPairs(contents string) []Pair {
	ranges := strings.Split(contents, ",")

	var allParts []Pair

	for _, r := range ranges {
		parts := strings.Split(r, "-")
		start, _ := strconv.Atoi(parts[0])
		end, _ := strconv.Atoi(parts[1])
		allParts = append(allParts, Pair{start, end})
	}

	return allParts
}

func partOne(allParts []Pair) int {
	sum := 0

	for _, pair := range allParts {
		for i := pair.start; i <= pair.end; i++ {
			val := strconv.Itoa(i)
			if len(val)%2 != 0 {
				continue
			}
			shouldAdd := true
			for j := 0; j < len(val)/2; j++ {
				a := val[j]
				b := val[len(val)/2+j]

				if a != b {
					shouldAdd = false
					break
				}
			}
			if shouldAdd {
				sum += i
			}
		}
	}

	return sum
}

func partTwo(allParts []Pair) int {
	sum := 0

	for _, pair := range allParts {
		for num := pair.start; num <= pair.end; num++ {
			val := strconv.Itoa(num)

			// We'll break the number into equal length strings starting from the longest possible string with matches.
			// If all parts match, then we add to the sum.
			for h := len(val) / 2; h > 0; h-- {
				if len(val)%h != 0 {
					continue
				}

				segments := len(val) / h
				var parts []string
				for j := 0; j < segments; j++ {
					parts = append(parts, val[j*h:(j+1)*h])
				}

				shouldAdd := true
				lastPart := parts[0]
				for _, part := range parts {
					if part != lastPart {
						shouldAdd = false
						break
					}
				}

				if shouldAdd {
					sum += num
					break // Don't count repeats
				}
			}
		}
	}

	return sum
}

func main() {
	d, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	path := fmt.Sprintf("%s/data/02.txt", d)

	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	b, err := io.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}

	contents := string(b)

	allParts := convertToPairs(contents)

	fmt.Printf("Part one: %v\n", partOne(allParts))
	fmt.Printf("Part two: %v\n", partTwo(allParts))
}

package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func partOne(lines []string) int {
	vals := make([]bool, len(lines[0]))

	splits := 0

	for _, line := range lines {
		for i, ch := range line {
			if ch == 'S' {
				vals[i] = true
			}
			if ch == '^' {
				if vals[i] {
					splits++
					vals[i] = false
					if i-1 >= 0 {
						vals[i-1] = true
					}
					if i+1 < len(vals) {
						vals[i+1] = true
					}
				}
			}
		}
	}

	return splits
}

func partTwo(lines []string) int {
	vals := make([]bool, len(lines[0]))
	counts := make([]int, len(lines[0]))

	splits := 0

	for x := 0; x < len(lines); x += 2 {
		line := lines[x]
		for i, ch := range line {
			if ch == 'S' {
				vals[i] = true
				counts[i] = 1
			}
			if ch == '^' {
				if vals[i] {

					if i-1 >= 0 {
						vals[i-1] = true
						counts[i-1] += counts[i]
					}
					if i+1 < len(vals) {
						vals[i+1] = true
						counts[i+1] += counts[i]
					}
				}
				vals[i] = false
				counts[i] = 0
			}
		}
	}

	for i := 0; i < len(counts); i++ {
		if vals[i] {
			splits += counts[i]
		}
	}

	return splits
}

func main() {
	d, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	path := fmt.Sprintf("%s/data/07.txt", d)

	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	b, err := io.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}

	contents := string(b)

	fmt.Printf("Part 1: %d\n", partOne(strings.Split(contents, "\n")))
	fmt.Printf("Part 2: %d\n", partTwo(strings.Split(contents, "\n")))
}

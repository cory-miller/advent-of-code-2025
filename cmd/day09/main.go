package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type pair struct {
	x int
	y int
}

func partOne(contents string) int {
	lines := strings.Split(contents, "\n")

	m := make(map[*pair]bool)

	for _, line := range lines {
		parts := strings.Split(line, ",")
		col, _ := strconv.Atoi(parts[0])
		row, _ := strconv.Atoi(parts[1])
		m[&pair{row, col}] = true
	}

	size := 0

	for p1, _ := range m {
		for p2, _ := range m {
			if p1 == p2 {
				continue
			}

			s := (p1.x - p2.x + 1) * (p1.y - p2.y + 1)
			if s < 0 {
				s = -s
			}

			if s > size {
				size = s
			}
		}
	}

	return size
}

func main() {
	d, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	path := fmt.Sprintf("%s/data/09.txt", d)

	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	b, err := io.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}

	contents := string(b)

	fmt.Printf("Part 1: %v\n", partOne(contents))
}

package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

type input struct {
	ranges       [][2]int
	availableIDs []int
}

func parseInput(contents string) input {
	lines := strings.Split(contents, "\n")

	in := input{
		ranges:       make([][2]int, 0),
		availableIDs: make([]int, 0),
	}

	parseRanges := true

	starts := make([]int, 0)
	ends := make([]int, 0)

	for _, line := range lines {
		if line == "" {
			parseRanges = false
			continue
		}

		if parseRanges {
			parts := strings.Split(line, "-")
			start, _ := strconv.Atoi(parts[0])
			end, _ := strconv.Atoi(parts[1])
			starts = append(starts, start)
			ends = append(ends, end)
		} else {
			id, _ := strconv.Atoi(line)
			in.availableIDs = append(in.availableIDs, id)
		}
	}

	slices.Sort(starts)
	slices.Sort(ends)

	in.ranges = append(in.ranges, [2]int{starts[0], ends[0]})

	rangeIndex := 0

	for i := 1; i < len(starts); i++ {
		nextStart := starts[i]
		nextEnd := ends[i]

		// No merge needed
		if nextStart > in.ranges[rangeIndex][1] {
			in.ranges = append(in.ranges, [2]int{nextStart, nextEnd})
			rangeIndex++
			continue
		}

		in.ranges[rangeIndex][0] = min(nextStart, in.ranges[rangeIndex][0])
		in.ranges[rangeIndex][1] = max(nextEnd, in.ranges[rangeIndex][1])
	}

	return in
}

func partOne(in input) int {
	fresh := 0

	for _, id := range in.availableIDs {
		for _, r := range in.ranges {
			if id >= r[0] && id <= r[1] {
				fresh++
			}
		}
	}

	return fresh
}

func partTwo(in input) int {
	fresh := 0

	for _, r := range in.ranges {
		fresh += r[1] - r[0] + 1
	}

	return fresh
}

func main() {
	d, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	path := fmt.Sprintf("%s/data/05.txt", d)

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
	fmt.Printf("Part 2: %d\n", partTwo(in))
}

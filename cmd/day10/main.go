package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type machineSolver struct {
	desired int
	buttons [][]int
	jolts   []int
	cache   map[int]int
}

func parseLine(line string) *machineSolver {
	parts := strings.Fields(line)

	machineSegment := parts[0]
	machineSegment = machineSegment[1 : len(machineSegment)-1]

	desired := 0

	for i, ch := range machineSegment {
		if ch == '#' {
			desired += 1 << i
		}
	}

	var buttons [][]int
	for i := 1; i < len(parts)-1; i++ {
		part := parts[i][1 : len(parts[i])-1]
		lightSwitches := strings.Split(part, ",")
		var button []int
		for _, lightSwitch := range lightSwitches {
			value, _ := strconv.Atoi(lightSwitch)
			button = append(button, value)
		}
		buttons = append(buttons, button)
	}

	lastIndex := len(buttons) - 1
	lastJoltIndex := len(parts[lastIndex]) - 1
	joltSegment := parts[lastIndex][1:lastJoltIndex]
	var joltValue []int
	for i := 0; i < len(joltSegment); i++ {
		value, _ := strconv.Atoi(string(joltSegment[i]))
		joltValue = append(joltValue, value)
	}

	return &machineSolver{
		desired: desired,
		buttons: buttons,
		jolts:   joltValue,
		cache:   make(map[int]int),
	}
}

func (m *machineSolver) solve(pressed2, count int) (int, bool) {
	if count >= len(m.buttons) {
		return count, false
	}

	var nextStates []int

	for _, button := range m.buttons {
		num := 0

		for _, b := range button {
			num += 1 << b
		}

		pressedState := num ^ pressed2

		if pressedState == m.desired {
			return count + 1, true
		}

		if val, has := m.cache[pressedState]; has && val <= count {
			continue
		}

		m.cache[pressedState] = count

		nextStates = append(nextStates, pressedState)
	}

	minCount := math.MaxInt
	hadMatch := false

	for _, nextState := range nextStates {
		thisCount, matched := m.solve(nextState, count+1)
		if matched {
			minCount = min(minCount, thisCount)
			hadMatch = true
		}
	}

	return minCount, hadMatch
}

func parseLinePartOne(line string) int {
	state := parseLine(line)

	count, ok := state.solve(0, 0)
	if ok {
		return count
	}

	return 0
}

func partOne(contents string) int {
	sum := 0

	for _, line := range strings.Split(contents, "\n") {
		sum += parseLinePartOne(line)
	}

	return sum
}

func main() {
	d, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	path := fmt.Sprintf("%s/data/10.txt", d)

	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	b, err := io.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}

	contents := string(b)

	fmt.Printf("Part 1: %v\n", partOne(contents))
}

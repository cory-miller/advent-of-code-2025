package main

import (
	"fmt"
	"testing"
)

func TestPartOne(t *testing.T) {
	input := `123 328  51 64 
 45 64  387 23 
  6 98  215 314
*   +   *   +  `

	parsedInput := parseInput(input)

	got := partOne(parsedInput)
	want := 4277556

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestPartTwo(t *testing.T) {
	input := `123 328  51 64 
 45 64  387 23 
  6 98  215 314
*   +   *   +  `

	parsedInput := parseInputPartTwo(input)

	got := partTwo(parsedInput)
	want := 3263827

	fmt.Println(parsedInput)

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

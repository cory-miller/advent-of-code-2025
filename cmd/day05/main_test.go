package main

import (
	"testing"
)

func TestPartOne(t *testing.T) {
	fileContents := `3-5
10-14
16-20
12-18

1
5
8
11
17
32`

	in := parseInput(fileContents)

	got := partOne(in)
	want := 3

	if got != want {
		t.Errorf("partOne() = %d, want %d", got, want)
	}
}

func TestPartTwo(t *testing.T) {
	fileContents := `3-5
10-14
16-20
12-18

1
5
8
11
17
32`

	in := parseInput(fileContents)

	got := partTwo(in)
	want := 14

	if got != want {
		t.Errorf("partTwo() = %d, want %d", got, want)
	}
}

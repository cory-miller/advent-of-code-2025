package main

import (
	"strings"
	"testing"
)

func TestPartOne(t *testing.T) {
	input := `[.##.] (3) (1,3) (2) (2,3) (0,2) (0,1) {3,5,4,7}
[...#.] (0,2,3,4) (2,3) (0,4) (0,1,2) (1,2,3,4) {7,5,12,7,2}
[.###.#] (0,1,2,3,4) (0,3,4) (0,1,2,4,5) (1,2) {10,11,11,5,10,5}`

	lines := strings.Split(input, "\n")

	sum := 0

	tests := []struct {
		want int
	}{
		{want: 2},
		{want: 3},
		{want: 2},
	}

	expectedSum := 7

	for i, tc := range tests {
		got := parseLinePartOne(lines[i])
		if got != tc.want {
			t.Errorf("Got %d, want %d", got, tc.want)
		}

		sum += got
	}

	if sum != expectedSum {
		t.Errorf("Got sum of %d, want %d", sum, expectedSum)
	}
}

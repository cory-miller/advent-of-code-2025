package main

import "testing"

func TestPartOne(t *testing.T) {
	input := `7,1
11,1
11,7
9,7
9,5
2,5
2,3
7,3`

	got := partOne(input)
	want := 50

	if got != want {
		t.Errorf("got: %v, want: %v", got, want)
	}

	t.Log("Not yet implemented")
}

func TestPartTwo(t *testing.T) {
	t.Log("Not yet implemented")
}

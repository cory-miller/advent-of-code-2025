package main

import "testing"

var (
	testInput = map[string]int{
		"987654321111111": 98,
		"811111111111119": 89,
		"234234234234278": 78,
		"818181911112111": 92,
	}
)

func TestPartOne(t *testing.T) {
	for battery, expectedPower := range testInput {
		power := partOne(battery)
		if power != expectedPower {
			t.Errorf("TestPartOne was incorrect, got: %d, want: %d.", power, expectedPower)
		}
	}
}

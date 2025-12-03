package main

import "testing"

func TestPartOne(t *testing.T) {
	testInput := []struct {
		battery       string
		expectedPower int
	}{
		{"987654321111111", 98},
		{"811111111111119", 89},
		{"234234234234278", 78},
		{"818181911112111", 92},
	}

	for _, tc := range testInput {
		power := partOne(tc.battery)
		if power != tc.expectedPower {
			t.Errorf("TestPartOne was incorrect, got: %d, want: %d.", power, tc.expectedPower)
		}
	}
}

func TestPartTwo(t *testing.T) {
	testInput := []struct {
		battery       string
		expectedPower int
	}{
		{"987654321111111", 987654321111},
		{"811111111111119", 811111111119},
		{"234234234234278", 434234234278},
		{"818181911112111", 888911112111},
		{"4123535244222342322334342233754335452333242522124322242423331132232242422443224231234323332243364522", 755554464522},
	}

	for _, tc := range testInput {
		power := partTwo(tc.battery)
		if power != tc.expectedPower {
			t.Errorf("partTwo(%v)=%v, want: %d.", tc.battery, power, tc.expectedPower)
		}
	}
}

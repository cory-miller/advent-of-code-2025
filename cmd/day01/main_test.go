package main

import (
	"fmt"
	"testing"
)

func TestPartOne(t *testing.T) {
	tests := []struct {
		initialPos  int
		expectedPos int
		command     string
	}{
		{
			initialPos:  0,
			expectedPos: 75,
			command:     "L25",
		},
		{
			initialPos:  0,
			expectedPos: 25,
			command:     "R25",
		},
		{
			initialPos:  50,
			expectedPos: 25,
			command:     "L25",
		},
		{
			initialPos:  50,
			expectedPos: 75,
			command:     "R25",
		},
		{
			initialPos:  50,
			expectedPos: 25,
			command:     "L125",
		},
		{
			initialPos:  50,
			expectedPos: 75,
			command:     "R125",
		},
	}

	for _, tc := range tests {
		got := turnDial(tc.initialPos, tc.command)
		if got != tc.expectedPos {
			t.Errorf("turnDial(%v, %v)=%v, want %v", tc.initialPos, tc.command, got, tc.expectedPos)
		}
	}
}

func TestPartTwo(t *testing.T) {
	tests := []struct {
		initialPos    int
		expectedPos   int
		expectedCount int
		command       string
	}{
		//L14
		//L46
		//L44
		{
			initialPos:    50,
			expectedPos:   36,
			expectedCount: 0,
			command:       "L14",
		},
		{
			initialPos:    36,
			expectedPos:   90,
			expectedCount: 1,
			command:       "L46",
		},
		{
			initialPos:    90,
			expectedPos:   46,
			expectedCount: 0,
			command:       "L44",
		},
		{
			initialPos:    90,
			expectedPos:   70,
			expectedCount: 5,
			command:       "L520",
		},
		{
			initialPos:    90,
			expectedPos:   10,
			expectedCount: 6,
			command:       "R520",
		},
		{
			initialPos:    0,
			expectedPos:   75,
			expectedCount: 0,
			command:       "L25",
		},
		{
			initialPos:    0,
			expectedPos:   25,
			expectedCount: 0,
			command:       "R25",
		},
		{
			initialPos:    50,
			expectedPos:   25,
			expectedCount: 0,
			command:       "L25",
		},
		{
			initialPos:    50,
			expectedPos:   75,
			expectedCount: 0,
			command:       "R25",
		},
		{
			initialPos:    50,
			expectedPos:   25,
			expectedCount: 1,
			command:       "L125",
		},
		{
			initialPos:    50,
			expectedPos:   75,
			expectedCount: 1,
			command:       "R125",
		},
		{
			initialPos:    50,
			expectedPos:   25,
			expectedCount: 11,
			command:       "L1125",
		},
		{
			initialPos:    50,
			expectedPos:   0,
			expectedCount: 2,
			command:       "L150",
		},
		{
			initialPos:    50,
			expectedPos:   75,
			expectedCount: 12,
			command:       "L1175",
		},
		{
			initialPos:    49,
			expectedPos:   74,
			expectedCount: 12,
			command:       "L1175",
		},
		{
			initialPos:    50,
			expectedPos:   75,
			expectedCount: 11,
			command:       "R1125",
		},
		{
			initialPos:    50,
			expectedPos:   0,
			expectedCount: 12,
			command:       "R1150",
		},
	}

	for _, tc := range tests {
		gotPos, got := turnDialCountPasses(tc.initialPos, tc.command)
		if got != tc.expectedCount || gotPos != tc.expectedPos {
			t.Errorf("turnDialCountPasses(%v, %v)=(%v, %v), want (%v, %v)", tc.initialPos, tc.command, gotPos, got, tc.expectedPos, tc.expectedCount)
		}
	}
}

func Test_PartTwoSim(t *testing.T) {
	position := 50
	count := 0
	position, got := turnDialCountPasses(position, "R50")
	count += got
	fmt.Printf("%v, %v\n", position, count)
	position, got = turnDialCountPasses(position, "R50")
	count += got
	fmt.Printf("%v, %v\n", position, count)
	position, got = turnDialCountPasses(position, "L50")
	count += got
	fmt.Printf("%v, %v\n", position, count)
	position, got = turnDialCountPasses(position, "L50")
	count += got
	fmt.Printf("%v, %v\n", position, count)
	position, got = turnDialCountPasses(position, "R75")
	count += got
	fmt.Printf("%v, %v\n", position, count)
	position, got = turnDialCountPasses(position, "L50")
	count += got
	fmt.Printf("%v, %v\n", position, count)
	position, got = turnDialCountPasses(position, "L25")
	count += got
	fmt.Printf("%v, %v\n", position, count)
	position, got = turnDialCountPasses(position, "L75")
	count += got
	fmt.Printf("%v, %v\n", position, count)
	position, got = turnDialCountPasses(position, "R50")
	count += got
	fmt.Printf("%v, %v\n", position, count)

	if count != 6 {
		t.Fatalf("%v, %v", position, count)
	}
}

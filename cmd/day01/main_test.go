package main

import "testing"

func Test(t *testing.T) {
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

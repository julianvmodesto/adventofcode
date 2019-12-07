package main

import "testing"

func TestReadInstruction(t *testing.T) {
	var tests = []struct {
		instruction string
		direction   string
		moves       int
	}{
		{
			instruction: "R6",
			direction:   "R",
			moves:       6,
		},
	}
	for i, test := range tests {
		var direction, moves = readInstruction(test.instruction)
		if direction != test.direction {
			t.Fatalf("failed test %d: expected direction %s, but got %s", i, test.direction, direction)
		}
		if moves != test.moves {
			t.Fatalf("failed test %d: expected moves %d, but got %d", i, test.moves, moves)
		}
	}
}

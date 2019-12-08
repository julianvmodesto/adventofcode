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
		{
			instruction: "R75",
			direction:   "R",
			moves:       75,
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

func TestGetDistanceToClosestIntersection(t *testing.T) {
	var tests = []struct {
		wire1    string
		wire2    string
		distance float64
	}{
		{
			wire1:    "R8,U5,L5,D3",
			wire2:    "U7,R6,D4,L4",
			distance: float64(6),
		},
		{
			wire1:    "R75,D30,R83,U83,L12,D49,R71,U7,L72",
			wire2:    "U62,R66,U55,R34,D71,R55,D58,R83",
			distance: float64(159),
		},
		{
			wire1:    "R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51",
			wire2:    "U98,R91,D20,R16,D67,R40,U7,R15,U6,R7",
			distance: float64(135),
		},
	}
	for i, test := range tests {
		var wire1, wire2 = NewPanels(test.wire1, test.wire2)
		var distance = wire1.GetDistanceToClosestIntersection(wire2)
		if distance != test.distance {
			t.Fatalf("failed test %d: expected distance %f, but got %f", i, test.distance, distance)
		}
	}
}

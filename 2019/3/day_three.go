package main

import (
	"log"
	"strconv"
	"strings"
)

type Wire struct {
	panel [][]bool
}

func (w *Wire) SetPath(instructions []string) {
	for _, i := range instructions {
		var direction, _ = readInstruction(i)
		switch direction {
		case "R":
		case "L":
		case "U":
		case "D":
		default:
		}
	}
}

func (w *Wire) GetDistance(i, j int) int {
	return 0
}

func readInstruction(instruction string) (string, int) {
	var parts = strings.Split(instruction, "")
	if len(parts) != 2 {
		log.Fatalf("expected 2 parts, but got %v parts for instruction %s", parts, instruction)
	}
	var direction = parts[0]
	if direction != "R" && direction != "L" && direction != "U" && direction != "D" {
		log.Fatalf("expected valid direction, but got %s", direction)
	}
	moves, err := strconv.Atoi(parts[1])
	if err != nil {
		log.Fatalf("expected valid number of moves, but got %s: %v", parts[1], err)
	}
	return direction, moves
}

func getPanelMaxSize(input string) int {
	var instructions = strings.Split(input, ",")
	var size int
	for _, i := range instructions {
		_, moves := readInstruction(i)
		size += moves
	}
	return size
}

func main() {

}

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	l := location{0, 0, 0}
	instructions := getInstructions(`R3, L5, R2, L1, L2, R5, L2, R2, L2, L2, L1,
	R2, L2, R4, R4, R1, L2, L3, R3, L1, R2, L2, L4, R4, R5, L3, R3, L3, L3, R4,
	R5, L3, R3, L5, L1, L2, R2, L1, R3, R1, L1, R187, L1, R2, R47, L5, L1, L2,
	R4, R3, L3, R3, R4, R1, R3, L1, L4, L1, R2, L1, R4, R5, L1, R77, L5, L4, R3,
	L2, R4, R5, R5, L2, L2, R2, R5, L2, R194, R5, L2, R4, L5, L4, L2, R5, L3, L2,
	L5, R5, R2, L3, R3, R1, L4, R2, L1, R5, L1, R5, L1, L1, R3, L1, R5, R2, R5,
	R5, L4, L5, L5, L5, R3, L2, L5, L4, R3, R1, R1, R4, L2, L4, R5, R5, R4, L2,
	L2, R5, R5, L5, L2, R4, R4, L4, R1, L3, R1, L1, L1, L1, L4, R5, R4, L4, L4,
	R5, R3, L2, L2, R3, R1, R4, L3, R1, L4, R3, L3, L2, R2, R2, R2, L1, L4, R3,
	R2, R2, L3, R2, L3, L2, R4, L2, R3,
	L4, R5, R4, R1, R5, R3`)

	for _, i := range instructions {
		l.doInstruction(i)
	}
	fmt.Printf("%+v\n", l)
}

func getInstructions(s string) []instruction {
	parts := strings.FieldsFunc(s, func(r rune) bool {
		switch r {
		case ' ', ',', '\t', '\n':
			return true
		}
		return false
	})
	instructions := make([]instruction, 10)
	for _, part := range parts {
		instructions = append(instructions, instruction{
			getIsLeft(part[0]),
			getSteps(part[1:len(part)]),
		})
	}
	return instructions
}

func getIsLeft(c byte) bool {
	return c == 'L'
}

func getSteps(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println(err)
	}
	return i
}

type location struct {
	x, y, dir int
}

type instruction struct {
	isLeft bool
	steps  int
}

func (l *location) doInstruction(i instruction) {
	l.turn(i.isLeft)
	l.walk(i.steps)
}

func (l *location) turn(isLeft bool) {
	// dir := 0 N, 1 E, 2 S, 3 W
	if isLeft {
		l.dir = l.dir - 1
	} else {
		l.dir = l.dir + 1
	}
	if l.dir < 0 {
		l.dir = 3
	}
	if l.dir > 3 {
		l.dir = 0
	}
}

func (l *location) walk(steps int) {
	if l.dir == 0 {
		l.x = l.x + steps
	} else if l.dir == 1 {
		l.y = l.y + steps
	} else if l.dir == 2 {
		l.x = l.x - steps
	} else if l.dir == 3 {
		l.y = l.y - steps
	}
}

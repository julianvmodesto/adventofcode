package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

type Instruction struct {
	Name  string
	Value int
}

func ParseInstruction(line string) (*Instruction, error) {
	i := Instruction{}
	parts := strings.Split(line, " ")
	i.Name = parts[0]
	v, err := strconv.Atoi(parts[1])
	if err != nil {
		return nil, fmt.Errorf("failed to parse value %q in line %q: %v", parts[1], line, err)
	}
	i.Value = v
	return &i, nil
}

func Solve(instructions []*Instruction) (int, int) {
	visited := make(map[int]bool)
	var i, acc, loop int
	for i < len(instructions) {
		if visited[i] {
			break
		}
		loop = i
		visited[i] = true
		instruction := instructions[i]
		switch instruction.Name {
		case "acc":
			acc += instruction.Value
			i++
		case "jmp":
			i += instruction.Value
		case "nop":
			i++
		default:
			log.Printf("unknown instruction %q in line %d\n", instruction.Name, i)
			break
		}
	}
	return acc, loop
}

func main() {
	if len(os.Args) < 2 {
		log.Println("missing input file, provide filename")
		return
	}
	data, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		log.Fatalf("failed to read file %s", os.Args[1])
	}
	var lines = strings.Split(string(data), "\n")
	var instructions []*Instruction
	for i, line := range lines {
		if len(line) == 0 {
			continue
		}
		instruction, err := ParseInstruction(line)
		if err != nil {
			log.Printf("failed to parse line %d %q: %v\n", i, line, err)
			break
		}
		instructions = append(instructions, instruction)
	}

	acc, loop := Solve(instructions)
	log.Printf("acc part 1: %d\n", acc)
	log.Printf("loop at %d: %q\n", loop, lines[loop])

	for i, instruction := range instructions {
		old := *instruction
		if instruction.Name == "nop" {
			instruction.Name = "jmp"
		} else if instruction.Name == "jmp" {
			instruction.Name = "nop"
		}
		acc, end := Solve(instructions)
		if end == len(instructions)-1 {
			log.Printf("acc part 2: %d, got %d, want %d\n", acc, end, len(instructions)-1)
			break
		}
		instructions[i] = &old
	}

}

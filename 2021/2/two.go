package main

import (
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func one(ins []int, fwd int) int {
	var depth int
	for _, i := range ins {
		depth += i
	}
	return depth * fwd
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
	var ins []int
	var fwd int

	var aim int
	var d int
	var h int
	for i, line := range lines {
		if line == "" {
			continue
		}

		parts := strings.Split(line, " ")
		if len(parts) != 2 {
			log.Fatalf("failed to parse line %d, got %d parts, '%s'\n", i, len(parts), line)
		}

		num, err := strconv.ParseInt(parts[1], 10, 64)
		if err != nil {
			log.Fatalf("failed to parse line %d, '%s'\n", i, line)
		}
		n := int(num)

		switch parts[0] {
		case "forward":
			fwd += n
			h += n
			d += n * aim
		case "up":
			ins = append(ins, n*-1)
			aim -= n
		case "down":
			ins = append(ins, n)
			aim += n
		}
	}

	part1 := one(ins, fwd)
	log.Printf("answer to part 1: %d", part1)

	log.Printf("answer to part 2: %d", d*h)
}

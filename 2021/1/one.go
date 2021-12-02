package main

import (
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func largerMeasurements(m []int) int {
	var ret int
	for i := 1; i < len(m); i++ {
		if m[i] > m[i-1] {
			ret++
		}
	}
	return ret
}

func largerMeasurementsSliding(m []int) int {
	var ret int
	last := m[0] + m[1] + m[2]
	cur := last
	for i := 3; i < len(m); i++ {
		cur -= m[i-3]
		cur += m[i]
		if cur > last {
			ret++
		}
		last = cur
	}
	return ret
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
	var nums []int
	for i, line := range lines {
		if line == "" {
			continue
		}
		num, err := strconv.ParseInt(line, 10, 64)
		if err != nil {
			log.Fatalf("failed to parse line %d, '%s'\n", i, line)
		}
		nums = append(nums, int(num))
	}

	part1 := largerMeasurements(nums)
	log.Printf("answer to part 1: %d", part1)
	part2 := largerMeasurementsSliding(nums)
	log.Printf("answer to part 2: %d", part2)
}

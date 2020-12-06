package main

import (
	"io/ioutil"
	"log"
	"os"
	"strings"
)

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

	var part1 int
	var part2 int
	answered1 := make([]bool, 26)
	first := true
	var answered2 []bool
	for _, line := range lines {
		if len(line) == 0 {
			first = true
			for i, a := range answered1 {
				if a {
					part1++
				}
				answered1[i] = false
			}
			for i, a := range answered2 {
				if a {
					part2++
				}
				answered2[i] = false
			}
			continue
		}
		curr := make([]bool, 26)
		for _, r := range line {
			answered1[r-'a'] = true
			curr[r-'a'] = true
		}
		if first {
			answered2 = append([]bool{}, curr...) // copy
			first = false
		}
		for i := range answered2 {
			answered2[i] = answered2[i] && curr[i]
		}
	}
	log.Printf("part 1: %d\n", part1)
	log.Printf("part 2: %d\n", part2)
}

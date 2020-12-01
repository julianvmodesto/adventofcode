package main

import (
	"io/ioutil"
	"log"
	"os"
	"strconv"
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

	var answer int64
	nums := make(map[int64]int)
	for i, line := range lines {
		if line == "" {
			continue
		}
		num, err := strconv.ParseInt(line, 10, 64)
		if err != nil {
			log.Fatalf("failed to parse line %d, '%s'\n", i, line)
		}
		if _, ok := nums[2020-num]; ok {
			answer = (2020 - num) * num
		}

		nums[num]++
	}
	log.Printf("answer: %d", answer)

	for k1, _ := range nums {
		nums[k1]--
		for k2, _ := range nums {
			nums[k2]--
			if c, _ := nums[2020-k1-k2]; c > 0 {
				answer = k1 * k2 * (2020 - k1 - k2)
			}
			nums[k2]++
		}
		nums[k1]++
	}
	log.Printf("answer: %d", answer)
}

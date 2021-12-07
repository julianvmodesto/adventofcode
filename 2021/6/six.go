package main

import (
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

/*
for i := 0; i < 80; i++
	newFish
	for f in fish
		if fish[f] == 0
			newFish = append(newFish, 8)
		else
			fish[f]--

count fish

fish 0 1 2 3 4 5 6 7 8

for 256 days
	newFish []int
	if i == 0
		newFish[6] += fish[i]
		newFish[8] += fish[i]
	newFish[i] = fish[i-1]
*/
func partOne(fish []int, days int) int {
	log.Printf("fish: %v", fish)
	log.Printf("day 0 fish count: %d", len(fish))
	for d := 1; d <= days; d++ {
		var newFish []int
		for f := range fish {
			if fish[f] == 0 {
				fish[f] = 6
				newFish = append(newFish, 8)
			} else {
				fish[f]--
			}
		}
		fish = append(fish, newFish...)
		log.Printf("day %d fish count: %d", d, len(fish))
	}
	return len(fish)
}

func count(fishC []int) int {
	var ret int
	for _, c := range fishC {
		ret += c
	}
	return ret
}

func partTwo(fish []int, days int) int {
	log.Printf("fish: %v", fish)
	fishC := make([]int, 9)
	for _, f := range fish {
		fishC[f]++
	}
	log.Printf("fish counts: %v", fishC)
	log.Printf("day 0 fish count: %d", count(fishC))

	for d := 1; d <= days; d++ {
		newFishC := make([]int, 9)
		for f := range fishC {
			if f == 0 {
				newFishC[6] += fishC[f]
				newFishC[8] += fishC[f]
			} else {
				newFishC[f-1] += fishC[f]
			}
		}
		fishC = newFishC
		log.Printf("day %d fish count: %d", d, count(fishC))
	}

	return count(fishC)
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
	var parts = strings.Split(strings.Trim(string(data), "\n"), ",")
	var fish []int
	for _, part := range parts {
		if part == "" {
			continue
		}

		num, err := strconv.ParseInt(part, 10, 64)
		if err != nil {
			log.Fatalf("failed to parse part '%s'\n", part)
		}

		fish = append(fish, int(num))
	}

	partTwo(fish, 256)
}

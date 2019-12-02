package main

import (
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

var cache map[int64]*int64 = make(map[int64]*int64)

func GetModuleFuel(mass int64) int64 {
	if cache[mass] != nil {
		return *cache[mass]
	}
	fuel := getModuleFuel(mass)
	cache[mass] = &fuel
	return fuel
}

func getModuleFuel(mass int64) int64 {
	return mass/3 - 2
}

func GetModuleFuelRecursive(mass, sum int64) int64 {
	fuel := GetModuleFuel(mass)
	if fuel <= 0 {
		return sum
	}
	return sum + fuel + GetModuleFuelRecursive(fuel, sum)
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
	var sum int64
	var sumFuel int64
	for i, line := range lines {
		if line == "" {
			log.Printf("skipping line %d", i)
			continue
		}
		mass, err := strconv.ParseInt(line, 10, 64)
		if err != nil {
			log.Fatalf("failed to parse line %d, '%s'\n", i, line)
		}
		fuel := GetModuleFuel(mass)
		fuelRecursive := GetModuleFuelRecursive(mass, 0)
		sum += fuel
		sumFuel += fuelRecursive
	}
	log.Printf("module mass sum part 1: %d", sum)
	log.Printf("module mass sum part 2: %d", sumFuel)
}

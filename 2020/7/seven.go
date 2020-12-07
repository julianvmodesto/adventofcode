package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

type Bag struct {
	Name     string
	Contains map[string]int
	visited  bool
	memo     bool
}

func ParseLine(line string) (*Bag, error) {
	bag := Bag{}
	parts := strings.Split(line, " contain ")
	bag.Name = strings.TrimSuffix(parts[0], " bags")

	if parts[1] == "no other bags." {
		return &bag, nil
	}

	bag.Contains = make(map[string]int)
	bags := strings.Split(
		strings.TrimSuffix(parts[1], "."),
		",",
	)
	for _, b := range bags {
		b = strings.TrimSuffix(b, "bags")
		b = strings.TrimSuffix(b, "bag")
		b = strings.TrimSpace(b)
		parts = strings.Split(b, " ")

		count, err := strconv.Atoi(parts[0])
		if err != nil {
			return nil, fmt.Errorf("failed to parse bag count in line %q: %v", line, err)
		}
		name := strings.Join(parts[1:], " ")
		bag.Contains[name] = count
	}

	return &bag, nil
}

func Contains(head *Bag, find string, bags map[string]*Bag, visitedMemo map[string]bool) bool {
	if head == nil {
		return false
	}

	if memo, visited := visitedMemo[head.Name]; visited {
		return memo
	}
	visitedMemo[head.Name] = false

	if head.Name == find {
		visitedMemo[head.Name] = true
		return true
	}

	for in := range head.Contains {
		bag, ok := bags[in]
		if !ok {
			log.Printf("failed to find bag %q\n", in)
			continue
		}
		if Contains(bag, find, bags, visitedMemo) {
			visitedMemo[head.Name] = true
			return true
		}
	}
	return false
}

func Count(head *Bag, bags map[string]*Bag) int {
	count := 0
	for in, c := range head.Contains {
		bag, ok := bags[in]
		if !ok {
			log.Printf("failed to find bag %q\n", in)
			continue
		}
		count += c * (1 + Count(bag, bags))
	}
	return count
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
	bags := make(map[string]*Bag)
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}

		bag, err := ParseLine(line)
		if err != nil {
			log.Fatalf("failed to parse line %q: %v", line, err)
			break
		}
		bags[bag.Name] = bag
	}

	var count1 int
	visitedMemo1 := make(map[string]bool)
	for _, bag := range bags {
		if bag.Name != "shiny gold" && Contains(bag, "shiny gold", bags, visitedMemo1) {
			count1++
		}
	}
	log.Printf("count part 1: %d of %d\n", count1, len(bags))

	count2 := Count(bags["shiny gold"], bags)
	log.Printf("count part 2: %d\n", count2)
}

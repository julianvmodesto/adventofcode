package main

import (
	"io/ioutil"
	"log"
	"os"
	"sort"
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
	var nums []int
	for i, line := range lines {
		if len(line) == 0 {
			continue
		}
		num, err := strconv.Atoi(line)
		if err != nil {
			log.Printf("failed to parse line %d %q\n", i, line)
			break
		}
		nums = append(nums, num)
	}

	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	var ones, threes int
	for i := 0; i <= len(nums); i++ {
		var diff int
		if i == 0 {
			diff = nums[i]
		} else if i < len(nums) {
			diff = nums[i] - nums[i-1]
		} else {
			diff = 3
		}
		switch diff {
		case 1:
			ones++
		case 3:
			threes++
		default:
			log.Printf("unexpected difference %d between nums[%d]=%d, nums[%d]=%d in line %q\n", diff, i, nums[i], i-1, nums[i-1], lines[i])
			break
		}
	}

	log.Printf("part 1: %d\n", ones*threes)
}

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

	all := make([]int, 0)
	stack := make([]int, 0)
	nums := make(map[int]int)
	var want int
	for i, line := range lines {
		if len(line) == 0 {
			continue
		}
		x, err := strconv.Atoi(line)
		if err != nil {
			log.Fatalf("failed to parse line %d %q: %v", i, line, err)
			break
		}

		if i > 25 {
			var found bool
			for num, _ := range nums {
				c := nums[x-num]
				if (x-num == num && c > 1) || c > 0 {
					found = true
					break
				}
			}
			if !found {
				log.Printf("part 1: %d\n", x)
				want = x
			}
		}

		all = append(all, x)
		stack = append(stack, x)
		nums[x]++
		if len(stack) > 25 {
			nums[stack[0]]--
			if nums[stack[0]] == 0 {
				delete(nums, stack[0])
			}
			stack = stack[1:]
		}
	}

	var left, right int
	var sum int
	for sum != want && right < len(all) {
		sum += all[right]
		for sum > want && left < right {
			sum -= all[left]
			left++
		}
		if sum == want {
			break
		}
		right++
	}
	min, max := all[left], all[left]
	for i := left + 1; i <= right; i++ {
		if all[i] < min {
			min = all[i]
		}
		if all[i] > max {
			max = all[i]
		}
	}
	log.Printf("part 2: left %d, right %d, min %d, max %d, ans %d\n", left, right, min, max, min+max)
	log.Printf("nums len %d, input len %d\n", len(all), len(lines)-1)
}

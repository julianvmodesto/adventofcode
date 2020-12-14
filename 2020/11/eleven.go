package main

import (
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func Model(seats [][]string) ([][]string, bool) {

	var ret [][]string
	same := true
	spots := [][]int{
		[]int{-1, -1},
		[]int{-1, 0},
		[]int{-1, 1},
		[]int{0, -1},
		[]int{0, 1},
		[]int{1, -1},
		[]int{1, 0},
		[]int{1, 1},
	}

	for row := range seats {
		var newRow []string
		for col := range seats[row] {
			if seats[row][col] == "." {
				newRow = append(newRow, ".")
				continue
			}

			var occupied int
			for _, spot := range spots {
				if row+spot[0] < 0 || row+spot[0] >= len(seats) {
					continue
				}
				if col+spot[1] < 0 || col+spot[1] >= len(seats[row]) {
					continue
				}
				if seats[row+spot[0]][col+spot[1]] == "#" {
					occupied++
				}
			}

			if seats[row][col] == "L" && occupied == 0 {
				newRow = append(newRow, "#")
				same = false
			} else if seats[row][col] == "#" && occupied >= 4 {
				newRow = append(newRow, "L")
				same = false
			} else {
				newRow = append(newRow, seats[row][col])
			}
		}
		ret = append(ret, newRow)
	}

	return ret, same
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
	var seats [][]string
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}

		row := make([]string, 0)
		for _, s := range strings.Split(line, "") {
			row = append(row, s)
		}
		seats = append(seats, row)
	}

	var same bool

	for !same {
		seats, same = Model(seats)
	}

	var count int
	for _, row := range seats {
		for _, seat := range row {
			if seat == "#" {
				count++
			}
		}
	}
	log.Printf("part 1: %d occupied seats\n", count)
}

package main

import (
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strings"
)

func GetSeatID(letters []string) int {
	_, _, seatID := GetRowColSeatID(letters)
	return seatID
}

func GetRowColSeatID(letters []string) (int, int, int) {
	var seatID int
	rowLo := 0
	rowHi := 127
	colLo := 0
	colHi := 7
	var row, col int
	for i, l := range letters {
		switch l {
		case "F":
			if i == 6 {
				row = (rowHi - 1)
				seatID = (rowHi - 1) * 8
				//log.Printf("row %d\n", seatID)
			} else {
				rowHi = rowHi - (rowHi-rowLo)/2 - 1
			}
		case "B":
			if i == 6 {
				row = (rowLo + 1)
				seatID = (rowLo + 1) * 8
				//log.Printf("row %d\n", seatID)
			} else {
				rowLo = rowLo + (rowHi-rowLo)/2 + 1
			}
		case "L":
			if i == 9 {
				col = colHi - 1
				seatID = seatID + (colHi - 1)
				//log.Printf("col %d\n", colHi-1)
			} else {
				colHi = colHi - (colHi-colLo)/2 - 1
			}
		case "R":
			if i == 9 {
				col = colLo + 1
				seatID = seatID + (colLo + 1)
				//log.Printf("col %d\n", colLo+1)
			} else {
				colLo = colLo + (colHi-colLo)/2 + 1
			}
		default:
			log.Printf("unknown dir %q in line %v\n", l, letters)
		}
	}
	return row, col, seatID
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
	var lo int
	var hi int
	var seats []int
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		letters := strings.Split(line, "")
		_, _, seatID := GetRowColSeatID(letters)
		if seatID > hi {
			hi = seatID
		}
		if lo == 0 {
			lo = seatID
		} else if lo > seatID {
			lo = seatID
		}
		seats = append(seats, seatID)
	}
	log.Printf("highest: %d\n", hi)

	sort.Slice(seats, func(i, j int) bool {
		return seats[i] < seats[j]
	})

	for i := range seats {
		if i > 0 && seats[i-1]+1 != seats[i] {
			log.Printf("seat: %d\n", seats[i]-1)
			break
		}
	}
}

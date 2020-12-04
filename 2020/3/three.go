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
	trees1 := 0
	trees3 := 0
	trees5 := 0
	trees7 := 0
	trees12 := 0
	col1 := 0
	col3 := 0
	col5 := 0
	col7 := 0
	row := 0
	for ; row < len(lines)-1; row++ {
		if strings.Compare(string(lines[row][col1]), "#") == 0 {
			trees1++
		}
		if strings.Compare(string(lines[row][col3]), "#") == 0 {
			trees3++
		}
		if strings.Compare(string(lines[row][col5]), "#") == 0 {
			trees5++
		}
		if strings.Compare(string(lines[row][col7]), "#") == 0 {
			trees7++
		}
		if row%2 == 0 && strings.Compare(string(lines[row][col1]), "#") == 0 {
			trees12++
		}
		col1 = (col1 + 1) % (len(lines[row]))
		col3 = (col3 + 3) % (len(lines[row]))
		col5 = (col5 + 5) % (len(lines[row]))
		col7 = (col7 + 7) % (len(lines[row]))
	}
	log.Printf("part 1 trees: %d\n", trees3)
	log.Printf("part 2 trees: %d\n", trees1*trees3*trees5*trees7*trees12)
	log.Printf("rows %d, total %d\n", row, len(lines)-1)
}

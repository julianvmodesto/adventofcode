package main

import (
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

type Board [][]int

var zero bool

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
	/*
			read first line

			read each board
				read each row
				add to last board
				if board has 5 rows
				add board to boards
				create new board

		    nums
			boards

			for each num in nums
				for each board in boards
					for each row in board
						for each col in row
							mark num
					if checkWin(board)
						return calculateScore(board, last num)

			mark(row, col, board)
				board[row][col] *= -1
			checkWin(board)
				for row in board
					if row > 0 break
					if row == len(board) return true
				for col in board
					if col > 0 break
					if col == len(board) return true
			calculateScore(board, last=nums[len(nums)-1])
			  sum unmarked numbers
			  sum
			  for row in board
			  multiply by last
	*/

	var nums []int
	var boards []Board
	var board Board
	for i, line := range lines {
		if line == "" {
			continue
		}
		if i == 0 {
			parts := strings.Split(line, ",")
			for _, n := range parts {
				num, err := strconv.ParseInt(n, 10, 64)
				if err != nil {
					log.Fatalf("failed to parse line %d: '%s'", i, line)
				}
				nums = append(nums, int(num))
			}
			continue
		}

		// Parse boards
		var row []int
		parts := strings.Fields(line)
		for _, n := range parts {
			num, err := strconv.ParseInt(n, 10, 64)
			if err != nil {
				log.Fatalf("failed to parse line %d: '%s'", i, line)
			}
			row = append(row, int(num))
		}
		if len(row) != 5 {
			log.Fatalf("expected 5 rows, got %d in line %d: '%s'", len(row), i, line)
		}
		board = append(board, row)
		if len(board) == 5 {
			boards = append(boards, board)
			board = make(Board, 0)
		}
	}

	log.Printf("nums: %v", nums)
	log.Printf("board count: %d", len(boards))

	winningBoards := make([]bool, 100)
	for _, num := range nums {
		for b, board := range boards {
			if winningBoards[b] {
				continue
			}
		boardloop:
			for row := 0; row < 5; row++ {
				for col := 0; col < 5; col++ {
					if board[row][col] != num {
						continue
					}
					board[row][col] *= -1
					if num == 0 {
						zero = true
					}
					if checkWin(board) {
						score := calculateScore(board, num)
						log.Printf("winning score for board %d: %d", b, score)
						winningBoards[b] = true
					}
					break boardloop
				}
			}
		}
	}
}

func checkWin(board Board) bool {
	for row := 0; row < 5; row++ {
		for col := 0; col < 5; col++ {
			if board[row][col] > 0 {
				break
			}
			if board[row][col] == 0 && !zero {
				break
			}
			if col == 4 {
				log.Printf("winning row %d board: %v", row, board)
				return true
			}
		}
	}
	for col := 0; col < 5; col++ {
		for row := 0; row < 5; row++ {
			if board[row][col] > 0 {
				break
			}
			if board[row][col] == 0 && !zero {
				break
			}
			if row == 4 {
				log.Printf("winning col %d board: %v", col, board)
				return true
			}
		}
	}
	return false
}

func calculateScore(board Board, num int) int {
	var sum int
	for row := 0; row < 5; row++ {
		for col := 0; col < 5; col++ {
			if num := board[row][col]; num > 0 {
				sum += num
			}
		}
	}
	log.Printf("sum %d * num %d", sum, num)
	return sum * num
}

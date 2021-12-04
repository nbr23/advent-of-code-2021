package main

import (
	"adventofcodego/utils/inputs"
	"adventofcodego/utils/utils"
	"strings"
)

var BOARD_SIZE int = 5

func inSlice(value int64, slice []int64) bool {
	for _, v := range slice {
		if value == v {
			return true
		}
	}
	return false
}

func isBoardWinner(called []int64, board []int64) bool {
	winner := false

	// check horizontal matches
	for y := 0; y < BOARD_SIZE; y++ {
		winner = true
		for x := 0; x < BOARD_SIZE; x++ {
			if !inSlice(board[y*BOARD_SIZE+x], called) {
				winner = false
			}
		}
		if winner == true {
			return true
		}
	}

	// check vertical matches
	for x := 0; x < BOARD_SIZE; x++ {
		winner = true
		for y := 0; y < BOARD_SIZE; y++ {
			if !inSlice(board[y*BOARD_SIZE+x], called) {
				winner = false
			}
		}
		if winner == true {
			return true
		}
	}
	return false
}

func parseBoards(input string) (calls []int64, boards [][]int64) {
	str_inputs := strings.Split(input, "\n")
	strcalls := strings.Split(strings.TrimSpace(str_inputs[0]), ",")

	calls = make([]int64, len(strcalls))
	for i, v := range strcalls {
		calls[i] = inputs.ParseDecInt(v)
	}

	boards = [][]int64{}
	boardscount := -1
	for i := 1; i < len(str_inputs); i++ {
		if len(str_inputs[i]) == 0 {
			boardscount += 1
			boards = append(boards, []int64{})
		} else {
			for _, v := range strings.Split(str_inputs[i], " ") {
				if len(v) > 0 {
					boards[boardscount] = append(boards[boardscount], inputs.ParseDecInt(v))
				}
			}
		}
	}

	return calls, boards
}

func getWinnerBoard(calls []int64, boards [][]int64) (lastcalled int, winner []int64) {
	for i := range calls {
		for _, board := range boards {
			if isBoardWinner(calls[:i+1], board) {
				return i, board
			}
		}
	}
	return 0, nil
}

func sumWinnerBoard(board []int64, called []int64) (sum int64) {
	for _, v := range board {
		if !inSlice(v, called) {
			sum += v
		}
	}
	return sum
}

func part1(input string) interface{} {
	calls, boards := parseBoards(input)

	lastcalledindex, winner := getWinnerBoard(calls, boards)

	return calls[lastcalledindex] * sumWinnerBoard(winner, calls[:lastcalledindex+1])
}

func part2(input string) interface{} {
	return nil
}

func main() {
	var day int64 = 4

	utils.Solve(part1, part2, day)
}
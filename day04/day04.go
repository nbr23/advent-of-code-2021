package main

import (
	"adventofcodego/utils/inputs"
	"adventofcodego/utils/utils"
	"strings"
)

var DAY int = 4

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
		calls[i] = inputs.ParseDecInt64(v)
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
					boards[boardscount] = append(boards[boardscount], inputs.ParseDecInt64(v))
				}
			}
		}
	}

	return calls, boards
}

func getWinnerBoard(calls []int64, boards [][]int64) (lastcalledindex int, winner []int64) {
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

func getBoardsWinning(called []int64, boards [][]int64, boards_index []int64) (winningboards []int64) {
	for _, i := range boards_index {
		if isBoardWinner(called, boards[i]) {
			winningboards = append(winningboards, i)
		}
	}
	return winningboards
}

func part2(input string) interface{} {
	calls, boards := parseBoards(input)
	var lastcalledindex = 0
	var winner []int64 = nil

	winningboards := []int64{}
	for i := range boards {
		winningboards = append(winningboards, int64(i))
	}

	for i := range calls {
		new_winningboards := getBoardsWinning(calls[0:len(calls)-i], boards, winningboards)

		if len(winningboards) != len(new_winningboards) {
			for _, v := range winningboards {
				if !inSlice(v, new_winningboards) {
					lastcalledindex = len(calls) - i
					winner = boards[v]
					break
				}
			}
		}
		if winner != nil {
			break
		}
		winningboards = new_winningboards
	}
	return calls[lastcalledindex] * sumWinnerBoard(winner, calls[:lastcalledindex+1])
}

func main() {
	utils.Solve(part1, part2, DAY)
}

package main

import (
	// "adventofcodego/utils/inputs"
	"adventofcodego/utils/utils"
	"strings"
)

var opening = map[rune]bool{
	'{': true,
	'[': true,
	'<': true,
	'(': true,
}

var pairs = map[rune]rune{
	'{': '}',
	'[': ']',
	'<': '>',
	'(': ')',
	'}': '{',
	']': '[',
	'>': '<',
	')': '(',
}

var scores = map[rune]int{
	')': 3,
	']': 57,
	'}': 1197,
	'>': 25137,
}

func getScoreForLine(line string) int {
	stack := make([]rune, 0)
	for _, c := range line {
		if opening[c] {
			stack = append(stack, c)
		} else {
			if pairs[c] == stack[len(stack)-1] {
				stack = stack[:len(stack)-1]
			} else {
				return scores[c]
			}
		}
	}
	return 0
}

func part1(input string) interface{} {
	score := 0
	for _, line := range strings.Split(input, "\n") {
		score += getScoreForLine(line)
	}
	return score
}

func part2(input string) interface{} {
	return nil
}

func main() {
	var day int64 = 10

	utils.Solve(part1, part2, day)
}

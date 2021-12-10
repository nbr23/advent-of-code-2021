package main

import (
	// "adventofcodego/utils/inputs"
	"adventofcodego/utils/utils"
	"sort"
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

var scorescorrect = map[rune]int{
	')': 1,
	']': 2,
	'}': 3,
	'>': 4,
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

func getCorrectScoreForLine(line string) int {
	stack := make([]rune, 0)
	score := 0
	for _, c := range line {
		if opening[c] {
			stack = append(stack, c)
		} else {
			if pairs[c] == stack[len(stack)-1] {
				stack = stack[:len(stack)-1]
			}
		}
	}
	for i := len(stack) - 1; i >= 0; i-- {
		score = score*5 + scorescorrect[pairs[stack[i]]]
	}
	return score
}

func part2(input string) interface{} {
	scores := make([]int, 0)
	for _, line := range strings.Split(input, "\n") {
		if getScoreForLine(line) == 0 {
			scores = append(scores, getCorrectScoreForLine(line))
		}
	}
	sort.Slice(scores, func(i, j int) bool { return scores[i] > scores[j] })
	return scores[len(scores)/2]
}

func main() {
	var day int64 = 10

	utils.Solve(part1, part2, day)
}

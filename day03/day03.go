package main

import (
	"adventofcodego/utils/utils"
	"strings"
)

func binToDec(binput []int64) (res int64) {
	for _, c := range binput {
		res = res << 1
		res += c
	}
	return res
}

func sumColumns(input []string) []int64 {
	sums := make([]int64, len(input[0]))
	for col := range input[0] {
		for line := range input {
			sums[col] += int64(input[line][col] - '0')
		}
	}
	return sums
}

func mostLeastCommonBit(sums []int64, size int64) ([]int64, []int64) {
	mostCommon := make([]int64, len(sums))
	leastCommon := make([]int64, len(sums))
	for i, v := range sums {
		if v > size/2 {
			mostCommon[i] = 1
		} else {
			leastCommon[i] = 1
		}
	}
	return mostCommon, leastCommon
}

func part1(input string) interface{} {
	inputlist := strings.Split(strings.Trim(input, "\n"), "\n")
	most, least := mostLeastCommonBit(sumColumns(inputlist), int64(len(inputlist)))
	return binToDec(most) * binToDec(least)
}

func part2(input string) interface{} {
	return nil //len(input)
}

func main() {
	var day int64 = 3

	utils.Solve(part1, part2, day)
}

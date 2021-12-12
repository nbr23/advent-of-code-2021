package main

import (
	"adventofcodego/utils/utils"
	"strings"
)

var DAY int = 3

func binToDec(binput []int64) (res int64) {
	for _, c := range binput {
		res = res << 1
		res += c
	}
	return res
}

func strBinToDec(binput string) (res int64) {
	for _, c := range binput {
		res = res << 1
		res += int64(c - '0')
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

func filterOut(inputlist []string, position int, untie int) string {
	criteria := sumColumns(inputlist)
	inputlen := int64(len(inputlist))

	if len(inputlist) <= 1 {
		return inputlist[0]
	}
	newlist := make([]string, 0, len(inputlist))
	for _, s := range inputlist {
		if (criteria[position]*2 >= inputlen && int(s[position]-'0') == (untie&1)) || (criteria[position]*2 < inputlen && int(s[position]-'0') == (untie^1)) {
			newlist = append(newlist, s)
		}
	}
	return filterOut(newlist, position+1, untie)
}

func part2(input string) interface{} {
	inputlist := strings.Split(strings.Trim(input, "\n"), "\n")

	ox := filterOut(inputlist, 0, 1)
	carb := filterOut(inputlist, 0, 0)

	return strBinToDec(carb) * strBinToDec(ox)
}

func main() {
	utils.Solve(part1, part2, DAY)
}

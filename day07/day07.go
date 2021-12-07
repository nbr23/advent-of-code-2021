package main

import (
	// "adventofcodego/utils/inputs"
	"adventofcodego/utils/inputs"
	"adventofcodego/utils/utils"
	"sort"
	"strings"
)

func inputToIntList(input string) (result []int) {
	input = strings.TrimSpace(strings.Trim(input, "\n"))

	for _, i := range strings.Split(input, ",") {
		result = append(result, int(inputs.ParseDecInt(i)))
	}
	return result
}

func getMedian(positions []int) int {
	sort.Ints(positions)
	middle := len(positions) / 2

	return positions[middle]
}

func part1(input string) interface{} {
	result := 0
	positions := inputToIntList(input)
	sort.Ints(positions)
	median := getMedian(positions)
	for _, v := range positions {
		if median > v {
			result += median - v
		} else {
			result += v - median
		}
	}
	return result
}

func part2(input string) interface{} {
	return nil
}

func main() {
	var day int64 = 7

	utils.Solve(part1, part2, day)
}

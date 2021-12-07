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

func getMean(positions []int) int {
	sum := 0
	for _, v := range positions {
		sum += v
	}

	return (sum + 1) / len(positions)
}

func part1(input string) interface{} {
	result := 0
	positions := inputToIntList(input)
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

func cumDiff(a int, b int) int {
	diff := a - b
	return diff * (diff + 1) / 2
}

func part2(input string) interface{} {
	result := 0
	positions := inputToIntList(input)
	mean := getMean(positions)

	for _, v := range positions {
		if mean > v {
			result += cumDiff(mean, v)
		} else {
			result += cumDiff(v, mean)
		}
	}
	return result
}

func main() {
	var day int64 = 7

	utils.Solve(part1, part2, day)
}

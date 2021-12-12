package main

import (
	"adventofcodego/utils/inputs"
	"adventofcodego/utils/utils"
)

var DAY int = 1

func getIncreasedCount(input []int64) (increased int64) {
	for i, v := range input {
		if i != 0 && v > input[i-1] {
			increased++
		}
	}
	return increased
}

func sumSlice(slice []int64) (res int64) {
	for _, v := range slice {
		res += v
	}
	return res
}

func part1(input string) interface{} {
	inputint := inputs.InputToIntList(input)
	return getIncreasedCount(inputint)
}

func part2(input string) interface{} {
	inputint := inputs.InputToIntList(input)
	cumsums := make([]int64, 0, len(inputint))

	for i := 0; i < len(inputint)-2; i++ {
		cumsums = append(cumsums, sumSlice(inputint[i:i+3]))
	}
	return getIncreasedCount(cumsums)
}

func main() {
	utils.Solve(part1, part2, DAY)
}

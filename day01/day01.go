package main

import (
	"fmt"

	"adventofcodego/utils/inputs"
)

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

func part1(day int64) {
	input := inputs.GetInputInt64(day)
	fmt.Printf("Result 1: %v\n", getIncreasedCount(input))
}

func part2(day int64) {
	input := inputs.GetInputInt64(day)
	cumsums := make([]int64, 0, len(input))

	for i := 0; i < len(input)-2; i++ {
		cumsums = append(cumsums, sumSlice(input[i:i+3]))
	}
	fmt.Printf("Result 2: %v\n", getIncreasedCount(cumsums))
}

func main() {
	var day int64 = 1
	part1(day)
	part2(day)
}

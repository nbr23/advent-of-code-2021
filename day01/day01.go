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

func part1(input string) {
	inputint := inputs.InputToIntList(input)
	fmt.Printf("Result 1: %v\n", getIncreasedCount(inputint))
}

func part2(input string) {
	inputint := inputs.InputToIntList(input)
	cumsums := make([]int64, 0, len(inputint))

	for i := 0; i < len(inputint)-2; i++ {
		cumsums = append(cumsums, sumSlice(inputint[i:i+3]))
	}
	fmt.Printf("Result 2: %v\n", getIncreasedCount(cumsums))
}

func main() {
	var day int64 = 1
	input := inputs.GetInput(day)

	part1(input)
	part2(input)
}

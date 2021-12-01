package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getInput() string {
	data, _ := os.ReadFile("./input")
	return string(data)
}

func strListToIntList(input []string) []int64 {
	intlist := make([]int64, 0, len(input))
	for _, s := range input {
		i, _ := strconv.ParseInt(s, 10, 64)
		intlist = append(intlist, i)
	}
	return intlist
}

func getInputInt64() []int64 {
	return strListToIntList(strings.Split(getInput(), "\n"))
}

func getIncreasedCount(input []int64) (increased int64) {
	for i, v := range input {
		if i != 0 && v > input[i-1] {
			increased++
		}
	}
	return increased
}

func part1() {
	input := getInputInt64()
	fmt.Printf("Result 1: %v\n", getIncreasedCount(input))
}

func sumSlice(slice []int64) (res int64) {
	for _, v := range slice {
		res += v
	}
	return res
}

func part2() {
	input := getInputInt64()
	cumsums := make([]int64, 0, len(input))

	for i := 0; i < len(input)-2; i++ {
		cumsums = append(cumsums, sumSlice(input[i:i+3]))
	}
	fmt.Printf("Result 2: %v\n", getIncreasedCount(cumsums))
}

func main() {
	part1()
	part2()
}

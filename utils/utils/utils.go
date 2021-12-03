package utils

import (
	"adventofcodego/utils/inputs"
	"fmt"
)

type Resolver func(string) interface{}

func Solve(part1 Resolver, part2 Resolver, day int64) {
	input := inputs.GetInput(day)
	res1 := part1(input)
	fmt.Printf("Result 1: %v\n", res1)

	res2 := part2(input)

	fmt.Printf("Result 2: %v\n", res2)
}

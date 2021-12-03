package utils

import (
	"adventofcodego/utils/inputs"
	"flag"
	"fmt"
)

type Resolver func(string) interface{}

func Solve(part1 Resolver, part2 Resolver, day int64) {
	var input string

	test_input := flag.Bool("test", false, "If set, uses ./dayX/testinput as input instead of the user specific input")
	flag.Parse()

	if *test_input {
		input = string(inputs.ReadFile(fmt.Sprintf("./day%02d/testinput", day)))
	} else {
		input = inputs.GetInput(day)
	}

	res1 := part1(input)
	fmt.Printf("Result 1: %v\n", res1)

	res2 := part2(input)

	fmt.Printf("Result 2: %v\n", res2)
}

package utils

import (
	"adventofcodego/utils/inputs"
	"fmt"
	"os"
)

type Resolver func(string) interface{}

func Solve(part1 Resolver, part2 Resolver, day int) {
	var input string

	binput, err := os.ReadFile(fmt.Sprintf("./inputs/day%02d.txt", day))
	if err != nil {
		fmt.Println("Fetching input")
		input = inputs.GetInput(day)
		err := os.WriteFile(fmt.Sprintf("./inputs/day%02d.txt", day), []byte(input), 0700)
		if err != nil {
			fmt.Println(err)
		}
	} else {
		input = string(binput)
	}

	fmt.Printf("*** DAY %d ***\n", day)

	res1 := part1(input)
	fmt.Printf("Result 1: %v\n", res1)

	res2 := part2(input)

	fmt.Printf("Result 2: %v\n", res2)
}

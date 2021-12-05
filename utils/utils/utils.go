package utils

import (
	"adventofcodego/utils/inputs"
	"flag"
	"fmt"
	"os"
)

type Resolver func(string) interface{}

func Solve(part1 Resolver, part2 Resolver, day int64) {
	var input string

	test_input := flag.Bool("test", false, "If set, uses ./dayX/testinput.txt as input instead of the user specific input")
	flag.Parse()

	if *test_input {
		binput, err := os.ReadFile(fmt.Sprintf("./day%02d/testinput.txt", day))
		if err != nil {
			panic(err)
		}
		input = string(binput)
	} else {
		binput, err := os.ReadFile(fmt.Sprintf("./day%02d/input.txt", day))
		if err != nil {
			fmt.Println("Fetching input")
			input = inputs.GetInput(day)
			err := os.WriteFile(fmt.Sprintf("./day%02d/input.txt", day), []byte(input), 0700)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			input = string(binput)
		}
	}

	fmt.Printf("*** DAY %d ***\n", day)

	res1 := part1(input)
	fmt.Printf("Result 1: %v\n", res1)

	res2 := part2(input)

	fmt.Printf("Result 2: %v\n", res2)
}

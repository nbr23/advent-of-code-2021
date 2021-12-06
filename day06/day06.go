package main

import (
	"adventofcodego/utils/inputs"
	"adventofcodego/utils/utils"
	"strings"
)

func inputToIntList(input string) (result []int64) {
	input = strings.TrimSpace(strings.Trim(input, "\n"))

	for _, i := range strings.Split(input, ",") {
		result = append(result, inputs.ParseDecInt(i))
	}
	return result
}

func iterateFishCycle(fish []int64) []int64 {
	fish_count := len(fish)

	for i := 0; i < fish_count; i++ {
		if fish[i] == 0 {
			fish[i] = 6
			fish = append(fish, 8)
		} else {
			fish[i]--
		}
	}
	return fish
}

func part1(input string) interface{} {
	fish := inputToIntList(input)

	for i := 0; i < 80; i++ {
		fish = iterateFishCycle(fish)
	}
	return len(fish)
}

func part2(input string) interface{} {
	return nil
}

func main() {
	var day int64 = 6

	utils.Solve(part1, part2, day)
}

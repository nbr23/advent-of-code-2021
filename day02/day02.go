package main

import (
	"strconv"
	"strings"

	"adventofcodego/utils/utils"
)

func part1(input string) interface{} {
	x, y := 0, 0
	for _, s := range strings.Split(input, "\n") {
		instructions := strings.Split(s, " ")
		switch instructions[0] {
		case "forward":
			i, _ := strconv.Atoi(instructions[1])
			x += i
		case "down":
			i, _ := strconv.Atoi(instructions[1])
			y += i
		case "up":
			i, _ := strconv.Atoi(instructions[1])
			y -= i
		}
	}
	return x * y
}

func part2(input string) interface{} {
	x, y, aim := 0, 0, 0
	for _, s := range strings.Split(input, "\n") {
		instructions := strings.Split(s, " ")
		switch instructions[0] {
		case "forward":
			i, _ := strconv.Atoi(instructions[1])
			x += i
			y += i * aim
		case "down":
			i, _ := strconv.Atoi(instructions[1])
			aim += i
		case "up":
			i, _ := strconv.Atoi(instructions[1])
			aim -= i
		}
	}
	return x * y
}

func main() {
	var day int64 = 2

	utils.Solve(part1, part2, day)
}

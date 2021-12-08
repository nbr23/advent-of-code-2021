package main

import (
	// "adventofcodego/utils/inputs"
	"adventofcodego/utils/utils"
	"strings"
)

func getDigits(line string) ([]string, []string) {
	linesplit := strings.Split(line, " | ")
	return strings.Split(linesplit[0], " "), strings.Split(linesplit[1], " ")
}

func part1(input string) interface{} {
	count := 0
	for _, line := range strings.Split(input, "\n") {
		_, output := getDigits(line)
		for _, d := range output {
			if len(d) >= 2 && len(d) <= 4 || len(d) == 7 {
				count++
			}
		}
	}
	return count
}

func part2(input string) interface{} {
	return nil
}

func main() {
	var day int64 = 8

	utils.Solve(part1, part2, day)
}

package main

import (
	// "adventofcodego/utils/inputs"
	"adventofcodego/utils/utils"
	"strings"
)

var DAY int = 14

func parseSubstitutions(inputs []string) map[string]rune {
	subs := make(map[string]rune)
	for _, input := range inputs {
		splits := strings.Split(input, " -> ")
		subs[splits[0]] = rune(splits[1][0])
	}
	return subs
}

func parseInput(input string) (polymer []rune, subsitutions map[string]rune) {
	splits := strings.Split(input, "\n")
	return []rune(splits[0]), parseSubstitutions(splits[2:])
}

func minMaxSlice(slice map[rune]int) (int, int) {
	min := 100000000000000000
	max := 0

	for i := range slice {
		if slice[i] > max {
			max = slice[i]
		}
		if slice[i] < min {
			min = slice[i]
		}
	}
	return min, max
}

func part1(input string) interface{} {
	poly, subs := parseInput(input)

	counts := make(map[rune]int)
	for _, i := range poly {
		counts[i]++
	}
	for i := 0; i < 10; i++ {
		newpoly := make([]rune, 0)
		for j := 0; j < len(poly)-1; j++ {
			newpoly = append(newpoly, poly[j])
			insert := subs[string(poly[j:j+2])]
			counts[insert]++
			newpoly = append(newpoly, insert)

		}
		poly = append(newpoly, poly[len(poly)-1])
	}
	min, max := minMaxSlice(counts)
	return max - min
}

func part2(input string) interface{} {
	return nil
}

func main() {
	utils.Solve(part1, part2, DAY)
}

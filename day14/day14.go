package main

import (
	// "adventofcodego/utils/inputs"
	"adventofcodego/utils/utils"
	"math"
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

func minMaxSlice(slice []int) (int, int) {
	min := math.MaxInt
	max := 0

	for i := range slice {
		if slice[i] > max {
			max = slice[i]
		}
		if slice[i] < min && slice[i] != 0 {
			min = slice[i]
		}
	}
	return min, max
}

func part1(input string) interface{} {
	poly, subs := parseInput(input)

	counts := make([]int, 26)
	for _, i := range poly {
		counts[runeValue(i)]++
	}

	cache := make(map[state][]int)
	for j := 0; j < len(poly)-1; j++ {
		counts = addCounts(counts, subsitutionTree(poly[j], poly[j+1], subs, cache, 10))
	}

	min, max := minMaxSlice(counts)
	return max - min
}

func runeValue(r rune) int {
	return int(r - 'A')
}

func addCounts(l1 []int, l2 []int) []int {
	l3 := make([]int, len(l1))
	for i := range l1 {
		l3[i] = l1[i] + l2[i]
	}
	return l3
}

type state struct {
	depth int
	pair  string
}

func subsitutionTree(left rune, right rune, subs map[string]rune, cache map[state][]int, depth int) []int {
	if depth == 0 {
		return make([]int, 26)
	}
	insert := subs[string([]rune{left, right})]
	var c1 []int
	var c2 []int
	var ok bool
	if c1, ok = cache[state{depth: depth - 1, pair: string([]rune{left, insert})}]; !ok {
		c1 = subsitutionTree(left, insert, subs, cache, depth-1)
		cache[state{depth: depth - 1, pair: string([]rune{left, insert})}] = c1
	}
	if c2, ok = cache[state{depth: depth - 1, pair: string([]rune{insert, right})}]; !ok {
		c2 = subsitutionTree(insert, right, subs, cache, depth-1)
		cache[state{depth: depth - 1, pair: string([]rune{insert, right})}] = c2
	}
	counts := addCounts(c1, c2)
	counts[runeValue(insert)]++
	return counts
}

func part2(input string) interface{} {
	poly, subs := parseInput(input)

	counts := make([]int, 26)
	for _, i := range poly {
		counts[runeValue(i)]++
	}

	cache := make(map[state][]int)
	for j := 0; j < len(poly)-1; j++ {
		counts = addCounts(counts, subsitutionTree(poly[j], poly[j+1], subs, cache, 40))
	}

	min, max := minMaxSlice(counts)
	return max - min
}

func main() {
	utils.Solve(part1, part2, DAY)
}

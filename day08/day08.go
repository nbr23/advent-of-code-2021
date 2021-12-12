package main

import (
	// "adventofcodego/utils/inputs"
	"adventofcodego/utils/utils"
	"fmt"
	"sort"
	"strings"
)

var DAY int = 8

var DIGITS = []string{"abcefg", "cf", "acdeg", "acdfg", "bcdf", "abdfg", "abdefg", "acf", "abcdefg", "abcdfg"}

func strSort(str string) string {
	s := []byte(str)
	sort.Slice(s, func(i int, j int) bool { return s[i] < s[j] })
	return string(s)
}

func strSliceSort(strs []string) []string {
	res := make([]string, 0)
	for _, s := range strs {
		res = append(res, strSort(s))
	}
	return res
}

func getDigits(line string) ([]string, []string) {
	linesplit := strings.Split(line, " | ")
	return strSliceSort(strings.Split(linesplit[0], " ")), strSliceSort(strings.Split(linesplit[1], " "))
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

type ituple struct {
	x int
	y int
}

func computeCommon(digits []string) map[ituple]string {
	common := make(map[ituple]string, 0)
	for i := range digits {
		for j := range digits {
			if j == i {
				continue
			}
			common[ituple{i, j}] = ""
			for _, c := range digits[i] {
				if strings.Contains(digits[j], string(c)) {
					common[ituple{i, j}] = fmt.Sprintf("%s%c", common[ituple{i, j}], c)
				}
			}
		}
	}
	return common
}

func getFootprint(commons map[ituple]string, pos int) (footprint int64) {
	i := 1
	for k, v := range commons {
		if k.x == pos {
			footprint += int64(len(v))
			i++
		}
	}
	return footprint
}

func part2(input string) interface{} {
	var result int64 = 0
	commons := computeCommon(DIGITS)
	footprintmap := make(map[int64]int)
	for i := 0; i < 10; i++ {
		footprintmap[getFootprint(commons, i)] = i
	}

	for _, line := range strings.Split(input, "\n") {
		var line_sum int64 = 0
		defs, display := getDigits(line)
		commons = computeCommon(defs)
		line_map := make(map[string]int)
		for i := 0; i < 10; i++ {
			fp := getFootprint(commons, i)
			line_map[defs[i]] = footprintmap[fp]
		}
		for _, s := range display {
			line_sum = line_sum*10 + int64(line_map[s])
		}
		result += line_sum
	}
	return result
}

func main() {
	utils.Solve(part1, part2, DAY)
}

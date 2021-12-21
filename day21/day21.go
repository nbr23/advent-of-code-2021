package main

import (
	// "adventofcodego/utils/inputs"
	"adventofcodego/utils/inputs"
	"adventofcodego/utils/utils"
	"strings"
	//	"github.com/pkg/profile"
)

var DAY int = 21

var DIEROLLS = 0
var DIEVAL = 1

func rollDie(max int) int {
	roll := DIEVAL
	DIEROLLS++
	DIEVAL++

	if DIEVAL > max {
		DIEVAL = 1
	}

	return roll
}

func parseStartingPositions(input []string) (int, int) {
	p1 := strings.Split(input[0], " ")
	p2 := strings.Split(input[1], " ")
	return inputs.ParseDecInt(p1[len(p1)-1]), inputs.ParseDecInt(p2[len(p2)-1])
}

func part1(input string) interface{} {
	p1, p2 := parseStartingPositions(strings.Split(input, "\n"))
	p1score, p2score := 0, 0

	for {
		for i := 0; i < 3; i++ {
			p1 += rollDie(100)
		}
		p1 = p1 % 10
		if p1 == 0 {
			p1 = 10
		}
		p1score += p1
		if p1score >= 1000 {
			return p2score * DIEROLLS
		}
		for i := 0; i < 3; i++ {
			p2 += rollDie(100)
		}
		p2 = p2 % 10
		if p2 == 0 {
			p2 = 10
		}
		p2score += p2
		if p2score >= 1000 {
			return p1score * DIEROLLS
		}
	}
}

func part2(input string) interface{} {
	return nil
}

func main() {
	// defer profile.Start().Stop()
	utils.Solve(part1, part2, DAY)
}

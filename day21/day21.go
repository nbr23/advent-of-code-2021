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

var POSSIBILITIES [][]int

func possibilities() [][]int {
	p := make([][]int, 3*3*3)
	c := 0
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			for k := 1; k <= 3; k++ {
				p[c] = []int{i, j, k}
				c++
			}
		}
	}
	return p
}

type state struct {
	p1score int
	p2score int
	p1pos   int
	p2pos   int
}

type wins struct {
	p1wins int
	p2wins int
}

var CACHE map[state]wins

func playTurn(p1score int, p2score int, pos1 int, pos2 int) (int, int) {
	if p1score >= 21 {
		return 1, 0
	}
	p1count := 0
	p2count := 0

	for _, d1 := range POSSIBILITIES {
		pos1b := pos1 + d1[0] + d1[1] + d1[2]
		pos1b = pos1b % 10
		if pos1b == 0 {
			pos1b = 10
		}
		p1bscore := p1score + pos1b

		if p1bscore >= 21 {
			p1count++
			continue
		}

		if s, ok := CACHE[state{p2score, p1bscore, pos2, pos1b}]; ok {
			p1count += s.p2wins
			p2count += s.p1wins
		} else {
			p2c, p1c := playTurn(p2score, p1bscore, pos2, pos1b)
			CACHE[state{p2score, p1bscore, pos2, pos1b}] = wins{p2c, p1c}
			p1count += p1c
			p2count += p2c
		}
	}
	return p1count, p2count
}

func part2(input string) interface{} {
	CACHE = make(map[state]wins)
	POSSIBILITIES = possibilities()
	p1, p2 := parseStartingPositions(strings.Split(input, "\n"))
	s1, s2 := playTurn(0, 0, p1, p2)

	if s1 > s2 {
		return s1
	}
	return s2
}

func main() {
	// defer profile.Start().Stop()
	utils.Solve(part1, part2, DAY)
}

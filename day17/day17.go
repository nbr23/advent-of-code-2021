package main

import (
	// "adventofcodego/utils/inputs"
	"adventofcodego/utils/inputs"
	"adventofcodego/utils/utils"
	"math"
	"regexp"
)

var DAY int = 17

type area struct {
	x1 int
	y1 int
	x2 int
	y2 int
}

type point struct {
	x int
	y int
}

func isInTarget(target area, x int, y int) bool {
	return x >= target.x1 && x <= target.x2 && y >= target.y1 && y <= target.y2
}

func parseTargetArea(input string) area {
	r, _ := regexp.Compile("target area: x=([-0-9]+)..([-0-9]+), y=([-0-9]+)..([-0-9]+)")
	match := r.FindStringSubmatch(input)
	return area{inputs.ParseDecInt(match[1]), inputs.ParseDecInt(match[3]), inputs.ParseDecInt(match[2]), inputs.ParseDecInt(match[4])}
}

func shootProbe(target area, vx int, vy int) int {
	position := point{0, 0}
	maxy := math.MinInt
	touched := false

	for {
		if position.x > target.x2 || position.y < target.y1 {
			if touched {
				return maxy
			} else {
				return math.MinInt
			}
		}
		if isInTarget(target, position.x, position.y) {
			touched = true
		}
		position.x += vx
		position.y += vy
		maxy = utils.IntMax(maxy, position.y)

		vy -= 1
		if vx > 0 {
			vx -= 1
		} else if vx < 0 {
			vx += 1
		}
	}
}

func part1(input string) interface{} {
	target := parseTargetArea(input)
	maxy := math.MinInt
	for vx := 0; vx < 1000; vx++ {
		for vy := 0; vy < 1000; vy++ {
			sh := shootProbe(target, vx, vy)
			maxy = utils.IntMax(maxy, sh)
		}
	}
	return maxy
}

func part2(input string) interface{} {
	target := parseTargetArea(input)
	count := 0
	for vx := 1; vx < 500; vx++ {
		for vy := -500; vy < 500; vy++ {
			sh := shootProbe(target, vx, vy)
			if sh != math.MinInt {
				count++
			}
		}
	}
	return count
}

func main() {
	utils.Solve(part1, part2, DAY)
}

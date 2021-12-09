package main

import (
	// "adventofcodego/utils/inputs"

	"adventofcodego/utils/utils"
	"fmt"
	"math"
	"strings"
)

var HEIGHT int
var WIDTH int

func getMap(input []string) []int {
	result := make([]int, 0)
	for _, line := range input {
		for _, c := range line {
			result = append(result, int(c-'0'))
		}
	}
	return result
}

func isLowest(heightmap []int, x int, y int) bool {
	min := 10

	for vx := math.Max(0, float64(x-1)); vx < math.Min(float64(x+2), float64(WIDTH)); vx++ {
		if int(vx) == x {
			continue
		}
		if getAtXY(heightmap, int(vx), int(y)) < min {
			min = getAtXY(heightmap, int(vx), int(y))
		}
	}
	for vy := math.Max(0, float64(y-1)); vy < math.Min(float64(y+2), float64(HEIGHT)); vy++ {
		if int(vy) == y {
			continue
		}
		if getAtXY(heightmap, int(x), int(vy)) < min {
			min = getAtXY(heightmap, int(x), int(vy))
		}
	}

	return min > getAtXY(heightmap, x, y)
}

func getAtXY(heightmap []int, x int, y int) int {
	return heightmap[x+y*WIDTH]
}

func part1(input string) interface{} {
	lines := strings.Split(input, "\n")
	mins := make([]int, 0)
	minsums := 0
	HEIGHT = len(lines)
	WIDTH = len(lines[0])
	hmap := getMap(lines)
	for x := 0; x < WIDTH; x++ {
		for y := 0; y < HEIGHT; y++ {
			if isLowest(hmap, x, y) {
				mins = append(mins, getAtXY(hmap, x, y)+1)
				minsums += getAtXY(hmap, x, y) + 1
			}
		}
	}
	fmt.Println(mins)
	return minsums
}

func part2(input string) interface{} {
	return nil
}

func main() {
	var day int64 = 9

	utils.Solve(part1, part2, day)
}

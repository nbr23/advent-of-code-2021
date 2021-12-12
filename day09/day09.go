package main

import (
	// "adventofcodego/utils/inputs"

	"adventofcodego/utils/utils"
	"math"
	"sort"
	"strings"
)

var DAY int = 9

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
	minsums := 0
	HEIGHT = len(lines)
	WIDTH = len(lines[0])
	hmap := getMap(lines)
	for x := 0; x < WIDTH; x++ {
		for y := 0; y < HEIGHT; y++ {
			if isLowest(hmap, x, y) {
				minsums += getAtXY(hmap, x, y) + 1
			}
		}
	}
	return minsums
}

type point struct {
	x int
	y int
}

func getBasinSize(heightmap []int, x int, y int, visited map[point]bool) int {
	here := getAtXY(heightmap, x, y)
	count := 0

	for vx := math.Max(0, float64(x-1)); vx < math.Min(float64(x+2), float64(WIDTH)); vx++ {
		if int(vx) == x || visited[point{x: int(vx), y: y}] {
			continue
		}
		if getAtXY(heightmap, int(vx), int(y)) >= here && getAtXY(heightmap, int(vx), int(y)) != 9 {
			visited[point{x: int(vx), y: y}] = true
			count += 1 + getBasinSize(heightmap, int(vx), y, visited)
		}
	}
	for vy := math.Max(0, float64(y-1)); vy < math.Min(float64(y+2), float64(HEIGHT)); vy++ {
		if int(vy) == y || visited[point{y: int(vy), x: x}] {
			continue
		}
		if getAtXY(heightmap, x, int(vy)) >= here && getAtXY(heightmap, x, int(vy)) != 9 {
			visited[point{y: int(vy), x: x}] = true
			count += 1 + getBasinSize(heightmap, x, int(vy), visited)
		}
	}
	return count
}

func part2(input string) interface{} {
	lines := strings.Split(input, "\n")
	hmap := getMap(lines)
	wells := make([]int, 0)
	for x := 0; x < WIDTH; x++ {
		for y := 0; y < HEIGHT; y++ {
			if isLowest(hmap, x, y) {
				wells = append(wells, 1+getBasinSize(hmap, x, y, make(map[point]bool)))
			}
		}
	}
	sort.Slice(wells, func(i int, j int) bool { return wells[i] > wells[j] })

	return wells[0] * wells[1] * wells[2]
}

func main() {
	utils.Solve(part1, part2, DAY)
}

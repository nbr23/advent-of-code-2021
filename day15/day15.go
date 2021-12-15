package main

import (
	// "adventofcodego/utils/inputs"

	"adventofcodego/utils/utils"
	"fmt"
	"strings"
)

var DAY int = 15
var HEIGHT int = 0
var WIDTH int = 0
var ACCU int = 0

type point struct {
	x int
	y int
}

func parseMatrix(input string) []int {
	lines := strings.Split(input, "\n")
	WIDTH = len(lines[0])
	HEIGHT = len(lines)
	matrix := make([]int, 0)

	for y := range lines {
		for x := range lines[y] {
			matrix = append(matrix, int(lines[y][x]-'0'))
		}
	}
	return matrix
}

func getMinCost(a, b int) int {
	if a == -1 || b < a {
		return b
	}
	return a
}

func computeCosts(matrix []int, size int) []int {
	costs := make([]int, HEIGHT*WIDTH)
	for x := 0; x < WIDTH*; x++ {
		for y := 0; y < HEIGHT; y++ {
			if x == 0 && y == 0 {
				costs[x+y*WIDTH] = 0
			} else if x == 0 {
				costs[x+y*WIDTH] = matrix[x+y*WIDTH] + costs[x+(y-1)*WIDTH]
			} else if y == 0 {
				costs[x+y*WIDTH] = matrix[x+y*WIDTH] + costs[x-1+y*WIDTH]
			} else {
				costs[x+y*WIDTH] = matrix[x+y*WIDTH] + getMinCost(costs[x-1+y*WIDTH], costs[x+(y-1)*WIDTH])
			}
		}
	}
	return costs
}

func printMatrix(matrix []int) {
	for y := 0; y < HEIGHT; y++ {
		fmt.Println(matrix[y*HEIGHT : y*HEIGHT+WIDTH])
	}
}

func part1(input string) interface{} {
	matrix := parseMatrix(input)
	costs := computeCosts(matrix)
	return costs[len(costs)-1]
}

func part2(input string) interface{} {
	return nil
}

func main() {
	utils.Solve(part1, part2, DAY)
}

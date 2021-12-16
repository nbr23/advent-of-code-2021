package main

import (
	// "adventofcodego/utils/inputs"

	"adventofcodego/utils/utils"
	"math"
	"strings"

	"github.com/pkg/profile"
)

var DAY int = 15
var HEIGHT int = 0
var WIDTH int = 0
var ACCU int = 0

type point struct {
	x int
	y int
}

type pointlist struct {
	next  *pointlist
	point point
	value int
}

func (l *pointlist) insertOrAdd(point point, val int) {
	current := l
	var insert *pointlist
	for {
		if current.next == nil {
			if insert == nil {
				current.next = &pointlist{next: current.next, point: point, value: val}
			} else {
				insert.next = &pointlist{next: insert.next, point: point, value: val}
			}
			return
		} else if current.point == point {
			current.value = val
			return
		} else if val < current.next.value && insert == nil {
			insert = current
		}
		current = current.next
	}
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
	if b < a {
		return b
	}
	return a
}

func Dijkstra(matrix []int, size int) []int {
	visited := make([]bool, HEIGHT*WIDTH*size*size)
	current := point{0, 0}
	costs, newmatrix := initCosts(size, matrix)

	active_nodes := &pointlist{next: nil, point: point{0, 0}, value: 0}

	for {
		for nx := utils.IntMax(0, current.x-1); nx < utils.IntMin(WIDTH*size, current.x+2); nx++ {
			for ny := utils.IntMax(0, current.y-1); ny < utils.IntMin(HEIGHT*size, current.y+2); ny++ {
				if nx == current.x && ny == current.y || nx != current.x && ny != current.y || visited[nx+ny*WIDTH*size] {
					continue
				}
				costs[nx+ny*WIDTH*size] = getMinCost(costs[nx+ny*WIDTH*size], newmatrix[nx+ny*WIDTH*size]+costs[current.x+current.y*WIDTH*size])
				active_nodes.insertOrAdd(point{nx, ny}, costs[nx+ny*WIDTH*size])
			}
		}

		// mark as visited
		visited[current.x+current.y*WIDTH*size] = true

		// search for next node
		if active_nodes == nil || visited[WIDTH*size*HEIGHT*size-1] {
			break
		}
		current = active_nodes.point
		active_nodes = active_nodes.next
	}
	return costs
}

func initCosts(size int, matrix []int) ([]int, []int) {
	costs := make([]int, HEIGHT*size*WIDTH*size)
	newmatrix := make([]int, HEIGHT*size*WIDTH*size)
	for y := 0; y < HEIGHT*size; y++ {
		for x := 0; x < WIDTH*size; x++ {
			costs[x+y*WIDTH*size] = math.MaxInt
			v := (matrix[(x%WIDTH)+(y%WIDTH)*WIDTH] + x/WIDTH + y/HEIGHT)
			newmatrix[x+y*WIDTH*size] = v%10 + v/10
		}
	}
	costs[0] = 0
	return costs, newmatrix
}

func part1(input string) interface{} {
	size := 1
	matrix := parseMatrix(input)
	costs := Dijkstra(matrix, size)
	return costs[WIDTH*HEIGHT*size*size-1]
}

func part2(input string) interface{} {
	size := 5
	matrix := parseMatrix(input)
	costs := Dijkstra(matrix, size)
	return costs[WIDTH*HEIGHT*size*size-1]
}

func main() {
	defer profile.Start().Stop()
	utils.Solve(part1, part2, DAY)
}

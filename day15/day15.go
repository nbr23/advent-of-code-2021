package main

import (
	// "adventofcodego/utils/inputs"

	"adventofcodego/utils/utils"
	"fmt"
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

func (l *pointlist) insert(point point, val int) {
	current := l
	for {
		if current.next == nil || val <= current.next.value {
			current.next = &pointlist{next: current.next, point: point, value: val}
			return
		}
		current = current.next
	}
}

func (l *pointlist) insertOrAdd(point point, val int) {
	current := l
	for {
		if current.next == nil || val < current.next.value {
			current.next = &pointlist{next: current.next, point: point, value: val}
			return
		}
		current = current.next
	}
}

func (l *pointlist) updateValue(point point, val int) bool {
	current := l
	for {
		if current == nil {
			return false
		}
		if current.point == point {
			current.value = val
			return true
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

func getCostXY(matrix []int, x int, y int) int {
	cost := (matrix[(x%WIDTH)+(y%WIDTH)*WIDTH] + x/WIDTH + y/HEIGHT)
	return cost%10 + cost/10
}

func computeCosts(matrix []int, size int, costs []int, visited []bool, x int, y int) {
	if visited[len(visited)-1] {
		return
	}
	visited[x+y*WIDTH*size] = true
	for nx := int(math.Max(0, float64(x-1))); nx < int(math.Min(float64(WIDTH*size), float64(x+2))); nx++ {
		for ny := int(math.Max(0, float64(y-1))); ny < int(math.Min(float64(HEIGHT*size), float64(y+2))); ny++ {
			if nx == x && ny == y || nx != x && ny != y || visited[nx+ny*WIDTH*size] {
				continue
			}
			costs[nx+ny*WIDTH*size] = getMinCost(costs[nx+ny*WIDTH*size], getCostXY(matrix, nx, ny)+costs[x+y*WIDTH*size])
			computeCosts(matrix, size, costs, visited, nx, ny)
		}
	}
}

func Dijkstra(matrix []int, size int) map[point]int {
	visited := make([]bool, HEIGHT*WIDTH*size*size)
	current := point{0, 0}
	costs := initCosts(size)

	active_nodes := &pointlist{next: nil, point: point{0, 0}, value: 0}

	for {
		for nx := int(math.Max(0, float64(current.x-1))); nx < int(math.Min(float64(WIDTH*size), float64(current.x+2))); nx++ {
			for ny := int(math.Max(0, float64(current.y-1))); ny < int(math.Min(float64(HEIGHT*size), float64(current.y+2))); ny++ {
				if nx == current.x && ny == current.y || nx != current.x && ny != current.y || visited[nx+ny*WIDTH*size] {
					continue
				}
				costs[point{nx, ny}] = getMinCost(costs[point{nx, ny}], getCostXY(matrix, nx, ny)+costs[point{current.x, current.y}])
				if !active_nodes.updateValue(point{nx, ny}, costs[point{nx, ny}]) {
					active_nodes.insert(point{nx, ny}, costs[point{nx, ny}])
				}
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

func printMatrix(matrix []int, size int) {
	for y := 0; y < HEIGHT*size; y++ {
		for x := 0; x < WIDTH*size; x++ {
			fmt.Printf("%v ", getCostXY(matrix, x, y))
		}
		fmt.Println()
	}
}

func printCosts(matrix map[point]int, size int) {
	for y := 0; y < HEIGHT*size; y++ {
		for x := 0; x < WIDTH*size; x++ {
			fmt.Printf("%d ", matrix[point{x, y}])
		}
		fmt.Println()
	}
}

func initCosts(size int) map[point]int {
	costs := make(map[point]int)
	for y := 0; y < HEIGHT*size; y++ {
		for x := 0; x < WIDTH*size; x++ {
			costs[point{x, y}] = math.MaxInt
		}
	}
	costs[point{0, 0}] = 0
	return costs
}

func part1(input string) interface{} {
	size := 1
	matrix := parseMatrix(input)
	costs := Dijkstra(matrix, size)
	return costs[point{WIDTH*size - 1, HEIGHT*size - 1}]
}

func part2(input string) interface{} {
	size := 5
	matrix := parseMatrix(input)
	costs := Dijkstra(matrix, size)
	return costs[point{WIDTH*size - 1, HEIGHT*size - 1}]
}

func main() {
	defer profile.Start().Stop()
	utils.Solve(part1, part2, DAY)
}

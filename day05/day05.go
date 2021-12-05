package main

import (
	"adventofcodego/utils/inputs"
	"adventofcodego/utils/utils"
	"math"
	"strings"
)

type vector struct {
	x1 int64
	y1 int64
	x2 int64
	y2 int64
}

type point struct {
	x int64
	y int64
}

func getVectorsFromInput(input string) (vectors []vector) {
	for _, line := range strings.Split(input, "\n") {
		sp := strings.Split(line, " ")
		v1 := strings.Split(sp[0], ",")
		v2 := strings.Split(sp[2], ",")
		vectors = append(vectors, vector{
			x1: inputs.ParseDecInt(v1[0]),
			y1: inputs.ParseDecInt(v1[1]),
			x2: inputs.ParseDecInt(v2[0]),
			y2: inputs.ParseDecInt(v2[1]),
		})
	}
	return vectors
}

func plotPoints(vectors []vector) (points map[point]int) {
	points = make(map[point]int)
	for _, v := range vectors {
		if v.y1 == v.y2 {
			for x := math.Min(float64(v.x1), float64(v.x2)); x <= math.Max(float64(v.x1), float64(v.x2)); x++ {
				points[point{x: int64(x), y: v.y1}]++
			}
		} else if v.x1 == v.x2 {
			for y := math.Min(float64(v.y1), float64(v.y2)); y <= math.Max(float64(v.y1), float64(v.y2)); y++ {
				points[point{x: v.x1, y: int64(y)}]++
			}
		}
	}
	return points
}

func part1(input string) interface{} {
	vectors := getVectorsFromInput(input)
	points := plotPoints(vectors)
	intersects := 0

	for p := range points {
		if points[p] > 1 {
			intersects++
		}
	}
	return intersects
}

func plotPointsDiagonals(vectors []vector) (points map[point]int) {
	points = make(map[point]int)
	for _, v := range vectors {
		if v.y1 == v.y2 {
			for x := math.Min(float64(v.x1), float64(v.x2)); x <= math.Max(float64(v.x1), float64(v.x2)); x++ {
				points[point{x: int64(x), y: v.y1}]++
			}
		} else if v.x1 == v.x2 {
			for y := math.Min(float64(v.y1), float64(v.y2)); y <= math.Max(float64(v.y1), float64(v.y2)); y++ {
				points[point{x: v.x1, y: int64(y)}]++
			}
		} else {
			ax, ay := int64(1), int64(1)
			x, y := v.x1, v.y1
			if v.x1 > v.x2 {
				ax = -1
			}
			if v.y1 > v.y2 {
				ay = -1
			}
			for {
				points[point{x: x, y: y}]++
				if x == v.x2 && y == v.y2 {
					break
				}
				x += ax
				y += ay
			}

		}
	}
	return points
}

func part2(input string) interface{} {
	vectors := getVectorsFromInput(input)
	points := plotPointsDiagonals(vectors)
	intersects := 0

	for p := range points {
		if points[p] > 1 {
			intersects++
		}
	}
	return intersects
}

func main() {
	var day int64 = 5

	utils.Solve(part1, part2, day)
}

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

func part2(input string) interface{} {
	return nil
}

func main() {
	var day int64 = 5

	utils.Solve(part1, part2, day)
}

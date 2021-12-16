package main

import (
	"adventofcodego/utils/inputs"
	"adventofcodego/utils/utils"
	"strings"
)

var DAY int = 5

type vector struct {
	x1 int
	y1 int
	x2 int
	y2 int
}

type point struct {
	x int
	y int
}

var HEIGHT int = 0
var WIDTH int = 0

func getVectorsFromInput(input string) (vectors []vector) {
	for _, line := range strings.Split(input, "\n") {
		sp := strings.Split(line, " ")
		v1 := strings.Split(sp[0], ",")
		v2 := strings.Split(sp[2], ",")
		x1 := inputs.ParseDecInt(v1[0])
		y1 := inputs.ParseDecInt(v1[1])
		x2 := inputs.ParseDecInt(v2[0])
		y2 := inputs.ParseDecInt(v2[1])
		vectors = append(vectors, vector{
			x1: x1,
			y1: y1,
			x2: x2,
			y2: y2,
		})
		WIDTH = utils.IntMax(WIDTH, utils.IntMax(x1, x2))
		HEIGHT = utils.IntMax(HEIGHT, utils.IntMax(y1, y2))
	}
	WIDTH++
	HEIGHT++
	return vectors
}

func plotPoints(vectors []vector) int {
	intersects := 0
	points := make([]int, WIDTH*HEIGHT)
	for _, v := range vectors {
		if v.y1 == v.y2 {
			for x := utils.IntMin(v.x1, v.x2); x <= utils.IntMax(v.x1, v.x2); x++ {
				points[int(x)+int(v.y1)*WIDTH]++
				if points[int(x)+int(v.y1)*WIDTH] == 2 {
					intersects++
				}
			}
		} else if v.x1 == v.x2 {
			for y := utils.IntMin(v.y1, v.y2); y <= utils.IntMax(v.y1, v.y2); y++ {
				points[int(v.x1)+int(y)*WIDTH]++
				if points[int(v.x1)+int(y)*WIDTH] == 2 {
					intersects++
				}
			}
		}
	}
	return intersects
}

func part1(input string) interface{} {
	vectors := getVectorsFromInput(input)

	return plotPoints(vectors)
}

func plotPointsDiagonals(vectors []vector) int {
	points := make([]int, WIDTH*HEIGHT)
	intersects := 0
	for _, v := range vectors {
		if v.y1 == v.y2 {
			for x := utils.IntMin(v.x1, v.x2); x <= utils.IntMax(v.x1, v.x2); x++ {
				points[int(x)+int(v.y1)*WIDTH]++
				if points[int(x)+int(v.y1)*WIDTH] == 2 {
					intersects++
				}
			}
		} else if v.x1 == v.x2 {
			for y := utils.IntMin(v.y1, v.y2); y <= utils.IntMax(v.y1, v.y2); y++ {
				points[v.x1+y*WIDTH]++
				if points[v.x1+y*WIDTH] == 2 {
					intersects++
				}
			}
		} else {
			ax, ay := 1, 1
			x, y := v.x1, v.y1
			if v.x1 > v.x2 {
				ax = -1
			}
			if v.y1 > v.y2 {
				ay = -1
			}
			for {
				points[x+y*WIDTH]++
				if points[x+y*WIDTH] == 2 {
					intersects++
				}
				if x == v.x2 && y == v.y2 {
					break
				}
				x += ax
				y += ay
			}

		}
	}
	return intersects
}

func part2(input string) interface{} {
	vectors := getVectorsFromInput(input)
	return plotPointsDiagonals(vectors)
}

func main() {
	utils.Solve(part1, part2, DAY)
}

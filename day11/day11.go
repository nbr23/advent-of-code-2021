package main

import (
	// "adventofcodego/utils/inputs"
	"adventofcodego/utils/utils"
	"math"
)

var HEIGHT int = 10
var WIDTH int = 10

type point struct {
	x int
	y int
}

func flashXY(matrix *[]int, flashed map[point]bool, x int, y int) int {
	flashcount := 0
	if (*matrix)[x+y*WIDTH] > 9 && !flashed[point{x: x, y: y}] {
		flashcount++
		flashed[point{x: x, y: y}] = true
		(*matrix)[x+y*WIDTH] = 0
		for nx := int(math.Max(0, float64(x-1))); nx < int(math.Min(float64(WIDTH), float64(x+2))); nx++ {
			for ny := int(math.Max(0, float64(y-1))); ny < int(math.Min(float64(HEIGHT), float64(y+2))); ny++ {
				if nx == x && ny == y || flashed[point{x: nx, y: ny}] {
					continue
				}
				(*matrix)[nx+ny*WIDTH]++
				if (*matrix)[nx+ny*WIDTH] > 9 {
					flashcount += flashXY(matrix, flashed, nx, ny)
				}
			}
		}
	}
	return flashcount
}

func getAtXY(matrix []int, x int, y int) int {
	return matrix[x+y*WIDTH]
}

func fillMatrix(matrix *[]int, input string) {
	for _, c := range input {
		if c == '\n' {
			continue
		}
		*matrix = append(*matrix, int(c-'0'))
	}
}

func increaseEach(matrix *[]int) []point {
	willFlash := make([]point, 0)
	for x := 0; x < WIDTH; x++ {
		for y := 0; y < HEIGHT; y++ {
			(*matrix)[x+y*WIDTH]++
			if (*matrix)[x+y*WIDTH] > 9 {
				willFlash = append(willFlash, point{x: x, y: y})
			}
		}
	}
	return willFlash
}

func part1(input string) interface{} {
	steps := 100
	flash_count := 0
	matrix := make([]int, 0)
	fillMatrix(&matrix, input)

	for i := 0; i < steps; i++ {
		willFlash := increaseEach(&matrix)
		flashed := make(map[point]bool)
		for _, flasher := range willFlash {
			flash_count += flashXY(&matrix, flashed, flasher.x, flasher.y)
		}
	}
	return flash_count
}

func part2(input string) interface{} {
	steps := 1000000
	matrix := make([]int, 0)
	fillMatrix(&matrix, input)

	for i := 0; i < steps; i++ {
		willFlash := increaseEach(&matrix)
		flashed := make(map[point]bool)

		for _, flasher := range willFlash {
			flashXY(&matrix, flashed, flasher.x, flasher.y)
		}

		if len(flashed) == 100 {
			return i + 1
		}
	}
	return 0
}

func main() {
	var day int64 = 11

	utils.Solve(part1, part2, day)
}

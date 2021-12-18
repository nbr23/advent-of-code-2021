package main

import (
	"adventofcodego/utils/characters"
	"adventofcodego/utils/inputs"
	"adventofcodego/utils/utils"

	"fmt"
	"strings"
)

var DAY int = 13

type dotmatrix struct {
	dots   []point
	width  int64
	height int64
}

type point struct {
	x int64
	y int64
}

func parseInput(input string) (matrix dotmatrix, folds []point) {
	matrix.dots = make([]point, 0)
	folds = make([]point, 0)
	parts := strings.Split(input, "\n\n")

	for _, d := range strings.Split(parts[0], "\n") {
		coords := strings.Split(d, ",")
		dot := point{x: inputs.ParseDecInt64(coords[0]), y: inputs.ParseDecInt64(coords[1])}
		matrix.dots = append(matrix.dots, dot)
		if dot.x > matrix.width {
			matrix.width = dot.x
		}
		if dot.y > matrix.height {
			matrix.height = dot.y
		}
	}
	for _, d := range strings.Split(parts[1], "\n") {
		operands := strings.Split(strings.Split(d, " ")[2], "=")
		if operands[0] == "x" {
			folds = append(folds, point{x: inputs.ParseDecInt64(operands[1]), y: 0})
		} else {
			folds = append(folds, point{y: inputs.ParseDecInt64(operands[1]), x: 0})
		}
	}
	return matrix, folds
}

func dotExists(matrix dotmatrix, dot point) bool {
	for _, d := range matrix.dots {
		if d.x == dot.x && d.y == dot.y {
			return true
		}
	}
	return false
}

func foldOn(matrix dotmatrix, x int64, y int64) dotmatrix {
	var newmatrix dotmatrix
	if y != 0 {
		newmatrix.width = matrix.width
		for _, dot := range matrix.dots {
			if dot.y > y {
				dot.y = (y*2 - dot.y)
			}
			if !dotExists(newmatrix, dot) {
				newmatrix.dots = append(newmatrix.dots, dot)
			}
		}
		newmatrix.height = y
	} else {
		newmatrix.height = matrix.height
		for _, dot := range matrix.dots {
			if dot.x > x {
				dot.x = (x*2 - dot.x)
			}
			if !dotExists(newmatrix, dot) {
				newmatrix.dots = append(newmatrix.dots, dot)
			}
		}
		newmatrix.width = x
	}
	return newmatrix
}

func printMatrix(matrix dotmatrix) {
	for y := int64(0); y < matrix.height; y++ {
		for x := int64(0); x < matrix.width; x++ {
			if dotExists(matrix, point{x: x, y: y}) {
				fmt.Print("X")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Print("\n")
	}
}

func (matrix dotmatrix) ToStrings() []string {
	result := make([]string, matrix.width/5+1)

	for y := int64(0); y < matrix.height; y++ {
		for x := int64(0); x < matrix.width; x++ {
			if dotExists(matrix, point{x: x, y: y}) {
				result[x/5] = fmt.Sprintf("%sX", result[x/5])
			} else {
				result[x/5] = fmt.Sprintf("%s ", result[x/5])
			}
		}
		for i := range result {
			result[i] = fmt.Sprintf("%s\n", result[i])
		}
	}
	return result
}

func part1(input string) interface{} {
	matrix, folds := parseInput(input)
	for _, fold := range folds {
		matrix = foldOn(matrix, fold.x, fold.y)
		break
	}
	return len(matrix.dots)
}

func resultTochars(matrix dotmatrix) (result string) {
	s := matrix.ToStrings()
	for i := range s {
		if _, ok := characters.CHARS[s[i]]; ok {
			result = fmt.Sprintf("%s%s", result, string(characters.CHARS[s[i]]))
		}
	}
	return result
}

func part2(input string) interface{} {
	matrix, folds := parseInput(input)
	for _, fold := range folds {
		matrix = foldOn(matrix, fold.x, fold.y)
	}
	//printMatrix(matrix)
	return resultTochars(matrix)
}

func main() {
	utils.Solve(part1, part2, DAY)
}

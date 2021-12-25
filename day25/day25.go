package main

import (
	// "adventofcodego/utils/inputs"
	"adventofcodego/utils/utils"
	"fmt"
	"strings"
	//	"github.com/pkg/profile"
)

var DAY int = 25
var HEIGHT int
var WIDTH int

func parseMap(input string) []rune {
	lines := strings.Split(input, "\n")
	WIDTH = len(lines[0])
	HEIGHT = len(lines)
	cmap := make([]rune, WIDTH*HEIGHT)
	for y := range lines {
		for x := range lines[y] {
			cmap[x+y*WIDTH] = rune(lines[y][x])
			if cmap[x+y*WIDTH] == '.' {
				cmap[x+y*WIDTH] = 0
			}
		}
	}
	return cmap
}

func down(y int) int {
	y++
	if y >= HEIGHT {
		y = 0
	}
	return y
}

func right(x int) int {
	x++
	if x >= WIDTH {
		x = 0
	}
	return x
}

func runStep(cmap []rune) ([]rune, bool) {
	updated := false
	newmap := make([]rune, WIDTH*HEIGHT)
	for y := 0; y < HEIGHT; y++ {
		for x := 0; x < WIDTH; x++ {
			switch cmap[x+y*WIDTH] {
			case '>':
				if cmap[right(x)+y*WIDTH] == 0 {
					newmap[right(x)+y*WIDTH] = '>'
					updated = true
				} else {
					newmap[x+y*WIDTH] = '>'
				}
			}
		}
	}

	for y := 0; y < HEIGHT; y++ {
		for x := 0; x < WIDTH; x++ {
			switch cmap[x+y*WIDTH] {
			case 'v':
				if newmap[x+down(y)*WIDTH] != '>' && cmap[x+down(y)*WIDTH] != 'v' {
					newmap[x+down(y)*WIDTH] = 'v'
					updated = true
				} else {
					newmap[x+y*WIDTH] = 'v'
				}
			}
		}
	}

	return newmap, updated
}

func printMap(cmap []rune) {
	for y := 0; y < HEIGHT; y++ {
		for x := 0; x < WIDTH; x++ {
			if cmap[x+y*WIDTH] == 0 {
				fmt.Print(".")
			} else {
				fmt.Print(string(cmap[x+y*WIDTH]))
			}
		}
		fmt.Println()
	}
}

func part1(input string) interface{} {
	cmap := parseMap(input)
	updated := true
	i := 0
	for i = 0; updated; i++ {
		cmap, updated = runStep(cmap)
	}
	return i
}

func part2(input string) interface{} {
	return nil
}

func main() {
	// defer profile.Start().Stop()
	utils.Solve(part1, part2, DAY)
}

package main

import (
	// "adventofcodego/utils/inputs"
	"adventofcodego/utils/utils"
	"fmt"
	"strings"
	//	"github.com/pkg/profile"
)

var DAY int = 20

func parsePixels(inputs []string, width int, height int) []int {
	pixels := make([]int, width*height)
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if inputs[y][x] == '#' {
				pixels[x+y*height] = 1
			}
		}
	}
	return pixels
}

func binToInt(bits []int) int {
	value := 0
	for i := range bits {
		value = value<<1 + bits[i]
	}
	return value
}

func getNeighborsNumber(pixels []int, x int, y int, width int, height int, default_pixel int) int {
	neighbors := make([]int, 9)
	i := 0
	for iy := -1; iy <= 1; iy++ {
		for ix := -1; ix <= 1; ix++ {
			nx := x + ix
			ny := y + iy
			if nx < 0 || ny < 0 || nx >= width || ny >= height {
				neighbors[i] = default_pixel
				i++
				continue
			} else {
				neighbors[i] = pixels[nx+ny*width]
				i++
			}
		}
	}
	return binToInt(neighbors)
}

func enhanceImage(pixels []int, width int, height int, iea string, default_pixel int) ([]int, int, int, int) {
	width += 2
	height += 2
	newpixels := make([]int, (width * height))
	count := 0
	for y := -1; y < height-1; y++ {
		for x := -1; x < width-1; x++ {
			if iea[getNeighborsNumber(pixels, x, y, width-2, height-2, default_pixel)] == '#' {
				newpixels[x+1+(y+1)*width] = 1
				count++
			}
		}
	}
	return newpixels, count, width, height
}

func printPixels(pixels []int, width int, height int) {
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if pixels[x+y*width] == 1 {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func part1(input string) interface{} {
	inputs := strings.Split(input, "\n")
	iea := inputs[0]
	inputs = inputs[2:]
	width := len(inputs[0])
	height := len(inputs)
	pixels := parsePixels(inputs, width, height)
	count := 0
	default_pixel := 0

	for i := 0; i < 2; i++ {
		pixels, count, width, height = enhanceImage(pixels, width, height, iea, default_pixel)
		if iea[0] == '#' && i%2 == 0 {
			default_pixel = 1
		} else {
			default_pixel = 0
		}
	}
	//printPixels(pixels, width, height)

	return count
}

func part2(input string) interface{} {
	inputs := strings.Split(input, "\n")
	iea := inputs[0]
	inputs = inputs[2:]
	width := len(inputs[0])
	height := len(inputs)
	pixels := parsePixels(inputs, width, height)
	count := 0
	default_pixel := 0

	for i := 0; i < 50; i++ {
		pixels, count, width, height = enhanceImage(pixels, width, height, iea, default_pixel)
		if iea[0] == '#' && i%2 == 0 {
			default_pixel = 1
		} else {
			default_pixel = 0
		}
	}
	//printPixels(pixels, width, height)

	return count
}

func main() {
	// defer profile.Start().Stop()
	utils.Solve(part1, part2, DAY)
}

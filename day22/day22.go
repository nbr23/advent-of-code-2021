package main

import (
	// "adventofcodego/utils/inputs"
	"adventofcodego/utils/inputs"
	"adventofcodego/utils/utils"
	"regexp"
	"strings"
	//	"github.com/pkg/profile"
)

var DAY int = 22

type point struct {
	x int
	y int
	z int
}

type cube struct {
	on bool
	x1 int
	x2 int
	y1 int
	y2 int
	z1 int
	z2 int
}

func parseCubes(input []string) []cube {
	cubes := make([]cube, len(input))
	r, _ := regexp.Compile("([onf]+) x=([-0-9]+)..([-0-9]+),y=([-0-9]+)..([-0-9]+),z=([-0-9]+)..([-0-9]+)")
	for i, s := range input {
		m := r.FindStringSubmatch(s)
		on := m[1] == "on"
		cubes[i] = cube{on, inputs.ParseDecInt(m[2]), inputs.ParseDecInt(m[3]), inputs.ParseDecInt(m[4]), inputs.ParseDecInt(m[5]), inputs.ParseDecInt(m[6]), inputs.ParseDecInt(m[7])}
	}
	return cubes
}

func processCube(c cube, litpoints map[point]bool) {
	for x := utils.IntMax(-50, c.x1); x <= utils.IntMin(50, c.x2); x++ {
		for y := utils.IntMax(-50, c.y1); y <= utils.IntMin(50, c.y2); y++ {
			for z := utils.IntMax(-50, c.z1); z <= utils.IntMin(50, c.z2); z++ {
				if c.on {
					litpoints[point{x, y, z}] = true
				} else {
					delete(litpoints, point{x, y, z})
				}
			}
		}
	}
}

func part1(input string) interface{} {
	litpoints := make(map[point]bool)
	cubes := parseCubes(strings.Split(input, "\n"))
	for _, c := range cubes {
		processCube(c, litpoints)
	}
	return len(litpoints)
}

func part2(input string) interface{} {
	return nil
}

func main() {
	// defer profile.Start().Stop()
	utils.Solve(part1, part2, DAY)
}

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

type cube struct {
	on bool
	x1 int
	x2 int
	y1 int
	y2 int
	z1 int
	z2 int
}

func parseCubes(input []string, bounds *cube) []cube {
	cubes := make([]cube, len(input))
	r, _ := regexp.Compile("([onf]+) x=([-0-9]+)..([-0-9]+),y=([-0-9]+)..([-0-9]+),z=([-0-9]+)..([-0-9]+)")
	for _, s := range input {
		m := r.FindStringSubmatch(s)
		on := m[1] == "on"
		if bounds != nil {
			if inputs.ParseDecInt(m[2]) >= bounds.x1 && inputs.ParseDecInt(m[3]) <= bounds.x2 &&
				inputs.ParseDecInt(m[4]) >= bounds.y1 && inputs.ParseDecInt(m[5]) <= bounds.y2 &&
				inputs.ParseDecInt(m[6]) >= bounds.z1 && inputs.ParseDecInt(m[7]) <= bounds.z2 {
				cubes = append(cubes, cube{on, inputs.ParseDecInt(m[2]), inputs.ParseDecInt(m[3]), inputs.ParseDecInt(m[4]), inputs.ParseDecInt(m[5]), inputs.ParseDecInt(m[6]), inputs.ParseDecInt(m[7])})
			}
		} else {
			cubes = append(cubes, cube{on, inputs.ParseDecInt(m[2]), inputs.ParseDecInt(m[3]), inputs.ParseDecInt(m[4]), inputs.ParseDecInt(m[5]), inputs.ParseDecInt(m[6]), inputs.ParseDecInt(m[7])})
		}
	}
	return cubes
}

func part1(input string) interface{} {
	cubes := parseCubes(strings.Split(input, "\n"), &cube{false, -50, 50, -50, 50, -50, 50})
	return processCubes(cubes)
}

func intersection(c1 cube, c2 cube) (cube, bool) {
	newon := c1.on
	if c1.on == c2.on {
		newon = !c1.on
	} else if !c2.on {
		newon = true
	}
	if utils.IntMax(c1.x1, c2.x1) <= utils.IntMin(c1.x2, c2.x2) &&
		utils.IntMax(c1.y1, c2.y1) <= utils.IntMin(c1.y2, c2.y2) &&
		utils.IntMax(c1.z1, c2.z1) <= utils.IntMin(c1.z2, c2.z2) {
		return cube{newon,
			utils.IntMax(c1.x1, c2.x1), utils.IntMin(c1.x2, c2.x2),
			utils.IntMax(c1.y1, c2.y1), utils.IntMin(c1.y2, c2.y2),
			utils.IntMax(c1.z1, c2.z1), utils.IntMin(c1.z2, c2.z2)}, true
	}
	return cube{}, false
}

func cubeSize(c cube) int {
	if c.on {
		return (c.x2 - c.x1 + 1) * (c.y2 - c.y1 + 1) * (c.z2 - c.z1 + 1)
	}
	return -(c.x2 - c.x1 + 1) * (c.y2 - c.y1 + 1) * (c.z2 - c.z1 + 1)
}

func processCubes(cubes []cube) int {
	intersections := make([]cube, 0)
	for i := 0; i < len(cubes); i++ {
		newinter := make([]cube, 0)
		for j := range intersections {
			inter, ok := intersection(cubes[i], intersections[j])
			if ok {
				newinter = append(newinter, inter)
			}
		}
		if cubes[i].on {
			intersections = append(intersections, cubes[i])
		}
		intersections = append(intersections, newinter...)
	}

	total := 0
	for _, inter := range intersections {
		total += cubeSize(inter)
	}
	return total
}

func part2(input string) interface{} {
	cubes := parseCubes(strings.Split(input, "\n"), nil)
	return processCubes(cubes)
}

func main() {
	// defer profile.Start().Stop()
	utils.Solve(part1, part2, DAY)
}

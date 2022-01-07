package main

import (
	// "adventofcodego/utils/inputs"
	"adventofcodego/utils/inputs"
	"adventofcodego/utils/utils"
	"fmt"
	"regexp"
	"strings"
	//	"github.com/pkg/profile"
)

var DAY int = 19

type pair struct {
	x int
	y int
}

type point struct {
	x int
	y int
	z int
}

type angle struct {
	sinus   int
	cosinus int
}

var ROTATIONS = [][]int{
	{0, 0, 0},
	{0, 0, 90},
	{0, 0, 180},
	{0, 0, 270},
	{0, 90, 270},
	{0, 90, 0},
	{0, 90, 90},
	{0, 90, 180},
	{0, 180, 270},
	{0, 180, 0},
	{0, 180, 90},
	{0, 180, 180},
	{0, 270, 0},
	{0, 270, 90},
	{0, 270, 180},
	{0, 270, 270},
	{90, 0, 180},
	{90, 0, 270},
	{90, 0, 0},
	{90, 0, 90},
	{90, 180, 270},
	{90, 180, 0},
	{90, 180, 90},
	{90, 180, 180},
}

func getCosSin(angle int) (int, int) {
	switch angle {
	case 0:
		return 1, 0
	case 90:
		return 0, 1
	case 180:
		return -1, 0
	case 270:
		return 0, -1
	}
	panic(fmt.Errorf("Unsupported angle %d", angle))
}

func rotateX(pt point, theta int) point {
	cos, sin := getCosSin(theta)
	return point{
		x: pt.x,
		y: pt.y*cos - pt.z*sin,
		z: pt.y*sin + pt.z*cos,
	}
}

func rotateY(pt point, theta int) point {
	cos, sin := getCosSin(theta)
	return point{
		x: pt.x*cos + pt.z*sin,
		y: pt.y,
		z: pt.z*cos - pt.x*sin,
	}
}

func rotateZ(pt point, theta int) point {
	cos, sin := getCosSin(theta)
	return point{
		x: pt.x*cos - pt.y*sin,
		y: pt.x*sin + pt.y*cos,
		z: pt.z,
	}
}

func (p point) String() string {
	return fmt.Sprintf("%d,%d,%d", p.x, p.y, p.z)
}

func parseScanners(input string) [][]point {
	scanner := -1
	scanners := make([][]point, 0)
	r, _ := regexp.Compile("--- scanner [0-9]+ ---")
	for _, line := range strings.Split(input, "\n") {
		if len(line) == 0 {
			continue
		} else if r.MatchString(line) {
			scanner++
			scanners = append(scanners, make([]point, 0))
		} else {
			coords := strings.Split(line, ",")
			scanners[scanner] = append(scanners[scanner], point{inputs.ParseDecInt(coords[0]), inputs.ParseDecInt(coords[1]), inputs.ParseDecInt(coords[2])})
		}
	}
	return scanners
}

func getRotations(scanner []point) [][]point {
	rotations := make([][]point, 24)
	for i, rot := range ROTATIONS {
		aX, aY, aZ := rot[0], rot[1], rot[2]
		newrot := make([]point, len(scanner))
		for j, beacon := range scanner {
			newrot[j] = rotateZ(rotateY(rotateX(beacon, aX), aY), aZ)
		}
		rotations[i] = newrot
	}
	return rotations
}

func rotationExists(rotations [][]point, newrot []point) bool {
	for _, rot := range rotations {
		same := true
		for i := range rot {
			if rot[i] != newrot[i] {
				same = false
				break
			}
		}
		if same {
			return true
		}
	}
	return false
}

func normalizeOn(beacons []point, pt point) []point {
	newbeacons := make([]point, len(beacons))
	for i := range beacons {
		newbeacons[i] = point{beacons[i].x - pt.x, beacons[i].y - pt.y, beacons[i].z - pt.z}
	}
	return newbeacons
}

func countCommonBeacons(a []point, b []point) (int, []pair) {
	res := 0
	commonlist := make([]pair, 0)
	for i := range a {
		for j := range b {
			if a[i] == b[j] {
				res++
				commonlist = append(commonlist, pair{i, j})
			}
		}
	}
	return res, commonlist
}

func pointSub(a point, b point) point {
	return point{a.x - b.x, a.y - b.y, a.z - b.z}
}
func pointAdd(a point, b point) point {
	return point{a.x + b.x, a.y + b.y, a.z + b.z}
}

func searchForBeacons(input string) (int, int) {
	scanners := parseScanners(input)
	scanners_count := len(scanners)

	max_manhattan := 0

	i := 0
	current_scanner := scanners[0]
	normalizer := point{0, 0, 0}
	goodscannersindex := make([]bool, scanners_count)
	goodscannersindex[0] = true
	goodscannerscount := 1
	next := make([][]point, scanners_count)
	nextnormalizer := make([]point, scanners_count)
	nextcoordnormalizer := make([]point, scanners_count)
	coordnormalizer := point{0, 0, 0}
	nextsize := 0

	scanner_coords := make([]point, scanners_count)
	scanner_coords[0] = point{0, 0, 0}

	goodbeacons := make(map[point]bool)

	for {
		for _, beacon := range current_scanner {
			normalized := normalizeOn(current_scanner, beacon)
			for j := range scanners {
				scanner2 := scanners[j]
				if i == j {
					continue
				}
				for _, rot := range getRotations(scanner2) {
					for _, beacon2 := range rot {
						normalized2 := normalizeOn(rot, beacon2)
						if count, pairs := countCommonBeacons(normalized, normalized2); count >= 12 {
							if !goodscannersindex[j] {
								if len(goodbeacons) == 0 {
									for _, b := range current_scanner {
										goodbeacons[b] = true
									}
									normalizer = point{0, 0, 0}
								}
								scanner_coords[goodscannerscount] = pointAdd(current_scanner[pairs[0].x], pointSub(coordnormalizer, rot[pairs[0].y]))
								max_manhattan = utils.IntMax(max_manhattan, getMaxManhattan(scanner_coords[goodscannerscount], scanner_coords[0:goodscannerscount]))
								goodscannersindex[j] = true
								goodscannerscount++
								i = j
								next[nextsize] = normalized2
								nextcoordnormalizer[nextsize] = pointAdd(coordnormalizer, current_scanner[pairs[0].x])
								nextnormalizer[nextsize] = point{normalizer.x - beacon.x, normalizer.y - beacon.y, normalizer.z - beacon.z}
								nextsize++
								for _, b := range normalizeOn(normalized2, point{normalizer.x - beacon.x, normalizer.y - beacon.y, normalizer.z - beacon.z}) {
									goodbeacons[b] = true
								}
								j = 0
							}
						}
					}
				}
			}
		}
		if nextsize > 0 {
			nextsize--
			current_scanner = next[nextsize]
			normalizer = nextnormalizer[nextsize]
			coordnormalizer = nextcoordnormalizer[nextsize]
		}
		if goodscannerscount == scanners_count {
			break
		}
	}
	return len(goodbeacons), max_manhattan
}

func getMaxManhattan(p point, coords []point) int {
	max := 0
	for i := range coords {
		man := abs(p.x-coords[i].x) + abs(p.y-coords[i].y) + abs(p.z-coords[i].z)
		if man > max {
			max = man
		}
	}
	return max
}

func abs(v int) int {
	if v >= 0 {
		return v
	}
	return -v
}

var P2 int

func part1(input string) interface{} {
	var p1 int
	p1, P2 = searchForBeacons(input)
	return p1
}

func part2(input string) interface{} {
	return P2
}

func main() {
	// defer profile.Start().Stop()
	utils.Solve(part1, part2, DAY)
}

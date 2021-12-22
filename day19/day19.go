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

var ANGLES = map[int]angle{
	0:   {cosinus: 1, sinus: 0},
	90:  {cosinus: 0, sinus: 1},
	180: {cosinus: -1, sinus: 0},
	270: {cosinus: 0, sinus: -1},
}

func rotateX(pt point, theta int) point {
	return point{
		x: pt.x,
		y: pt.y*ANGLES[theta].cosinus - pt.z*ANGLES[theta].sinus,
		z: pt.y*ANGLES[theta].sinus + pt.z*ANGLES[theta].cosinus,
	}
}

func rotateY(pt point, theta int) point {
	return point{
		x: pt.x*ANGLES[theta].cosinus + pt.z*ANGLES[theta].sinus,
		y: pt.y,
		z: pt.z*ANGLES[theta].cosinus - pt.x*ANGLES[theta].sinus,
	}
}

func rotateZ(pt point, theta int) point {
	return point{
		x: pt.x*ANGLES[theta].cosinus - pt.y*ANGLES[theta].sinus,
		y: pt.x*ANGLES[theta].sinus + pt.y*ANGLES[theta].cosinus,
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
	rotations := make([][]point, 0)
	i := 0
	for aX := range ANGLES {
		for aY := range ANGLES {
			for aZ := range ANGLES {
				newrot := make([]point, 0)
				for _, beacon := range scanner {
					newrot = append(newrot, rotateZ(rotateY(rotateX(beacon, aX), aY), aZ))
				}
				if !rotationExists(rotations, newrot) {
					rotations = append(rotations, newrot)
					i++
				}
			}
		}
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

func searchForBeacons(input string) int {
	scanners := parseScanners(input)
	scanners_rotated := make([][][]point, len(scanners))
	for i, scanner := range scanners {
		scanners_rotated[i] = getRotations(scanner)
	}

	i := 0
	current_scanner := scanners[i]
	normalizer := point{0, 0, 0}
	goodscanners := make([][]point, 0)
	goodscanners = append(goodscanners, current_scanner)
	goodscannersindex := make(map[int]bool)
	goodscannersindex[0] = true
	goodscannerscount := 1
	next := make([][]point, len(scanners))
	nextnormalizer := make([]point, len(scanners))
	nextsize := 0

	scanner_coords := make([]point, len(scanners))
	scanner_coords[0] = point{0, 0, 0}

	goodbeacons := make(map[point]bool)

	for {
		for _, beacon := range current_scanner {
			normalized := normalizeOn(current_scanner, beacon)
			for j, scanner2 := range scanners {
				if i == j {
					continue
				}
				for _, rot := range getRotations(scanner2) {
					for _, beacon2 := range rot {
						normalized2 := normalizeOn(rot, beacon2)
						if count, pairs := countCommonBeacons(normalized, normalized2); count >= 12 {
							if _, ok := goodscannersindex[j]; !ok {
								if len(goodbeacons) == 0 {
									for _, b := range current_scanner {
										goodbeacons[b] = true
									}
									normalizer = point{0, 0, 0}
								}
								scanner_coords[j] = pointAdd(current_scanner[pairs[0].x], pointSub(normalizer, rot[pairs[0].y]))
								goodscannersindex[j] = true
								goodscannerscount++
								i = j
								goodscanners = append(goodscanners, normalized2)
								next[nextsize] = normalized2
								nextnormalizer[nextsize] = current_scanner[pairs[0].x]
								nextsize++
								for _, b := range normalizeOn(normalized2, point{normalizer.x - beacon.x, normalizer.y - beacon.y, normalizer.z - beacon.z}) {
									goodbeacons[b] = true
								}
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
		}
		if goodscannerscount == len(scanners) {
			break
		}
	}

	return len(goodbeacons)
}

func getMaxManhattan(coords []point) int {
	max := 0
	for i, _ := range coords {
		for j, _ := range coords {
			if i == j {
				continue
			}
			man := abs(coords[i].x-coords[j].x) + abs(coords[i].y-coords[j].y) + abs(coords[i].z-coords[j].z)
			if man > max {
				max = man
			}
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

func part1(input string) interface{} {
	return searchForBeacons(input)
}

func part2(input string) interface{} {
	return nil
}

func main() {
	// defer profile.Start().Stop()
	utils.Solve(part1, part2, DAY)
}

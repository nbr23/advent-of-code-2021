package main

import (
	// "adventofcodego/utils/inputs"
	"adventofcodego/utils/utils"
	"fmt"
	"math"
	"regexp"
	"strings"
	//	"github.com/pkg/profile"
)

var DAY int = 23

var HEIGHT = 3
var WIDTH = 11
var PODSCOUNT = 8

const (
	Amber  = 1
	Bronze = 10
	Copper = 100
	Desert = 1000
)

type Pod struct {
	x       int
	y       int
	element int
}

type Move struct {
	cost int
	x    int
	y    int
}

func parsePods(input string) []int {
	lines := strings.Split(input, "\n")
	r1, _ := regexp.Compile("###([A-Z])#([A-Z])#([A-Z])#([A-Z])###")
	r2, _ := regexp.Compile("  #([A-Z])#([A-Z])#([A-Z])#([A-Z])#")
	m1 := r1.FindStringSubmatch(lines[2])
	m2 := r2.FindStringSubmatch(lines[3])

	podsmap := make([]int, HEIGHT*WIDTH)

	for i := 1; i <= 4; i++ {
		x := i * 2
		podsmap[x+1*WIDTH] = getCost(m1[i][0])
		podsmap[x+2*WIDTH] = getCost(m2[i][0])
	}
	return podsmap
}

func podToString(p int) string {
	switch p {
	case Amber:
		return "A"
	case Bronze:
		return "B"
	case Copper:
		return "C"
	case Desert:
		return "D"
	}
	panic("Unknown element")
}

func printPods(state State) {
	fmt.Printf("State cost: %d\n", state.cost)
	fmt.Println("############")
	for y := 0; y < HEIGHT; y++ {
		fmt.Print(" ")
		for x := 0; x < WIDTH; x++ {
			if y > 0 && x%2 == 1 {
				fmt.Print("#")
			} else {
				pod := podAt(x, y, state.pods)
				if pod == 0 {
					fmt.Print(" ")
				} else {
					fmt.Print(podToString(pod))
				}
			}
		}
		fmt.Println()
	}
	fmt.Println("  #########")
}

func getCost(letter byte) int {
	switch letter {
	case 'A':
		return Amber
	case 'B':
		return Bronze
	case 'C':
		return Copper
	case 'D':
		return Desert
	}
	panic(fmt.Errorf("Unknown element %s", string(letter)))
}

func isRoom(x, y int) bool {
	return y > 0 && (x == 2 || x == 4 || x == 6 || x == 8)
}

func getRoomX(element int) int {
	switch element {
	case Amber:
		return 2
	case Bronze:
		return 4
	case Copper:
		return 6
	case Desert:
		return 8
	}
	return 0
}

func podAt(x int, y int, pods []int) int {
	return pods[x+y*WIDTH]
}

func isProperRoom(element int, x int, y int) bool {
	return y > 0 && (x == 2 && element == Amber ||
		x == 4 && element == Bronze ||
		x == 6 && element == Copper ||
		x == 8 && element == Desert)
}

func isFinalPosition(element int, x int, y int, pods []int) bool {

	return y == 2 && (x == 2 && element == Amber ||
		x == 4 && element == Bronze ||
		x == 6 && element == Copper ||
		x == 8 && element == Desert) ||
		y == 1 && (x == 2 && element == Amber ||
			x == 4 && element == Bronze ||
			x == 6 && element == Copper ||
			x == 8 && element == Desert) && podAt(x, 2, pods) == element
}

type State struct {
	pods []int
	cost int
}

func getPossibleMoves(podx int, pody int, pods []int, cost_acc int, solution *int) []State {
	moves := make([]State, 0)
	current_element := podAt(podx, pody, pods)

	// Currently in the hallway
	if !isRoom(podx, pody) {
		roomx := getRoomX(current_element)
		p1 := podAt(roomx, 1, pods)
		p2 := podAt(roomx, 2, pods)

		// The room is empty or contains one element already in the right place
		if p2 == 0 || (p2 == current_element && p1 == 0) {
			// We need to go left
			newy := 1
			if p2 == 0 {
				newy = 2
			}
			if roomx < podx {
				freeway := true
				for x := podx - 1; x >= roomx; x-- {
					if podAt(x, 0, pods) != 0 {
						freeway = false
						break
					}
				}
				if freeway {
					if cost_acc+(newy+podx-roomx)*current_element < *solution {
						newpods := make([]int, WIDTH*HEIGHT)
						copy(newpods, pods)
						newpods[podx+pody*WIDTH] = 0
						newpods[roomx+newy*WIDTH] = current_element
						moves = append(moves, State{pods: newpods, cost: cost_acc + (newy+podx-roomx)*current_element})
					}
				}
			} else { // We need to go right
				freeway := true
				for x := podx + 1; x <= roomx; x++ {
					if podAt(x, 0, pods) != 0 {
						freeway = false
						break
					}
				}
				if freeway {
					if cost_acc+(newy+roomx-podx)*current_element < *solution {
						newpods := make([]int, WIDTH*HEIGHT)
						copy(newpods, pods)
						newpods[podx+pody*WIDTH] = 0
						newpods[roomx+newy*WIDTH] = current_element
						moves = append(moves, State{pods: newpods, cost: cost_acc + (newy+roomx-podx)*current_element})
					}
				}
			}
		}
	} else { // In a room
		if podAt(podx, pody-1, pods) != 0 {
			return moves
		}

		// Check left
		for x := podx - 1; x >= 0; x-- {
			// if we're in front of a room, we skip
			if isRoom(x, 1) {
				continue
			}
			if podAt(x, 0, pods) == 0 {
				if cost_acc+(pody+podx-x)*current_element < *solution {
					newpods := make([]int, WIDTH*HEIGHT)
					copy(newpods, pods)
					newpods[podx+pody*WIDTH] = 0
					newpods[x] = current_element
					moves = append(moves, State{pods: newpods, cost: cost_acc + (pody+podx-x)*current_element})
				}
			} else {
				break
			}
		}

		// Check right
		for x := podx + 1; x <= 10; x++ {
			// if we're in front of a room, we skip
			if isRoom(x, 1) {
				continue
			}
			if podAt(x, 0, pods) == 0 {
				if cost_acc+(pody+x-podx)*current_element < *solution {
					newpods := make([]int, WIDTH*HEIGHT)
					copy(newpods, pods)
					newpods[podx+pody*WIDTH] = 0
					newpods[x] = current_element
					moves = append(moves, State{pods: newpods, cost: cost_acc + (pody+x-podx)*current_element})
				}
			} else {
				break
			}
		}
	}
	return moves
}

func playStep(state State, solution *int) {
	if state.cost >= *solution {
		return
	}
	arrived := make(map[Pod]bool)
	for y := 0; y < HEIGHT; y++ {
		for x := 0; x < WIDTH; x++ {
			if isFinalPosition(podAt(x, y, state.pods), x, y, state.pods) {
				arrived[Pod{x: x, y: y, element: podAt(x, y, state.pods)}] = true
			}
		}
	}

	if len(arrived) == PODSCOUNT {
		if state.cost < *solution {
			*solution = state.cost
			fmt.Println(*solution)
		}
		return
	}

	for y := 0; y < HEIGHT; y++ {
		for x := 0; x < WIDTH; x++ {
			if _, exists := arrived[Pod{x: x, y: y, element: podAt(x, y, state.pods)}]; exists {
				continue
			}

			if podAt(x, y, state.pods) != 0 {
				moves := getPossibleMoves(x, y, state.pods, state.cost, solution)
				for _, move := range moves {
					playStep(move, solution)
				}
			}
		}
	}
}

func part1(input string) interface{} {
	pods := parsePods(input)
	printPods(State{pods: pods, cost: 0})
	solution := math.MaxInt
	playStep(State{pods, 0}, &solution)
	return solution
}

func part2(input string) interface{} {
	HEIGHT = 5
	return nil
}

func main() {
	// defer profile.Start().Stop()
	utils.Solve(part1, part2, DAY)
}

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

type State struct {
	pods []int
	cost int
}

func parsePods(input string) []int {
	lines := strings.Split(input, "\n")
	r1, _ := regexp.Compile("###([A-Z])#([A-Z])#([A-Z])#([A-Z])###")
	r2, _ := regexp.Compile("  #([A-Z])#([A-Z])#([A-Z])#([A-Z])#")
	m1 := r1.FindStringSubmatch(lines[2])
	m2 := r2.FindStringSubmatch(lines[3])

	podsarray := make([]int, HEIGHT*WIDTH)

	for i := 1; i <= 4; i++ {
		x := i * 2
		podsarray[x+1*WIDTH] = getCost(m1[i][0])
		podsarray[x+2*WIDTH] = getCost(m2[i][0])
	}
	return podsarray
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
	return (x == 2 && element == Amber ||
		x == 4 && element == Bronze ||
		x == 6 && element == Copper ||
		x == 8 && element == Desert)
}

func isFinalPosition(element int, x int, y int, pods []int) bool {
	if element == 0 || !isProperRoom(element, x, y) {
		return false
	}
	for py := HEIGHT - 1; py > y; py-- {
		if element != podAt(x, py, pods) {
			return false
		}
	}
	return true
}

func getAvailableYPositionInRoom(roomx int, pods []int) int {
	for y := HEIGHT - 1; y > 0; y-- {
		elt := podAt(roomx, y, pods)
		if elt == 0 {
			return y
		} else if !isProperRoom(elt, roomx, y) {
			return -1
		}
	}
	return -1
}

func getPossibleMoves(podx int, pody int, pods []int, cost_acc int, solution *int) []State {
	moves := make([]State, 0)
	current_element := podAt(podx, pody, pods)

	// Currently in the hallway
	if pody == 0 {
		roomx := getRoomX(current_element)
		newy := getAvailableYPositionInRoom(roomx, pods)

		// The room is empty or contains elements already in the right place
		if newy != -1 {
			if roomx < podx {
				// Check if an element is on our way
				freeway := true
				for x := podx - 1; x >= roomx; x-- {
					if pods[x] != 0 {
						freeway = false
						break
					}
				}
				if freeway && cost_acc+(newy+podx-roomx)*current_element < *solution {
					newpods := append(make([]int, 0, len(pods)), pods...)
					newpods[podx+pody*WIDTH] = 0
					newpods[roomx+newy*WIDTH] = current_element
					moves = append(moves, State{pods: newpods, cost: cost_acc + (newy+podx-roomx)*current_element})
				}
			} else { // We need to go right
				// Check if an element is on our way
				freeway := true
				for x := podx + 1; x <= roomx; x++ {
					if pods[x] != 0 {
						freeway = false
						break
					}
				}
				if freeway && cost_acc+(newy+roomx-podx)*current_element < *solution {
					newpods := append(make([]int, 0, len(pods)), pods...)
					newpods[podx+pody*WIDTH] = 0
					newpods[roomx+newy*WIDTH] = current_element
					moves = append(moves, State{pods: newpods, cost: cost_acc + (newy+roomx-podx)*current_element})
				}
			}
		}
	} else { // In a room
		// Check if there's something above us
		for y := pody - 1; y > 0; y-- {
			if podAt(podx, y, pods) != 0 {
				return moves
			}
		}

		// Check left
		for x := podx - 1; x >= 0; x-- {
			// if we're in front of a room, we skip
			if x%2 == 0 && x >= 2 {
				continue
			}
			// An element on our way, we cannot go further
			if pods[x] != 0 {
				break
			}
			if cost_acc+(pody+podx-x)*current_element < *solution {
				newpods := append(make([]int, 0, len(pods)), pods...)
				newpods[podx+pody*WIDTH] = 0
				newpods[x] = current_element
				moves = append(moves, State{pods: newpods, cost: cost_acc + (pody+podx-x)*current_element})
			}
		}

		// Check right
		for x := podx + 1; x <= 10; x++ {
			// if we're in front of a room, we skip
			if x%2 == 0 && x <= 8 {
				continue
			}
			// An element on our way, we cannot go further
			if pods[x] != 0 {
				break
			}
			if cost_acc+(pody+x-podx)*current_element < *solution {
				newpods := append(make([]int, 0, len(pods)), pods...)
				newpods[podx+pody*WIDTH] = 0
				newpods[x] = current_element
				moves = append(moves, State{pods: newpods, cost: cost_acc + (pody+x-podx)*current_element})
			}
		}
	}
	return moves
}

func playStep(state State, solution *int) {
	if state.cost >= *solution {
		return
	}
	arrived := make([]bool, PODSCOUNT, PODSCOUNT)
	arrived_count := 0
	for y := 0; y < HEIGHT; y++ {
		for x := 0; x < WIDTH; x++ {
			element := podAt(x, y, state.pods)
			if y > 0 && isFinalPosition(element, x, y, state.pods) {
				arrived[x/2-1+(y-1)*4] = true
				arrived_count++
			}
		}
	}

	if arrived_count == PODSCOUNT {
		if state.cost < *solution {
			*solution = state.cost
		}
		return
	}

	for y := 0; y < HEIGHT; y++ {
		for x := 0; x < WIDTH; x++ {
			element := podAt(x, y, state.pods)
			if idx := x/2 - 1 + (y-1)*4; idx >= 0 && idx < PODSCOUNT && arrived[idx] {
				continue
			}

			if element != 0 {
				for _, move := range getPossibleMoves(x, y, state.pods, state.cost, solution) {
					playStep(move, solution)
				}
			}
		}
	}
}

func part1(input string) interface{} {
	pods := parsePods(input)
	solution := math.MaxInt
	playStep(State{pods, 0}, &solution)
	return solution
}

func part2(input string) interface{} {
	HEIGHT = 5
	PODSCOUNT = 16

	pods := parsePods(input)
	insert1 := []int{1000, 100, 10, 1}
	insert2 := []int{1000, 10, 1, 100}
	for x := 1; x <= 4; x++ {
		pods[x*2+4*WIDTH] = pods[x*2+2*WIDTH]
		pods[x*2+2*WIDTH] = insert1[x-1]
		pods[x*2+3*WIDTH] = insert2[x-1]
	}

	solution := math.MaxInt
	playStep(State{pods, 0}, &solution)
	return solution
}

func main() {
	// defer profile.Start().Stop()
	utils.Solve(part1, part2, DAY)
}

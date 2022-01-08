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

var DAY int = 24

const (
	W = 0
	X = 1
	Y = 2
	Z = 3
)

type state []int

type Operator func(operand, operand, *state)

func getOp(op string) Operator {
	switch op {
	case "mul":
		return mul
	case "add":
		return add
	case "eql":
		return eql
	case "div":
		return div
	case "mod":
		return mod
	}
	return nil
}

func isVar(s byte) bool {
	return s == 'w' || s == 'x' || s == 'y' || s == 'z'
}

func mul(a, b operand, st *state) {
	if b.isvar {
		(*st)[a.varname] = (*st)[a.varname] * (*st)[b.varname]
	} else {
		(*st)[a.varname] = (*st)[a.varname] * b.intval
	}
}

func add(a, b operand, st *state) {
	if b.isvar {
		(*st)[a.varname] = (*st)[a.varname] + (*st)[b.varname]
	} else {
		(*st)[a.varname] = (*st)[a.varname] + b.intval
	}
}

func div(a, b operand, st *state) {
	if b.isvar {
		(*st)[a.varname] = (*st)[a.varname] / (*st)[b.varname]
	} else {
		(*st)[a.varname] = (*st)[a.varname] / b.intval
	}
}

func mod(a, b operand, st *state) {
	if b.isvar {
		(*st)[a.varname] = (*st)[a.varname] % (*st)[b.varname]
	} else {
		(*st)[a.varname] = (*st)[a.varname] % b.intval
	}
}

func eql(a, b operand, st *state) {
	vb := 0
	if b.isvar {
		vb = (*st)[b.varname]
	} else {
		vb = b.intval
	}
	if vb == (*st)[a.varname] {
		(*st)[a.varname] = 1
	} else {
		(*st)[a.varname] = 0
	}
}

type instruction struct {
	op string
	a  operand
	b  operand
}

type operand struct {
	isvar   bool
	intval  int
	varname byte
}

func stringToOperand(s string) operand {
	if isVar(s[0]) {
		return operand{true, 0, s[0] - 'w'}
	} else {
		return operand{false, inputs.ParseDecInt(s), 0}
	}
}

func parseInstructions(input string) []instruction {
	instructions := make([]instruction, 0)
	for _, s := range strings.Split(input, "\n") {
		if len(s) == 0 {
			continue
		}
		sp := strings.Split(s, " ")
		instructions = append(instructions, instruction{op: sp[0], a: stringToOperand(sp[1]), b: stringToOperand(sp[2])})
	}
	return instructions
}

type cacheobj struct {
	depth int
	w     int
	z     int
}

func getInstructionBlocks(input string) [][]instruction {
	r, _ := regexp.Compile("inp [a-z]")
	input = r.ReplaceAllString(input, "inp")
	blocks := make([][]instruction, 0)

	for _, b := range strings.Split(input, "inp\n") {
		if len(b) == 0 {
			continue
		}
		blocks = append(blocks, parseInstructions(b))
	}
	return blocks
}

func execBlock(block []instruction, curst state) state {
	for _, instr := range block {
		getOp(instr.op)(instr.a, instr.b, &curst)
	}
	return curst
}

func tryCombination(blocks [][]instruction, index int, z int, iterations []int, cache map[cacheobj]bool) (string, bool) {
	if index == len(blocks) {
		return "", z == 0
	}

	for _, i := range iterations {
		var newz int

		newz = execBlock(blocks[index], state{i, 0, 0, z})[Z]

		if _, ok := cache[cacheobj{index + 1, i, newz}]; !ok {
			res, found := tryCombination(blocks, index+1, newz, iterations, cache)
			if found {
				return fmt.Sprintf("%d%s", i, res), true
			} else {
				cache[cacheobj{index + 1, i, newz}] = true
			}
		}
	}
	return "", false
}

func part1(input string) interface{} {
	blocks := getInstructionBlocks(input)

	iterations := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}

	cache := make(map[cacheobj]bool)
	res, found := tryCombination(blocks, 0, 0, iterations, cache)
	if found {
		return res
	}
	return nil
}

func part2(input string) interface{} {
	blocks := getInstructionBlocks(input)

	iterations := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	cache := make(map[cacheobj]bool)
	res, found := tryCombination(blocks, 0, 0, iterations, cache)
	if found {
		return res
	}
	return nil
}

func main() {
	// defer profile.Start().Stop()
	utils.Solve(part1, part2, DAY)
}

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

type state map[string]int

type Operator func(string, string, *state)

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

func isVar(s string) bool {
	return s == "w" || s == "x" || s == "y" || s == "z"
}

func mul(a, b string, st *state) {
	if isVar(b) {
		(*st)[a] = (*st)[a] * (*st)[b]
	} else {
		(*st)[a] = (*st)[a] * inputs.ParseDecInt(b)
	}
}

func add(a, b string, st *state) {
	if isVar(b) {
		(*st)[a] = (*st)[a] + (*st)[b]
	} else {
		(*st)[a] = (*st)[a] + inputs.ParseDecInt(b)
	}
}

func div(a, b string, st *state) {
	if isVar(b) {
		(*st)[a] = (*st)[a] / (*st)[b]
	} else {
		(*st)[a] = (*st)[a] / inputs.ParseDecInt(b)
	}
}

func mod(a, b string, st *state) {
	if isVar(b) {
		(*st)[a] = (*st)[a] % (*st)[b]
	} else {
		(*st)[a] = (*st)[a] % inputs.ParseDecInt(b)
	}
}

func eql(a, b string, st *state) {
	vb := 0
	if isVar(b) {
		vb = (*st)[b]
	} else {
		vb = inputs.ParseDecInt(b)
	}
	if vb == (*st)[a] {
		(*st)[a] = 1
	} else {
		(*st)[a] = 0
	}
}

type instruction struct {
	op string
	a  string
	b  string
}

func parseInstructions(input string) []instruction {
	instructions := make([]instruction, 0)
	for _, s := range strings.Split(input, "\n") {
		if len(s) == 0 {
			continue
		}
		sp := strings.Split(s, " ")
		instructions = append(instructions, instruction{op: sp[0], a: sp[1], b: sp[2]})
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

func tryCombination(blocks [][]instruction, index int, z int, initmin int, initmax int, cache map[cacheobj]bool) (string, bool) {
	if index == len(blocks) {
		return "", z == 0
	}

	for i := initmax; i >= initmin; i-- {
		var newz int

		newz = execBlock(blocks[index], state{"w": i, "x": 0, "y": 0, "z": z})["z"]

		if _, ok := cache[cacheobj{index + 1, i, newz}]; !ok {
			res, found := tryCombination(blocks, index+1, newz, 1, 9, cache)
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

	for i := 9; i > 0; i-- {
		cache := make(map[cacheobj]bool)
		res, found := tryCombination(blocks, 0, 0, i, i, cache)
		if found {
			return res
		}
	}
	return nil
}

func part2(input string) interface{} {
	return nil
}

func main() {
	// defer profile.Start().Stop()
	utils.Solve(part1, part2, DAY)
}

package main

import (
	// "adventofcodego/utils/inputs"
	"adventofcodego/utils/utils"
	"fmt"
	"strings"
	//	"github.com/pkg/profile"
)

var DAY int = 18

type pair struct {
	left  *pair
	right *pair
	value int
}

func parsePairs(line string) *pair {
	depth := 0
	middle := -1
	for i := 0; i < len(line); i++ {
		if line[i] == ']' {
			depth--
		} else if line[i] == '[' {
			depth++
		} else if line[i] == ',' && depth == 0 {
			middle = i
		}
	}
	if middle == 1 {
		if len(line) == 3 {
			return &pair{left: &pair{value: int(line[0] - '0')}, right: &pair{value: int(line[2] - '0')}}
		} else {
			return &pair{left: &pair{value: int(line[0] - '0')}, right: parsePairs(line[middle+2 : len(line)-1])}
		}
	} else if middle+2 == len(line) {
		return &pair{left: parsePairs(line[1 : middle-1]), right: &pair{value: int(line[middle+1] - '0')}}
	} else {
		return &pair{left: parsePairs(line[1 : middle-1]), right: parsePairs(line[middle+2 : len(line)-1])}
	}
}

func getLeftRegular(tree *pair) *pair {
	if tree == nil {
		return nil
	}
	if tree.left == nil && tree.right == nil {
		return tree
	}
	return getLeftRegular(tree.left)
}

func getRightRegular(tree *pair) *pair {
	if tree == nil {
		return nil
	}
	if tree.left == nil && tree.right == nil {
		return tree
	}
	return getRightRegular(tree.right)
}

func (tree *pair) String() string {
	if tree.left == nil && tree.right == nil {
		return fmt.Sprintf("%d", tree.value)
	}
	return fmt.Sprintf("[%s,%s]", tree.left, tree.right)
}

func dosplit(tree *pair, leftmost *pair, rightmost *pair, count int) (*pair, bool) {
	operated := false
	if tree.left == nil && tree.right == nil {
		if tree.value >= 10 {
			return &pair{left: &pair{value: tree.value / 2}, right: &pair{value: tree.value - tree.value/2}}, true
		}
		return tree, operated
	}
	if tree.left != nil {
		tree.left, operated = dosplit(tree.left, leftmost, tree.right, count+1)
	}
	if tree.right != nil && !operated {
		tree.right, operated = dosplit(tree.right, tree.left, rightmost, count+1)
	}
	return tree, operated
}

func reduce(tree *pair, leftmost *pair, rightmost *pair, count int) (*pair, bool) {
	operated := false

	if count >= 4 && tree.left != nil {
		closestRight := getLeftRegular(rightmost)
		if closestRight != nil {
			closestRight.value += tree.right.value
		}
		closestLeft := getRightRegular(leftmost)
		if closestLeft != nil {
			closestLeft.value += tree.left.value
		}
		return &pair{value: 0}, true
	}
	if tree.left != nil {
		tree.left, operated = reduce(tree.left, leftmost, tree.right, count+1)
	}
	if tree.right != nil && !operated {
		tree.right, operated = reduce(tree.right, tree.left, rightmost, count+1)
	}
	return tree, operated
}

func reduceFully(tree *pair) *pair {
	operated := true
	for operated {
		exploded := true
		operated = false
		for exploded {
			tree, exploded = reduce(tree, nil, nil, 0)
			if exploded {
				operated = true
			}
		}
		split := false
		tree, split = dosplit(tree, nil, nil, 0)
		if split {
			operated = true
		}
	}
	return tree
}

func magnitude(tree *pair) int {
	if tree.left == nil {
		return tree.value
	}
	return 3*magnitude(tree.left) + 2*magnitude(tree.right)
}

func part1(input string) interface{} {
	var prevtree *pair
	for _, line := range strings.Split(input, "\n") {
		if prevtree == nil {
			prevtree = parsePairs(line[1 : len(line)-1])
		} else {
			prevtree = reduceFully(&pair{left: prevtree, right: parsePairs(line[1 : len(line)-1])})
		}
	}
	return magnitude(prevtree)
}

// low 4808
func part2(input string) interface{} {
	inputs := strings.Split(input, "\n")
	trees := make([]*pair, len(inputs))
	for i, line := range inputs {
		trees[i] = parsePairs(line[1 : len(line)-1])
	}
	res := 0
	for i, linei := range inputs {
		for j, linej := range inputs {
			if i == j {
				continue
			}
			mag := magnitude(reduceFully(&pair{left: parsePairs(linei[1 : len(linei)-1]), right: parsePairs(linej[1 : len(linej)-1])}))
			res = utils.IntMax(mag, res)
		}
	}
	return res
}

func main() {
	// defer profile.Start().Stop()
	utils.Solve(part1, part2, DAY)
}

package main

import (
	"adventofcodego/utils/testsuite"
	"testing"
)

func TestPart1(t *testing.T) {
	testsuite.RunTestforPart(t, part1, part2, 1, DAY)
}

func TestPart2(t *testing.T) {
	testsuite.RunTestforPart(t, part1, part2, 2, DAY)
}

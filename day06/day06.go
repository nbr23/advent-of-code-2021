package main

import (
	"adventofcodego/utils/inputs"
	"adventofcodego/utils/utils"
	"strings"
)

var COUNT_1 = 80
var COUNT_2 = 256

func inputToIntList(input string) (result []int64) {
	input = strings.TrimSpace(strings.Trim(input, "\n"))

	for _, i := range strings.Split(input, ",") {
		result = append(result, inputs.ParseDecInt(i))
	}
	return result
}

func iterateFishCycle(fish []int64) []int64 {
	fish_count := len(fish)

	for i := 0; i < fish_count; i++ {
		if fish[i] == 0 {
			fish[i] = 6
			fish = append(fish, 8)
		} else {
			fish[i]--
		}
	}
	return fish
}

func part1(input string) interface{} {
	fish := inputToIntList(input)

	for i := 0; i < COUNT_1; i++ {
		fish = iterateFishCycle(fish)
	}
	return len(fish)
}

type fishcache struct{
	fish int64
	count int64
}

func fishCycleRec(fish int64, count int64, cache map[fishcache]int64) int64 {
	if count <= fish {
		return 0
	}
	if fish == 0 {
		if _, ok := cache[fishcache{fish: 6, count: count-1}]; !ok {
			cache[fishcache{fish: 6, count: count-1}] = fishCycleRec(6, count - 1, cache)
		}
		if _, ok := cache[fishcache{fish: 8, count: count-1}]; !ok {
			cache[fishcache{fish: 8, count: count-1}] = fishCycleRec(8, count - 1,cache)
		}
		return 1 + cache[fishcache{fish: 6, count: count-1}] + cache[fishcache{fish: 8, count: count-1}]
	}
	if _, ok := cache[fishcache{fish: 0, count: count-fish}]; !ok {
		cache[fishcache{fish: 0, count: count-fish}] = fishCycleRec(0, count - fish, cache)
	}
	return cache[fishcache{fish: 0, count: count-fish}]
}

func part2(input string) interface{} {
	fish := inputToIntList(input)
	fish_count := int64(len(fish))
	cache :=  make(map[fishcache]int64)

	for i := range fish {
		fish_count += fishCycleRec(fish[i], int64(COUNT_2), cache)
	}
	return fish_count
}

func main() {
	var day int64 = 6

	utils.Solve(part1, part2, day)
}

package main

import (
	// "adventofcodego/utils/inputs"
	"adventofcodego/utils/utils"
	"strings"
)

func makeLinks(input string) map[string][]string {
	links := make(map[string][]string)
	for _, line := range strings.Split(input, "\n") {
		sp := strings.Split(line, "-")
		if _, ok := links[sp[0]]; ok {
			links[sp[0]] = append(links[sp[0]], sp[1])
		} else {
			links[sp[0]] = make([]string, 1)
			links[sp[0]][0] = sp[1]
		}
		if _, ok := links[sp[1]]; ok {
			links[sp[1]] = append(links[sp[1]], sp[0])
		} else {
			links[sp[1]] = make([]string, 1)
			links[sp[1]][0] = sp[0]
		}
	}
	return links
}

func followPath(links map[string][]string, visited []string, current string) int {
	score := 0
	if current == "end" {
		return 1
	}
	if current != "start" && strings.ToLower(current) == current {
		for i := range visited {
			if current == visited[i] {
				return 0
			}
		}
	}
	visited = append(visited, current)
	for _, next := range links[current] {
		if next == "start" {
			continue
		}
		score += followPath(links, visited, next)
	}
	return score
}

func part1(input string) interface{} {
	links := makeLinks(input)
	return followPath(links, make([]string, 0), "start")
}

func followPath2(links map[string][]string, visited []string, current string) int {
	if current == "end" {
		return 1
	}
	score := 0
	if current != "start" && strings.ToLower(current) == current {
		visited_count := make(map[string]int)
		var visited_twice string
		for i := range visited {
			if strings.ToLower(visited[i]) == visited[i] {
				visited_count[visited[i]]++
				if visited_count[visited[i]] >= 2 {
					visited_twice = visited[i]
				}
			}
		}
		if visited_twice != "" && visited_count[current] > 0 {
			return 0
		}
	}
	visited = append(visited, current)
	for _, next := range links[current] {
		if next == "start" {
			continue
		}
		score += followPath2(links, visited, next)
	}
	return score
}

func part2(input string) interface{} {
	links := makeLinks(input)
	return followPath2(links, make([]string, 0), "start")
}

func main() {
	var day int64 = 12

	utils.Solve(part1, part2, day)
}

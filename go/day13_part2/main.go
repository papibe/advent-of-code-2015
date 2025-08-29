package main

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type AdjacencyMatrix [][]int

func parse(filename string) AdjacencyMatrix {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic("No file found!")
	}
	lines := strings.Split(strings.Trim(string(data), "\n"), "\n")
	re_line := regexp.MustCompile(`(\w+) would (\w+) (\d+) happiness units by sitting next to (\w+)\.`)

	// build basic rules
	happiness := make(map[string]map[string]int)

	for _, line := range lines {
		matches := re_line.FindStringSubmatch(line)

		person := matches[1]
		str_action := matches[2]

		var action int
		if str_action == "gain" {
			action = 1
		} else {
			action = -1
		}
		happines_units, _ := strconv.Atoi(matches[3])
		units := action * happines_units
		partner := matches[4]

		_, person_in_happiness := happiness[person]
		if !person_in_happiness {
			happiness[person] = make(map[string]int)
		}
		happiness[person][partner] = units
	}

	// patch for part 2: add yourself
	happiness["yourself"] = make(map[string]int)
	for person := range happiness {
		happiness["yourself"][person] = 0
	}

	// bijection
	names := []string{}
	indexes := make(map[string]int)
	index := 0
	for person := range happiness {
		names = append(names, person)
		indexes[person] = index
		index++
	}
	_ = names

	// create adjancency matrix
	am := AdjacencyMatrix{}
	for range len(happiness) {
		new_row := make([]int, len(happiness))
		am = append(am, new_row)
	}

	for person, rules := range happiness {
		for partner, units := range rules {
			am[indexes[person]][indexes[partner]] = units
		}
	}
	return am
}

func all_visited(visited []bool) bool {
	for _, value := range visited {
		if !value {
			return false
		}
	}
	return true
}

func solve(adjacency_matrix AdjacencyMatrix) int {
	max_happiness := math.MinInt
	visited := make([]bool, len(adjacency_matrix))

	var dfs func(int, int)

	dfs = func(person, happiness int) {
		if all_visited(visited) {
			happiness += adjacency_matrix[person][0] + adjacency_matrix[0][person]
			max_happiness = max(max_happiness, happiness)
		}

		for partner, is_visited := range visited {
			if !is_visited {
				visited[partner] = true
				new_happines := happiness + adjacency_matrix[person][partner] + adjacency_matrix[partner][person]
				dfs(partner, new_happines)
				visited[partner] = false
			}
		}
	}

	visited[0] = true
	dfs(0, 0)

	return max_happiness
}

func solution(filename string) int {
	adjacency_matrix := parse(filename)
	return solve(adjacency_matrix)
}

func main() {
	fmt.Println(solution("./input.txt")) // 640
}

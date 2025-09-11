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
		panic("file error")
	}
	lines := strings.Split(strings.Trim(string(data), "\n"), "\n")
	re_line := regexp.MustCompile(`(\w+) to (\w+) = (\d+)`)

	// first pass: read all data into a dictionary
	nodes := make(map[string]map[string]int)
	for _, line := range lines {
		matches := re_line.FindStringSubmatch(line)

		node1 := matches[1]
		node2 := matches[2]
		distance, _ := strconv.Atoi(matches[3])

		_, node1_in_nodes := nodes[node1]
		if !node1_in_nodes {
			nodes[node1] = map[string]int{node1: distance}
		} else {
			nodes[node1][node2] = distance
		}

		_, node2_in_nodes := nodes[node2]
		if !node2_in_nodes {
			nodes[node2] = map[string]int{node1: distance}
		} else {
			nodes[node2][node1] = distance
		}
	}

	// bijection
	node_list := make([]string, len(nodes))
	node_map := make(map[string]int)
	index := 0
	for node_name := range nodes {
		node_list[index] = node_name
		node_map[node_name] = index
		index++
	}

	// create adjacency matrix
	am := make(AdjacencyMatrix, len(nodes))
	for row := range len(nodes) {
		am[row] = make([]int, len(nodes))
		for col := range len(nodes) {
			am[row][col] = math.MaxInt
		}
	}

	// set distance to itself
	for index := range len(nodes) {
		am[index][index] = 0
	}

	// set other distances
	for node1, distances := range nodes {
		for node2, distance := range distances {
			index1 := node_map[node1]
			index2 := node_map[node2]

			am[index1][index2] = distance
			am[index2][index1] = distance
		}
	}

	return am
}

func all_visited(visited []bool) bool {
	for _, is_visited := range visited {
		if !is_visited {
			return false
		}
	}
	return true
}

func solve(adjancency_matrix AdjacencyMatrix) int {
	min_distance := math.MaxInt
	visited := make([]bool, len(adjancency_matrix))
	for index := range len(adjancency_matrix) {
		visited[index] = false
	}

	var dfs func(int, int)

	dfs = func(node, steps int) {
		if all_visited(visited) {
			min_distance = min(min_distance, steps)
			return
		}

		for next_node, is_visited := range visited {
			if !is_visited {
				visited[next_node] = true
				dfs(next_node, steps+adjancency_matrix[node][next_node])
				visited[next_node] = false
			}
		}
	}

	// start at every node
	for initial_node, is_visited := range visited {
		if !is_visited {
			visited[initial_node] = true

			// travel starting at initial node
			dfs(initial_node, 0)
			visited[initial_node] = false
		}
	}

	return min_distance
}

func solution(filename string) int {
	adjacency_matrix := parse(filename)
	return solve(adjacency_matrix)
}

func main() {
	fmt.Println(solution("./example.txt")) // 605
	fmt.Println(solution("./input.txt"))   // 141
}

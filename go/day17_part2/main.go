package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parse(filename string) []int {
	data, err := os.ReadFile(filename)

	if err != nil {
		panic("file error")
	}
	lines := strings.Split(strings.Trim(string(data), "\n"), "\n")

	containers := []int{}

	for _, line := range lines {
		liters, _ := strconv.Atoi(line)
		containers = append(containers, liters)
	}

	return containers
}

func solve(containers []int, capacity int) int {
	combinations := 0
	min_combinations := len(containers) - 1

	var dfs func(int, int, int)
	dfs = func(start_index, remaining_capacity, selected_containers int) {
		if remaining_capacity == 0 {
			if selected_containers < min_combinations {
				min_combinations = selected_containers
				combinations = 1
			} else if selected_containers == min_combinations {
				combinations++
			}
			return
		}

		for container := start_index; container < len(containers); container++ {
			container_capacity := containers[container]
			if remaining_capacity-container_capacity >= 0 {
				dfs(
					container+1,
					remaining_capacity-container_capacity,
					selected_containers+1,
				)
			}
		}

	}
	dfs(0, capacity, 0)
	return combinations
}

func solution(filename string, capacity int) int {
	containers := parse(filename)
	return solve(containers, capacity)
}

func main() {
	fmt.Println(solution("./example.txt", 25)) // 3
	fmt.Println(solution("./input.txt", 150))  // 57
}

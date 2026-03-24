package main

import (
	"fmt"
	"math"
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
	packets := []int{}

	for _, line := range lines {
		packet, _ := strconv.Atoi(line)
		packets = append(packets, packet)
	}
	return packets
}

func get_entanglement(packets []int) int {
	entanglement := 1
	for _, packet := range packets {
		entanglement *= packet
	}
	return entanglement
}

func solve(packets []int, group_size int) int {
	packets_sum := 0
	for _, packet := range packets {
		packets_sum += packet
	}
	target_weight := packets_sum / group_size

	min_size := math.MaxInt
	size_winners := []int{}

	var n_sum_k func(int, int, []int)
	n_sum_k = func(cs, index int, selection []int) {
		if index < 0 {
			return
		}
		if cs > target_weight {
			return
		}
		if len(selection) > min_size {
			return
		}

		if cs == target_weight {
			if len(selection) < min_size {
				min_size = len(selection)
				size_winners = []int{get_entanglement(selection)}
			} else if len(selection) == min_size {
				size_winners = append(size_winners, get_entanglement(selection))
			}
		}

		next_index := index - 1
		for next_index >= 0 {
			packet := packets[next_index]
			if cs+packet <= target_weight {
				new_selection := make([]int, len(selection))
				copy(new_selection, selection)
				new_selection = append(new_selection, packet)
				n_sum_k(cs+packet, next_index, new_selection)
			}
			next_index--
		}

	}

	n_sum_k(0, len(packets), []int{})

	min_entanglement := math.MaxInt
	for _, entanglement := range size_winners {
		min_entanglement = min(min_entanglement, entanglement)
	}
	return min_entanglement
}

func solution(filename string, group_size int) int {
	packets := parse(filename)
	return solve(packets, group_size)
}

func main() {
	fmt.Println("Part 1:")
	fmt.Println("  Example:", solution("./example.txt", 3)) // 99
	fmt.Println("  Input  :", solution("./input.txt", 3))   // 11266889531

	fmt.Println("Part 2:")
	fmt.Println("  Example:", solution("./example.txt", 4)) // 44
	fmt.Println("  Input  :", solution("./input.txt", 4))   // 77387711
}

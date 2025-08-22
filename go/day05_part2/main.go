package main

import (
	"fmt"
	"os"
	"strings"
)

func parse(filename string) []string {
	data, err := os.ReadFile(filename)

	if err != nil {
		panic("file error")
	}
	return strings.Split(strings.Trim(string(data), "\n"), "\n")
}

func is_nice(s string) bool {
	// check rule 1
	rule1 := false
	pairs := make(map[string][]int)

rule1_loop:
	for i := range len(s) - 1 {
		pair := s[i : i+2]
		_, is_in_pairs := pairs[pair]
		if !is_in_pairs {
			pairs[pair] = []int{}
		}
		pairs[pair] = append(pairs[pair], i)

		n := len(pairs[pair])
		if n > 1 {
			for j := range n {
				for k := j + 1; k < n; k++ {
					if pairs[pair][k] > pairs[pair][j]+1 {
						rule1 = true
						break rule1_loop
					}
				}
			}
		}
	}

	// check rule 2
	rule2 := false
	for i := range len(s) - 2 {
		if s[i] == s[i+2] {
			rule2 = true
			break
		}
	}

	return rule1 && rule2
}

func solve(strs []string) int {
	total := 0

	for _, str := range strs {
		if is_nice(str) {
			total++
		}
	}
	return total
}

func solution(filename string) int {
	strs := parse(filename)
	return solve(strs)
}

func main() {
	fmt.Println(solution("./input.txt")) // 43
}

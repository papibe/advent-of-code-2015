package main

import (
	"fmt"
	"os"
	"strings"
)

const BASEMENT = -1
const INVALID = 0

var VALUE = map[rune]int{
	'(': 1, ')': -1,
}

func parse(filename string) string {
	data, err := os.ReadFile(filename)

	if err != nil {
		panic("file error")
	}

	return strings.Trim(string(data), "\n")
}

func solve(data string) int {
	floor := 0

	for index, char := range data {
		floor += VALUE[char]
		if floor == BASEMENT {
			return index + 1
		}
	}
	return INVALID
}

func solution(filename string) int {
	data := parse(filename)
	return solve(data)
}

func main() {
	fmt.Println(solution("./input.txt")) // 1783
}

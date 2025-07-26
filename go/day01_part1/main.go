package main

import (
	"fmt"
	"os"
	"strings"
)

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

	for _, char := range data {
		floor += VALUE[char]
	}
	return floor
}

func solution(filename string) int {
	data := parse(filename)
	return solve(data)
}

func main() {
	fmt.Println(solution("./input.txt")) // 232
}

package main

import (
	"fmt"
	"os"
	"strings"
)

type Position struct {
	row int
	col int
}

var DIR_STEPS = map[rune]Position{
	'^': Position{-1, 0},
	'v': Position{1, 0},
	'>': Position{0, 1},
	'<': Position{0, -1},
}

func parse(filename string) string {
	data, err := os.ReadFile(filename)

	if err != nil {
		panic("file error")
	}
	return strings.Trim(string(data), "\n")
}

func solve(directions string) int {
	houses := NewSet[Position]()

	current := Position{0, 0}
	houses.add(current)

	for _, direction := range directions {
		current.row += DIR_STEPS[direction].row
		current.col += DIR_STEPS[direction].col
		houses.add(current)
	}

	return houses.len()
}

func solution(filename string) int {
	directions := parse(filename)
	return solve(directions)
}

func main() {
	fmt.Println(solution("./input.txt")) // 2081
}

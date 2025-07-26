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
	'^': {-1, 0},
	'v': {1, 0},
	'>': {0, 1},
	'<': {0, -1},
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

	santa := Position{0, 0}
	robo := Position{0, 0}
	houses.add(santa)

	for index, direction := range directions {
		if (index+1)%2 == 0 {
			santa.row += DIR_STEPS[direction].row
			santa.col += DIR_STEPS[direction].col
			houses.add(santa)
		} else {
			robo.row += DIR_STEPS[direction].row
			robo.col += DIR_STEPS[direction].col
			houses.add(robo)
		}
	}

	return houses.len()
}

func solution(filename string) int {
	directions := parse(filename)
	return solve(directions)
}

func main() {
	fmt.Println(solution("./input.txt")) // 2341
}

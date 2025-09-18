package main

import (
	"fmt"
	"os"
	"strings"
)

const ON = '#'
const OFF = '.'

func parse(filename string) [][]rune {
	data, err := os.ReadFile(filename)

	if err != nil {
		panic("file error")
	}
	lines := strings.Split(strings.Trim(string(data), "\n"), "\n")

	grid := [][]rune{}

	for _, line := range lines {
		grid_row := []rune{}
		for _, char := range line {
			grid_row = append(grid_row, char)
		}
		grid = append(grid, grid_row)
	}
	return grid
}

func number_of_on_neighbors(row, col int, grid [][]rune) int {
	rows := len(grid)
	cols := len(grid[0])
	on_neighbors := 0

	for _, coords := range [][]int{
		{row - 1, col - 1},
		{row - 1, col},
		{row - 1, col + 1},
		{row, col - 1},
		{row, col + 1},
		{row + 1, col - 1},
		{row + 1, col},
		{row + 1, col + 1},
	} {
		nrow := coords[0]
		ncol := coords[1]
		if 0 <= nrow && nrow < rows && 0 <= ncol && ncol < cols {
			if grid[nrow][ncol] == ON {
				on_neighbors++
			}
		}
	}

	return on_neighbors
}

func solve(grid [][]rune, cycles int) int {
	rows := len(grid)
	cols := len(grid[0])

	for range cycles {
		// generate next grid
		next_grid := make([][]rune, rows)
		for i := range rows {
			grid_row := make([]rune, cols)
			for j := range cols {
				grid_row[j] = '+'
			}
			next_grid[i] = grid_row
		}

		// cycle
		for row, line := range grid {
			for col, light := range line {
				noon := number_of_on_neighbors(row, col, grid)

				if light == '+' {
					panic("what?")
				}

				if light == ON {
					if noon == 2 || noon == 3 {
						next_grid[row][col] = ON
					} else {
						next_grid[row][col] = OFF
					}

					// light OFF
				} else {
					if noon == 3 {
						next_grid[row][col] = ON
					} else {
						next_grid[row][col] = OFF
					}
				}
			}
		}
		grid = next_grid
	}
	lights_on := 0
	for _, line := range grid {
		for _, char := range line {
			if char == ON {
				lights_on++
			}
		}
	}
	return lights_on
}

func solution(filename string, cycles int) int {
	grid := parse(filename)
	return solve(grid, cycles)
}

func main() {
	fmt.Println(solution("./example1.txt", 4)) // 4
	fmt.Println(solution("./input.txt", 100))  // 814
}

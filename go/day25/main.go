package main

import (
	"fmt"
)

func get_order(row, col int) int {
	return ((col+row)*(col+row) - col - 3*row + 2) / 2
}

func solution(row, col int) int {
	position := get_order(row, col)

	code := 20151125
	for range position - 1 {
		code = (code * 252533) % 33554393
	}
	return code
}

func main() {
	fmt.Println(solution(2978, 3083)) // 2650453
}

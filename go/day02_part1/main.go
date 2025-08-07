package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Dimension struct {
	L int
	w int
	h int
}

func parse(filename string) []Dimension {
	data, err := os.ReadFile(filename)

	if err != nil {
		panic("file error")
	}
	lines := strings.Split(strings.Trim(string(data), "\n"), "\n")

	dimensions := []Dimension{}
	for _, line := range lines {
		parts := strings.Split(line, "x")
		L, _ := strconv.Atoi(parts[0])
		w, _ := strconv.Atoi(parts[1])
		h, _ := strconv.Atoi(parts[2])

		dimensions = append(dimensions, Dimension{L, w, h})
	}
	return dimensions
}

func solve(dimensions []Dimension) int {
	wrapping_paper := 0

	for _, dim := range dimensions {
		area := (2 * dim.L * dim.w) + (2 * dim.w * dim.h) + (2 * dim.h * dim.L)
		small_side := min(dim.L*dim.w, dim.w*dim.h, dim.h*dim.L)
		wrapping_paper += area + small_side
	}
	return wrapping_paper
}

func solution(filename string) int {
	dimensions := parse(filename)
	return solve(dimensions)
}

func main() {
	fmt.Println(solution("./input.txt")) // 1598415
}

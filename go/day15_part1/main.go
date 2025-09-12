package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Ingredient struct {
	name       string
	capacity   int
	durability int
	flavor     int
	texture    int
	calories   int
}

func parse(filename string) []Ingredient {
	data, err := os.ReadFile(filename)

	if err != nil {
		panic("file error")
	}
	lines := strings.Split(strings.Trim(string(data), "\n"), "\n")

	re_line := regexp.MustCompile(`(\w+): capacity ([0-9-]+), durability ([0-9-]+), flavor ([0-9-]+), texture ([0-9-]+), calories ([0-9-]+)`)
	ingredients := []Ingredient{}

	for _, line := range lines {
		matches := re_line.FindStringSubmatch(line)

		name := matches[1]
		capacity, _ := strconv.Atoi(matches[2])
		durability, _ := strconv.Atoi(matches[3])
		flavor, _ := strconv.Atoi(matches[4])
		texture, _ := strconv.Atoi(matches[5])
		calories, _ := strconv.Atoi(matches[6])

		ingredients = append(ingredients, Ingredient{name, capacity, durability, flavor, texture, calories})
	}

	return ingredients
}

func partition(n, k int) [][]int {
	solutions := [][]int{}

	var _partitions func(n, k int, numbers []int)
	_partitions = func(n, k int, numbers []int) {
		if k == 0 {
			if n == 0 {
				solutions = append(solutions, numbers)
			}
			return
		}
		if n == 0 {
			return
		}
		for i := 1; i < n-k+1+1; i++ {
			new_numbers := append([]int{i}, numbers...)
			_partitions(n-i, k-1, new_numbers)
		}

	}
	_partitions(n, k, []int{})
	return solutions
}

func solve(ingredients []Ingredient) int {
	highest_scoring := 0
	n := len(ingredients)
	_ = n

	for _, part_list := range partition(100, n) {
		capacity_score := 0
		durability_score := 0
		flavor_score := 0
		texture_score := 0

		for index, ingredient := range ingredients {
			capacity_score += ingredient.capacity * part_list[index]
			durability_score += ingredient.durability * part_list[index]
			flavor_score += ingredient.flavor * part_list[index]
			texture_score += ingredient.texture * part_list[index]
		}
		var score int
		if capacity_score < 0 || durability_score < 0 || flavor_score < 0 || texture_score < 0 {
			score = 0
		} else {
			score = capacity_score * durability_score * flavor_score * texture_score
		}
		highest_scoring = max(highest_scoring, score)

	}

	return highest_scoring
}

func solution(filename string, timeframe int) int {
	ingredients := parse(filename)
	return solve(ingredients)
}

func main() {
	fmt.Println(solution("./example.txt", 1000)) // 62842880
	fmt.Println(solution("./input.txt", 2503))   // 18965440
}

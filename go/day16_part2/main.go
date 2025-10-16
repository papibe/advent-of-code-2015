package main

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var AUNT_SUE_EXACTLY = map[string]int{
	"children": 3,
	"samoyeds": 2,
	"akitas":   0,
	"vizslas":  0,
	"cars":     2,
	"perfumes": 1,
}

var AUNT_SUE_GREATER = map[string]int{
	"cats":  7,
	"trees": 3,
}

var AUNT_SUE_LOWER = map[string]int{
	"pomeranians": 3,
	"goldfish":    5,
}

func parse(filename string) []map[string]int {
	data, err := os.ReadFile(filename)

	if err != nil {
		panic("file error")
	}
	lines := strings.Split(strings.Trim(string(data), "\n"), "\n")

	re_line := regexp.MustCompile(`Sue (\d+): (.*)`)
	re_thing := regexp.MustCompile(`(\w+): (\d+)`)

	aunts := []map[string]int{}

	for _, line := range lines {
		matches := re_line.FindStringSubmatch(line)

		things := matches[2]
		aunt_things := make(map[string]int)

		for _, thing := range strings.Split(things, ", ") {
			thing_match := re_thing.FindStringSubmatch(thing)

			name := thing_match[1]
			amount, _ := strconv.Atoi(thing_match[2])
			aunt_things[name] = amount
		}
		aunts = append(aunts, aunt_things)
	}

	return aunts
}

func solve(aunts []map[string]int) int {
	max_score := math.MinInt
	max_index := -1
	var max_scores *Set[int]

	for i, aunt := range aunts {
		aunt_score := 0

		for ingredient_name, ingredient_count := range aunt {
			// exact count
			value, ok := AUNT_SUE_EXACTLY[ingredient_name]
			if ok && value == ingredient_count {
				aunt_score++
			}

			// greater count
			value, ok = AUNT_SUE_GREATER[ingredient_name]
			if ok && ingredient_count > value {
				aunt_score++
			}

			// fewer count
			value, ok = AUNT_SUE_LOWER[ingredient_name]
			if ok && ingredient_count < value {
				aunt_score++
			}

		}

		if aunt_score > max_score {
			max_score = aunt_score
			max_scores = NewSet[int]()
			max_scores.add(aunt_score)
			max_index = i + 1
		} else if aunt_score == max_score {
			max_scores.add(aunt_score)
		}
	}

	return max_index
}

func solution(filename string) int {
	aunts := parse(filename)
	return solve(aunts)
}

func main() {
	fmt.Println(solution("./input.txt")) // 323
}

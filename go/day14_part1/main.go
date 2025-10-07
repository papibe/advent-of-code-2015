package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)


type ReinDeer struct {
	speed  int
	sprint int
	rest   int
}

func parse(filename string) []ReinDeer {
	data, err := os.ReadFile(filename)

	if err != nil {
		panic("file error")
	}
	lines := strings.Split(strings.Trim(string(data), "\n"), "\n")

	re_line := regexp.MustCompile(`(\w+) can fly (\d+) km/s for (\d+) seconds, but then must rest for (\d+) seconds.`)
	reindeers := []ReinDeer{}

	for _, line := range lines {
		matches := re_line.FindStringSubmatch(line)

		speed, _ := strconv.Atoi(matches[2])
		sprint, _ := strconv.Atoi(matches[3])
		rest, _ := strconv.Atoi(matches[4])

		reindeers = append(reindeers, ReinDeer{speed, sprint, rest})
	}

	return reindeers
}

func solve(reindeers []ReinDeer, timeframe int) int {
	max_distance := 0

	for _, r := range reindeers {
		cycles := timeframe / (r.sprint + r.rest)
		total_time_remaining := timeframe % (r.sprint + r.rest)

		distance := cycles * r.speed * r.sprint
		seconds_to_sprint_available := min(r.sprint, total_time_remaining)
		distance += r.speed * seconds_to_sprint_available

		max_distance = max(max_distance, distance)
	}

	return max_distance
}

func solution(filename string, timeframe int) int {
	reindeers := parse(filename)
	return solve(reindeers, timeframe)
}

func main() {
	fmt.Println(solution("./example.txt", 1000)) // 1120
	fmt.Println(solution("./input.txt", 2503))   // 2655
}

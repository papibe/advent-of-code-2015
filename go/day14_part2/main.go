package main

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const REST = 0
const SPRINT = 1

type ReinDeer struct {
	speed    int
	sprint   int
	rest     int
	points   int
	position int
	state    int
	counter  int
}

func (r *ReinDeer) move() int {
	if r.counter == 0 {
		if r.state == REST {
			r.state = SPRINT
			r.counter = r.sprint
		} else {
			r.state = REST
			r.counter = r.rest
		}
	}
	r.counter -= 1

	if r.state == SPRINT {
		r.position += r.speed
	}

	return r.position
}

func (r *ReinDeer) add_point() {
	r.points += 1
}

func parse(filename string) []*ReinDeer {
	data, err := os.ReadFile(filename)

	if err != nil {
		panic("file error")
	}
	lines := strings.Split(strings.Trim(string(data), "\n"), "\n")

	re_line := regexp.MustCompile(`(\w+) can fly (\d+) km/s for (\d+) seconds, but then must rest for (\d+) seconds.`)
	reindeers := []*ReinDeer{}

	for _, line := range lines {
		matches := re_line.FindStringSubmatch(line)

		speed, _ := strconv.Atoi(matches[2])
		sprint, _ := strconv.Atoi(matches[3])
		rest, _ := strconv.Atoi(matches[4])

		reindeers = append(reindeers, &ReinDeer{speed, sprint, rest, 0, 0, REST, 0})
	}

	return reindeers
}

func solve(reindeers []*ReinDeer, timeframe int) int {

	for range timeframe {
		max_distance := 0
		var max_reindeers []*ReinDeer

		for _, reindeer := range reindeers {
			distance := reindeer.move()
			if distance > max_distance {
				max_distance = distance
				max_reindeers = []*ReinDeer{reindeer}
			} else if distance == max_distance {
				max_reindeers = append(max_reindeers, reindeer)
			}
		}
		for _, reindeer := range max_reindeers {
			reindeer.add_point()
		}
	}

	max_points := math.MinInt
	for _, reindeer := range reindeers {
		max_points = max(max_points, reindeer.points)
	}

	return max_points
}

func solution(filename string, timeframe int) int {
	reindeers := parse(filename)
	return solve(reindeers, timeframe)
}

func main() {
	fmt.Println(solution("./example.txt", 1000)) // 689
	fmt.Println(solution("./input.txt", 2503))   // 1059
}

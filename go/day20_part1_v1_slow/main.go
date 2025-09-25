package main

import (
	"fmt"
	"math"
)

func get_divisors(n int) []int {
	divisors := []int{}
	limit := int(math.Sqrt(float64(n))) + 1

	for divisor := 1; divisor < limit; divisor++ {
		if n%divisor == 0 {
			divisors = append(divisors, divisor)
			if divisor*divisor != n {
				divisors = append(divisors, n/divisor)
			}
		}
	}
	return divisors
}

func get_presents(n int) int {
	gifts := 0
	for _, divisor := range get_divisors(n) {
		gifts += divisor
	}
	return 10 * gifts
}

func solution(goal int) int {
	house_number := 0
	presents := 0

	for presents < goal {
		house_number++
		presents = get_presents(house_number)
	}
	return house_number
}

func main() {
	fmt.Println(solution(34_000_000)) // 786240
}

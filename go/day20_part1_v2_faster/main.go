package main

import (
	"fmt"
)

func get_prime_factors(n int) [][2]int {
	prime_factors := [][2]int{}
	for d := 2; d*d <= n; d++ {
		times := 0
		for n%d == 0 {
			times += 1
			n /= d
		}
		if times > 0 {
			prime_factors = append(prime_factors, [2]int{d, times})
		}
	}
	if n > 1 {
		prime_factors = append(prime_factors, [2]int{n, 1})
	}
	return prime_factors
}

func get_divisors(prime_factors [][2]int) []int {
	divisors := []int{}
	var _get_divisors func(int, int)
	_get_divisors = func(index, divisor int) {
		if index == len(prime_factors) {
			divisors = append(divisors, divisor)
			return
		}
		prime := prime_factors[index][0]
		exponent := prime_factors[index][1]
		value := 1
		for range exponent + 1 {
			_get_divisors(index+1, divisor*value)
			value *= prime
		}
	}
	_get_divisors(0, 1)
	return divisors
}

func get_presents(n int) int {
	prime_factors := get_prime_factors(n)
	gifts := 0
	for _, divisor := range get_divisors(prime_factors) {
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

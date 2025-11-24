package main

import (
	"fmt"
	"math"
)

func sieve_of_eratosthenes(limit int) []int {
	prime := make([]bool, limit+1)
	for i := range limit + 1 {
		prime[i] = true
	}

	for p := 2; p*p <= limit; p++ {
		if prime[p] {
			for i := p * p; i < limit+1; i += p {
				prime[i] = false
			}
		}
	}

	primes := []int{}
	for p := 2; p < limit+1; p++ {
		if prime[p] {
			primes = append(primes, p)
		}
	}
	return primes
}

func get_prime_factors(primes []int, n int) [][2]int {
	if n == 1 {
		return [][2]int{{1, 2}}
	}
	prime_factors := [][2]int{}

	for _, p := range primes {
		if p*p > n {
			break
		}
		times := 0
		for n%p == 0 {
			times += 1
			n /= p
		}
		if times > 0 {
			prime_factors = append(prime_factors, [2]int{p, times})
		}
	}
	if n > 1 {
		prime_factors = append(prime_factors, [2]int{n, 1})
	}
	return prime_factors
}

func sum_of_divisors(prime_factors [][2]int) int {
	result := 1.0
	for _, prime_factor := range prime_factors {
		value := prime_factor[0]
		exponent := prime_factor[1]
		result *= (math.Pow(float64(value), float64(exponent+1)) - 1) / (float64(value) - 1)
	}
	return int(result)
}

func get_presents(primes []int, n int) int {
	prime_factors := get_prime_factors(primes, n)
	gifts := sum_of_divisors(prime_factors)
	return 10 * gifts
}

func solution(goal int) int {
	house_number := 0
	presents := 0
	primes := sieve_of_eratosthenes(goal)

	for presents < goal {
		house_number++
		presents = get_presents(primes, house_number)
	}
	return house_number
}

func main() {
	fmt.Println(solution(34_000_000)) // 786240
}

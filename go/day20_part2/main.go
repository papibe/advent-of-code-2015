package main

import (
	"fmt"
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

func get_divisors(prime_factors [][2]int, n int) []int {
	divisors := []int{}
	var _get_divisors func(int, int)
	_get_divisors = func(index, divisor int) {
		if index == len(prime_factors) {
			if divisor*50 >= n {
				divisors = append(divisors, divisor)
			}
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

func get_presents(primes []int, n int) int {
	prime_factors := get_prime_factors(primes, n)
	gifts := 0
	for _, divisor := range get_divisors(prime_factors, n) {
		gifts += divisor
	}
	return 11 * gifts
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

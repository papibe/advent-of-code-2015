from typing import List, Tuple


def sieve_of_eratosthenes(limit: int) -> List[int]:
    """
    Finds all prime numbers up to 'limit' using the Sieve of Eratosthenes.
    """
    # Create a boolean array "prime[0..limit]" and initialize all entries it as true.
    # A value prime[i] will be false if i is not a prime, else true.
    prime: List[bool] = [True] * (limit + 1)

    p: int = 2
    while p * p <= limit:
        # If prime[p] is still true, then it is a prime
        if prime[p]:
            # Update all multiples of p as not prime
            for i in range(p * p, limit + 1, p):
                prime[i] = False
        p += 1

    # Collect all prime numbers
    primes: List[int] = []
    for p in range(2, limit + 1):
        if prime[p]:
            primes.append(p)
    return primes


def get_prime_factors(primes: List[int], n: int) -> List[Tuple[int, int]]:
    prime_factors: List[Tuple[int, int]] = []

    for p in primes:
        if p * p > n:
            break
        times: int = 0
        while n % p == 0:
            times += 1
            n //= p
        if times > 0:
            prime_factors.append((p, times))
    if n > 1:
        prime_factors.append((n, 1))

    return prime_factors


def sum_of_divisors(prime_factors: List[Tuple[int, int]]) -> int:
    result: int = 1

    for prime, exponent in prime_factors:
        result *= (prime ** (exponent + 1) - 1) / (prime - 1)
    return result


def get_presents(primes: List[int], n: int) -> int:
    prime_factors: List[Tuple[int, int]] = get_prime_factors(primes, n)
    presents: int = 10 * sum_of_divisors(prime_factors)

    return presents


def solution(goal: int) -> int:
    house_number: int = 1
    presents: int = 10
    primes: List[int] = sieve_of_eratosthenes(goal)

    while presents < goal:
        house_number += 1
        presents = get_presents(primes, house_number)

    return house_number


if __name__ == "__main__":
    # it takes 3.7s
    print(solution(34_000_000))

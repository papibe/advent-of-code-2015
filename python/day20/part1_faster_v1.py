from typing import List, Tuple


def get_prime_factors(n: int) -> List[Tuple[int, int]]:
    prime_factors: List[Tuple[int, int]] = []
    d: int = 2
    while d * d <= n:
        times: int = 0
        while n % d == 0:
            times += 1
            n //= d
        if times > 0:
            prime_factors.append((d, times))
        d += 1
    if n > 1:
        prime_factors.append((n, 1))

    return prime_factors


def get_divisors(prime_factors: List[Tuple[int, int]]) -> List[int]:
    divisors: List[int] = []

    def _get_divisors(index: int, divisor: int) -> None:
        if index == len(prime_factors):
            divisors.append(divisor)
            return

        prime: int = prime_factors[index][0]
        exponent: int = prime_factors[index][1]
        value: int = 1
        for _ in range(exponent + 1):
            _get_divisors(index + 1, divisor * value)
            value *= prime

    _get_divisors(0, 1)
    return divisors


def get_presents(n: int) -> int:
    prime_factors: List[Tuple[int, int]] = get_prime_factors(n)
    divisors: List[int] = get_divisors(prime_factors)
    presents: int = 10 * sum(divisors)

    return presents


def solution(goal: int) -> int:
    house_number: int = 0
    presents: int = 0

    while presents < goal:
        house_number += 1
        presents = get_presents(house_number)

    return house_number


if __name__ == "__main__":
    # it takes 9s
    print(solution(34_000_000))  # 786240

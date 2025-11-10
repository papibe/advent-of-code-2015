import math
from typing import List


def get_divisors(n: int) -> List[int]:
    divisors: List[int] = []
    divisor: int = 1

    for divisor in range(1, int(math.sqrt(n)) + 1):
        if n % divisor == 0:
            divisors.append(divisor)
            if divisor * divisor != n:
                divisors.append(n // divisor)

    return divisors


def get_presents(n: int) -> int:
    return 10 * sum(get_divisors(n))


def solution(goal: int) -> int:
    house_number: int = 0
    presents: int = 0

    while presents < goal:
        house_number += 1
        presents = get_presents(house_number)

    return house_number


if __name__ == "__main__":
    # it takes 12s
    print(solution(34_000_000))  # 786240

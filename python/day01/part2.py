from typing import Dict

VALUE: Dict[str, int] = {"(": 1, ")": -1}
BASEMENT: int = -1
INVALID: int = 0


def parse(filename: str) -> str:
    with open(filename, "r") as fp:
        data: str = fp.read().strip()

    return data


def solve(data: str) -> int:
    floor: int = 0

    for position, char in enumerate(data, 1):
        floor += VALUE[char]
        if floor == BASEMENT:
            return position

    return INVALID


def solution(filename: str) -> int:
    data: str = parse(filename)
    return solve(data)


if __name__ == "__main__":
    print(solution("./input.txt"))  # 1783

import re
from collections import namedtuple
from typing import List, Match, Optional

ReinDeer = namedtuple("ReinDeer", ["speed", "sprint", "rest"])


def parse(filename: str) -> List[ReinDeer]:
    with open(filename, "r") as fp:
        data: List[str] = fp.read().splitlines()

    re_line: str = (
        r"(\w+) can fly (\d+) km/s for (\d+) seconds, "
        r"but then must rest for (\d+) seconds."
    )

    reindeers: List[ReinDeer] = []

    for line in data:
        matches: Optional[Match[str]] = re.match(re_line, line)
        assert matches is not None

        reindeers.append(
            ReinDeer(
                speed=int(matches.group(2)),
                sprint=int(matches.group(3)),
                rest=int(matches.group(4)),
            )
        )

    return reindeers


def solve(reindeers: List[ReinDeer], timeframe: int) -> int:
    max_distance: int = 0

    for r in reindeers:
        cycles: int = timeframe // (r.sprint + r.rest)
        total_time_remaining: int = timeframe % (r.sprint + r.rest)

        distance: int = cycles * r.speed * r.sprint
        seconds_to_sprint_available: int = min(r.sprint, total_time_remaining)
        distance += r.speed * seconds_to_sprint_available

        max_distance = max(max_distance, distance)

    return max_distance


def solution(filename: str, timeframe: int) -> int:
    reindeers: List[ReinDeer] = parse(filename)
    return solve(reindeers, timeframe)


if __name__ == "__main__":
    print(solution("./example.txt", 1000))  # 1120
    print(solution("./input.txt", 2503))  # 2655

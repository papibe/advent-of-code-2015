import re
from typing import List, Match, Optional

REST: int = 0
SPRINT: int = 1


class ReinDeer:
    def __init__(self, speed: int, sprint: int, rest: int) -> None:
        self.speed: int = speed
        self.sprint: int = sprint
        self.rest: int = rest

        self.points: int = 0
        self.position: int = 0
        self.state: int = REST
        self.counter: int = 0

    def move(self) -> int:
        if self.counter == 0:
            if self.state == REST:
                self.state = SPRINT
                self.counter = self.sprint
            else:
                self.state = REST
                self.counter = self.rest

        self.counter -= 1

        if self.state == SPRINT:
            self.position += self.speed

        return self.position

    def add_point(self) -> None:
        self.points += 1


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
    for _second in range(1, timeframe + 1):

        max_distance: int = 0
        max_reindeers: List[ReinDeer] = []

        for reindeer in reindeers:
            distance: int = reindeer.move()
            if distance > max_distance:
                max_distance = distance
                max_reindeers = [reindeer]
            elif distance == max_distance:
                max_reindeers.append(reindeer)

        for reindeer in max_reindeers:
            reindeer.add_point()

    return max(r.points for r in reindeers)


def solution(filename: str, timeframe: int) -> int:
    reindeers: List[ReinDeer] = parse(filename)
    return solve(reindeers, timeframe)


if __name__ == "__main__":
    print(solution("./example.txt", 1000))  # 689
    print(solution("./input.txt", 2503))  # 1059

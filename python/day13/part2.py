import re
from typing import Dict, List, Match, Optional

AdjacencyMatrix = List[List[int]]


def parse(filename: str) -> AdjacencyMatrix:
    with open(filename, "r") as fp:
        data: List[str] = fp.read().splitlines()

    re_line: str = r"(\w+) would (\w+) (\d+) happiness units by sitting next to (\w+)\."

    # build basic rules
    happiness: Dict[str, Dict[str, int]] = {}

    for line in data:
        matches: Optional[Match[str]] = re.match(re_line, line)
        assert matches is not None

        person: str = matches.group(1)

        str_action: str = matches.group(2)
        assert str_action in ["gain", "lose"]

        action: int = 1 if matches.group(2) == "gain" else -1
        units: int = action * int(matches.group(3))
        partner: str = matches.group(4)

        if person not in happiness:
            happiness[person] = {}

        happiness[person][partner] = units

    # patch for part 2: add yourself
    current_happiness_participants: List[str] = list(happiness.keys())
    happiness["yourself"] = {}
    for person in current_happiness_participants:
        happiness["yourself"][person] = 0

    # bijection
    names: List[str] = []
    indexes: Dict[str, int] = {}
    for index, person in enumerate(happiness):
        names.append(person)
        indexes[person] = index

    # create adjacency matrix
    am: AdjacencyMatrix = [[0] * len(happiness) for _ in range(len(happiness))]

    for person, rules in happiness.items():
        for partner, units in rules.items():
            am[indexes[person]][indexes[partner]] = units

    return am


def solve(adjacency_matrix: AdjacencyMatrix) -> int:
    max_happiness: int = float("-inf")  # type: ignore
    visited: List[bool] = [False] * len(adjacency_matrix)

    def dfs(person: int, happiness: int) -> None:
        nonlocal max_happiness

        if all(visited):
            happiness += adjacency_matrix[person][0] + adjacency_matrix[0][person]
            max_happiness = max(max_happiness, happiness)
            return

        for partner, is_visited in enumerate(visited):
            if not is_visited:
                visited[partner] = True
                new_happiness = (
                    happiness
                    + adjacency_matrix[person][partner]
                    + adjacency_matrix[partner][person]
                )
                dfs(partner, new_happiness)
                visited[partner] = False

    visited[0] = True
    dfs(0, 0)

    return max_happiness


def solution(filename: str) -> int:
    adjacency_matrix: AdjacencyMatrix = parse(filename)
    return solve(adjacency_matrix)


if __name__ == "__main__":
    print(solution("./input.txt"))  # 664

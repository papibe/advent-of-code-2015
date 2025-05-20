from typing import List


def parse(filename: str) -> List[int]:
    with open(filename, "r") as fp:
        data: List[str] = fp.read().splitlines()

    containers: List[int] = []
    for line in data:
        containers.append(int(line))

    return containers


def solve(containers: List[int], capacity: int) -> int:
    combinations = 1
    min_combinations: int = len(containers) + 1

    def dfs(
        start_index: int, remaining_capacity: int, selected_containers: int
    ) -> None:
        nonlocal combinations
        nonlocal min_combinations

        if remaining_capacity == 0:
            if selected_containers < min_combinations:
                min_combinations = selected_containers
                combinations = 1

            elif selected_containers == min_combinations:
                combinations += 1

            return

        for container in range(start_index, len(containers)):
            container_capacity = containers[container]
            if remaining_capacity - container_capacity >= 0:
                dfs(
                    container + 1,
                    remaining_capacity - container_capacity,
                    selected_containers + 1,
                )

    dfs(0, capacity, 0)

    return combinations


def solution(filename: str, capacity: int) -> int:
    data: List[int] = parse(filename)
    return solve(data, capacity)


if __name__ == "__main__":
    print(solution("./example.txt", 25))  # 3
    print(solution("./input.txt", 150))  # 57

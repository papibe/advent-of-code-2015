import re
from typing import Dict, List, Match, Optional

AdjacencyMatrix = List[List[int]]


def parse(filename: str) -> AdjacencyMatrix:
    with open(filename, "r") as fp:
        data: List[str] = fp.read().splitlines()

    re_line: str = r"(\w+) to (\w+) = (\d+)"

    # first pass: read all data into a
    nodes: Dict[str, Dict[str, int]] = {}
    for line in data:
        matches: Optional[Match[str]] = re.match(re_line, line)
        assert matches is not None
        node1: str = matches.group(1)
        node2: str = matches.group(2)
        distance: int = int(matches.group(3))
        if node1 not in nodes:
            nodes[node1] = {node2: distance}
        else:
            nodes[node1][node2] = distance

        if node2 not in nodes:
            nodes[node2] = {node1: distance}
        else:
            nodes[node2][node1] = distance

    # bijection
    node_list: List[str] = [""] * len(nodes)
    node_map: Dict[str, int] = {}
    for index, node_name in enumerate(nodes):
        node_list[index] = node_name
        node_map[node_name] = index

    # create adjacency matrix
    am: AdjacencyMatrix = [[float("inf")] * len(nodes) for _ in range(len(nodes))]  # type: ignore

    # set distance to itself
    for index in range(len(am)):
        am[index][index] = 0

    for node1, distances in nodes.items():
        for node2, distance in distances.items():
            index1: int = node_map[node1]
            index2: int = node_map[node2]

            am[index1][index2] = distance
            am[index2][index1] = distance

    return am


def solve(adjacency_matrix: AdjacencyMatrix) -> int:
    min_distance: int = float("inf")  # type: ignore
    visited: List[bool] = [False] * len(adjacency_matrix)

    def dfs(node: int, steps: int) -> None:
        nonlocal min_distance

        if all(visited):
            min_distance = min(min_distance, steps)
            return

        for next_node, is_visited in enumerate(visited):
            if not is_visited:
                visited[next_node] = True
                dfs(next_node, steps + adjacency_matrix[node][next_node])
                visited[next_node] = False

    # start at every node
    for initial_node, is_visited in enumerate(visited):
        if not is_visited:
            visited[initial_node] = True
            # travel starting at initial_node
            dfs(initial_node, 0)
            visited[initial_node] = False

    return min_distance


def solution(filename: str) -> int:
    adjacency_matrix: List[List[int]] = parse(filename)
    return solve(adjacency_matrix)


if __name__ == "__main__":
    print(solution("./example.txt"))  # 605
    print(solution("./input.txt"))  # 141

from typing import List

ON: str = "#"
OFF: str = "."


def parse(filename: str) -> List[List[str]]:
    with open(filename, "r") as fp:
        data: List[str] = fp.read().splitlines()

    grid: List[List[str]] = []

    for row, line in enumerate(data):
        grid.append([char for char in line])

    # patch grid for part 2
    rows: int = len(grid)
    cols: int = len(grid[0])

    grid[0][0] = ON
    grid[0][cols - 1] = ON
    grid[rows - 1][0] = ON
    grid[rows - 1][cols - 1] = ON

    return grid


def number_of_on_neighbors(row: int, col: int, grid: List[List[str]]) -> int:
    rows: int = len(grid)
    cols: int = len(grid[0])
    on_neighbors: int = 0

    for nrow, ncol in [
        (row - 1, col - 1),
        (row - 1, col),
        (row - 1, col + 1),
        (row, col - 1),
        (row, col + 1),
        (row + 1, col - 1),
        (row + 1, col),
        (row + 1, col + 1),
    ]:
        if 0 <= nrow < rows and 0 <= ncol < cols:
            if grid[nrow][ncol] == ON:
                on_neighbors += 1

    return on_neighbors


def solve(grid: List[List[str]], cycles: int) -> int:
    rows: int = len(grid)
    cols: int = len(grid[0])

    for _ in range(cycles):
        next_grid: List[List[str]] = [["+"] * cols for _ in range(rows)]

        for row, line in enumerate(grid):
            for col, light in enumerate(line):
                noon: int = number_of_on_neighbors(row, col, grid)

                assert light != "+"

                if light == ON:
                    if noon == 2 or noon == 3:
                        next_grid[row][col] = ON
                    else:
                        next_grid[row][col] = OFF

                # light is off
                else:
                    if noon == 3:
                        next_grid[row][col] = ON
                    else:
                        next_grid[row][col] = OFF

        # patch bad always on corner lights
        next_grid[0][0] = ON
        next_grid[0][cols - 1] = ON
        next_grid[rows - 1][0] = ON
        next_grid[rows - 1][cols - 1] = ON

        grid = next_grid

    lights_on: int = 0
    for line in grid:
        for char in line:
            if char == ON:
                lights_on += 1

    return lights_on


def solution(filename: str, cycles: int) -> int:
    grid: List[List[str]] = parse(filename)
    return solve(grid, cycles)


if __name__ == "__main__":
    print(solution("./example2.txt", 5))  # 17
    print(solution("./input.txt", 100))  # 924

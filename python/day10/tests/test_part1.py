from typing import List

import pytest

from part1_n_2 import cycle, solution


@pytest.mark.parametrize(
    "puzzle_input,times,expected",
    [
        ("1", 1, "11"),
        ("11", 1, "21"),
        ("21", 1, "1211"),
        ("1211", 1, "111221"),
        ("111221", 1, "312211"),
        ("1", 5, "312211"),
        ("1", 4, "111221"),
        ("1", 3, "1211"),
        ("1", 2, "21"),
    ],
    ids=[
        "1_1_times_should_be_11",
        "11_1_times_should_be_21",
        "21_1_times_should_be_1211",
        "1211_1_times_should_be_111221",
        "111221_1_times_should_be_312211",
        "1_5_times_should_be_312211",
        "1_4_times_should_be_111221",
        "1_3_times_should_be_1211",
        "1_2_times_should_be_21",
    ],
)
def test_cycle(puzzle_input: str, times: int, expected: str) -> None:
    list_result: List[str] = cycle(list(puzzle_input), times)
    result: str = "".join(list_result)
    assert result == expected, f"got {result}, needs {expected}"


@pytest.mark.parametrize(
    "puzzle_input,times,expected",
    [
        ("1", 1, 2),
        ("11", 1, 2),
        ("21", 1, 4),
        ("1211", 1, 6),
        ("111221", 1, 6),
        ("1", 5, 6),
        ("1", 4, 6),
        ("1", 3, 4),
        ("1", 2, 2),
    ],
    ids=[
        "1_1_times_should_be_2",
        "11_1_times_should_be_2",
        "21_1_times_should_be_4",
        "1211_1_times_should_be_6",
        "111221_1_times_should_be_6",
        "1_5_times_should_be_6",
        "1_4_times_should_be_6",
        "1_3_times_should_be_4",
        "1_2_times_should_be_2",
    ],
)
def test_solution(puzzle_input: str, times: int, expected: int) -> None:
    result: int = solution(puzzle_input, times)
    assert result == expected, f"got {result}, needs {expected}"

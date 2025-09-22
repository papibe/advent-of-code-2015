from typing import Dict

import pytest

from part2 import solve


@pytest.mark.parametrize(
    "json_data,expected",
    [
        ([1, 2, 3], 6),
        ({"a": 2, "b": 4}, 6),
        ([[[3]]], 3),
        ({"a": {"b": 4}, "c": -1}, 3),
        ({"a": [-1, 1]}, 0),
        ([-1, {"a": 1}], 0),
        ([], 0),
        ({}, 0),
        ([1, {"c": "red", "b": 2}, 3], 4),
        ({"d": "red", "e": [1, 2, 3, 4], "f": 5}, 0),
        ([1, "red", 5], 6),
    ],
    ids=[
        "[1, 2, 3]_should_be_6",
        '{"a": 2, "b": 4}_should_be_6',
        "[[[3]]]_should_be_3",
        '{"a": {"b": 4}, "c": -1}_should_be_3',
        '{"a": [-1, 1]}_should_be_0',
        '[-1, {"a": 1}]_should_be_0',
        "[]_should_be_0",
        "{}_should_be_0",
        '[1, {"c": "red", "b": 2}, 3]_should_be_4',
        '{"d": "red", "e": [1, 2, 3, 4], "f": 5}_should_be_0',
        '[1, "red", 5]_should_be_6',
    ],
)
def test_part2(json_data: Dict[str, int], expected: int) -> None:
    result: int = solve(json_data)
    assert result == expected, f"got {result}, needs {expected}"

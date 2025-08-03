from typing import List

import pytest

from part1 import Dimension, solve


@pytest.mark.parametrize(
    "dimension,expected",
    [
        ([Dimension(2, 3, 4)], 58),
        ([Dimension(1, 1, 10)], 43),
    ],
    ids=[
        "2x3x4_should_be_58",
        "1x1x10_should_be_43",
    ],
)
def test_part1(dimension: List[Dimension], expected: int) -> None:
    result: int = solve(dimension)
    assert result == expected, f"got {result}, needs {expected}"

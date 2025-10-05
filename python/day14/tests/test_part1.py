from typing import List

import pytest

from part1 import ReinDeer, solve

comet: ReinDeer = ReinDeer(14, 10, 127)
dancer: ReinDeer = ReinDeer(16, 11, 162)


@pytest.mark.parametrize(
    "reindeers,timeframe,expected",
    [
        ([comet], 1, 14),
        ([dancer], 1, 16),
        ([comet], 11, 140),
        ([dancer], 11, 176),
        ([comet], 12, 140),
        ([dancer], 12, 176),
    ],
    ids=[
        "comet at 1 sec reaches 14 km",
        "dance at 1 sec reaches 16 km",
        "comet at 11 sec reaches 140 km",
        "dance at 11 sec reaches 176 km",
        "comet at 12 sec reaches 140 km",
        "dance at 12 sec reaches 176 km",
    ],
)
def test_part1(reindeers: List[ReinDeer], timeframe: int, expected: int) -> None:
    result: int = solve(reindeers, timeframe)
    assert result == expected, f"got {result}, needs {expected}"

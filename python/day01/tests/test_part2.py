import pytest

from part2 import solve


@pytest.mark.parametrize(
    "directions,expected",
    [
        (")", 1),
        ("()())", 5),
    ],
    ids=[
        ")_should_be_1",
        "()())_should_be_5",
    ],
)
def test_part2(directions: str, expected: int) -> None:
    result: int = solve(directions)
    assert result == expected, f"got {result}, needs {expected}"

import pytest

from part1 import solve


@pytest.mark.parametrize(
    "directions,expected",
    [
        (">", 2),
        ("^>v<", 4),
        ("^v^v^v^v^v", 2),
    ],
    ids=[
        ">_2",
        "^>v<_4",
        "^v^v^v^v^v_2",
    ],
)
def test_part1(directions: str, expected: int) -> None:
    result: int = solve(directions)
    assert result == expected, f"got {result}, needs {expected}"

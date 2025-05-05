import pytest

from part2 import solve


@pytest.mark.parametrize(
    "directions,expected",
    [
        ("^v", 3),
        ("^>v<", 3),
        ("^v^v^v^v^v", 11),
    ],
    ids=[
        ">_2",
        "^>v<_4",
        "^v^v^v^v^v_2",
    ],
)
def test_part2(directions: str, expected: int) -> None:
    result: int = solve(directions)
    assert result == expected, f"got {result}, needs {expected}"

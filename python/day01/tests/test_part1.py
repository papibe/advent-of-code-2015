import pytest

from part1 import solve


@pytest.mark.parametrize(
    "directions,expected",
    [
        ("(())", 0),
        ("()()", 0),
        ("(((", 3),
        ("(()(()(", 3),
        ("))(((((", 3),
        ("())", -1),
        ("))(", -1),
        (")))", -3),
        (")())())", -3),
    ],
    ids=[
        "(())_should_be_0",
        "()()_should_be_0",
        "(((_should_be_3",
        "(()(()(_should_be_3",
        "))(((((_should_be_3",
        "())_should_be_-1",
        "))(_should_be_-1",
        ")))_should_be_-3",
        ")())())_should_be_-3",
    ],
)
def test_part1(directions: str, expected: int) -> None:
    result: int = solve(directions)
    assert result == expected, f"got {result}, needs {expected}"

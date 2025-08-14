import pytest

from part1_n_2 import solve


@pytest.mark.parametrize(
    "input_,zeros,expected",
    [
        ("abcdef", 5, 609043),
        ("pqrstuv", 5, 1048970),
    ],
    ids=[
        "abcdef_should_be_609043",
        "pqrstuv_should_be_1048970",
    ],
)
def test_part1_examples(input_: str, zeros: int, expected: int) -> None:
    result: int = solve(input_, zeros)
    assert result == expected, f"got {result}, needs {expected}"

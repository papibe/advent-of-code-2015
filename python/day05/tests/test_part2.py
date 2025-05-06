import pytest

from part2 import is_nice


@pytest.mark.parametrize(
    "string,expected",
    [
        ("qjhvhtzxzqqjkmpb", True),
        ("xxyxx", True),
        ("uurcxstgmygtbstg", False),
        ("ieodomkazucvgmuy", False),
    ],
    # ids=[
    #     "aa_bb_cc_dd_ee_should_be_True",
    #     "aa_bb_cc_dd_aa_should_be_False",
    #     "aa_bb_cc_dd_aaa_should_be_True",
    # ],
)
def test_part1(string: str, expected: int) -> None:
    result: bool = is_nice(string)
    assert result == expected, f"got {result}, needs {expected}"

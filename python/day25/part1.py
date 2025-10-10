def get_order(row: int, col: int) -> int:
    return ((col + row) ** 2 - col - 3 * row + 2) // 2


def solution(row: int, col: int) -> int:
    position: int = get_order(row, col)

    code: int = 20151125
    for _ in range(position - 1):
        code = (code * 252533) % 33554393

    return code


if __name__ == "__main__":
    print(solution(2978, 3083))  # 2650453

from typing import List


def parse(filename: str) -> List[int]:
    with open(filename, "r") as fp:
        data: List[str] = fp.read().splitlines()

    packets: List[int] = []
    for str_number in data:
        packets.append(int(str_number))

    return packets


def get_entanglement(packets: List[int]) -> int:
    entanglement: int = 1
    for p in packets:
        entanglement *= p
    return entanglement


def solve(packets: List[int], group_size: int) -> int:
    target_weight: int = sum(packets) // group_size

    min_size: int = float("inf")  # type: ignore
    size_winners: List[int] = []

    def n_sum_k(cs: int, index: int, selection: List[int]) -> None:
        nonlocal min_size, size_winners

        if index < 0:
            return

        if cs > target_weight:
            return

        if len(selection) > min_size:
            return

        if cs == target_weight:
            if len(selection) < min_size:
                min_size = len(selection)
                size_winners = [get_entanglement(selection)]
            elif len(selection) == min_size:
                size_winners.append(get_entanglement(selection))

        for next_index in range(index - 1, -1, -1):
            packet: int = packets[next_index]
            if cs + packet <= target_weight:
                new_selection = selection.copy()
                new_selection.append(packet)
                n_sum_k(cs + packet, next_index, new_selection)

    n_sum_k(0, len(packets), [])

    return min(size_winners)


def solution(filename: str, group_size: int) -> int:
    packets: List[int] = parse(filename)
    return solve(packets, group_size)


if __name__ == "__main__":
    print("Part 1:")
    print("  Example:", solution("./example.txt", 3))  # 99
    print("  Input  :", solution("./input.txt", 3))  # 11266889531

    print("Part 2:")
    print("  Example:", solution("./example.txt", 4))  # 44
    print("  Input  :", solution("./input.txt", 4))  # 77387711

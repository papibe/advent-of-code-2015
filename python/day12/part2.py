from json import load
from typing import Any, Dict, List

Json = Dict[str, Any] | List[Any]


def parse(filename: str) -> Json:
    with open(filename, "r") as fp:
        data: Json = load(fp)

    return data


def solve(json_data: Json) -> int:

    def dfs(data: Json) -> int:
        if isinstance(data, int):
            return data

        total_sum: int = 0
        if isinstance(data, list):
            for item in data:
                total_sum += dfs(item)

        if isinstance(data, dict):
            local_dict_sum: int = 0
            for _key, value in data.items():
                if value == "red":
                    break
                local_dict_sum += dfs(value)
            else:
                total_sum += local_dict_sum

        return total_sum

    return dfs(json_data)


def solution(filename: str) -> int:
    json_data: Json = parse(filename)
    return solve(json_data)


if __name__ == "__main__":
    print(solution("./input.txt"))  # 68466

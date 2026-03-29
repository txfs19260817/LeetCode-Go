from __future__ import annotations


def are_connected(
    arr: list[int],
    diff: int,
    queries: list[list[int]],
) -> list[bool]:
    n = len(arr)
    if n == 0:
        return [False] * len(queries)

    value_index_pairs = sorted((value, index) for index, value in enumerate(arr))
    component_of = [0] * n

    component_id = 0
    first_value, first_index = value_index_pairs[0]
    component_of[first_index] = component_id

    for i in range(1, n):
        value, index = value_index_pairs[i]
        previous_value, _ = value_index_pairs[i - 1]

        # A gap larger than diff breaks connectivity, so we start
        # a new component in sorted-value order.
        if value - previous_value > diff:
            component_id += 1

        component_of[index] = component_id

    return [component_of[left] == component_of[right] for left, right in queries]


if __name__ == "__main__":
    assert are_connected([1, 2, 3, 6], 2, [[0, 2], [1, 3]]) == [True, False]
    assert are_connected([5, 1, 3, 2], 1, [[1, 3], [1, 0], [2, 3]]) == [True, False, True]
    assert are_connected([10], 0, [[0, 0]]) == [True]
    assert are_connected([1, 10, 20], 3, [[0, 1], [1, 2]]) == [False, False]

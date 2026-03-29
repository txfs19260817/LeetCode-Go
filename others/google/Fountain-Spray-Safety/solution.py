from __future__ import annotations


def _previous_ge_indices(terrain: list[int]) -> list[int]:
    prev_ge = [-1] * len(terrain)
    stack: list[int] = []

    for i, height in enumerate(terrain):
        while stack and terrain[stack[-1]] < height:
            stack.pop()
        if stack:
            prev_ge[i] = stack[-1]
        stack.append(i)

    return prev_ge


def _next_ge_indices(terrain: list[int]) -> list[int]:
    n = len(terrain)
    next_ge = [n] * n
    stack: list[int] = []

    for i in range(n - 1, -1, -1):
        while stack and terrain[stack[-1]] < terrain[i]:
            stack.pop()
        if stack:
            next_ge[i] = stack[-1]
        stack.append(i)

    return next_ge


def find_unsafe_positions(terrain: list[int], fountains: list[int]) -> list[int]:
    n = len(terrain)
    prev_ge = _previous_ge_indices(terrain)
    next_ge = _next_ge_indices(terrain)
    diff = [0] * (n + 1)

    for fountain in fountains:
        if not 0 <= fountain < n:
            raise ValueError(f"fountain index out of bounds: {fountain}")

        # The nearest blocking wall on the left is prev_ge[fountain].
        # Every cell after that wall up to fountain - 1 is strictly lower,
        # so the fountain sprays the closed interval [left_start, left_end].
        left_start = prev_ge[fountain] + 1
        left_end = fountain - 1
        if left_start <= left_end:
            # Difference-array range add for the sprayed left interval.
            diff[left_start] += 1
            diff[left_end + 1] -= 1

        # Symmetrically, the fountain sprays from fountain + 1 up to the
        # cell just before the first blocking wall on the right.
        right_start = fountain + 1
        right_end = next_ge[fountain] - 1
        if right_start <= right_end:
            # Difference-array range add for the sprayed right interval.
            diff[right_start] += 1
            diff[right_end + 1] -= 1

    unsafe = [0] * n
    active = 0
    for i in range(n):
        active += diff[i]
        unsafe[i] = 1 if active > 0 else 0

    return unsafe


if __name__ == "__main__":
    assert find_unsafe_positions([2, 1, 3, 2, 1, 1], [0, 3]) == [0, 1, 0, 0, 1, 1]
    assert find_unsafe_positions([5, 4, 3, 2, 1], [0]) == [0, 1, 1, 1, 1]
    assert find_unsafe_positions([1, 2, 3, 2, 1], [2]) == [1, 1, 0, 1, 1]
    assert find_unsafe_positions([2, 2, 2], [1]) == [0, 0, 0]

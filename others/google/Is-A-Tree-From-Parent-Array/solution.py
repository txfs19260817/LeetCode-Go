from __future__ import annotations


def is_tree(parents: list[int]) -> bool:
    n = len(parents)
    root_count = 0

    for parent in parents:
        if parent == -1:
            root_count += 1
        elif not 0 <= parent < n:
            return False

    if root_count != 1:
        return False

    # state[i]:
    # 0 = unvisited
    # 1 = currently exploring this chain
    # 2 = already verified to lead to the unique root
    state = [0] * n

    def dfs(node: int) -> bool:
        if parents[node] == -1:
            state[node] = 2
            return True
        if state[node] == 1:
            return False
        if state[node] == 2:
            return True

        state[node] = 1
        if not dfs(parents[node]):
            return False
        state[node] = 2
        return True

    for node in range(n):
        if state[node] == 0 and not dfs(node):
            return False

    return True


if __name__ == "__main__":
    assert is_tree([1, -1, 1, 2, 5, 2]) is True
    assert is_tree([-1, 0, -1, 2, 3]) is False
    assert is_tree([1, -1, 3, 4, 5, 2]) is False
    assert is_tree([-1]) is True
    assert is_tree([0]) is False

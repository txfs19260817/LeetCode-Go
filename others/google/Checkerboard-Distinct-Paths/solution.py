from __future__ import annotations


def count_paths(rows: int, cols: int) -> int:
    if rows <= 0 or cols <= 0:
        return 0

    # After processing a column, dp[row] is the number of ways
    # to reach cell (row, current_column) from the bottom-left start cell.
    dp = [0] * rows
    dp[rows - 1] = 1

    for _ in range(1, cols):
        next_dp = [0] * rows
        for row in range(rows):
            # We can arrive at (row, col) from the previous column using
            # the same row, the row above, or the row below.
            next_dp[row] += dp[row]
            if row - 1 >= 0:
                next_dp[row] += dp[row - 1]
            if row + 1 < rows:
                next_dp[row] += dp[row + 1]
        dp = next_dp

    return dp[rows - 1]


if __name__ == "__main__":
    assert count_paths(1, 1) == 1
    assert count_paths(2, 2) == 1
    assert count_paths(2, 3) == 2
    assert count_paths(3, 3) == 2
    assert count_paths(3, 4) == 4

from __future__ import annotations

from typing import List


class Solution:
    def deleteCoveredPoint(
        self, intervals: List[List[int]], idx: int
    ) -> List[List[int]]:
        remaining = idx
        for i, (start, end) in enumerate(intervals):
            size = end - start
            if remaining < size:
                point = start + remaining
                if size == 1:
                    # Remove interval entirely
                    return intervals[:i] + intervals[i + 1 :]
                if point == start:
                    # Shrink left
                    return intervals[:i] + [[start + 1, end]] + intervals[i + 1 :]
                if point == end - 1:
                    # Shrink right
                    return intervals[:i] + [[start, end - 1]] + intervals[i + 1 :]
                # Split into two intervals
                return (
                    intervals[:i]
                    + [[start, point], [point + 1, end]]
                    + intervals[i + 1 :]
                )
            remaining -= size
        return intervals  # idx out of range


if __name__ == "__main__":
    sol = Solution()

    # Example 1: split in middle
    assert sol.deleteCoveredPoint([[10, 12], [13, 16], [4, 8]], 3) == [
        [10, 12],
        [13, 14],
        [15, 16],
        [4, 8],
    ]

    # Example 2: shrink left
    assert sol.deleteCoveredPoint([[4, 8], [13, 16], [10, 12]], 0) == [
        [5, 8],
        [13, 16],
        [10, 12],
    ]

    # Example 3: shrink right
    assert sol.deleteCoveredPoint([[2, 6], [8, 10], [15, 18]], 3) == [
        [2, 5],
        [8, 10],
        [15, 18],
    ]

    # Edge: single-element interval removed entirely
    assert sol.deleteCoveredPoint([[5, 6]], 0) == []

    # Split in middle of large interval
    assert sol.deleteCoveredPoint([[1, 10]], 4) == [[1, 5], [6, 10]]

    print("All tests passed!")

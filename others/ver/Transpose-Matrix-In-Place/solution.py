from __future__ import annotations

from typing import List


def transpose_in_place(matrix: List[List[int]]) -> None:
    """Transpose square matrix in place."""
    if not matrix:
        raise ValueError("matrix must be non-empty")

    n = len(matrix)
    if any(len(row) != n for row in matrix):
        raise ValueError("matrix must be square")

    # Swap only upper triangle with lower triangle; diagonal stays unchanged.
    for i in range(n):
        for j in range(i + 1, n):
            matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]


def transpose_worker(matrix: List[List[int]], num_workers: int, worker_index: int) -> None:
    """Run one worker's share of in-place transpose.

    Partition rule: row i belongs to worker (i % num_workers).
    """
    if not matrix:
        raise ValueError("matrix must be non-empty")
    if num_workers <= 0:
        raise ValueError("num_workers must be positive")
    if not (0 <= worker_index < num_workers):
        raise ValueError("worker_index out of range")

    n = len(matrix)
    if any(len(row) != n for row in matrix):
        raise ValueError("matrix must be square")

    # Row ownership avoids double-swapping between workers.
    for i in range(worker_index, n, num_workers):
        for j in range(i + 1, n):
            matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]


def transpose_with_workers(matrix: List[List[int]], num_workers: int) -> None:
    """Simulate all workers; useful for local validation/interview demo."""
    for w in range(num_workers):
        transpose_worker(matrix, num_workers, w)


if __name__ == "__main__":
    m1 = [
        [1, 2, 3],
        [4, 5, 6],
        [7, 8, 9],
    ]
    transpose_in_place(m1)
    assert m1 == [
        [1, 4, 7],
        [2, 5, 8],
        [3, 6, 9],
    ]

    m2 = [
        [1, 2, 3, 4],
        [5, 6, 7, 8],
        [9, 10, 11, 12],
        [13, 14, 15, 16],
    ]
    transpose_with_workers(m2, 3)
    assert m2 == [
        [1, 5, 9, 13],
        [2, 6, 10, 14],
        [3, 7, 11, 15],
        [4, 8, 12, 16],
    ]

    try:
        transpose_worker([[1, 2]], 2, 0)
        raise AssertionError("Expected ValueError for non-square matrix")
    except ValueError:
        pass

    print("All assertions passed.")

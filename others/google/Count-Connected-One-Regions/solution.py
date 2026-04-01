from __future__ import annotations

from collections import deque


EIGHT_DIRECTIONS = [
    (-1, -1), (-1, 0), (-1, 1),
    (0, -1),           (0, 1),
    (1, -1),  (1, 0),  (1, 1),
]

FOUR_DIRECTIONS = [
    (-1, 0),
    (1, 0),
    (0, -1),
    (0, 1),
]


def count_connected_one_regions(grid: list[list[int]], eight_connected: bool = True) -> int:
    if not grid or not grid[0]:
        return 0

    rows, cols = len(grid), len(grid[0])
    visited = [[False] * cols for _ in range(rows)]
    directions = EIGHT_DIRECTIONS if eight_connected else FOUR_DIRECTIONS
    regions = 0

    def bfs(start_row: int, start_col: int) -> None:
        queue = deque([(start_row, start_col)])
        visited[start_row][start_col] = True

        while queue:
            row, col = queue.popleft()

            for dr, dc in directions:
                nr, nc = row + dr, col + dc
                if not (0 <= nr < rows and 0 <= nc < cols):
                    continue
                if visited[nr][nc] or grid[nr][nc] != 1:
                    continue
                visited[nr][nc] = True
                queue.append((nr, nc))

    for row in range(rows):
        for col in range(cols):
            if grid[row][col] == 1 and not visited[row][col]:
                regions += 1
                bfs(row, col)

    return regions


if __name__ == "__main__":
    grid = [
        [0, 1, 1, 0],
        [0, 1, 0, 0],
        [0, 0, 1, 1],
        [0, 0, 0, 1],
    ]
    assert count_connected_one_regions(grid, eight_connected=True) == 1
    assert count_connected_one_regions(grid, eight_connected=False) == 2

    mouth_shape = [
        [0, 0, 0, 0, 0],
        [0, 1, 1, 1, 0],
        [0, 1, 0, 1, 0],
        [0, 1, 1, 1, 0],
        [0, 0, 0, 0, 0],
    ]
    assert count_connected_one_regions(mouth_shape) == 1

    sun_shape = [
        [0, 0, 0, 0, 0],
        [0, 1, 1, 1, 0],
        [0, 1, 1, 1, 0],
        [0, 0, 0, 0, 0],
        [0, 1, 1, 1, 0],
        [0, 1, 1, 1, 0],
        [0, 0, 0, 0, 0],
    ]
    assert count_connected_one_regions(sun_shape) == 2

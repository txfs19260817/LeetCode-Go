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


def count_surrounded_regions(grid: list[list[int]], eight_connected: bool = True) -> int:
    if not grid or not grid[0]:
        return 0

    rows, cols = len(grid), len(grid[0])
    visited = [[False] * cols for _ in range(rows)]
    directions = EIGHT_DIRECTIONS if eight_connected else FOUR_DIRECTIONS
    regions = 0

    def bfs(start_row: int, start_col: int) -> bool:
        queue = deque([(start_row, start_col)])
        visited[start_row][start_col] = True
        touches_border = (
            start_row == 0
            or start_row == rows - 1
            or start_col == 0
            or start_col == cols - 1
        )

        while queue:
            row, col = queue.popleft()

            for dr, dc in directions:
                nr, nc = row + dr, col + dc
                if not (0 <= nr < rows and 0 <= nc < cols):
                    continue
                if visited[nr][nc] or grid[nr][nc] != 1:
                    continue
                visited[nr][nc] = True
                if nr == 0 or nr == rows - 1 or nc == 0 or nc == cols - 1:
                    touches_border = True
                queue.append((nr, nc))

        # 只有整个 1 连通块完全不接触矩阵边界，才算被 0 包围。
        return not touches_border

    for row in range(rows):
        for col in range(cols):
            if grid[row][col] == 1 and not visited[row][col]:
                if bfs(row, col):
                    regions += 1

    return regions

if __name__ == "__main__":
    grid = [
        [0, 1, 1, 0],
        [0, 1, 0, 0],
        [0, 0, 1, 1],
        [0, 0, 0, 1],
    ]
    assert count_surrounded_regions(grid, eight_connected=True) == 0
    assert count_surrounded_regions(grid, eight_connected=False) == 0

    mouth_shape = [
        [0, 0, 0, 0, 0],
        [0, 1, 1, 1, 0],
        [0, 1, 0, 1, 0],
        [0, 1, 1, 1, 0],
        [0, 0, 0, 0, 0],
    ]
    assert count_surrounded_regions(mouth_shape) == 1

    sun_shape = [
        [0, 0, 0, 0, 0],
        [0, 1, 1, 1, 0],
        [0, 1, 1, 1, 0],
        [0, 0, 0, 0, 0],
        [0, 1, 1, 1, 0],
        [0, 1, 1, 1, 0],
        [0, 0, 0, 0, 0],
    ]
    assert count_surrounded_regions(sun_shape) == 2

    diagonal_only = [
        [0, 0, 0, 0, 0],
        [0, 1, 0, 0, 0],
        [0, 0, 1, 0, 0],
        [0, 0, 0, 1, 0],
        [0, 0, 0, 0, 0],
    ]
    assert count_surrounded_regions(diagonal_only, eight_connected=True) == 1
    assert count_surrounded_regions(diagonal_only, eight_connected=False) == 3

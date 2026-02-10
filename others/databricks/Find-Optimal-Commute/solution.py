from __future__ import annotations

from collections import deque
from typing import List


class Solution:
    def findOptimalCommute(
        self,
        grid: List[List[str]],
        costs: List[int],
        times: List[int],
    ) -> str:
        """Return the name of the fastest transport mode from S to D.

        For each mode (1-4) run BFS on the grid counting only cells of
        that mode (plus the S/D endpoints).  Pick the mode with the
        smallest total time; break ties by smallest total cost.
        """
        rows, cols = len(grid), len(grid[0])

        sr = sc = dr = dc = 0
        for r in range(rows):
            for c in range(cols):
                if grid[r][c] == "S":
                    sr, sc = r, c
                elif grid[r][c] == "D":
                    dr, dc = r, c

        directions = [(0, 1), (0, -1), (1, 0), (-1, 0)]

        def bfs(mode_str: str) -> int:
            visited = [[False] * cols for _ in range(rows)]
            visited[sr][sc] = True
            queue: deque[tuple[int, int, int]] = deque([(sr, sc, 0)])
            while queue:
                r, c, steps = queue.popleft()
                if r == dr and c == dc:
                    return steps
                for delta_r, delta_c in directions:
                    nr, nc = r + delta_r, c + delta_c
                    if 0 <= nr < rows and 0 <= nc < cols and not visited[nr][nc]:
                        if grid[nr][nc] == mode_str or grid[nr][nc] == "D":
                            visited[nr][nc] = True
                            queue.append((nr, nc, steps + 1))
            return -1

        mode_names = ["Walk", "Bike", "Car", "Train"]
        best_mode = ""
        best_time = best_cost = float("inf")

        for i in range(4):
            dist = bfs(str(i + 1))
            if dist < 0:
                continue
            t = dist * times[i]
            c = dist * costs[i]
            if t < best_time or (t == best_time and c < best_cost):
                best_time = t
                best_cost = c
                best_mode = mode_names[i]

        return best_mode


if __name__ == "__main__":
    sol = Solution()

    # Sample: Bike reaches D in 5 blocks (time 10), Walk in 5 blocks (time 15).
    grid1 = [
        ["3", "3", "S", "2", "X"],
        ["3", "1", "1", "2", "X"],
        ["3", "1", "1", "2", "2"],
        ["3", "1", "1", "1", "D"],
        ["3", "3", "3", "3", "4"],
        ["4", "4", "4", "4", "4"],
    ]
    assert sol.findOptimalCommute(grid1, [0, 1, 3, 2], [3, 2, 1, 1]) == "Bike"

    # Tie on time → pick lower cost (Walk beats Bike).
    grid2 = [
        ["1", "S", "2"],
        ["1", "X", "2"],
        ["1", "D", "2"],
    ]
    assert sol.findOptimalCommute(grid2, [1, 2, 0, 0], [1, 1, 0, 0]) == "Walk"

    # No path for any mode.
    grid3 = [
        ["S", "X"],
        ["X", "D"],
    ]
    assert sol.findOptimalCommute(grid3, [1, 1, 1, 1], [1, 1, 1, 1]) == ""

    # Single mode (Walk) straight line.
    grid4 = [["S", "1", "1", "D"]]
    assert sol.findOptimalCommute(grid4, [2, 5, 5, 5], [3, 9, 9, 9]) == "Walk"

    print("All tests passed!")

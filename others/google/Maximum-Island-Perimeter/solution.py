class Solution:
    def maxPerimeter(self, grid: list[list[int]]) -> int:
        m, n = len(grid), len(grid[0])
        ans = 0

        def dfs(i, j):
            if i < 0 or j < 0 or i >= m or j >= n or grid[i][j] == 0:
                return 1
            if grid[i][j] == 2:
                return 0
            grid[i][j] = 2
            return dfs(i + 1, j) + dfs(i, j + 1) + dfs(i - 1, j) + dfs(i, j - 1)

        for i, row in enumerate(grid):
            for j, v in enumerate(row):
                if v == 1:
                    ans = max(ans, dfs(i, j))
        return ans


if __name__ == "__main__":
    solution = Solution()


    def run(grid: list[list[int]]) -> int:
        return solution.maxPerimeter([row[:] for row in grid])


    assert run([[0, 1, 0, 0], [1, 1, 1, 0], [0, 1, 0, 0], [1, 1, 0, 0]]) == 16
    assert run([[1, 0], [0, 1]]) == 4
    assert run([[1]]) == 4
    assert run([[0, 0], [0, 0]]) == 0

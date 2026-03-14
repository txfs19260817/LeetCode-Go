# Maximum Island Perimeter

- Difficulty: Medium
- Companies: Google, Snapchat
- Stages: Screening, Onsite
- Asked By: Google, Snapchat
- Source: https://www.hack2hire.com/companies/google/coding-questions/6933c5bb5f306485cc5f07bf/practice?questionId=6933c5c85f306485cc5f07c0

## Problem

(This question is a variation of the LeetCode question 200. Number of Islands. If you haven't completed that question yet, it is recommended to solve it first.)

You are given a matrix grid of size m x n where each element is either land ('1') or water ('0'). A group of connected '1's (land) forms an island. Two land cells are considered connected if they are adjacent vertically or horizontally (not diagonally).

Each land cell has up to four edges. An edge contributes to the island's perimeter if it is either adjacent to water or lies on the boundary of the matrix.

Return the maximum perimeter among all islands in the grid. If there is no island, return 0.

Constraints:

1 ≤ m ≤ 100
1 ≤ n ≤ 100
Each grid[i][j] is either '0'(water) or '1'(land)

Example 1:

Input: grid = [[0, 1, 0, 0], [1, 1, 1, 0], [0, 1, 0, 0], [1, 1, 0, 0]]
Output: 16
Explanation: The grid is shown below:

Each edge of a land cell that is adjacent to water or lies on the grid boundary adds to the island's perimeter. Summing all such edges gives a total perimeter of 16.

Example 2:

Input: grid = [[1, 0], [0, 1]]
Output: 4

Example 3:

Input: grid = [[1]]
Output: 4

## Python Template

```python
from typing import List, Optional

class Solution:
    def maxPerimeter(self, grid: List[List[int]]) -> int:
        # TODO: Implement maxPerimeter logic
```

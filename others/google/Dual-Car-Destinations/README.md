# Dual Car Destinations

- Difficulty: Medium
- Companies: Google
- Stages: OA
- Asked By: Google
- Source: https://www.hack2hire.com/companies/google/coding-questions/67eaea5f4c9bb801bfd0ef4c/practice?questionId=67eaec184c9bb801bfd0ef4d

## Problem

Given an N x M grid representing a map with two cars, "a" and "b". Each car can move up, down, left, right, or stay in place at each step. The grid contains the following symbols:

".": Road (movable).
"#": Wall (impassable).
"a": Starting position of car a.
"b": Starting position of car b.
"A": Destination for car a.
"B": Destination for car b.

Your task is to determine whether both cars can reach their respective destinations (A for a, and B for b) at the same time, without colliding or blocking each other’s path during their movements.

Constraints:

Cars cannot occupy the same cell at the same time.
Cars cannot swap positions in one move (e.g., car a moves into car b’s position while b moves into a’s).
1
≤
𝑁
,
𝑀
≤
20
1≤N,M≤20.

Example 1:

Input: N = 1, M = 4, grid = [["a","B","A","b"]]

Output: false
Explanation: Cars "a" and "b" block each other’s paths.

Example 2:

Input: N = 3, M = 3, grid = [["b",".","A"],[".","a","."],[".",".","B"]]

Output: true

Example 3:

Input: N = 3, M = 3, grid = [["a","#","B"],["#","b","#"],["A","#","."]]

Output: false

## Python Template

```python
from typing import List, Optional

class Solution:
    def canBothReachDestinations(self, N: int, M: int, grid: List[List[str]]) -> bool:
        # TODO: Implement canBothReachDestinations logic
```

# Ball Picking Game

- Difficulty: Hard
- Companies: Google
- Stages: Onsite
- Asked By: Google
- Source: https://www.hack2hire.com/companies/google/coding-questions/67775b1c8df5ec401f786daf/practice?questionId=67775b978df5ec401f786db0

## Problem

(This question is a variation of the LeetCode question 486. Predict the Winner. If you haven't completed that question yet, it is recommended to solve it first.)

Two players take turns selecting balls from a line. At the start of the game, both players have a score of 0. On each turn, a player can choose to pick either one or two balls from the remaining line. Each ball has an integer value, which can be either positive or negative. Once a player picks a ball, it cannot be selected again by either player. Both players aim to maximize their own total profit by the end of the game. The game ends when all the balls have been picked.

After all selections are made, calculate and return the maximum difference between Player One's total profit and Player Two's total profit.

Constraints:

1 ≤ nums.length ≤
10
5
10
5
−
10
4
−10
4
≤ nums[i] ≤
10
4
10
4

Example 1:

Input: [1, -1, -3, 1, 2, 4]
Output: 2
Explanation:

Player One picks 1 and -1, total = 0.
Player Two picks -3, total = -3.
Player One picks 1 and 2, total = 3.
Player Two picks 4, total = 1.
After all selections, player one total = 3, player two total = 1. The maximum difference is 2.

Example 2:

Input: [1, 2, 3, 4]
Output: 0

Example 3:

Input: [4, -1, 2, -3, 5]
Output: 5

## Python Template

```python
from typing import List, Optional

class Solution:
    def maxDifference(self, nums: List[int]) -> int:
        # TODO: Implement maxDifference logic
```

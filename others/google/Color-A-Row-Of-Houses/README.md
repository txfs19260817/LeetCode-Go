# Color a Row of Houses

- Difficulty: Hard
- Companies: Google
- Stages: Screening
- Asked By: Google
- Source: https://www.hack2hire.com/companies/google/coding-questions/6818cc0233adaac2299bc9f0/practice?questionId=6818d209aa9eb23a3dba07f0

## Problem

A builder has a project to repaint a row of houses on a street. Determine the number of ways to paint a row of n houses using k available colors (represented by integers 0 to k - 1). The painting must follow two rules:

No two adjacent houses can have the same color.
The color of the first house (firstHouseColor) and the color of the last house (lastHouseColor) are fixed and provided as input.

Calculate the total number of distinct ways to color the houses from index 1 to n - 2 (the intermediate houses) while adhering to these rules. Return 0 if no valid coloring exists.

Constraints:

1 <= n <= 10
1 <= k <= 10
k <= n

Example 1:

Input: n = 4, k = 3, firstHouseColor = 0, lastHouseColor = 2
Output: 3
Explanation: Given houses H₀, H₁, H₂, H₃ and available colors [0, 1, 2], where H₀ is fixed as color 0 and H₃ is fixed as color 2, the valid ways to paint the houses are [0, 1, 0, 2], [0, 2, 0, 2], and [0, 2, 1, 2].

Example 2:

Input: n = 5, k = 3, firstHouseColor = 1, lastHouseColor= 1
Output: 6

Example 3:

Input: n = 3, k = 2, firstHouseColor = 0, lastHouseColor= 1
Output: 0

## Python Template

```python
from typing import List, Optional

class Solution:
    def countColorings(self, n: int, k: int, firstHouseColor: int, lastHouseColor: int) -> int:
        # TODO: Implement countColorings logic
```

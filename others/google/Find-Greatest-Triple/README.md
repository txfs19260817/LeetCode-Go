# Find Greatest Triple

- Difficulty: Easy
- Companies: Google
- Stages: OA
- Asked By: Google
- Source: https://www.hack2hire.com/companies/google/coding-questions/6789743211318cc872dd83eb/practice?questionId=6789754f11318cc872dd83ec

## Problem

You are given an array digits consisting of N single-digit integers. You can select up to three digits (keeping their original order) to form a new integer.

Your goal is to create the largest possible number from these up to three chosen digits. If you pick fewer than three digits (for example, only 2 digits), that result must still be as large as possible under the given constraints.

Constraints:

3 ≤ N ≤ 50
Each element of digits is between 0 and 9 (inclusive).

Example 1:

Input: digits = [7, 2, 3, 3, 4, 9]
Output: 749
Explanation:
We can select (7, 4, 9) to form 749, which is the largest 3-digit number possible while respecting the order.

Example 2:

Input: digits = [0, 0, 5, 7]
Output: 57

Example 3:

Input: digits = [3, 1, 9, 1, 5, 9]
Output: 959

## Python Template

```python
from typing import List, Optional

class Solution:
    def findMaxNumber(self, digits: List[int]) -> int:
        # TODO: Implement findMaxNumber logic
```

# Find Optimal Moves

- Difficulty: Easy
- Companies: Google
- Stages: OA
- Asked By: Google
- Source: https://www.hack2hire.com/companies/google/coding-questions/678e7198fa7ce9616e190fbb/practice?questionId=678e72f4fa7ce9616e190fbc

## Problem

You are given an array arr, initially represented by an array of the same length filled with zeros. In each move, you can select a continuous subarray and increase every element in that subarray by 1.

Find the minimum number of moves required to transform the initial zero-filled array into the target array arr.

Constraints:

1 ≤ N ≤ 100,000
Elements of arr range from 0 to 1,000,000,000
The required number of moves is guaranteed not to exceed 1,000,000,000

Example 1:

Input: arr = [2, 1, 3]
Output: 4
Explanation: A possible sequence of moves is: [0, 0, 0] → [1, 1, 1] → [2, 1, 1] → [2, 1, 2] → [2, 1, 3].

Example 2:

Input: arr = [2, 2, 0, 0, 1]
Output: 3

Example 3:

Input: arr = [5, 4, 2, 4, 1]
Output: 7

## Python Template

```python
from typing import List, Optional

class Solution:
    def calculateMoves(self, arr: List[int]) -> int:
        # TODO: Implement calculateMoves logic
```

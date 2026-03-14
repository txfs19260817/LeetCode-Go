# Count Distinct Elements in Window

- Difficulty: Easy
- Companies: Google, Microsoft, Meta
- Stages: Screening, Onsite
- Asked By: Google, Meta, Microsoft
- Source: https://www.hack2hire.com/companies/google/coding-questions/67e08f3ea59c3d249d3625df/practice?questionId=67e08fffa59c3d249d3625e0

## Problem

Given an array of integers nums and an integer k, compute the number of distinct numbers in every contiguous subarray (window) of size k. For each window, determine how many unique elements it contains and return these counts as an array.

Constraints:

1 ≤ k ≤ nums.length
The array may include any integers (positive, negative, or zero).

Example 1:

Input: arr = [1, 2, 1, 3, 4, 2, 3], k = 4
Output: [3, 4, 4, 3]
Explanation: Below is the calculation of the result for a fixed window size of 4:

Window [1, 2, 1, 3] → Unique elements: {1, 2, 3} → Count = 3
Window [2, 1, 3, 4] → Unique elements: {1, 2, 3, 4} → Count = 4
Window [1, 3, 4, 2] → Unique elements: {1, 2, 3, 4} → Count = 4
Window [3, 4, 2, 3] → Unique elements: {2, 3, 4} → Count = 3

Example 2:

Input: nums = [1, 2, 4, 4], k = 2
Output: [2, 2, 1]

Example 3:

Input: [1, 2, 4, 4], k = 3
Output: [3, 2]

## Python Template

```python
from typing import List, Optional

class Solution:
    def countWindowDistinct(self, nums: List[int], k: int) -> List[int]:
        # TODO: Implement countWindowDistinct logic
```

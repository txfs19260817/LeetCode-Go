# Windowed Average excluding Largest K

- Difficulty: Medium
- Companies: Google
- Stages: Screening
- Asked By: Google
- Source: https://www.hack2hire.com/companies/google/coding-questions/67886cfb402dc47b67a1493f/practice?questionId=67886e3c402dc47b67a14940

## Problem

Given an integer array nums, a window size windowSize, and an integer k, return a list of averages for each sliding window of size windowSize as the window moves from left to right across the array. When calculating the average for each window, ignore the largest k numbers within that window.

Constraints:

1 <= windowSize <= nums.length <=
10
5
10
5
0 <= k < windowSize
−
10
4
−10
4
 <= nums[i] <=
10
4
10
4

Example 1:

Input: nums = [10, 20, 30, 40, 50, 60], windowSize = 3, k = 1
Output: [15.0, 25.0, 35.0, 45.0]
Explanation:

Window [10, 20, 30]: Ignoring the largest number 30, the average is (10 + 20) / 2 = 15.0
Window [20, 30, 40]: Ignoring the largest number 40, the average is (20 + 30) / 2 = 25.0
Window [30, 40, 50]: Ignoring the largest number 50, the average is (30 + 40) / 2 = 35.0
Window [40, 50, 60]: Ignoring the largest number 60, the average is (40 + 50) / 2 = 45.0

Example 2:

Input: nums = [5, 1, 3, 8, 7], windowSize = 2, k = 0
Output: [3.0, 2.0, 5.5, 7.5]

Example 3:

Input: nums = [4, 4, 4, 4], windowSize = 2, k = 1
Output: [4.0, 4.0, 4.0]

## Python Template

```python
from typing import List, Optional

class Solution:
    def SlidingWindowAverage(self, nums: List[int], windowSize: int, k: int) -> List[float]:
        # TODO: Implement SlidingWindowAverage logic
```

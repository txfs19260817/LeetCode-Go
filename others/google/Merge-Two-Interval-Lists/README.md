# Merge Two Interval Lists

- Difficulty: Medium
- Companies: Google, Microsoft, Meta
- Stages: Onsite
- Asked By: Google, Meta, Microsoft
- Source: https://www.hack2hire.com/companies/google/coding-questions/67df682a06b62d81733ab892/practice?questionId=67df6ae706b62d81733ab893

## Problem

(This question is a variation of the LeetCode question 56. Merge Intervals. If you haven't completed that question yet, it is recommended to solve it first.)

Given two lists of intervals, where each interval is represented as a two-element array [start, end]. Each list contains non-overlapping intervals that are sorted by their start times.

You are asked to merge these two lists into a single list of non-overlapping intervals. If intervals from different lists overlap, they should be combined into a single continuous interval.

Constraints:

Both lists are sorted in non-descending order by the start time.
There is no overlap within each individual list.

Example 1:

Input: list1 = [[1, 5], [10, 14], [16, 18]], list2 = [[2, 6], [8, 10], [11, 20]]
Output: [[1, 6], [8, 20]]
Explanation: The intervals [1, 5] and [2, 6] are merged into [1, 6], while the remaining intervals combine to [8, 20].

Example 2:

Input: list1 = [[0, 2], [5, 10], [13, 23], [24, 25]], list2 = [[1, 5], [8, 12], [15, 24], [25, 26]]
Output: [[0, 12], [13, 26]]

Example 3:

Input: list1 = [[0, 10]], list2 = [[20, 100]]
Output: [[0, 10], [20, 100]]

## Python Template

```python
from typing import List, Optional

class Solution:
    def mergeLists(self, list1: List[List[int]], list2: List[List[int]]) -> List[List[int]]:
        # TODO: Implement mergeLists logic
```

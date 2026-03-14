# Find Median In Large Array

- Difficulty: Medium
- Companies: Google, Meta, Airbnb
- Stages: Screening
- Asked By: Google, Meta, Airbnb
- Source: https://www.hack2hire.com/companies/google/coding-questions/68c2d9081f64f4312b0a3df1/practice?questionId=68c2d9291f64f4312b0a3df2

## Problem

(This question is a variation of the LeetCode question 295. Find Median from Data Stream. If you haven't completed that question yet, it is recommended to solve it first.)

In the world of big data, analysts often work with massive, unsorted datasets. Imagine you are given a very large and unsorted array of integers nums. Your task is to develop an efficient method to find its median.

The median is the middle value in an ordered dataset, which is defined as:

If the array contains an odd number of elements, the median is the single middle element after sorting.
If the array contains an even number of elements, the median is the average of the two middle elements after sorting.

Since the array can be extremely large, your solution must run in the time complexity of amortized
𝑂
(
𝑁
)
O(N) .

Constraints:

1 ≤ nums.length ≤
10
5
10
5
−
2
31
−2
31
 ≤ nums[i] ≤
2
31
−
1
2
31
−1

Example 1:

Input: nums = [3, 1, 2, 4, 5]
Output: 3.0
Explanation: The array has 5 elements (an odd number). After sorting, the array becomes [1, 2, 3, 4, 5]. The middle element is at the 3rd position, which is 3.

Example 2:

Input: nums = [7, 4, 1, 2]
Output: 3.0

Example 3:

Input: nums = [9, 2, 5, 3, 5, 8, 9, 7, 9, 3, 2]
Output: 5.0

## Python Template

```python
from typing import List, Optional

class Solution:
    def findMedian(self, nums: List[int]) -> float:
        # TODO: Implement findMedian logic
```

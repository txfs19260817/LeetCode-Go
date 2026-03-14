# Subset Sum Equals K+2

- Difficulty: Hard
- Companies: Google, Confluent
- Stages: Onsite, Screening
- Asked By: Google, Confluent
- Source: https://www.hack2hire.com/companies/google/coding-questions/677c87c55cbaff553b53b06d/practice?questionId=677c8aba5cbaff553b53b06e

## Problem

Subset Sum Equals K
+2
Hard
Dynamic Programming
Memoization
Interview Stages
Onsite
Screening
Frequency
Asked By
Last Reported
1 weeks ago

(This question is a variation of the LeetCode question 39. Combination Sum. If you haven't completed that question yet, it is recommended to solve it first.)

Given a list of positive integers and a target number k, write a function that returns true if there exists a subset of nums that adds up to k, and false otherwise. Note that numbers can appear more than once in the list.

Constraints:

1 <= nums.length <= 1000
1 <= nums[i] <=
10
6
10
6
1 <= k <=
10
9
10
9

Example 1:

Input: nums = [12, 1, 61, 5, 9, 2], k = 24
Output: true
Explanation: There exists a subset [12, 9, 2, 1] that sums up to 24.

Example 2:

Input: nums = [3, 34, 4, 12, 5, 2, 2], k = 9
Output: true
Explanation: There exists a subset [4, 5] that sums up to 9.

Example 3:

Input: nums = [5, 3, 9, 2, 7], k = 6
Output: false

## Python Template

```python
from typing import List, Optional

class Solution:
    def findSubsetSum(self, nums: List[int], k: int) -> bool:
        # TODO: Implement findSubsetSum logic
```

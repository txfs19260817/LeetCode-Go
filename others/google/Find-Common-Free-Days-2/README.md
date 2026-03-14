# Find Common Free Days+2

- Difficulty: Medium
- Companies: Google, Microsoft
- Stages: OA
- Asked By: Google, Microsoft
- Source: https://www.hack2hire.com/companies/google/coding-questions/680fba682590daeb41a329b4/practice?questionId=680fc0442590daeb41a329b8

## Problem

Find Common Free Days
+2
Medium
Array
Prefix Sum
Interview Stages
OA
Frequency
Asked By
Last Reported
1 months ago

Given a list of time blocks [person_id, start_day, end_day] where each block indicates a person's busy days (inclusive), generate an array of days where all people are free.

Each person_id is used only to group intervals for the same person, which should be merged accordingly. Days are 1-indexed, meaning the count starts from day 1.

Find all days when everyone is simultaneously available.

Example 1:

Input: intervals = [[1, 1, 2], [1, 4, 5]]
Output: [3]
Explanation: Person 1 is busy on days [1, 2] and [4, 5]. Therefore, free days should be [3].

Example 2:

Input: intervals = [[1, 1, 3], [2, 2, 4], [3, 3, 5]]
Output: []

Example 3:

Input: intervals = [[1, 1, 2], [2, 3, 4], [3, 5, 6]]
Output: []

## Python Template

```python
from typing import List, Optional

class Solution:
    def everyoneAvailable(self, blocks: List[List[int]]) -> List[int]:
        # TODO: Implement everyoneAvailable logic
```

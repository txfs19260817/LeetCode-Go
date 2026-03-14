# Employee Shift Timeline Table

- Difficulty: Medium
- Companies: Google, Microsoft
- Stages: Onsite
- Asked By: Google, Microsoft
- Source: https://www.hack2hire.com/companies/google/coding-questions/67fd4d7dd2adf72ba62b0ec3/practice?questionId=67fd508ed2adf72ba62b0ec4

## Problem

You are given a 2D list schedules representing employees' shift schedules. Each element is represented as a triplet containing the employee’s name, a start time, and an end time.

Your task is to merge overlapping shifts and generate a timeline of intervals. Each interval should display the start time, end time, and the list of employees working during that period. In the final output, employee names must appear in the order they first appeared in the input.

Constraints:

Each shift is given as [<name>, <start_time>, <end_time>] with start_time and end_time as integers.
Shifts may overlap, be adjacent, or separate.
An employee may have multiple shifts.
The output should exclude intervals with no active employees.

Example 1:

Input: schedules = [["Alice","1","5"],["Bob","2","6"],["Charlie","4","7"]]
Output: [["1","2","Alice"],["2","4","Alice","Bob"],["4","5","Alice","Bob","Charlie"],["5","6","Bob","Charlie"],["6","7","Charlie"]]
Explanation: The merged intervals are as follows:

The first interval [1,2] starts with only Alice working.
At time 2, Bob’s shift begins, so the interval [2,4] has both Alice and Bob.
At time 4, Charlie starts, which the interval [4,5] with all three employees.
When Alice’s shift ends at 5, the interval [5,6] consists of Bob and Charlie.
Finally, when Bob leaves at 6, the interval [6,7] shows only Charlie working.

Example 2:

Input: schedules = [["Alice","1","5"],["Bob","2","6"],["Alice","7","8"]]
Output: [["1","2","Alice"],["2","5","Alice","Bob"],["5","6","Bob"],["7","8","Alice"]]

Example 3:

Input: schedules = [["Charlie","1","5"],["Alice","2","6"],["Bob","4","7"]]
Output: [["1","2","Charlie"],["2","4","Charlie","Alice"],["4","5","Charlie","Alice","Bob"],["5","6","Alice","Bob"],["6","7","Bob"]]

## Python Template

```python
from typing import List, Optional

class Solution:
    def processShifts(self, schedules: List[List[str]]) -> List[List[str]]:
        # TODO: Implement processShifts logic
```

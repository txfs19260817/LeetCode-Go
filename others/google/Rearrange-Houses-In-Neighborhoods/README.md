# Rearrange Houses in Neighborhoods

- Difficulty: Medium
- Companies: Google
- Stages: Screening
- Asked By: Google
- Source: https://www.hack2hire.com/companies/google/coding-questions/67eae4964c9bb801bfd0ef47/practice?questionId=67eae60e4c9bb801bfd0ef48

## Problem

You are given a 2D array of neighborhoods houses in a city, where each neighborhood (houses[i]) contains house numbers represented by integers. Rearrange the houses such that:

Each neighborhood is sorted in ascending order.
No two houses in the same neighborhood have the same number.
The capacity of each neighborhood (number of houses per row) remains unchanged.

Return the rearranged neighborhoods houses. If it is impossible to rearrange the houses so that each neighborhood meets the above conditions, return the original input unchanged. If there are multiple valid arrangements, returning any one of them is acceptable.

Constraints:

1
<
=
1<= houses.length
<
=
100
<=100
1
<
=
1<= houses[i].length
<
=
100
<=100
0
<
=
0<= houses[i][j]
<
=
10
4
<=10
4

Example 1:

Input: houses = [[1,2], [4,4,7,8], [4,9,9,9]]
Output: [[4,9], [1,4,8,9], [2,4,7,9]]
Explanation: This is one valid arrangement. Other possible answers are:

[[4,9],[1,2,4,9],[4,7,8,9]]
[[4,9],[1,4,7,9],[2,4,8,9]]
[[4,9],[2,4,7,9],[1,4,8,9]]
[[4,9],[2,4,8,9],[1,4,7,9]]
[[4,9],[4,7,8,9],[1,2,4,9]]

Example 2:

Input: houses = [[1,1,2], [2,3]]
Output: [[1,2,3], [1,2]]

Example 3:

Input: houses = [[5,5,5,4], [4,4,3], [3,2]]
Output: [[2,3,4,5], [3,4,5], [4,5]]

## Python Template

```python
from typing import List, Optional

class Solution:
    def rearrangeHouses(self, houses: List[List[int]]) -> List[List[int]]:
        # TODO: Implement rearrangeHouses logic
```

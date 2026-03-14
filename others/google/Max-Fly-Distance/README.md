# Max Fly Distance

- Difficulty: Hard
- Companies: Google
- Stages: Onsite
- Asked By: Google
- Source: https://www.hack2hire.com/companies/google/coding-questions/67fd53c9d2adf72ba62b0ec7/practice?questionId=67fd5511d2adf72ba62b0ec8

## Problem

Given an array arr, where each element represents the distance you can fly on the
𝑖
th
i
th
 day. You start with k units of energy, which is also your maximum energy capacity. Each day, you can either:

Fly: Use 1 unit of energy to fly arr[i] distance.
Rest: Recover 1 unit of energy, up to the maximum of k.

Determine the maximum total distance you can fly over all days.

Constraints:

1 <= arr.length <= 1000
1 <= arr[i] <= 1000
1 <= k <= 100

Example 1:

Input: arr = [5, 2, 8, 4, 3], k = 2
Output: 17
Explanation: The maximum fly distance is achieved by:

Day 0: Fly (energy 2 → 1, distance = 5)
Day 1: Rest (energy 1 → 2, distance = 5)
Day 2: Fly (energy 2 → 1, distance = 5 + 8 = 13)
Day 3: Fly (energy 1 → 0, distance = 13 + 4 = 17)
Day 4: Rest (energy 0 → 1, distance = 17)

Example 2:

Input: arr = [1, 2, 3, 4, 5], k = 3
Output: 13

Example 3:

Input: arr = [[10,1,1,10],2], k = 2
Output: 21

## Python Template

```python
from typing import List, Optional

class Solution:
    def maxFlyDistance(self, arr: List[int], k: int) -> int:
        # TODO: Implement maxFlyDistance logic
```

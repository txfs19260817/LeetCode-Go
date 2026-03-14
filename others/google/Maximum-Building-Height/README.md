# Maximum Building Height

- Difficulty: Medium
- Companies: Google
- Stages: Onsite
- Asked By: Google
- Source: https://www.hack2hire.com/companies/google/coding-questions/67c763516c89e288437c3025/practice?questionId=67c766dc6c89e288437c3026

## Problem

The city plans to allocate land for buildings and parking lots on a grid of equal-sized tiles. Each tile is either fixed as a parking lot ('P') or available for building ('-'). Some tiles may be marked as unusable (represented by 'X').

The planning rule requires that the height difference between any two adjacent tiles (up, down, left, or right) does not exceed one unit. Consequently, a building tile can be constructed taller if it is further from any parking lot—the building's maximum height is determined by the minimum number of moves needed to reach a parking lot.

Your task is to determine the maximum building height that can be planned, which is the greatest number of steps (or moves) required to reach any building tile from the nearest parking lot.

Constraints:

Adjacency includes only up, down, left, and right (diagonals are not considered).
The parking lot locations are fixed.
Unusable tiles ('X') cannot be used for buildings.

Example 1:

Input: [["P","-","-","X","-","-","-"],["-","X","-","-","-","X","-"],["-","-","X","-","-","-","-"],["X","-","-","-","X","-","P"]]
Output: 5
Explanation: The answer is 5 because it represents the maximum distance from any buildable tile ('-') such as (3,2), (3,3), or (0,4) to the nearest parking lot ('P'), and this is the highest height that the building can be built to meet all requirements.

Example 2:

Input: [[["P","X","-","-","-"],["-","X","X","X","-"],["-","-","X","-","-"],["-","X","X","X","-"],["-","-","-","X","P"]]]
Output: 6

Example 3:

Input: [[["P","-","-","X","-","-","-","-"],["-","X","-","X","-","X","-","-"],["-","-","-","-","X","-","-","-"],["-","X","X","-","-","X","X","-"],["-","-","-","-","-","-","-","-"],["-","X","-","X","-","-","X","-"],["-","-","-","-","X","-","-","-"],["-","-","-","-","-","-","-","P"]]]
Output: 11

## Python Template

```python
from typing import List, Optional

class Solution:
    def bfs(self, grid: List[List[str]]) -> int:
        # TODO: Implement bfs logic
```

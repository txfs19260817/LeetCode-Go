# Group Elements by Shared Properties

- Difficulty: Medium
- Companies: Google, Microsoft
- Stages: Screening
- Asked By: Google, Microsoft
- Source: https://www.hack2hire.com/companies/google/coding-questions/678969df11318cc872dd83e9/practice?questionId=67896bfb11318cc872dd83ea

## Problem

Given a fixed-length list of data, where each element has a unique ID and three distinct properties. Two elements are considered duplicates if they share any of these properties.

Group elements based on shared properties, ensuring that each group contains all IDs connected through common properties. If an element has no shared properties with others, it forms its own separate group. You may return the groups in any order.

Constraints:

Each element has exactly one unique ID and three string properties.
The given input list should not be empty.

Example 1:

Input: data = [["id1","p1","p2","p3"], ["id2","p1","p4","p5"], ["id3","p6","p7","p8"]]
Output: [["id2","id1"],["id3"]]
Explanation:

E1 and E2 share the property "p1", so they form one group: ["id2","id1"].
E3 does not share any property with E1 or E2, so it stands alone as ["id3"].

Example 2:

Input: data = [["id1","p1","p2","p3"], ["id2","p1","p4","p5"], ["id3","p5","p7","p8"]]
Output: [["id1","id2","id3"]]

Example 3:

Input: data = [["id1","p1","p2","p3"],["id2","p3","p4","p5"],["id3","p6","p7","p8"],["id4","p8","p9","p10"],["id5","p11","p12","p13"]]
Output: [["id2","id1"],["id4","id3"],["id5"]]

## Python Template

```python
from typing import List, Optional

class Solution:
    def findDuplicates(self, data: List[List[str]]) -> List[List[str]]:
        # TODO: Implement findDuplicates logic
```

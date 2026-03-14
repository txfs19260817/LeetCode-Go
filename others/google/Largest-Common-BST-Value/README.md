# Largest Common BST Value

- Difficulty: Medium
- Companies: Google, Amazon, Microsoft
- Stages: Onsite, Screening
- Asked By: Google, Amazon, Microsoft
- Source: https://www.hack2hire.com/companies/google/coding-questions/67708803366afdc1233e4ffa/practice?questionId=67709b6d366afdc1233e4ffc

## Problem

Given the roots of two Binary Search Trees (BSTs) containing integer values, find the greatest common integer present in both trees. If there is no common integer, return -1.

Constraints:

The number of nodes in both trees is in the range
[
1
,
10
5
]
[1,10
5
].
−
10
9
≤
Node.val
≤
10
9
−10
9
≤Node.val≤10
9
Both trees are valid BSTs.

Example 1:

Input: root1 = [5, 3, 7], root2 = [9, 5, 12, null, 7]
Output: 7
Explanation:  The common integers are 5 and 7. The greatest common integer is 7.

Tree 1:

Tree 2:

Example 2:

Input: root1 = [10, 5, 15, 3, 7, 12, 18], root2 = [20, 15, 25, 10, 18]
Output: 18

Example 3:

Input: root1 = [9, 5, 12, 3, 7, 10], root2 = [7, 3, 10, null, 5, null, 13]
Output: 10

## Python Template

```python
# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
from typing import List, Optional

class Solution:
    def findLargestNumber(self, root1: Optional[TreeNode], root2: Optional[TreeNode]) -> int:
        # TODO: Implement findLargestNumber logic
```

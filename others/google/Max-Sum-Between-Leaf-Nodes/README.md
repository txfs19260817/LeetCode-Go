# Max Sum Between Leaf Nodes

- Difficulty: Easy
- Companies: Google
- Stages: OA
- Asked By: Google
- Source: https://www.hack2hire.com/companies/google/coding-questions/6769d41711bbd512d60ce760/practice?questionId=6769dd5b11bbd512d60ce762

## Problem

Given a binary tree in which each node holds an integer value, determine the maximum path sum that connects two leaf nodes. A leaf node is defined as a node that does not have any left or right children. The path, which must begin and end at leaf nodes, can pass through intermediate (parent) nodes. The path sum is calculated as the sum of the node values along the traversal.

Constraints:

The number of nodes in tree is in the range
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
A tree must be a binary tree.

Example 1:

Input: [10, 2, 10, 20, 1, null, -25, null, null, null, null, 3, 4]

Output: 23
Explanation: The path with the highest sum is 20 → 2 → 1, yielding a total of 23.

Example 2:

Input: [5, 4, 8, 11, null, 13, 4, 7, 2, null, null, 5, 1]

Output: 48
Explanation: In this scenario, the path 7 → 11→ 4 → 5 → 8 → 13 produces the maximum sum of 48.

Example 3:

Input: [1, -2, 3, 4, 5, -6, 2]

Output: 9
Explanation: The path 5 → -2→ 1→ 3 → 2 produces the maximum sum of 9.

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
    def findMaxLeafToLeafSum(self, root: Optional[TreeNode]) -> int:
        # TODO: Implement findMaxLeafToLeafSum logic
```

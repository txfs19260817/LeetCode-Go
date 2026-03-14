# Build Binary Tree From Character Counts

- Difficulty: Medium
- Companies: Google, Square
- Stages: Onsite
- Asked By: Google, Square
- Source: https://www.hack2hire.com/companies/google/coding-questions/68f0129c742c37be9dc830dc/practice?questionId=68f012a4742c37be9dc830dd

## Problem

You are given a string s consisting of alphabetic characters ('A'-'Z', 'a'-'z). Build a binary tree that represents the frequency distribution of characters in s under the following conditions:

Leaf Nodes:
Each distinct character in s must appear exactly once as a leaf node.
A leaf node contains:
ch: the character itself.
count: the number of times that character appears in s.
Internal Nodes:
Every internal node must have:
ch set to '#'.
count equal to the sum of the count values of all leaves in its subtree.
Tree Construction Rules:
The tree must be built such that less frequent characters are combined earlier, and more frequent characters appear closer to the root.
If two characters have the same frequency, the one with the smaller lexicographic value is considered less frequent and must be combined earlier.
The first internal node is formed by combining the two characters with the lowest frequency, with the lower one placed on the left and the higher one on the right.
Every subsequent internal node is formed by combining:
The next character (based on frequency and tie-break rules) as the left child.
The previously constructed subtree as the right child.

You are provided with the following class definition:

class TreeNode {
    char ch;
    int count;
    TreeNode left;
    TreeNode right;
}

Return this node as the root of the tree.

Constraints:

1 ≤ s.length ≤ 10⁴
s consists only of uppercase and lowercase English letters.

Example 1:

Input: s = "aabbbbbcDDD"
Output: ["#, 11", "b, 5", "#, 6",  "D, 3", "#, 3", "c, 1", "a, 2"]
Explanation: The string has four distinct characters with frequencies: {'a': 2, 'b': 5, 'c': 1, 'D': 3}. These are merged step by step, always combining the two nodes with the smallest counts, until one root remains with a count 11.
The tree can be visualized as follows:

Example 2:

Input: s = "AABBC"
Output: [ "#, 5", "B, 2", "#, 3", "C, 1", "A, 2"]

Example 3:

Input: s = "AAAaaaBBBbbb"
Output: ["#, 12", "b, 3", "#, 9", "a, 3", "#, 6", "A, 3", "B, 3"]

## Python Template

```python
from collections import deque

class TreeNode:
    def __init__(self, ch, count):
        self.ch = ch
        self.count = count
        self.left = None
        self.right = None

    def __str__(self):
        return f"{self.ch}, {self.count}"

class Solution:
    def buildTree(self, s):
        # TODO: Implement buildTree logic.
        pass
```

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

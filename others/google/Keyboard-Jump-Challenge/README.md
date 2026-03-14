# Keyboard Jump Challenge

- Difficulty: Medium
- Companies: Google
- Stages: Screening
- Asked By: Google
- Source: https://www.hack2hire.com/companies/google/coding-questions/677c3ffa5cbaff553b53b069/practice?questionId=677c44295cbaff553b53b06a

## Problem

Given a 2D keyboard represented by a m × n character matrix, a maximum jump distance, and a target word, determine if you can type the word by moving between characters without exceeding this distance. The distance here is calculated as the Manhattan distance:
∣
𝑥
1
−
𝑥
2
∣
+
∣
𝑦
1
−
𝑦
2
∣
∣x1−x2∣+∣y1−y2∣.

Constraints:

Jump distance ≥ 0
The word can have up to 1000 characters.

Example 1:

Input:
jumpDistance = 2
keyboard = [["Q","X","P","L","E"],
            ["W","A","C","I","N"]]
word = "PENCIL"
Output: true
Explanation:

Start from 'P' at position (0, 2).
Jump to 'E' at (0, 4), distance = |0-0| + |4-2| = 2.
Jump to 'N' at (1, 4), distance = |1-0| + |4-4| = 1.
Jump to 'C' at (1, 2), distance = |1-1| + |2-4| = 2.
Jump to 'I' at (1, 3), distance = |1-1| + |3-2| = 1.
Finally, jump to 'L' at (0, 3), distance = |0-1| + |3-3| = 1.
All jumps stay within the maximum distance of 2.

Example 2:

Input:
jumpDistance = 4
keyboard = [["T","E","C","H"],
            ["W","A","V","E"],
            ["X","Y","Z","M"],
            ["N","O","P","Q"]]
word = "TECHWAVE"
Output: true

Example 3:

Input:
jumpDistance = 2
keyboard = [["C","A","T","X","O","J","W","Z"],
            ["O","M","B","V","R","T","U","P"],
            ["M","U","R","P","Q","G","D","R"],
            ["P","I","S","R","A","C","T","G"],
            ["U","O","L","E","M","S","P","V"],
            ["T","U","Q","H","I","C","R","A"],
            ["E","X","Z","B","T","D","Y","N"],
            ["F","G","R","G","A","M","P","O"]]
word = "COMPUTERGRAPH"
Output: false

## Python Template

```python
from typing import List, Optional

class Solution:
    def canTypeWord(self, keyboard: List[List[str]], jumpDistance: int, word: str) -> bool:
        # TODO: Implement canTypeWord logic
```

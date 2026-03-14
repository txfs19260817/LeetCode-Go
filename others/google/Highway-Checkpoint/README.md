# Highway Checkpoint

- Difficulty: Medium
- Companies: Google
- Stages: Onsite
- Asked By: Google
- Source: https://www.hack2hire.com/companies/google/coding-questions/677f36a9edc5bbd98ac0ea59/practice?questionId=677f3733edc5bbd98ac0ea5a

## Problem

A highway has multiple toll checkpoints. As vehicles pass each checkpoint, their license plates and the names of the checkpoints are recorded. The toll fee for traveling between two checkpoints is calculated by taking the absolute difference between their positions and multiplying it by 10 (excluding any characters). For example, the toll fee for a vehicle passing through checkpoints "A1" and "A5" is calculated as |1 - 5| * 10 = 40.

At the end of the day, determine the total bill for each vehicle based on the distances it has traveled between the checkpoints it has passed. Assume all test cases are valid and that every vehicle passes through at least two checkpoints. You should return a list of entries formatted as ["License: <license>, Fee: <total_fee>"] in any order.

Constraints:

Each log entry is formatted as ["license", "letter + position", "timestamp"].
The checkpoint name contains alphanumeric characters, and its numeric part represents the checkpoint's position.

Example 1:

Input: logs = ["CAR123,A1,1000", "CAR123,A5,2000"]
Output: ["License: CAR123, Fee: 40"]
Explanation: The car traveled from checkpoint A1 (position 1) to A5 (position 5). Fee = |5 - 1| * 10 = 40.

Example 2:

Input: logs = ["CAR111,C2,1100", "CAR111,C4,1300", "CAR222,C1,1000", "CAR222,C3,1500", "CAR222,C7,2000"]
Output: ["License: CAR111, Fee: 20", "License: CAR222, Fee: 60"]
Explanation:
For CAR111: Fee = |4 - 2| * 10 = 20.
For CAR222: Fee = |3 - 1| * 10 + |7 - 3| * 10 = 20 + 40 = 60.

Example 3:

Input: logs = ["CAR999,D10,3000", "CAR999,D1,1000", "CAR999,D5,2000"]
Output: ["License: CAR999, Fee: 90"]
Explanation: After sorting logs by timestamp: CAR999 travels from D1 (1) → D5 (5) → D10 (10).
Fee = |5 - 1| * 10 + |10 - 5| * 10 = 40 + 50 = 90.

## Python Template

```python
from typing import List, Optional

class Solution:
    def calculateFees(self, logEntries: List[str]) -> List[str]:
        # TODO: Implement calculateFees logic
```

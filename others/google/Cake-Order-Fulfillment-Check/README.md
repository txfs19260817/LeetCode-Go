# Cake Order Fulfillment Check

- Difficulty: Easy
- Companies: Google, Amazon
- Stages: OA
- Asked By: Google, Amazon
- Source: https://www.hack2hire.com/companies/google/coding-questions/67c75ead6c89e288437c301f/practice?questionId=67c7604e6c89e288437c3024

## Problem

You have three types of cakes in stock, represented by an array stock of length 3. You receive multiple orders (each order is also described by three integers for the required quantities of each cake type). An order is fully completed only if the stock has enough cakes for all requested types.

However, regardless of whether the order is fully completed or not, you deduct from your stock as much as possible for each cake type in the order. Your task is to determine how many orders are fully completed under these conditions.

Note: You must process orders in the given order, and may not skip or reorder them.

Constraints:

The stock array has exactly 3 non-negative integers.
Each element of orders is also an array of length 3, representing the required cakes of each type.
An order contributes to the final count only if it is completely fulfilled.

Example 1:

Input: stock = [3, 3, 3], orders = [[1, 0, 1], [0, 2, 0], [1, 1, 1], [2, 2, 2]]
Output: 3
Explanation:

Order 1: [1, 0, 1]
Stock [3, 3, 3] is sufficient (3 ≥ 1, 3 ≥ 0, 3 ≥ 1).
New stock: [2, 3, 2].
Order is fully completed.
Order 2: [0, 2, 0]
Stock [2, 3, 2] is sufficient (2 ≥ 0, 3 ≥ 2, 2 ≥ 0).
New stock: [2, 1, 2].
Order is fully completed.
Order 3: [1, 1, 1]
Stock [2, 1, 2] is sufficient (2 ≥ 1, 1 ≥ 1, 2 ≥ 1).
New stock: [1, 0, 1].
Order is fully completed.
Order 4: [2, 2, 2]
Stock [1, 0, 1] is insufficient (1 < 2, 0 < 2, 1 < 2).
New stock: [0, 0, 0].
Order is not completed.

Thus, 3 orders are fully completed.

Example 2:

Input: stock = [3, 2, 4], orders = [[2, 1, 2], [1, 2, 3], [1, 1, 1]]
Output: 1

Example 3:

Input: stock = [3, 3, 3], orders = [[1, 2, 1], [2, 2, 2]]
Output: 1

## Python Template

```python
from typing import List, Optional

class Solution:
    def getCompletedOrders(self, stock: List[int], orders: List[List[int]]) -> int:
        # TODO: Implement getCompletedOrders logic
```

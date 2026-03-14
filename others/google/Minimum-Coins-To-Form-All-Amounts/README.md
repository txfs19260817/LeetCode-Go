# Minimum Coins to Form All Amounts

- Difficulty: Medium
- Companies: Google, Microsoft, Meta
- Stages: Onsite
- Asked By: Google, Meta, Microsoft
- Source: https://www.hack2hire.com/companies/google/coding-questions/67803a256fee41e51af73ddb/practice?questionId=67803efe6fee41e51af73ddc

## Problem

You are given an array of coin denominations and an integer max, and you are asked to form every integer amount from 1 up to max using these coins.

Find the minimal number of coins used while still being able to create every value in the range [1, max]. Return -1 if these coin denominations cannot form any amount in that range.

Constraints:

1
≤
max
≤
10
5
1≤max≤10
5
.
Coin denominations are positive integers.

Example 1:

Input: coins = [1, 2, 5], max = 11
Output: 5
Explanation: With a total of 5 coins (at least 1 $1 coin, 2 $2 coins, and 2 $5 coins), we can cover all amounts from 1 to 11.

$1: 1 * $1
$2: 1 * $2
$3: 1 * $1 + 1 * $2
$4: 2 * $2
$5: 1 * $1 + 2 * $2
$6: 1 * $5 + 1 * $1
$7: 1 * $5 + 1 * $2
$8: 1 * $1 + 1 * $2 + 1 * $5
$9: 1 * $5 + 2 * $2
$10: 2 * $5
$11: 2 * $5 + 1 * $1

Across all these amounts, we only need 5 coins in total.

Example 2:

Input: coins = [1, 3, 4], max = 6
Output: 4

Example 3:

Input: coins = [2, 5], max = 10
Output: -1

## Python Template

```python
from typing import List, Optional

class Solution:
    def findMinimalCoins(self, coins: List[int], max: int) -> int:
        # TODO: Implement findMinimalCoins logic
```

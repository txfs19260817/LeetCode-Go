# Evaluate String Expression

- Difficulty: Medium
- Companies: Google, Uber
- Stages: Onsite
- Asked By: Google, Uber
- Source: https://www.hack2hire.com/companies/google/coding-questions/67872156e20f8f324213c296/practice?questionId=67872240e20f8f324213c297

## Problem

(This question is a variation of the LeetCode question 224. Basic Calculator. If you haven't completed that question yet, it is recommended to solve it first.)

Implement a function to evaluate a string expression containing nested "add" and "sub" functions. The "add" function adds its arguments, while the "sub" function subtracts the second argument from the first. Parse the expression and compute the correct integer result.

Note that the expression may include negative integers and spaces.

Constraints:

The input expression consists of only "add", "sub" functions, integers, parentheses "()", commas ",", and spaces.
The expression is always valid.
The numbers in the expression are all integers and in the range of [-105, 105].

Example 1:

Input: "add(add(1,3), sub(1,3))"
Output: 2
Explanation: It computes (1 + 3) + (1 - 3) = 2.

Example 2:

Input: "sub(1,3)"
Output: -2

Example 3:

Input: "add(-1, 3)"
Output: 2

## Python Template

```python
from typing import List, Optional

class Solution:
    def evaluate(self, expression: str) -> int:
        # TODO: Implement evaluate logic
```

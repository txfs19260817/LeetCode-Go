# Simplify Expression

- Difficulty: Hard
- Companies: Google, Meta, Oracle
- Stages: Screening, Onsite
- Asked By: Google, Meta, Oracle
- Source: https://www.hack2hire.com/companies/google/coding-questions/677b078878175fce2ecdfc0d/practice?questionId=677b092f78175fce2ecdfc0e

## Problem

Given an algebraic expression containing variables, parentheses, addition ("+"), and subtraction ("-") operators, simplify the expression by removing all unnecessary parentheses and correctly adjusting the signs of the terms. The simplified expression should maintain the correct order of operations and accurately represent the original expression's value.

Constraints:

1 <= expression.length <=
10
5
10
5
The expression consists of lowercase English letters, "+", "-", "(", ")", and no other characters.

Example 1:

Input: "a-(b+c)"
Output: "a-b-c"
Explanation: Removing the parentheses and distributing the negative sign results in "a-b-c".

Example 2:

Input: "a-(-b-c)"
Output: "a+b+c"

Example 3:

Input: "(x+y)-z"
Output: "x+y-z"

## Python Template

```python
from typing import List, Optional

class Solution:
    def simplifyExpression(self, expr: str) -> str:
        # TODO: Implement simplifyExpression logic
```

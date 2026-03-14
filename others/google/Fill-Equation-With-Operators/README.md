# Fill Equation with Operators

- Difficulty: Hard
- Companies: Google, Amazon
- Stages: Onsite
- Asked By: Google, Amazon
- Source: https://www.hack2hire.com/companies/google/coding-questions/67c76c906c89e288437c3030/practice?questionId=67c76f386c89e288437c3031

## Problem

(This question is a variation of the LeetCode question 282. Expression Add Operators. If you haven't completed that question yet, it is recommended to solve it first.)

Given an equation with no operators, you need to insert operators '+', '*', and parentheses '(', ')' only where necessary to form a valid expression that evaluates to the target value, and following standard arithmetic rules (parentheses, then multiplication, then addition).

The numbers must remain in their original order, and you can insert the operators anywhere between them. If at least one valid expression exists, return any one of them in the format of "<expression>=<target>". The output expression should be formatted with minimal parentheses (do not include unnecessary brackets). If no valid arrangement of operators yields the target, return "".

Constraints:

The left-hand side consists of one or more numbers.
You may use only the symbols '(', ')', '+', and '*' in addition to the given numbers.
The original order of numbers must not be altered.
Parentheses '(' and ')' may be added to override standard operator precedence, but they should only be included when required.

Example 1:

Input: numbers = [1,2,3,4,5], target = 105
Output: "(1+2)*(3+4)*5=105"
Explanation: (1 + 2) * (3 + 4) * 5 = 3 * 7 * 5 = 105. Other answers like "(1+((2+3)*4))*5=105" and "(1+2)*((3+4)*5)=105" are also correct.

Example 2:

Input: numbers = [2,3,4], target = 20
Output: "(2+3)*4=20"

Example 3:

Input: numbers = [2,3,5], target = 100
Output: ""

## Python Template

```python
from typing import List, Optional

class Solution:
    def solveEquation(self, numbers: List[int], target: int) -> str:
        # TODO: Implement solveEquation logic
```

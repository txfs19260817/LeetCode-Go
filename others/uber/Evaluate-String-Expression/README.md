# Evaluate String Expression

Implement a function to evaluate a string expression containing nested "add" and "sub" functions. The "add" function adds its arguments, while the "sub" function subtracts the second argument from the first. Parse the expression and compute the correct integer result.

Note that the expression may include negative integers and spaces.

## Constraints

- The input expression consists of only "add", "sub" functions, integers, parentheses "()", commas ",", and spaces.
- The expression is always valid.
- The numbers in the expression are all integers and in the range of `[-10^5, 10^5]`.

## Example 1

**Input:** "add(add(1,3), sub(1,3))"
**Output:** 2
**Explanation:** It computes (1 + 3) + (1 - 3) = 2.

## Example 2

**Input:** "sub(1,3)"
**Output:** -2

## Example 3

**Input:** "add(-1, 3)"
**Output:** 2

## Similar Questions
- [224. Basic Calculator](https://leetcode.com/problems/basic-calculator/)

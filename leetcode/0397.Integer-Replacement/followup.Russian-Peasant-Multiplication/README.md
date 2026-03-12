# Russian Peasant Multiplication

## Problem

Follow-up for `397. Integer Replacement`.

Given two integers `a` and `b`, return their product using the Russian peasant multiplication process:

- repeatedly halve one operand,
- double the other operand,
- add the doubled value into the answer whenever the halved operand is odd.

Implement both an iterative and a recursive solution.

This is a natural follow-up to `Integer Replacement` because both problems rely on the same ideas:

- checking parity,
- halving with shifts,
- using binary structure to avoid brute force.

## Examples

```text
Input: a = 13, b = 12
Output: 156
Explanation:
13  12   keep 12
 6  24   skip
 3  48   keep 48
 1  96   keep 96
Sum = 12 + 48 + 96 = 156
```

```text
Input: a = -13, b = 12
Output: -156
```

```text
Input: a = 0, b = 7
Output: 0
```

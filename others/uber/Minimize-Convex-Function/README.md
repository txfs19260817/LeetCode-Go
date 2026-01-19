# Minimize a Convex Function

## Description

You are given access to a black-box function, `F(x)`, which is convex over a given closed interval `[a, b]`. You can only interact with this function by calling a provided `Evaluate(x)` method, which returns the value of `F(x)`.

Your task is to find a point `x*` within the interval `[a, b]` that is within a specified precision, `eps`, of the true location of the minimum. In other words, if `x_min` is the point where the function's minimum occurs, find an `x*` such that `|x* - x_min| <= eps`.

Hint: A key property of a convex function on an interval is that it is unimodal, meaning it has a single minimum. This property allows for efficient search algorithms.

### Constraints

- `a < b`
- `0 < eps`

## Examples

Example 1:

Input: `F(x) = (x - 3)^2 + 5` (hidden), `a = -10`, `b = 10`, `eps = 0.01`  
Output: `3.0`  
Explanation: The function's global minimum occurs at `x = 3`. Any value in `[2.99, 3.01]` is acceptable.

Example 2:

Input: `F(x) = x^2`, `a = -100`, `b = 50`, `eps = 0.001`  
Output: `0.0`

Example 3:

Input: `F(x) = |x - 123.456|`, `a = 0`, `b = 1000`, `eps = 0.0001`  
Output: `123.456`

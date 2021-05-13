# 1269. Number of Ways to Stay in the Same Place After Some Steps

## LeetCode [1269. Number of Ways to Stay in the Same Place After Some Steps](title)

### Description

You have a pointer at index `0` in an array of size `arrLen`. At each step, you can move 1 position to the left, 1 position to the right in the array or stay in the same place \(The pointer should not be placed outside the array at any time\).

Given two integers `steps` and `arrLen`, return the number of ways such that your pointer still at index `0` after **exactly**`steps` steps.

Since the answer may be too large, return it **modulo** `10^9 + 7`.

**Example 1:**

```text
Input: steps = 3, arrLen = 2
Output: 4
Explanation: There are 4 differents ways to stay at index 0 after 3 steps.
Right, Left, Stay
Stay, Right, Left
Right, Stay, Left
Stay, Stay, Stay
```

**Example 2:**

```text
Input: steps = 2, arrLen = 4
Output: 2
Explanation: There are 2 differents ways to stay at index 0 after 2 steps
Right, Left
Stay, Stay
```

**Constraints:**

* `1 <= steps <= 500`
* `1 <= arrLen <= 10^6`

### Tags

Dynamic Programming

### Solution

We define `dp[steps][idx]` as the number of ways to back to index 0 with exactly `steps` moves. The number of rows is `steps`, yet that of columns is `min(arrLen, steps/2+1)`. If a pointer from index 0 wants to go back with `steps` moves, it can only reach as far as index `steps/2`. We initialize `dp[0][*] = 1`. For each step `i` from 1 to `steps`, we have the transition function `dp[i][j] = dp[i-1][j-1] + dp[i-1][j]+dp[i-1][j+1]`. Finally, we return `dp[steps][0]`. We notice that the state of the current step is only related to its pervious step, thus, we can compress the 2D array into 1D with the length of `min(arrLen, steps/2+1)`.

### Complexity

* Time complexity: $$O(steps×min(arrLen,steps))$$
* Space complexity: $$O(min(arrLen,steps))$$

### Code

```go
func numWays(steps int, arrLen int) int {
	const MOD = 1e9 + 7
	minCols := arrLen
  // If a pointer from index 0 wants to go back with `steps` steps,
	// it can only reach as far as index steps/2.
	if c := steps/2 + 1; c < minCols {
		minCols = c
	}
	dp := make([]int, minCols)
	dp[0] = 1
	for i := 1; i <= steps; i++ {
		dpNext := make([]int, minCols)
		for j := 0; j < minCols; j++ {
			dpNext[j] = dp[j]
			if j-1 >= 0 {
				dpNext[j] = (dpNext[j] + dp[j-1]) % MOD
			}
			if j+1 < minCols {
				dpNext[j] = (dpNext[j] + dp[j+1]) % MOD
			}
		}
		dp = dpNext
	}
	return dp[0]
}
```

## Reference

1. [停在原地的方案数](https://leetcode-cn.com/problems/number-of-ways-to-stay-in-the-same-place-after-some-steps/solution/ting-zai-yuan-di-de-fang-an-shu-by-leetcode-soluti/)


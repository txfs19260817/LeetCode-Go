# Maximum Equal Endpoint Subarray Sum

给定数组 `a[0] ... a[n-1]`，找出一段子数组 `a[i] ... a[j]`，满足：

- `i <= j`
- `a[i] = a[j]`

在所有满足条件的子数组中，返回最大的子数组和。

## Example
**Input:**
```text
nums = [1, -2, 3, 4, -1, 3]
```

**Output:**
```text
9
```

解释：

- 选择子数组 `[3, 4, -1, 3]`
- 它的首尾元素都等于 `3`
- 子数组和为 `3 + 4 + (-1) + 3 = 9`

## Follow-up
能否做到 `O(n)`？

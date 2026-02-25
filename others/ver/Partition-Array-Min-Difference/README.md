# Partition Array Min Difference

## 题目描述
给定整数数组，将其划分为两部分，使：

`abs(sum(part1) - sum(part2))`

最小，返回这个最小差值。

## 解法思路
- 设总和为 `S`，若一部分和为 `x`，差值是 `|S - 2x|`。
- 因此目标变为：找一个“可达子集和” `x`，尽量接近 `S / 2`。
- 使用 0/1 背包布尔 DP：`dp[s]` 表示和 `s` 是否可达。

## 算法步骤
1. 计算 `S = sum(nums)`，令 `target = S // 2`。
2. 初始化 `dp[0] = True`，其余为 `False`。
3. 遍历每个数字 `num`，对 `s` 从 `target` 逆序到 `num` 更新：
   - 若 `dp[s - num]` 为真，则 `dp[s] = True`
4. 从 `target` 向下找最大的可达和 `best`。
5. 返回 `S - 2 * best`。

## 复杂度分析
- 时间：`O(n * target)`，其中 `target = sum(nums) // 2`
- 空间：`O(target)`

## 边界与易错点
- 输入不能为空（实现中空数组抛 `ValueError`）。
- 内层循环必须逆序，否则同一元素会被重复使用（变成完全背包）。
- 单元素数组时答案就是该元素本身。

## 示例
```text
Input:
[2, 3, 10, 7, 5]

Output:
1
```

# Block Chain Mining

给定 `blockSize` 和交易列表 `transactions`，每个交易为：

- `id`
- `size`
- `fee`
- `parentId`（仅 Follow-up 使用）

目标：选择一个子集，使总 `size <= blockSize`，总 `fee` 最大。

## Constraints

- `blockSize` 常见为固定常量 `100`（这是关键）
- `1 <= size`
- `0 <= fee`

## Part 1 (Scalable, not always optimal)

按 `fee / size` 从高到低排序，贪心装入。

- 优点：快，工程上容易并发化
- 缺点：不保证最优

## Part 2 (Optimal baseline, required here)

使用 `0-1 Knapsack DP` 求精确最优解。

- 复杂度 `O(N * blockSize)`
- 当 `blockSize = 100` 时，通常非常可控

## Part 3 (Parent-child follow-up)

新增约束：

- child 被打包时，parent 必须在同一个 block 中
- 同一父链下分支互斥：  
  若 `1 -> 2 -> 3` 与 `1 -> 2 -> 4`，不能同时选 `3` 和 `4`

思路：

1. 用 DFS 枚举每棵树的 root-to-node 路径（每条路径是一个可选方案）
2. 每棵树视为一个 group（组内最多选一条路径）
3. 对 group 做 multiple-choice knapsack DP

## Example

**Input:**

```text
blockSize = 10
transactions = [
  ("a", 6, 13),
  ("b", 5, 10),
  ("c", 5, 10),
]
```

**Output:**

```text
optimalFee = 20
pickedIds = ["b", "c"]
```

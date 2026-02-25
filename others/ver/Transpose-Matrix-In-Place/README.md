# Transpose Matrix In-Place

## 题目描述
实现方阵的原地转置（in-place transpose）。

Follow-up：给定 `num_workers` 和 `worker_index`，让多个 worker 分工完成全部转置交换。

## 解法思路
- 原地转置只需要交换上三角和下三角对应元素：
  - 对所有 `0 <= i < j < n`，交换 `(i, j)` 与 `(j, i)`。
- 并行分工采用“按行取模”：
  - worker `w` 处理所有满足 `i % num_workers == w` 的行
  - 对该行只处理 `j > i` 的位置
- 这样每对元素只会被交换一次，不会冲突或重复。

## 算法步骤
1. 校验输入非空且为 `n x n` 方阵。
2. 单线程版本：
   - 双层循环 `i in [0..n-1]`，`j in [i+1..n-1]`
   - 执行 swap
3. worker 版本：
   - 对每个 worker 按 `i = worker_index, worker_index + num_workers, ...` 遍历
   - 同样仅遍历 `j > i` 并 swap

## 复杂度分析
- 单线程：
  - 时间：`O(n^2)`（准确地说约 `n(n-1)/2` 次交换）
  - 空间：`O(1)`
- 多 worker（理想均衡）：
  - 总工作量仍为 `O(n^2)`
  - 单 worker 工作量约为 `O(n^2 / num_workers)`

## 边界与易错点
- 仅适用于方阵，非方阵应报错。
- `num_workers` 必须为正，`worker_index` 必须在合法范围内。
- 循环必须从 `j = i + 1` 开始，避免对角线和重复交换。

## 示例
```text
Input:
matrix = [
  [1, 2, 3],
  [4, 5, 6],
  [7, 8, 9]
]

Output:
[
  [1, 4, 7],
  [2, 5, 8],
  [3, 6, 9]
]
```

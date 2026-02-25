# Find Common Free Intervals

## 题目描述
给定 `n` 个人的日程 `calendars`，每个人都有若干忙碌区间 `(start, end)`，语义为左闭右开 `[start, end)`。  
给定查询边界 `[day_start, day_end)`，返回所有人都空闲的公共时间区间。

## 解法思路
- 公共空闲区间 = 查询范围内的补集(所有人的忙碌并集)。
- 先把所有人的 busy 区间拍平并裁剪到 `[day_start, day_end)`。
- 对 busy 区间排序并合并重叠段，得到全局 busy 并集。
- 再在边界内求补集得到公共 free。

## 算法步骤
1. 校验 `day_start < day_end`。
2. 遍历所有人的区间：
   - 忽略非法区间（`start >= end`）
   - 裁剪到当天边界：`(max(start, day_start), min(end, day_end))`
   - 只保留有效裁剪结果（`s < e`）
3. 若 busy 为空，直接返回整段 `[(day_start, day_end)]`。
4. 按起点排序并线性合并 busy。
5. 用 `cursor` 从 `day_start` 扫到 `day_end`，在 merged busy 之间收集空档。

## 复杂度分析
- 设所有人区间总数为 `m`。
- 时间：`O(m log m)`（排序主导）
- 空间：`O(m)`（存储拍平区间与合并结果）

## 边界与易错点
- 区间语义是左闭右开，端点相接（如 `[9,10)` 与 `[10,11)`）不产生空档。
- 要先裁剪到 `[day_start, day_end)`，否则会引入边界外噪声。
- 空输入或所有区间都无效时，应返回整段可用时间。
- `day_start >= day_end` 应报错。

## 示例
```text
Input:
calendars = [
  [(9, 10), (12, 13), (16, 18)],
  [(8, 9), (11, 12), (14, 17)],
  [(9, 10), (13, 14), (15, 16)],
]
day_start = 8
day_end = 18

Output:
[(10, 11)]
```

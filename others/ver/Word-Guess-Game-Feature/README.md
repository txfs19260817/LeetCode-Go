# WordGuess Game Feature

## 题目描述
实现 WordGuess 打分功能，颜色编码如下：
- `GREEN = 2`：字符和位置都匹配
- `YELLOW = 1`：字符存在于 secret 中，但位置不匹配
- `GRAY = 0`：字符不存在于 secret 的剩余可用字符中

输入 `secret` 与 `guess`，返回每个位置的整数评分数组。输入非法时抛异常。

## 解法思路
- 采用两遍扫描，正确处理重复字符：
  - 第一遍先锁定所有 `GREEN`
  - 同时统计 secret 中“未被 GREEN 消耗”的字符库存
  - 第二遍对非 GREEN 位置尝试打 `YELLOW`，命中后库存减一
- 这样可避免重复字母被错误重复计分。

## 算法步骤
1. 校验 `secret` 和 `guess` 都是字符串，且非空、长度相同。
2. 初始化结果数组为全 `GRAY`，并准备计数器 `remaining_secret`。
3. 第一遍遍历：
   - `guess[i] == secret[i]` 则标记 `GREEN`
   - 否则把 `secret[i]` 计入 `remaining_secret`
4. 第二遍遍历所有非 GREEN 位置：
   - 若 `guess[i]` 在 `remaining_secret` 中计数大于 0，标 `YELLOW` 并减一
5. 返回结果数组。

## 复杂度分析
- 时间：`O(n)`
- 空间：`O(k)`，`k` 为字符集大小（上界可视作 `O(n)`）

## 边界与易错点
- 长度不同必须报错，不能继续打分。
- 处理重复字母时不能“一次命中全部标黄”，必须基于剩余库存。
- 必须先处理 GREEN，再处理 YELLOW，否则结果会偏大。

## 示例
```text
Input:
secret = "hello"
guess  = "hable"

Output:
[2, 0, 0, 2, 1]
```

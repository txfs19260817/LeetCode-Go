# Find Path in Fibonacci Tree

**Difficulty:** Hard
**Company:** Databricks
**Interview Stages:** Screening
**Topics:** Binary Tree, Recursion

(Variation of LeetCode 2096. Step-By-Step Directions From a Binary Tree Node to Another.)

## Description

A Fibonacci tree of order `n` (denoted `Fn(n)`) is defined recursively:

- **Base cases:** `Fn(0)` and `Fn(1)` are each a single node (leaf).
- **Recursive:** `Fn(n)` has a root, left subtree `Fn(n-2)`, and right subtree `Fn(n-1)`.

Nodes are labeled in **pre-order** traversal starting from 0.

Given `order`, `source`, and `dest`, return the path from `source` to `dest` as a string of moves:
- `'L'`: move to left child
- `'R'`: move to right child
- `'U'`: move to parent

## Constraints

- 2 ≤ order ≤ 10
- 0 ≤ source, dest ≤ n - 1

## Example

**Input:** order = 5, source = 5, dest = 7
**Output:** "UUURL"

**Explanation:**
5 → parent 3 ("U") → parent 1 ("U") → parent 0 ("U") → right child 6 ("R") → left child 7 ("L")

```
order=5, source=5, dest=7  → "UUURL"
order=4, source=8, dest=3  → "UUULR"
order=5, source=4, dest=13 → "UUURRRL"
```

## Approach (不建树, O(k) time, O(k) space)

### 1) 预处理每阶树的节点数 `size[k]`

```
size[0] = 1
size[1] = 1
size[k] = 1 + size[k-2] + size[k-1]
```

### 2) 利用 preorder 的连续区间定位节点在哪个子树

假设当前子树是 `F(k)`，根编号是 `r`：
- 左子树根 = `r + 1`，范围 `[r+1, r + size[k-2]]`
- 右子树根 = `r + 1 + size[k-2]`，范围是后面那一段

写一个 `getPath(k, r, x)` 函数：一路判断 `x` 落在左子树还是右子树，递归/迭代下去，拿到 root → x 的路径（L/R 序列），复杂度 O(k)。

### 3) 找 LCA，拼路径

- 分别拿到 root → a 和 root → b 两条路径
- 找 LCA：从头一起走，最后一个相同的分叉点
- a → b 路径 = a 往上到 LCA（全是 `'U'`）+ LCA 往下到 b（对应 L/R 序列后半段）
- 拼接时注意 LCA 不重复

### 注意点

- 区间边界别写错（左子树大小用 `size[k-2]` 不是 `size[k-1]`）
- size 用 `int64` / `long long`（k 大时可能溢出，面试可以提一嘴 BigInt/截断比较）

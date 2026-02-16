# Balanced Tree Without Builtins

实现一个不依赖任何内建数据结构（如 map/set/heap）来维护节点关系的平衡二叉搜索树（AVL Tree）。

支持以下操作：
- `Insert(key int)`
- `Delete(key int)`
- `Contains(key int) bool`
- `InOrder() []int`

重复插入同一个 key 时忽略该次插入。

## Example
**Input:**
```
Insert: 10, 20, 30, 40, 50, 25
InOrder
Delete: 40
Contains: 40
Contains: 25
InOrder
```

**Output:**
```
[10 20 25 30 40 50]
false
true
[10 20 25 30 50]
```

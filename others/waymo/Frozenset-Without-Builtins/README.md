# Frozenset Without Builtins

实现一个不可变集合（frozenset），且不允许使用任何 built-in 集合结构（如 `map` / `set`）维护成员。

要求：
- 集合元素为整数 `int`
- 每次更新操作都返回新集合，原集合不变
- 推荐用面试可手写方案：`有序去重数组 + 二分查找 + 双指针集合运算`
- 支持以下接口：
  - `With(value int) -> new set`
  - `Without(value int) -> new set`
  - `Contains(value int) bool`
  - `Union(other)`, `Intersection(other)`, `Difference(other)`
  - `Elements() []int`（升序）

## Example
**Input:**
```
s1 = frozenset([5, 1, 3, 3])
s2 = s1.with(4)
s3 = s2.without(1)
s4 = frozenset([3, 7])
union = s3.union(s4)
intersection = s3.intersection(s4)
difference = s3.difference(s4)
```

**Output:**
```
s1.elements() = [1, 3, 5]
s2.elements() = [1, 3, 4, 5]
s3.elements() = [3, 4, 5]
union.elements() = [3, 4, 5, 7]
intersection.elements() = [3]
difference.elements() = [4, 5]
```

# Find Super Admin User

## 题目描述
给定权限三元组 `(entity_id, relation, object_id)`，其中：
- `camera_owner`：实体直接拥有某摄像头权限
- `group_member`：实体属于某组，并继承该组权限

组可以多级嵌套且保证无环。题目保证存在且仅存在一个用户，能访问组织内全部摄像头，返回该用户 ID。

## 解法思路
- 建图表示继承关系：`entity -> parent_group`。
- 记录每个实体直接拥有的摄像头集合。
- 对每个用户做 DFS，沿继承边向上汇总可访问摄像头。
- 使用记忆化避免重复计算同一实体的可访问集合。
- 能覆盖全部摄像头的唯一用户即答案。

## 算法步骤
1. 一次遍历输入，构建：
   - `member_graph[entity] = [parent_group...]`
   - `owned_cameras[entity] = {camera...}`
   - `all_cameras`（全局摄像头集合）
   - `users`（所有用户实体）
2. 定义 `get_accessible_cameras(entity)`：
   - 基础值：实体的直接摄像头集合
   - 递归合并所有父组的可访问集合
   - 结果写入 memo
3. 遍历 `users`，找出可访问集合大小等于 `len(all_cameras)` 的用户并返回。

## 复杂度分析
- 设实体数为 `V`、继承关系数为 `E`、摄像头总数为 `C`。
- 时间：图遍历本身近似 `O(V + E)`，集合并集成本与数据分布有关。
- 空间：`O(V + E + C)`（图、集合、memo）。

## 边界与易错点
- 多级组嵌套必须递归处理，不能只看一层。
- 同一实体可属于多个组，需并集合并。
- 题目给出“无环”保证；若线上不保证，应加环检测。
- 非 `user_` 前缀实体不参与最终候选。

## 示例
```text
Input:
[
  ("user_1", "camera_owner", "camera_1"),
  ("user_1", "group_member", "group_1"),
  ("group_1", "camera_owner", "camera_2"),
  ("group_1", "group_member", "group_2"),
  ("group_2", "camera_owner", "camera_3"),
  ("user_2", "camera_owner", "camera_3"),
  ("user_2", "group_member", "group_1"),
  ("user_3", "group_member", "group_2"),
  ("user_3", "camera_owner", "camera_1")
]

Output:
user_1
```

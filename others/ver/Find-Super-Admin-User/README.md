# Find Super Admin User

## 题目描述
给定权限三元组 `(entity_id, relation, target_id)`，其中：
- `camera_owner`：`entity_id` 直接拥有 `target_id` 这台摄像头。
- `group_member`：`entity_id` 是 `target_id` 这个组的成员，并继承该组权限。

组允许多级嵌套（且无环）。题目保证存在且仅存在一个用户，可以访问组织内所有摄像头，返回该用户 ID。

## 解法思路
- 建立“继承图”：`entity -> parent_groups`。
- 记录每个实体直接拥有的摄像头集合。
- 对每个用户执行 DFS，向上递归收集所有可访问摄像头。
- 用缓存（memo）保存每个实体的“可访问摄像头集合”，避免重复遍历。
- 可访问摄像头数量等于“全局摄像头总数”的用户，就是 super admin。

## 算法步骤
1. 一次遍历输入，构建：
   - `parent_groups_by_entity[entity] = [parent_group...]`
   - `direct_cameras_by_entity[entity] = {camera...}`
   - `all_camera_ids`（全局摄像头集合）
   - `candidate_user_ids`（所有 `user_` 前缀实体）
2. 定义 `collect_accessible_cameras(entity)`：
   - 基础值：实体的直接摄像头集合
   - 递归合并所有父组的可访问集合
   - 结果写入缓存
3. 遍历 `candidate_user_ids`，找出可访问集合大小等于 `len(all_camera_ids)` 的用户并返回。

## 复杂度分析
- 设实体数为 `V`、继承关系数为 `E`、摄像头总数为 `C`。
- 时间：图遍历本身近似 `O(V + E)`，集合并集成本与数据分布有关。
- 空间：`O(V + E + C)`（图、集合、缓存）。

## 边界与易错点
- 多级组嵌套必须递归处理，不能只看一层。
- 同一实体可属于多个组，需并集合并。
- 题目给出“无环”保证；若线上不保证，应加环检测。
- 非 `user_` 前缀实体不参与最终候选用户。
- 如果输入不满足“唯一解保证”，当前实现会抛出 `RuntimeError`。

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

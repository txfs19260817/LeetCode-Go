from __future__ import annotations

from typing import Dict, Iterable, List, Set, Tuple

Permission = Tuple[str, str, str]


def find_super_admin(permissions_data: Iterable[Permission]) -> str:
    """Return the unique user ID that can access every camera."""
    member_graph: Dict[str, List[str]] = {}
    owned_cameras: Dict[str, Set[str]] = {}
    all_cameras: Set[str] = set()
    users: Set[str] = set()

    for entity, relation, obj in permissions_data:
        if relation == "group_member":
            # entity inherits permissions from parent group "obj".
            member_graph.setdefault(entity, []).append(obj)
        elif relation == "camera_owner":
            owned_cameras.setdefault(entity, set()).add(obj)
            all_cameras.add(obj)
        else:
            continue

        if entity.startswith("user_"):
            users.add(entity)

    memo: Dict[str, Set[str]] = {}

    def get_accessible_cameras(entity: str) -> Set[str]:
        # DFS + memoization over inheritance graph to avoid repeated traversal.
        if entity in memo:
            return memo[entity]

        cameras = set(owned_cameras.get(entity, set()))
        for parent_group in member_graph.get(entity, []):
            cameras |= get_accessible_cameras(parent_group)

        memo[entity] = cameras
        return cameras

    target_camera_count = len(all_cameras)
    # Guarantee says exactly one user can access all cameras.
    for user in users:
        if len(get_accessible_cameras(user)) == target_camera_count:
            return user

    raise RuntimeError("No admin user found; input may violate guarantees.")


if __name__ == "__main__":
    sample_permissions = [
        ("user_1", "camera_owner", "camera_1"),
        ("user_1", "group_member", "group_1"),
        ("group_1", "camera_owner", "camera_2"),
        ("group_1", "group_member", "group_2"),
        ("group_2", "camera_owner", "camera_3"),
        ("user_2", "camera_owner", "camera_3"),
        ("user_2", "group_member", "group_1"),
        ("user_3", "group_member", "group_2"),
        ("user_3", "camera_owner", "camera_1"),
    ]

    assert find_super_admin(sample_permissions) == "user_1"

    extra_case = [
        ("user_a", "group_member", "group_x"),
        ("group_x", "group_member", "group_y"),
        ("group_y", "camera_owner", "camera_1"),
        ("group_x", "camera_owner", "camera_2"),
        ("user_b", "camera_owner", "camera_1"),
    ]

    assert find_super_admin(extra_case) == "user_a"
    print("All assertions passed.")

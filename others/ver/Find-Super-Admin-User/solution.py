from __future__ import annotations

from collections.abc import Iterable

Permission = tuple[str, str, str]


def find_super_admin(permission_records: Iterable[Permission]) -> str:
    """Return the unique user ID that can access every camera."""
    # entity -> direct parent groups (entity inherits permissions from these groups)
    parent_groups_by_entity: dict[str, list[str]] = {}
    # entity -> cameras owned directly by this entity
    direct_cameras_by_entity: dict[str, set[str]] = {}
    all_camera_ids: set[str] = set()
    candidate_user_ids: set[str] = set()

    for entity_id, relation, target_id in permission_records:
        if relation == "group_member":
            parent_groups_by_entity.setdefault(entity_id, []).append(target_id)
        elif relation == "camera_owner":
            direct_cameras_by_entity.setdefault(entity_id, set()).add(target_id)
            all_camera_ids.add(target_id)
        else:
            continue

        if entity_id.startswith("user_"):
            candidate_user_ids.add(entity_id)

    accessible_camera_cache: dict[str, set[str]] = {}

    def collect_accessible_cameras(entity_id: str) -> set[str]:
        # DFS + memoization over inheritance graph to avoid repeated traversal.
        if entity_id in accessible_camera_cache:
            return accessible_camera_cache[entity_id]

        accessible_cameras = set(direct_cameras_by_entity.get(entity_id, set()))
        for parent_group_id in parent_groups_by_entity.get(entity_id, []):
            accessible_cameras |= collect_accessible_cameras(parent_group_id)

        accessible_camera_cache[entity_id] = accessible_cameras
        return accessible_cameras

    total_camera_count = len(all_camera_ids)
    # Guarantee says exactly one user can access all cameras.
    for user_id in candidate_user_ids:
        if len(collect_accessible_cameras(user_id)) == total_camera_count:
            return user_id

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

    deep_group_chain = [
        ("user_a", "group_member", "group_x"),
        ("group_x", "group_member", "group_y"),
        ("group_y", "camera_owner", "camera_1"),
        ("group_x", "camera_owner", "camera_2"),
        ("group_x", "group_member", "group_z"),
        ("group_z", "camera_owner", "camera_3"),
        ("user_b", "camera_owner", "camera_1"),
    ]
    assert find_super_admin(deep_group_chain) == "user_a"

    multi_parent_group_inheritance = [
        ("user_main", "group_member", "group_x"),
        ("user_main", "group_member", "group_y"),
        ("group_x", "camera_owner", "camera_1"),
        ("group_y", "camera_owner", "camera_2"),
        ("user_other", "camera_owner", "camera_1"),
    ]
    assert find_super_admin(multi_parent_group_inheritance) == "user_main"

    ignores_unknown_relation = [
        ("user_1", "group_member", "group_1"),
        ("group_1", "camera_owner", "camera_1"),
        ("group_1", "camera_owner", "camera_2"),
        ("user_1", "unknown_relation", "camera_999"),
        ("user_2", "camera_owner", "camera_1"),
    ]
    assert find_super_admin(ignores_unknown_relation) == "user_1"

    breaks_guarantee = [
        ("user_1", "camera_owner", "camera_1"),
        ("user_2", "camera_owner", "camera_2"),
    ]
    try:
        find_super_admin(breaks_guarantee)
        raise AssertionError("Expected RuntimeError when input violates guarantee")
    except RuntimeError:
        pass

    print("All assertions passed.")

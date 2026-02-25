from __future__ import annotations

from typing import Dict, List, Sequence, Tuple


PathTuple = Tuple[str, ...]


def _parse_paths(entries: Sequence[str], sep: str) -> List[PathTuple]:
    paths: List[PathTuple] = []
    for entry in entries:
        # Normalize spaces and drop empty segments from malformed inputs.
        parts = tuple(part.strip() for part in entry.split(sep) if part.strip())
        if parts:
            paths.append(parts)
    return paths


def _is_prefix(parent: PathTuple, child: PathTuple) -> bool:
    return len(parent) <= len(child) and child[: len(parent)] == parent


def merge_failures_sort(entries: Sequence[str], sep: str = ",") -> List[str]:
    """
    Merge failures using sorting.
    Parent path keeps, descendant paths are removed.
    """
    paths = _parse_paths(entries, sep)
    paths.sort()

    merged: List[PathTuple] = []
    for path in paths:
        # After lexicographic sort, descendants follow their parent prefix.
        if merged and _is_prefix(merged[-1], path):
            continue
        merged.append(path)

    return [sep.join(path) for path in merged]


class TrieNode:
    __slots__ = ("children", "terminal")

    def __init__(self) -> None:
        self.children: Dict[str, TrieNode] = {}
        self.terminal = False


def merge_failures_trie(entries: Sequence[str], sep: str = ",") -> List[str]:
    """
    Merge failures using trie.
    Parent path keeps, descendant paths are removed.
    """
    root = TrieNode()

    def insert(parts: Sequence[str]) -> None:
        node = root
        for part in parts:
            # Existing terminal parent already covers this full path.
            if node.terminal:
                return
            node = node.children.setdefault(part, TrieNode())

        node.terminal = True
        # Drop descendants: once parent is terminal, children are redundant.
        node.children.clear()

    for parts in _parse_paths(entries, sep):
        insert(parts)

    result: List[str] = []

    def collect(node: TrieNode, path: List[str]) -> None:
        if node.terminal:
            result.append(sep.join(path))
            return

        # Sort for deterministic output independent of insertion order.
        for key in sorted(node.children):
            collect(node.children[key], path + [key])

    collect(root, [])
    return result


if __name__ == "__main__":
    sample = ["o1,s1,c1", "o1,s1", "o1,s2,c1"]
    expected = ["o1,s1", "o1,s2,c1"]

    assert merge_failures_sort(sample) == expected
    assert merge_failures_trie(sample) == expected

    case_parent_covers_all = ["o1", "o1,s1", "o1,s1,c1", "o1,s2,c2"]
    assert merge_failures_sort(case_parent_covers_all) == ["o1"]
    assert merge_failures_trie(case_parent_covers_all) == ["o1"]

    case_no_parent = ["o1,s2,c1", "o1,s1,c1", "o2,s1,c1", "o1,s1,c2"]
    expected_no_parent = ["o1,s1,c1", "o1,s1,c2", "o1,s2,c1", "o2,s1,c1"]
    assert merge_failures_sort(case_no_parent) == expected_no_parent
    assert merge_failures_trie(case_no_parent) == expected_no_parent

    case_duplicates_and_spaces = [" o1,s1 ", "o1,s1,c1", "o1,s1", "o1,s2,c1 "]
    expected_dup = ["o1,s1", "o1,s2,c1"]
    assert merge_failures_sort(case_duplicates_and_spaces) == expected_dup
    assert merge_failures_trie(case_duplicates_and_spaces) == expected_dup

    print("All assertions passed.")

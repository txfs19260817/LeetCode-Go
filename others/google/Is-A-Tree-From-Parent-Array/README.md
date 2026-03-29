# Is A Tree From Parent Array

You are given an integer array `parents` where `parents[i]` is the parent of node `i`.

- `parents[i] = -1` means node `i` is the root.
- Otherwise `parents[i]` must be a valid node index.

Return `True` if the array describes a valid tree, otherwise return `False`.

A valid tree in this representation must satisfy:

- There is exactly one root
- Every non-root node points to a valid parent index
- There is no cycle

## Example
**Input:**
```text
parents = [1, -1, 1, 2, 5, 2]
parents = [-1, 0, -1, 2, 3]
parents = [1, -1, 3, 4, 5, 2]
```

**Output:**
```text
True
False
False
```

Explanation:

- `[1, -1, 1, 2, 5, 2]` is a valid tree with root `1`
- `[-1, 0, -1, 2, 3]` has two roots, so it is not a tree
- `[1, -1, 3, 4, 5, 2]` contains a cycle `2 -> 3 -> 4 -> 5 -> 2`

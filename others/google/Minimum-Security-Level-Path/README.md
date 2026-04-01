# Minimum Security Level Path

You are given a directed graph with `n` nodes. Each edge has a security level.

You may traverse a path only if your allowed security level `S` is at least the security level of every edge on that path.

The goal is to find the minimum security level required to travel from node `u` to node `v`.

Equivalent formulation:

- Each path has a cost equal to the maximum edge security level used on that path
- Return the minimum such cost among all paths from `u` to `v`

Assumptions used by this implementation:

- The graph is directed
- If `u == v`, the answer is `0`
- If `v` is unreachable from `u`, return `-1`

## Example
**Input:**
```text
n = 4
u = 0
v = 3
graph = [
    [(1, 3), (2, 1)],
    [(3, 5)],
    [(3, 2)],
    [],
]
```

**Output:**
```text
2
```

Explanation:

- Path `0 -> 1 -> 3` needs security `max(3, 5) = 5`
- Path `0 -> 2 -> 3` needs security `max(1, 2) = 2`
- The minimum required security level is `2`

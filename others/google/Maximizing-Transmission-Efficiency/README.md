# Maximizing Transmission Efficiency

You are given a directed graph with `n` nodes. Each edge has a transmission efficiency in the range `[0, 1]`.

The efficiency of a path is the product of the efficiencies of its edges.

Implement algorithms for the following tasks:

1. Basic: find the maximum transmission efficiency from `start` to `destination`
2. Follow-up 1: one edge on the chosen path can be upgraded to efficiency `1.0`
3. Follow-up 2: find the maximum transmission efficiency among paths that use an even number of edges
4. Follow-up 3: each edge also has a positive integer cost, and the total path cost must not exceed `B`

Assumptions used by this implementation:

- The graph is directed
- If `start == destination`, the empty path has efficiency `1.0`
- If no valid path exists, return `0.0`

## Example
**Input:**
```text
n = 4
start = 0
destination = 3
graph = [
    [(1, 0.5), (2, 0.2)],
    [(2, 0.9), (3, 0.5)],
    [(3, 0.8)],
    [],
]
```

**Output:**
```text
basic = 0.36
one_upgrade = 0.8
even_edges = 0.25
```

Explanation:

- Basic best path is `0 -> 1 -> 2 -> 3`, with efficiency `0.5 * 0.9 * 0.8 = 0.36`
- If edge `0 -> 2` is upgraded to `1.0`, path `0 -> 2 -> 3` becomes `1.0 * 0.8 = 0.8`
- Among even-edge paths, the best is `0 -> 1 -> 3`, with efficiency `0.5 * 0.5 = 0.25`

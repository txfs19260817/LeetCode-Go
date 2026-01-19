# Earliest Time All Riders Are Connected

## Description

You are given a sequence of log entries of the form:

```
time, riderA, riderB
```

Each entry means rider `A` and rider `B` share a ride at `time`, creating a connection between them. Connections are transitive: if `A` is connected to `B` and `B` to `C`, then `A` is connected to `C`.

Your task is to return the earliest time when **all riders are connected**.

If it never happens, return `-1`.

## Requirements

- Implement a function that takes:
  - `n`: total number of riders, IDs are `0..n-1`
  - `logs`: a list of connections with timestamps
- Return the earliest timestamp when all riders are connected.

## Follow-up

If riders can **block** other riders, connections are no longer monotonic and Union-Find alone is insufficient. Discuss an alternative approach (e.g., dynamic graph with BFS/DFS per query, or maintaining adjacency and re-evaluating connectivity).

## Example

Input:

```
n = 4
logs = [
  (20190101, 0, 1),
  (20190104, 3, 2),
  (20190107, 2, 0),
  (20190211, 1, 2),
  (20190224, 2, 3),
]
```

Output:

```
20190107
```

Explanation:
By time 20190107, all 4 riders become connected through transitive connections.

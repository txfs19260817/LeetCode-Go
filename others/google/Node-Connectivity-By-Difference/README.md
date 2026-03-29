# Node Connectivity By Difference

You are given:

- an integer array `arr` of length `N`
- an integer `diff`
- a list of queries `queries`, where each query is a pair of indices `[u, v]`

Treat each index as a node. Two nodes belong to the same connected component if their values can be linked through a chain of values where each adjacent pair in sorted order differs by at most `diff`.

Return a list of booleans where each answer tells whether the two queried indices are connected.

An `O(N log N + Q)` approach is expected:

- Sort pairs `(value, index)`
- Scan the sorted list and assign component ids
- Start a new component whenever the gap between adjacent sorted values is greater than `diff`
- Answer each query by checking whether both indices have the same component id

## Example
**Input:**
```text
arr = [1, 2, 3, 6]
diff = 2
queries = [[0, 2], [1, 3]]
```

**Output:**
```text
[True, False]
```

Explanation:

- Values `1, 2, 3` form one component because adjacent sorted gaps are `1` and `1`.
- Value `6` starts a new component because `6 - 3 = 3 > diff`.
- So index `0` is connected to index `2`, but index `1` is not connected to index `3`.

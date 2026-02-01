# N Ary Tree Path Sum

This problem involves three parts related to an N-ary tree:

1.  **Sum of All Nodes**: Return the sum of all node values in the tree.
2.  **Max Path Sum**: Return the maximum path value from root to leaf.
3.  **Max Path Nodes**: Return the values of the nodes in the maximum path found in Part 2.

The `Node` class/struct needs to be defined.

## Example

Consider the following N-ary tree:

```
      1
    /   \
   2     3
       /   \
      4     5
```

**Input:**
Root of the tree.

**Output:**

- **Part 1**: `15` (1 + 2 + 3 + 4 + 5)
- **Part 2**: `9` (Path 1 -> 3 -> 5 gives 1 + 3 + 5 = 9, which is greater than 1->2 (3) and 1->3->4 (8))
- **Part 3**: `[1, 3, 5]`

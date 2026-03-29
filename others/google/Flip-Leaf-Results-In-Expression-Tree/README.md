# Flip Leaf Results In Expression Tree

## Problem

You are given a boolean expression tree. Each node is one of:

- `leaf`: stores a boolean value
- `not`: unary operator
- `and`
- `or`
- `xor`

For each leaf, in left-to-right order, flip only that leaf's boolean value while keeping every other leaf unchanged. Return the result of the whole expression tree after each such flip.

## Example

The expression:

```text
Or(And(True, False), Not(False))
```

can be represented as:

```python
root = Node(
    "or",
    left=Node(
        "and",
        left=Node("leaf", True),
        right=Node("leaf", False),
    ),
    right=Node(
        "not",
        left=Node("leaf", False),
    ),
)
```

**Output:**

```text
[True, True, False]
```

## Notes

An `O(N)` solution is possible:

1. Evaluate the current value of every node with a postorder traversal.
2. Propagate whether flipping a subtree would flip the root value.
3. For each leaf, append `root_value XOR impact`.

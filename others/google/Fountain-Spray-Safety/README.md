# Fountain Spray Safety

You are given a terrain elevation array `terrain`, where `terrain[i]` is the height at position `i`, and a list of fountain positions `fountains`.

Each fountain at index `i` sprays water to the left and right, but only across positions that are strictly lower than `terrain[i]`.

For a fountain at position `i`:

- Walk left from `i - 1` while `terrain[j] < terrain[i]` and mark those positions as unsafe.
- Walk right from `i + 1` while `terrain[j] < terrain[i]` and mark those positions as unsafe.
- Stop in a direction as soon as you reach a position whose height is greater than or equal to the fountain height.
- The fountain's own position is not marked unsafe.

Return a binary array where:

- `1` means the position is unsafe to stand on because it gets sprayed by at least one fountain.
- `0` means the position is safe.

An `O(n)` solution is possible:

- For each index, compute the nearest position on the left with height greater than or equal to the current height.
- Compute the analogous nearest blocking position on the right.
- Each fountain then sprays one contiguous interval on the left and one on the right.
- Mark those intervals with a difference array and build the final binary result with one prefix pass.

## Example
**Input:**
```text
terrain = [2, 1, 3, 2, 1, 1]
fountains = [0, 3]
```

**Output:**
```text
[0, 1, 0, 0, 1, 1]
```

Explanation:

- Fountain at index `0` has height `2`, so it sprays position `1` and stops before position `2` because `terrain[2] = 3`.
- Fountain at index `3` has height `2`, so it sprays positions `4` and `5`, and it cannot spray left because `terrain[2] = 3`.

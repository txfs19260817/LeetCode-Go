# Remove Covered Point

Given a list of non-overlapping intervals `[start, end)` (start inclusive, end exclusive), and an integer `idx` (0-based index in the flattened sequence of all covered integers ordered by the input intervals), remove the integer at position `idx` and return the updated list of intervals.

Removing a point may: eliminate an interval (if it contains only 1 point), shrink it (adjust a boundary), or split it into two parts. The original interval order is preserved.

## Examples

**Example 1:**
```
Input:  intervals = [[10,12],[13,16],[4,8]], idx = 3
Flattened: [10,11,13,14,15,4,5,6,7]
Point at idx 3 is 14.
Output: [[10,12],[13,14],[15,16],[4,8]]
```

**Example 2:**
```
Input:  intervals = [[4,8],[13,16],[10,12]], idx = 0
Flattened: [4,5,6,7,13,14,15,10,11]
Point at idx 0 is 4.
Output: [[5,8],[13,16],[10,12]]
```

**Example 3:**
```
Input:  intervals = [[2,6],[8,10],[15,18]], idx = 3
Flattened: [2,3,4,5,8,9,15,16,17]
Point at idx 3 is 5.
Output: [[2,5],[8,10],[15,18]]
```

## Algorithm

1. Walk through intervals accumulating total point count.
2. Find which interval contains the `idx`-th point: for each interval, `size = end - start`. If `idx < size`, the point is `start + idx`.
3. Modify the interval:
   - Size 1 → remove the interval entirely.
   - Point at start → shrink left: `[start+1, end]`.
   - Point at end−1 → shrink right: `[start, end-1]`.
   - Otherwise → split: `[start, point]` and `[point+1, end]`.
4. Return modified intervals list preserving order.

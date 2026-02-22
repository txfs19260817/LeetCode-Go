# K-Means Clustering

Given `n` points in `d` dimensions, an integer `k`, and `maxIter`, implement K-Means clustering.

Assume valid inputs:
- `1 <= k <= n`
- all points have the same dimension `d`
- `maxIter >= 1`

Use the first `k` points as initial centroids (deterministic). For each iteration:
- assign each point to the nearest centroid (tie: smaller centroid index),
- recompute each centroid as the mean of assigned points,
- if a centroid gets no points, keep it unchanged.

Return final centroids and assignment index for each point.

## Example
**Input:**
```text
points = [
  [1,1], [1.5,2], [3,4], [5,7], [3.5,5], [4.5,5], [3.5,4.5]
]
k = 2
maxIter = 10
```

**Output:**
```text
centroids ~= [[1.25, 1.5], [3.9, 5.1]]
assignments = [0, 0, 1, 1, 1, 1, 1]
```

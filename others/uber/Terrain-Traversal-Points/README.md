# Terrain Traversal Points

You are given:

- an `m x n` integer matrix `terrain`
- an array `limits` of size `k`

You must compute an array `answer` of size `k` such that for each `limits[i]`:

1. Start at the top-left cell `(0,0)`.
2. Repeat:
   - If `limits[i]` is strictly greater than the value of the current cell, then:
     - You gain 1 point the first time you visit this cell.
     - You may move to any adjacent cell (up, down, left, right).
   - Otherwise:
     - You gain no points.
     - The process ends for this path.

After the process, `answer[i]` is the maximum points achievable. You are allowed to revisit cells multiple times, but points are only earned on the first visit.

Return `answer`.

## Example

**Input:**

```
terrain = [[1,4,2,8], [0,4,0,8], [1,2,0,8]]
limits = [8, 2]
```

**Output:**

```
[9, 3]
```

## Explanation

- For `limit = 8`:

  - Cells < 8 reachable from (0,0):
    - (0,0)=1
    - (0,1)=4
    - (0,2)=2
    - (1,0)=0
    - (1,1)=4
    - (1,2)=0
    - (2,0)=1
    - (2,1)=2
    - (2,2)=0
    - (0,3)=8 (stop)
    - (1,3)=8 (stop)
    - (2,3)=8 (stop)
  - All 9 cells in the first 3 columns are strictly less than 8 and connected. Total points = 9.

- For `limit = 2`:
  - Start (0,0)=1 (< 2). Points = 1.
  - Neighbors:
    - (0,1)=4 (>= 2, stop)
    - (1,0)=0 (< 2, continue). Points = 1+1=2.
      - Neighbors of (1,0):
        - (1,1)=4 (>= 2, stop)
        - (2,0)=1 (< 2, continue). Points = 2+1=3.
          - Neighbors of (2,0):
            - (2,1)=2 (>= 2, stop)
            - (1,0) (visited)
  - Total points = 3.

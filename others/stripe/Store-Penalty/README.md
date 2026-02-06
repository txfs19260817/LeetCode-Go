# Store Penalty

Each hour, a store can either have customers (`Y`) or not (`N`). If the store
is open when there are no customers, or closed when there are customers, a
penalty is incurred.

Given a log of hours, find the best closing time (hour index) that minimizes
penalty. Also parse aggregated logs that contain segments delimited by
`BEGIN ... END` and compute best closing times for each segment.

## Input

- `store_log`: space-separated string of `Y` and `N`.
- `agg_log`: text containing multiple segments.

## Output

- Best closing time per segment.

## Example

**Input**

```
Y Y N N
```

**Output**

```
2
```

## Constraints

- Hours are 0-based indices.
- Logs may contain multiple segments.

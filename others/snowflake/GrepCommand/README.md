# Grep Command

This problem simulates the behavior of the `grep` command with context control (`-C` or `-A`/`-B` functionality).

## Problem Description

Suppose you are given an array where each element is a line in a document. Implement a grep function that prints lines matching a search target. Additionally, the function accepts a `linesAround` parameter, which specifies the number of lines to print before and after each match.

**Constraints:**
- No line should be printed more than once.
- Overlapping contexts should be merged.

### Example

Input:
```
lines = [
  'good morning',
  'hello there',
  'my name is Alex',
  'my friend is albert',
  'it is nice to meet you Alex'
]
search target: "Alex"
linesAround = 1
```

Output:
```
'hello there'
'my name is Alex'
'my friend is Albert'
'it is nice to meet you Alex'
```

Note: 'my friend is albert' is printed only once, even though it is within 1 line of two different matches.

## Implementations & Complexity

Let $N$ be the number of lines, $L$ be the average length of a line, and $K$ be the context window size (`linesAround`).

### 1. Basic Grep (`Grep`)
Standard array-based implementation.

- **Time Complexity:** $O(N \cdot L)$
  - We scan every line once to find matches ($O(N \cdot L)$).
  - We iterate through the boolean array to collect results ($O(N)$).
- **Space Complexity:** $O(N)$
  - To store the boolean `shouldPrint` array.

### 2. Parallel Grep (`GrepParallel`)
Optimized for large datasets using goroutines to parallelize the search.

- **Time Complexity:** $O(\frac{N \cdot L}{P} + N)$ where $P$ is number of CPUs.
  - The search phase is parallelized.
  - The aggregation phase is sequential $O(N)$.
- **Space Complexity:** $O(N)$
  - For `shouldPrint` array and intermediate match indices.

### 3. Streaming Grep (`StreamingGrep`)
Processes lines one by one, maintaining a buffer. Useful when the entire file cannot fit in memory or data is arriving in real-time.

- **Time Complexity:** $O(N \cdot L)$
  - Each line is processed exactly once.
- **Space Complexity:** $O(K \cdot L)$
  - We only store up to `linesAround` lines in the buffer.

### 4. Parallel Streaming Grep (`GrepStreamingParallel`)
Combines streaming with parallel processing. Uses a worker pool to parallelize the matching logic while preserving output order.

- **Time Complexity:** Approximately $O(\frac{N \cdot L}{P})$ throughput.
  - Latency for a single line might be higher due to channel overhead and reordering.
- **Space Complexity:** $O((K + W) \cdot L)$
  - Buffers for reordering (proportional to worker count $W$) and context window ($K$).

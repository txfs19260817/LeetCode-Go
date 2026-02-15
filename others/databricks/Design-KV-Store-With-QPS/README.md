# Design KV Store With QPS

**Difficulty:** Medium  
**Company:** Databricks  
**Related:** [362. Design Hit Counter](https://leetcode.com/problems/design-hit-counter/)

Design a key-value store that supports `Put` and `Get` operations, and
additionally exposes a `GetQPS` API that returns the average
queries-per-second over the last 5 minutes (300 seconds).

The problem is intentionally open-ended — requirements like the exact time
window, how QPS is defined, and whether `Get` counts as a hit are things
you should clarify with the interviewer.

## Interface

```
KVStoreWithQPS()                          // constructor
Put(key string, value string, ts int)     // store key-value, counts as a hit
Get(key string, ts int) -> string         // retrieve value, counts as a hit
GetQPS(ts int) -> float64                 // hits in (ts-300, ts] / 300
```

Timestamps are in seconds and are guaranteed non-decreasing across calls.

## Example

```
store := NewKVStoreWithQPS()
store.Put("foo", "bar", 1)       // hit at t=1
store.Put("baz", "qux", 2)      // hit at t=2
store.Get("foo", 3)       -> "bar"  // hit at t=3
store.GetQPS(3)            -> 0.01  // 3 hits / 300 = 0.01
store.GetQPS(301)          -> 0.0067 // 2 hits / 300; t=1 expired
```

## Follow-up

- **Space optimisation:** Instead of storing every timestamp, use a
  circular buffer of size 300 (one slot per second). Each slot tracks
  the timestamp it belongs to and the count for that second. This gives
  O(W) space where W = window size, regardless of throughput.
- **Concurrency:** How would you make this thread-safe?
- **Distributed:** How would you aggregate QPS across multiple nodes?

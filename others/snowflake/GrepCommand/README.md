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

## Follow-up: Streaming

Instead of being given the entire array of lines upfront, suppose you are fed a stream of lines one by one. Implement the function to process lines as they arrive.

- The function should maintain internal state.
- It should output lines as soon as it can determine they should be printed.

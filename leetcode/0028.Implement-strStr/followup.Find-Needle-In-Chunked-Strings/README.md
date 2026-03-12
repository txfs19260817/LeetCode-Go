# Find Needle In Chunked Strings

## Problem

Follow-up for `28. Implement strStr()`.

Given a string array `chunks` and a string `needle`, treat `chunks` as one continuous string in order.
Return the start position of the first occurrence of `needle` as `[chunkIndex, offset]`.
The match is allowed to cross chunk boundaries.

If `needle` does not exist, return `[-1, -1]`.
If `needle` is empty, return `[0, 0]`.

## Example

Input:

```text
chunks = ["a", "bcd", "ea", "de"]
needle = "de"
```

Concatenated string:

```text
"abcdeade"
```

The first `"de"` starts at `chunks[1][2]`, so the answer is:

```text
[1, 2]
```

## Notes

- This directory is only a scaffold.
- The Go and Python solutions are intentionally left unimplemented.

# License Plate By Index

Write a function `f(index)` that maps a non-negative integer index to a 5-character license plate string ranging from `00000` to `ZZZZZ`.

Assumptions used by this implementation:

- Plates are generated in blocks:
  `00000..99999`, then `A0000..Z9999`, then `AA000..ZZ999`, ..., until `ZZZZZ`
- A plate always has total length 5
- A prefix of letters is followed by a suffix of digits
- `0` maps to `00000`
- The largest valid index maps to `ZZZZZ`

## Example
**Input:**
```text
index = 0
index = 99999
index = 100000
index = max_index
```

**Output:**
```text
00000
99999
A0000
ZZZZZ
```

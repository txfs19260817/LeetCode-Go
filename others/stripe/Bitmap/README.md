# Bitmap Utilities

You are given a bitmap as a 2D array of `0/1` integers. Implement three helpers:

- Part 1: **Compress** each row with run-length encoding (RLE).
- Part 2: **Decompress** the RLE representation back into the bitmap.
- Part 3: **Flip** the bitmap horizontally (mirror left-right).

## Input

- `bitmap`: 2D integer array with values `0` or `1`.
- `compressed_bitmap`: list of rows of `(value, count)` pairs.

## Output

- Part 1: `compress_bitmap` returns list of rows of `(value, count)` pairs.
- Part 2: `decompress_bitmap` returns a 2D integer array.
- Part 3: `flip_bitmap_horizontally` returns a 2D integer array.

## Example

**Input bitmap row**
```
[0, 0, 0, 1, 1]
```

**Compressed row**
```
[(0, 3), (1, 2)]
```

**Flipped row**
```
[1, 1, 0, 0, 0]
```

## Constraints

- All rows have the same length.
- Values are only `0` or `1`.

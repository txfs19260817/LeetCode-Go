# Bitmap Character to Image Conversion

Implement a reusable bitmap pipeline for character rendering.

## Parts Mapping (Code)

- `Part 1`: `BitFontSolver.part1_print_character`
- `Part 2`:
- `BitFontSolver.part2_compress_lookup_table`
- `BitFontSolver.part2_decompress_lookup_table`
- `Part 3`: `BitFontSolver.part3_invert_from_compressed_and_print`

`Part 3` explicitly reuses `Part 2` by calling `decompress_bitmap` first.

## Phase 1: Basic Printing

Given a lookup table (`dict`) where:

- key: character, such as `"J"`
- value: 2D array of `0/1` bitmap pixels

print a character's bitmap to console in row order.

## Phase 2: Compression and Decompression

Implement compression and decompression for the lookup table.

- `compress`: convert each bitmap row into a compact representation
- `decompress`: restore original rows
- verify integrity: decompressed result must equal original bitmap

This solution uses row-wise run-length encoding (RLE):

- raw row: `[0, 0, 0, 0, 1, 0]`
- compressed: `[(0, 4), (1, 1), (0, 1)]`

## Phase 3: Manipulation and Reuse

Apply additional processing (example: invert `0 <-> 1`) and print result.

Requirement: reuse Phase 2 code.
In this solution:

- read compressed bitmap from lookup table
- reuse `decompress_bitmap` from Phase 2
- apply manipulation (`invert_bitmap`)
- print manipulated bitmap

## Input

- `lookup_table`: `Dict[str, List[List[int]]]`
- optional `text`: a string to render character-by-character

## Output

- console rows using `0/1` values
- verification via `assert` in demo:
- decompressed bitmap equals original bitmap
- double invert returns original bitmap

## Example Bitmap (`J`)

```text
0 0 0 0 1 0
0 0 0 0 1 0
0 0 0 0 1 0
0 0 0 0 1 0
0 0 0 0 1 0
0 0 0 0 1 0
1 0 0 0 1 0
0 1 1 1 0 0
0 0 0 0 0 0
0 0 0 0 0 0
```

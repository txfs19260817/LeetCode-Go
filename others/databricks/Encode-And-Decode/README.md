# Encode And Decode

Given an array of non-negative integers, design an encoder and decoder using Run-Length Encoding (RLE) and Bit-Packing (BP).

**RLE Run:** Consecutive identical values → `"RLE[value,count]"`. Used for runs of 8+ repeats (except the last RLE run may be shorter).

**BP Run:** Groups of up to 8 non-RLE values → `"BP[v1,v2,...,vk]"`. Exactly 8 per run except the last may have fewer.

## Encoding Algorithm

1. Process left to right, identify maximal runs of consecutive identical values.
2. If run length >= 8 → RLE.
3. If run length < 8 AND it's the last run AND the BP buffer is empty → RLE (last run exception).
4. Otherwise → add to BP buffer, flush BP buffer every 8 values.
5. When encountering an RLE-eligible run, flush any accumulated BP buffer first.

## Example 1

**Input:**
```
[5,5,5,5,5,5,5,5,1,2,3]
```

**Output:**
```
["RLE[5,8]", "BP[1,2,3]"]
```

## Example 2

**Input:**
```
[1,1,1]
```

**Output:**
```
["RLE[1,3]"]
```

## Example 3

**Input:**
```
[1,1,1,1,2,3,4,5]
```

**Output:**
```
["BP[1,1,1,1,2,3,4,5]"]
```

## Decode

Parse each string: `"RLE[val,count]"` → repeat val count times. `"BP[v1,v2,...,vk]"` → extract values.

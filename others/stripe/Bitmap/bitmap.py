from typing import List, Tuple

Bitmap = List[List[int]]
CompressedBitmap = List[List[Tuple[int, int]]]


def print_bitmap(bitmap: Bitmap | CompressedBitmap) -> None:
    """Pretty-print a bitmap for quick visual inspection."""
    for row in bitmap:
        print(" ".join(str(cell) for cell in row))


# Part 1: compress each row with run-length encoding (RLE)
def compress_bitmap(bitmap: Bitmap) -> CompressedBitmap:
    """
    Convert each row into RLE pairs: (value, count).
    Empty rows are encoded as empty lists.
    """
    compressed: CompressedBitmap = []
    for row in bitmap:
        if not row:
            compressed.append([])
            continue

        compressed_row: List[Tuple[int, int]] = []
        prev = row[0]
        count = 1
        for i in range(1, len(row)):
            if row[i] == prev:
                count += 1
            else:
                compressed_row.append((prev, count))
                prev = row[i]
                count = 1
        compressed_row.append((prev, count))
        compressed.append(compressed_row)

    return compressed


# Part 2: decompress RLE back into the original bitmap
def decompress_bitmap(compressed_bitmap: CompressedBitmap) -> Bitmap:
    """Expand each (value, count) pair into its original row."""
    decompressed: Bitmap = []
    for row in compressed_bitmap:
        decompressed_row: List[int] = []
        for value, count in row:
            decompressed_row.extend([value] * count)
        decompressed.append(decompressed_row)

    return decompressed


# Part 3: flip the bitmap horizontally (mirror left-right)
def flip_bitmap_horizontally(bitmap: Bitmap) -> Bitmap:
    """Reverse each row to produce a horizontal flip."""
    return [list(reversed(row)) for row in bitmap]


if __name__ == "__main__":
    # Example bitmap for "J"
    j_bitmap: Bitmap = [
        [0, 0, 0, 0, 1, 0],
        [0, 0, 0, 0, 1, 0],
        [0, 0, 0, 0, 1, 0],
        [0, 0, 0, 0, 1, 0],
        [0, 0, 0, 0, 1, 0],
        [0, 0, 0, 0, 1, 0],
        [1, 0, 0, 0, 1, 0],
        [0, 1, 1, 1, 0, 0],
        [0, 0, 0, 0, 0, 0],
        [0, 0, 0, 0, 0, 0],
    ]

    compressed_j = compress_bitmap(j_bitmap)
    decompressed_j = decompress_bitmap(compressed_j)
    flipped_j = flip_bitmap_horizontally(decompressed_j)

    # Basic verification
    assert decompressed_j == j_bitmap
    assert flip_bitmap_horizontally(flipped_j) == j_bitmap

    # Optional visualization
    print("Original Bitmap:")
    print_bitmap(j_bitmap)
    print("Compressed Bitmap:")
    print_bitmap(compressed_j)
    print("Flipped Bitmap:")
    print_bitmap(flipped_j)

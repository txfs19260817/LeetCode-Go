from typing import Dict, List, Tuple

Bitmap = List[List[int]]
CompressedRow = List[Tuple[int, int]]
CompressedBitmap = List[CompressedRow]
LookupTable = Dict[str, Bitmap]
CompressedLookupTable = Dict[str, CompressedBitmap]


def validate_bitmap(bitmap: Bitmap) -> None:
    if not bitmap:
        return

    width = len(bitmap[0])
    for row in bitmap:
        if len(row) != width:
            raise ValueError("All bitmap rows must have equal width.")
        for bit in row:
            if bit not in (0, 1):
                raise ValueError("Bitmap values must be 0 or 1.")


def print_bitmap(bitmap: Bitmap) -> None:
    for row in bitmap:
        print(" ".join(str(bit) for bit in row))


class BitFontSolver:
    # -------- Part 1: Basic Printing --------
    def part1_print_character(self, lookup_table: LookupTable, ch: str) -> None:
        if ch not in lookup_table:
            raise KeyError(f"Character '{ch}' not found in lookup table.")
        validate_bitmap(lookup_table[ch])
        print_bitmap(lookup_table[ch])

    # -------- Part 2: Compression / Decompression --------
    def compress_row_rle(self, row: List[int]) -> CompressedRow:
        if not row:
            return []

        output: CompressedRow = []
        current = row[0]
        count = 1

        for value in row[1:]:
            if value == current:
                count += 1
            else:
                output.append((current, count))
                current = value
                count = 1

        output.append((current, count))
        return output

    def decompress_row_rle(self, compressed_row: CompressedRow) -> List[int]:
        row: List[int] = []
        for value, count in compressed_row:
            if value not in (0, 1):
                raise ValueError("Compressed value must be 0 or 1.")
            if count < 0:
                raise ValueError("Run length must be non-negative.")
            row.extend([value] * count)
        return row

    def compress_bitmap(self, bitmap: Bitmap) -> CompressedBitmap:
        validate_bitmap(bitmap)
        return [self.compress_row_rle(row) for row in bitmap]

    def decompress_bitmap(self, compressed_bitmap: CompressedBitmap) -> Bitmap:
        bitmap = [self.decompress_row_rle(row) for row in compressed_bitmap]
        validate_bitmap(bitmap)
        return bitmap

    def part2_compress_lookup_table(
        self, lookup_table: LookupTable
    ) -> CompressedLookupTable:
        return {ch: self.compress_bitmap(bitmap) for ch, bitmap in lookup_table.items()}

    def part2_decompress_lookup_table(
        self, compressed_lookup: CompressedLookupTable
    ) -> LookupTable:
        return {ch: self.decompress_bitmap(bitmap) for ch, bitmap in compressed_lookup.items()}

    # -------- Part 3: Manipulation + Reuse Part 2 --------
    def invert_bitmap(self, bitmap: Bitmap) -> Bitmap:
        validate_bitmap(bitmap)
        return [[1 - bit for bit in row] for row in bitmap]

    def part3_invert_from_compressed_and_print(
        self, compressed_lookup: CompressedLookupTable, ch: str
    ) -> None:
        if ch not in compressed_lookup:
            raise KeyError(f"Character '{ch}' not found in compressed lookup table.")

        # Reuse Part 2 code: decompress first.
        bitmap = self.decompress_bitmap(compressed_lookup[ch])
        inverted = self.invert_bitmap(bitmap)
        print_bitmap(inverted)


def build_demo_lookup_table() -> LookupTable:
    return {
        "J": [
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
    }


if __name__ == "__main__":
    solver = BitFontSolver()
    lookup_table = build_demo_lookup_table()

    print("Part 1: print original bitmap for 'J'")
    solver.part1_print_character(lookup_table, "J")

    compressed_lookup = solver.part2_compress_lookup_table(lookup_table)
    restored_lookup = solver.part2_decompress_lookup_table(compressed_lookup)

    assert restored_lookup == lookup_table, "Part 2 failed: decompressed data mismatch."

    print("\nPart 2: print decompressed bitmap for 'J'")
    solver.part1_print_character(restored_lookup, "J")

    print("\nPart 3: invert bitmap for 'J' (reuse Part 2 decompress)")
    solver.part3_invert_from_compressed_and_print(compressed_lookup, "J")

    original_j = lookup_table["J"]
    assert solver.invert_bitmap(solver.invert_bitmap(original_j)) == original_j

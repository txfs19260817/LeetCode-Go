from typing import List, Tuple, Dict, Callable, Optional

class Solution:
    # -------------------------
    # Phase 1: Compress (RLE)
    # -------------------------
    def compressBitmap(self, bitmap: List[List[int]]) -> List[List[Tuple[int, int]]]:
        """
        Compress each row with run-length encoding.
        Return: list of rows of (value, count).
        """
        compressed: List[List[Tuple[int, int]]] = []

        for row in bitmap:
            if not row:
                compressed.append([])
                continue

            out_row: List[Tuple[int, int]] = []
            cur = row[0]
            cnt = 1

            for x in row[1:]:
                if x == cur:
                    cnt += 1
                else:
                    out_row.append((cur, cnt))
                    cur = x
                    cnt = 1

            out_row.append((cur, cnt))
            compressed.append(out_row)

        return compressed

    # -------------------------
    # Phase 2: Decompress + Print (Reusable)
    # -------------------------
    def decompressBitmap(self, compressed_bitmap: List[List[Tuple[int, int]]]) -> List[List[int]]:
        """
        Decompress RLE back to bitmap.
        """
        bitmap: List[List[int]] = []
        expected_len: Optional[int] = None

        for rle_row in compressed_bitmap:
            row: List[int] = []
            for value, count in rle_row:
                if value not in (0, 1):
                    raise ValueError(f"Invalid value {value}, expected 0 or 1.")
                if count < 0:
                    raise ValueError(f"Invalid count {count}, expected non-negative.")
                row.extend([value] * count)

            if expected_len is None:
                expected_len = len(row)
            else:
                if len(row) != expected_len:
                    raise ValueError("All rows must have the same length after decompression.")

            bitmap.append(row)

        return bitmap

    def printBitmap(self, bitmap: List[List[int]]) -> None:
        """
        Print bitmap with spaces.
        Phase 4 will reuse this printing logic.
        """
        for row in bitmap:
            print(" ".join(str(x) for x in row))

    # -------------------------
    # Phase 3: Flip horizontally (mirror left-right)
    # -------------------------
    def flipBitmapHorizontally(self, bitmap: List[List[int]]) -> List[List[int]]:
        return [row[::-1] for row in bitmap]

    # -------------------------
    # Phase 4: Manipulation & Reuse (Invert example)
    # -------------------------
    def invertBitmap(self, bitmap: List[List[int]]) -> List[List[int]]:
        """
        Invert bitmap: 0->1, 1->0
        """
        return [[1 - x for x in row] for row in bitmap]

    def phase4InvertFromLookupAndPrint(
            self,
            lookup_table: Dict[str, List[List[Tuple[int, int]]]],
            key: str
    ) -> None:
        """
        Phase 4 requirement: MUST reuse Phase 2 code.
        Strategy:
          1) get compressed from lookup_table
          2) reuse Phase 2 decompressBitmap()
          3) invert
          4) reuse Phase 2 printBitmap()
        """
        if key not in lookup_table:
            raise KeyError(f"Key '{key}' not found in lookup_table.")

        compressed = lookup_table[key]

        # ✅ reuse Phase 2: decompress
        bitmap = self.decompressBitmap(compressed)

        # Phase 4: manipulate
        inverted = self.invertBitmap(bitmap)

        # ✅ reuse Phase 2: print
        self.printBitmap(inverted)

    # （可选）更通用的 Phase 4：传入任意操作函数，仍然复用 Phase 2
    def phase4ManipulateFromLookupAndPrint(
            self,
            lookup_table: Dict[str, List[List[Tuple[int, int]]]],
            key: str,
            op: Callable[[List[List[int]]], List[List[int]]]
    ) -> None:
        if key not in lookup_table:
            raise KeyError(f"Key '{key}' not found in lookup_table.")

        bitmap = self.decompressBitmap(lookup_table[key])  # ✅ reuse Phase 2
        manipulated = op(bitmap)
        self.printBitmap(manipulated)  # ✅ reuse Phase 2


# -------------------------
# Demo usage
# -------------------------
if __name__ == "__main__":
    sol = Solution()

    bitmap = [
        [0, 0, 0, 0, 1, 0],
        [0, 1, 1, 1, 0, 0],
    ]

    # Phase 1
    compressed = sol.compressBitmap(bitmap)
    print("Compressed:", compressed)

    # Phase 2
    decompressed = sol.decompressBitmap(compressed)
    print("\nDecompressed:")
    sol.printBitmap(decompressed)

    # Phase 3
    flipped = sol.flipBitmapHorizontally(bitmap)
    print("\nFlipped:")
    sol.printBitmap(flipped)

    # Phase 4: lookup table -> decompress -> invert -> print (reuse Phase 2)
    lookup_table = {"imgA": compressed}
    print("\nPhase 4 Inverted (from lookup):")
    sol.phase4InvertFromLookupAndPrint(lookup_table, "imgA")

class Solution:
    def encode(self, values: list[int]) -> list[str]:
        if not values:
            return []

        # Identify maximal runs of consecutive identical values.
        runs: list[tuple[int, int]] = []
        i = 0
        while i < len(values):
            val = values[i]
            cnt = 1
            while i + cnt < len(values) and values[i + cnt] == val:
                cnt += 1
            runs.append((val, cnt))
            i += cnt

        result: list[str] = []
        bp_buf: list[int] = []

        def flush_bp() -> None:
            nonlocal bp_buf
            if not bp_buf:
                return
            result.append(f"BP[{','.join(str(v) for v in bp_buf)}]")
            bp_buf = []

        for idx, (val, cnt) in enumerate(runs):
            is_last = idx == len(runs) - 1

            if cnt >= 8:
                flush_bp()
                result.append(f"RLE[{val},{cnt}]")
            elif is_last and not bp_buf:
                # Last run exception
                result.append(f"RLE[{val},{cnt}]")
            else:
                for _ in range(cnt):
                    bp_buf.append(val)
                    if len(bp_buf) == 8:
                        flush_bp()

        flush_bp()
        return result

    def decode(self, runs: list[str]) -> list[int]:
        result: list[int] = []
        for s in runs:
            if s.startswith("RLE["):
                inner = s[4:-1]
                parts = inner.split(",")
                value, count = int(parts[0]), int(parts[1])
                result.extend([value] * count)
            elif s.startswith("BP["):
                inner = s[3:-1]
                parts = inner.split(",")
                result.extend(int(p) for p in parts)
        return result


if __name__ == "__main__":
    sol = Solution()

    # Example 1
    inp1 = [5, 5, 5, 5, 5, 5, 5, 5, 1, 2, 3]
    enc1 = sol.encode(inp1)
    assert enc1 == ["RLE[5,8]", "BP[1,2,3]"], f"Got {enc1}"
    assert sol.decode(enc1) == inp1

    # Example 2
    inp2 = [1, 1, 1]
    enc2 = sol.encode(inp2)
    assert enc2 == ["RLE[1,3]"], f"Got {enc2}"
    assert sol.decode(enc2) == inp2

    # Example 3
    inp3 = [1, 1, 1, 1, 2, 3, 4, 5]
    enc3 = sol.encode(inp3)
    assert enc3 == ["BP[1,1,1,1,2,3,4,5]"], f"Got {enc3}"
    assert sol.decode(enc3) == inp3

    # Single element
    inp4 = [42]
    enc4 = sol.encode(inp4)
    assert enc4 == ["RLE[42,1]"], f"Got {enc4}"
    assert sol.decode(enc4) == inp4

    # Long mixed
    inp5 = [7]*10 + [1, 2, 3] + [2]*9
    enc5 = sol.encode(inp5)
    assert enc5 == ["RLE[7,10]", "BP[1,2,3]", "RLE[2,9]"], f"Got {enc5}"
    assert sol.decode(enc5) == inp5

    print("All tests passed!")

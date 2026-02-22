class Solution:
    def encode(self, values: list[int]) -> list[str]:
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
            if bp_buf:
                result.append(f"BP[{','.join(str(v) for v in bp_buf)}]")
                bp_buf.clear()

        for idx, (val, cnt) in enumerate(runs):
            if cnt >= 8:
                flush_bp()
                result.append(f"RLE[{val},{cnt}]")
            elif idx == len(runs) - 1 and not bp_buf:
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


class StreamEncoder:
    def __init__(self) -> None:
        self.cur_val: int = 0
        self.cur_cnt: int = 0
        self.bp_buf: list[int] = []
        self.started: bool = False

    def write(self, value: int) -> list[str]:
        if not self.started:
            self.started = True
            self.cur_val = value
            self.cur_cnt = 1
            return []

        if value == self.cur_val:
            self.cur_cnt += 1
            return []

        result = self._finalize_run(is_last=False)
        self.cur_val = value
        self.cur_cnt = 1
        return result

    def flush(self) -> list[str]:
        if not self.started:
            return []
        result = self._finalize_run(is_last=True)
        self.started = False
        self.cur_cnt = 0
        return result

    def _finalize_run(self, is_last: bool) -> list[str]:
        result: list[str] = []

        if self.cur_cnt >= 8:
            result.extend(self._flush_bp())
            result.append(f"RLE[{self.cur_val},{self.cur_cnt}]")
        elif is_last and not self.bp_buf:
            result.append(f"RLE[{self.cur_val},{self.cur_cnt}]")
        else:
            for _ in range(self.cur_cnt):
                self.bp_buf.append(self.cur_val)
                if len(self.bp_buf) == 8:
                    result.extend(self._flush_bp())
            if is_last:
                result.extend(self._flush_bp())

        return result

    def _flush_bp(self) -> list[str]:
        if not self.bp_buf:
            return []
        run = f"BP[{','.join(str(v) for v in self.bp_buf)}]"
        self.bp_buf.clear()
        return [run]


class StreamDecoder:
    def write(self, run: str) -> list[int]:
        if run.startswith("RLE["):
            inner = run[4:-1]
            parts = inner.split(",")
            value, count = int(parts[0]), int(parts[1])
            return [value] * count
        if run.startswith("BP["):
            inner = run[3:-1]
            return [int(p) for p in inner.split(",")]
        return []


def encode_via_stream(values: list[int]) -> list[str]:
    enc = StreamEncoder()
    runs: list[str] = []
    for v in values:
        runs.extend(enc.write(v))
    runs.extend(enc.flush())
    return runs


def decode_via_stream(runs: list[str]) -> list[int]:
    dec = StreamDecoder()
    values: list[int] = []
    for r in runs:
        values.extend(dec.write(r))
    return values


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
    inp5 = [7] * 10 + [1, 2, 3] + [2] * 9
    enc5 = sol.encode(inp5)
    assert enc5 == ["RLE[7,10]", "BP[1,2,3]", "RLE[2,9]"], f"Got {enc5}"
    assert sol.decode(enc5) == inp5

    # Stream Example 1
    s_enc1 = encode_via_stream(inp1)
    assert s_enc1 == ["RLE[5,8]", "BP[1,2,3]"], f"Got {s_enc1}"
    assert decode_via_stream(s_enc1) == inp1

    # Stream Example 2
    s_enc2 = encode_via_stream(inp2)
    assert s_enc2 == ["RLE[1,3]"], f"Got {s_enc2}"
    assert decode_via_stream(s_enc2) == inp2

    # Stream Example 3
    s_enc3 = encode_via_stream(inp3)
    assert s_enc3 == ["BP[1,1,1,1,2,3,4,5]"], f"Got {s_enc3}"
    assert decode_via_stream(s_enc3) == inp3

    # Stream Single element
    s_enc4 = encode_via_stream(inp4)
    assert s_enc4 == ["RLE[42,1]"], f"Got {s_enc4}"
    assert decode_via_stream(s_enc4) == inp4

    # Stream Long mixed
    s_enc5 = encode_via_stream(inp5)
    assert s_enc5 == ["RLE[7,10]", "BP[1,2,3]", "RLE[2,9]"], f"Got {s_enc5}"
    assert decode_via_stream(s_enc5) == inp5

    # Stream Empty
    assert encode_via_stream([]) == []

    # Stream Incremental emission
    enc = StreamEncoder()
    for _ in range(8):
        assert enc.write(5) == []
    assert enc.write(1) == ["RLE[5,8]"]
    assert enc.write(2) == []
    assert enc.write(3) == []
    assert enc.flush() == ["BP[1,2,3]"]

    # Stream BP flush at 8
    enc = StreamEncoder()
    for _ in range(7):
        assert enc.write(1) == []
    assert enc.write(2) == []
    assert enc.write(3) == ["BP[1,1,1,1,1,1,1,2]"]
    assert enc.flush() == ["RLE[3,1]"]

    print("All tests passed!")

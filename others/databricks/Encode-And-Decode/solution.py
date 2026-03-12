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


class Run:
    def encode(self) -> str:
        raise NotImplementedError


class RLERun(Run):
    def __init__(self, value: int, count: int) -> None:
        self.value = value
        self.count = count

    def encode(self) -> str:
        return f"RLE[{self.value}, {self.count}]"


class BPRun(Run):
    def __init__(self, values: list[int]) -> None:
        self.values = list(values)

    def encode(self) -> str:
        return f"BP{self.values}"


class Encoder:
    def __init__(self) -> None:
        self.cur_val: int = 0
        self.cur_cnt: int = 0
        self.bp_buf: list[int] = []
        self.started: bool = False
        self.runs: list[Run] = []

    def append(self, value: int) -> None:
        if not self.started:
            self.started = True
            self.cur_val = value
            self.cur_cnt = 1
            return

        if value == self.cur_val:
            self.cur_cnt += 1
            return

        self._finalize_run(is_last=False)
        self.cur_val = value
        self.cur_cnt = 1

    def finish(self) -> list[Run]:
        if not self.started:
            return []
        self._finalize_run(is_last=True)
        out = self.runs

        # reset
        self.started = False
        self.cur_cnt = 0
        self.runs = []
        self.bp_buf = []
        return out

    def _finalize_run(self, is_last: bool) -> None:
        if self.cur_cnt >= 8:
            self._flush_bp()
            self.runs.append(RLERun(self.cur_val, self.cur_cnt))
        elif is_last and not self.bp_buf:
            self.runs.append(RLERun(self.cur_val, self.cur_cnt))
        else:
            for _ in range(self.cur_cnt):
                self.bp_buf.append(self.cur_val)
                if len(self.bp_buf) == 8:
                    self._flush_bp()
            if is_last:
                self._flush_bp()

    def _flush_bp(self) -> None:
        if not self.bp_buf:
            return
        run = BPRun(self.bp_buf)
        self.bp_buf.clear()
        self.runs.append(run)


class Decoder:
    def decode(self, runs: list[Run]) -> list[int]:
        values: list[int] = []
        for run in runs:
            if isinstance(run, RLERun):
                values.extend([run.value] * run.count)
            elif isinstance(run, BPRun):
                values.extend(run.values)
        return values


def encode_via_stream(values: list[int]) -> list[Run]:
    enc = Encoder()
    for v in values:
        enc.append(v)
    return enc.finish()


def decode_via_stream(runs: list[Run]) -> list[int]:
    dec = Decoder()
    return dec.decode(runs)


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
    assert [r.encode() for r in s_enc1] == ["RLE[5, 8]", "BP[1, 2, 3]"], f"Got {s_enc1}"
    assert decode_via_stream(s_enc1) == inp1

    # Stream Example 2
    s_enc2 = encode_via_stream(inp2)
    assert [r.encode() for r in s_enc2] == ["RLE[1, 3]"], f"Got {s_enc2}"
    assert decode_via_stream(s_enc2) == inp2

    # Stream Example 3
    s_enc3 = encode_via_stream(inp3)
    assert [r.encode() for r in s_enc3] == [
        "BP[1, 1, 1, 1, 2, 3, 4, 5]"
    ], f"Got {s_enc3}"
    assert decode_via_stream(s_enc3) == inp3

    # Stream Single element
    s_enc4 = encode_via_stream(inp4)
    assert [r.encode() for r in s_enc4] == ["RLE[42, 1]"], f"Got {s_enc4}"
    assert decode_via_stream(s_enc4) == inp4

    # Stream Long mixed
    s_enc5 = encode_via_stream(inp5)
    assert [r.encode() for r in s_enc5] == [
        "RLE[7, 10]",
        "BP[1, 2, 3]",
        "RLE[2, 9]",
    ], f"Got {s_enc5}"
    assert decode_via_stream(s_enc5) == inp5

    # Stream Empty
    assert encode_via_stream([]) == []

    # Stream Incremental emission
    enc = Encoder()
    for _ in range(8):
        enc.append(5)
    enc.append(1)
    enc.append(2)
    enc.append(3)
    assert [r.encode() for r in enc.finish()] == ["RLE[5, 8]", "BP[1, 2, 3]"]

    # Stream BP flush at 8
    enc = Encoder()
    for _ in range(7):
        enc.append(1)
    enc.append(2)
    enc.append(3)
    assert [r.encode() for r in enc.finish()] == [
        "BP[1, 1, 1, 1, 1, 1, 1, 2]",
        "RLE[3, 1]",
    ]

    print("All tests passed!")

import random

# ---------------- Mock filesystem ----------------
filesystem = {}

def fs_read(name: str) -> str:
    return filesystem[name]

def fs_write(name: str, contents: str) -> None:
    filesystem[name] = contents

def fs_append(name: str, contents: str) -> None:
    filesystem[name] = filesystem.get(name, "") + contents


# ---------------- Game of Life: file-backed rows ----------------
def row_name(gen: int, r: int) -> str:
    return f"gen{gen}_row_{r}"

def read_row(gen: int, r: int, rows: int, cols: int) -> str:
    if r < 0 or r >= rows:
        return "0" * cols
    return fs_read(row_name(gen, r))

def write_row(gen: int, r: int, s: str) -> None:
    fs_write(row_name(gen, r), s)

def initialize(p: float, rows: int, cols: int, gen: int = 0, seed: int = 0) -> None:
    random.seed(seed)
    for r in range(rows):
        # '1' = alive, '0' = dead
        line = "".join("1" if random.random() <= p else "0" for _ in range(cols))
        write_row(gen, r, line)

def next_gen(rows: int, cols: int, in_gen: int, out_gen: int) -> None:
    for r in range(rows):
        prev = read_row(in_gen, r - 1, rows, cols)
        curr = read_row(in_gen, r, rows, cols)
        nxt  = read_row(in_gen, r + 1, rows, cols)

        # pad left/right to avoid per-neighbor col bounds checks
        p = "0" + prev + "0"
        c = "0" + curr + "0"
        n = "0" + nxt  + "0"

        out = []
        for j in range(1, cols + 1):
            live = (p[j-1] == "1") + (p[j] == "1") + (p[j+1] == "1") + \
                   (c[j-1] == "1")                 + (c[j+1] == "1") + \
                   (n[j-1] == "1") + (n[j] == "1") + (n[j+1] == "1")
            alive = (c[j] == "1")
            out.append("1" if (live == 3 or (live == 2 and alive)) else "0")

        write_row(out_gen, r, "".join(out))


# ---------------- Helpers for tests ----------------
def load_board(gen: int, rows: int, cols: int):
    return [fs_read(row_name(gen, r)) for r in range(rows)]

def set_board(gen: int, board):
    for r, line in enumerate(board):
        fs_write(row_name(gen, r), line)

def print_board(gen: int, rows: int):
    for r in range(rows):
        print(fs_read(row_name(gen, r)).replace("0", ".").replace("1", "#"))
    print()


# ---------------- Tests ----------------
def run_tests():
    # Test 1: Still life (2x2 block) should remain
    filesystem.clear()
    rows, cols = 4, 4
    b0 = [
        "0000",
        "0110",
        "0110",
        "0000",
    ]
    set_board(0, b0)
    next_gen(rows, cols, 0, 1)
    assert load_board(1, rows, cols) == b0

    # Test 2: Blinker oscillator (period 2)
    filesystem.clear()
    rows, cols = 5, 5
    b0 = [
        "00000",
        "00000",
        "01110",
        "00000",
        "00000",
    ]
    b1 = [
        "00000",
        "00100",
        "00100",
        "00100",
        "00000",
    ]
    set_board(0, b0)
    next_gen(rows, cols, 0, 1)
    assert load_board(1, rows, cols) == b1
    next_gen(rows, cols, 1, 2)
    assert load_board(2, rows, cols) == b0

    # Test 3: Edge behavior (no wrap). Single live cell dies.
    filesystem.clear()
    rows, cols = 3, 3
    b0 = [
        "100",
        "000",
        "000",
    ]
    b1 = [
        "000",
        "000",
        "000",
    ]
    set_board(0, b0)
    next_gen(rows, cols, 0, 1)
    assert load_board(1, rows, cols) == b1

    print("All tests passed.")

if __name__ == "__main__":
    run_tests()

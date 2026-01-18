# Game of Life at Scale

You are given a Python implementation of Conway's Game of Life that assumes
the entire grid fits in memory. The interviewer asks you to scale it to handle
very large boards, e.g., more than 1,000,000 x 1,000,000 cells, without OOM.

Your task is to redesign the algorithm to support external storage. In
particular:

- The board is too large to fit in memory.
- You must offload rows to a filesystem-like interface.
- A correct solution should read only a small window (e.g., 3 rows) at a time.

## Input

- `rows`, `cols`: dimensions of the grid.
- `p`: initialization probability of a cell being alive.

## Output

- A correct next-generation board stored in the filesystem interface.

## Follow Up

Design an efficient solution that:

- Uses O(cols) memory by streaming rows.
- Writes each row to disk as a file or as a fixed-length segment.
- Correctly computes neighbors using only the previous, current, and next rows.

## Provided Mock Filesystem API

```python
filesystem = {}

def read(name: str) -> str:
    return filesystem[name]

def write(name: str, contents: str) -> None:
    filesystem[name] = contents

# read/append mock for the seek-style file storage. Mocks the file content as a string.
def read(name: str, offset: int, length: int) -> str:
    return filesystem[name][offset:offset+length]

def append(name: str, contents: str) -> None:
    if name not in filesystem:
        filesystem[name] = ""
    filesystem[name] += contents
```

## Baseline Code (In-Memory)

```python
import random

def initialize(p: float, rows: int, cols: int) -> None:
    """
    Initialize a grid of size `rows` by `cols`.
    Each cell has `p` probability of being alive.
    """
    for i in range(rows):
        for j in range(cols):
            # Generate a random number between 0 and 1, and compare to p.
            if random.random() <= p:
                board[j] = True

def _process_cell(board: list[list[bool]], new_board: list[list[bool]], row: int, col: int) -> None:
    """Update a single cell (row, col) in new_board."""
    live_neighbors = 0
    for i in (row-1, row, row+1):
        for j in (col-1, col, col+1):
            # Bounds check
            if not (0 <= i < len(board) and 0 <= j < len(board[0])):
                continue

            # Don't count the cell itself
            if row == i and col == j:
                continue

            if board[j]:
                live_neighbors += 1

    if live_neighbors == 3:
        new_board[row][col] = True
    elif live_neighbors == 2:
        new_board[row][col] = board[row][col]
    else:
        new_board[row][col] = False

def next_gen(rows, cols) -> None:
    """Return the next generation."""
    new_board = [[False] * len(board[0]) for _ in range(len(board))]

    for i in range(len(board)):
        for j in range(len(board[j])):
            _process_cell(board, new_board, i, j)
    return new_board
```

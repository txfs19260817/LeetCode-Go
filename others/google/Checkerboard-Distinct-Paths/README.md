# Checkerboard Distinct Paths

Given a checkerboard with `rows` rows and `cols` columns, count the number of distinct paths from the bottom-left cell to the bottom-right cell.

From a cell `(row, col)`, you may move only to the next column using one of these steps:

- `(row + 1, col + 1)`
- `(row, col + 1)`
- `(row - 1, col + 1)`

Moves that leave the board are not allowed.

Use dynamic programming to compute the number of distinct valid paths.

# Tic-Tac-Toe II

Design a generalized Tic-Tac-Toe on an n × m board for 2 players. A player wins by placing k consecutive marks in a horizontal, vertical, or diagonal line.

## API

- `TicTacToe(n, m, k)` — Initializes an n × m board with win condition k in a row.
- `Move(row, col, player) → int` — Player places a mark at (row, col). Returns 0 if no winner yet, 1 if player 1 wins, or 2 if player 2 wins.

A move is always valid and on an empty cell. Once someone wins, no more moves are made.

## Constraints

- 1 ≤ n, m ≤ 10⁴
- 1 ≤ k ≤ 10⁴
- At most 10⁵ calls to `Move`

## Example

```
game = TicTacToe(4, 6, 4)
game.Move(0, 2, 1)  → 0
game.Move(0, 3, 2)  → 0
game.Move(1, 2, 1)  → 0
game.Move(1, 3, 2)  → 0
game.Move(2, 2, 1)  → 0
game.Move(2, 3, 2)  → 0
game.Move(3, 2, 1)  → 1   (player 1 has 4 in column 2)
```

## Approach

The board can be up to 10⁴ × 10⁴, so we cannot allocate a full grid. Instead, store placed marks in a hash map `(row, col) → player`. On each move, check 4 direction pairs (horizontal, vertical, two diagonals) centered on the placed cell and count consecutive same-player marks. If any direction reaches k, that player wins.

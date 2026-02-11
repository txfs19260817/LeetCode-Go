# ---------------------------------------------------------------------------
# Solution 1: O(k) per move — scan up to k cells in each direction
# ---------------------------------------------------------------------------

class TicTacToe:
    def __init__(self, n: int, m: int, k: int):
        self.n = n
        self.m = m
        self.k = k
        self.board: dict[tuple[int, int], int] = {}
        self.winner = 0

    # Direction pairs: horizontal, vertical, main diagonal, anti-diagonal
    _directions = [(0, 1), (1, 0), (1, 1), (1, -1)]

    def move(self, row: int, col: int, player: int) -> int:
        if self.winner != 0:
            return self.winner
        self.board[(row, col)] = player

        for dx, dy in self._directions:
            count = 1
            # Extend in the positive direction
            for i in range(1, self.k):
                nr, nc = row + dx * i, col + dy * i
                if nr < 0 or nr >= self.n or nc < 0 or nc >= self.m:
                    break
                if self.board.get((nr, nc)) != player:
                    break
                count += 1
            # Extend in the negative direction
            for i in range(1, self.k):
                nr, nc = row - dx * i, col - dy * i
                if nr < 0 or nr >= self.n or nc < 0 or nc >= self.m:
                    break
                if self.board.get((nr, nc)) != player:
                    break
                count += 1
            if count >= self.k:
                self.winner = player
                return player
        return 0


# ---------------------------------------------------------------------------
# Solution 2: O(1) per move — endpoint run-length counters
# ---------------------------------------------------------------------------
#
# run[(row, col, half_dir_idx)] = length of the maximal same-player run
#   starting at (row, col) and extending in that half-direction.
#
# Only the two ENDPOINTS of each run are kept current.  New placements
# can only land on empty cells (outside every existing run), so only
# endpoint values are ever queried.
#
# On each move, for each of 4 direction pairs:
#   a = backward run via neighbor   b = forward run via neighbor
#   total = a + 1 + b               (merged run through new cell)
#   → update far-back endpoint      → update far-forward endpoint

class TicTacToeO1:
    # 8 half-directions in opposite pairs (even index ↔ odd index).
    _half_dirs = [
        (0, 1), (0, -1),    # pair 0: right  / left
        (1, 0), (-1, 0),    # pair 1: down   / up
        (1, 1), (-1, -1),   # pair 2: ↘      / ↖
        (1, -1), (-1, 1),   # pair 3: ↙      / ↗
    ]

    def __init__(self, n: int, m: int, k: int):
        self.n = n
        self.m = m
        self.k = k
        self.board: dict[tuple[int, int], int] = {}
        self.run: dict[tuple[int, int, int], int] = {}
        self.winner = 0

    def move(self, row: int, col: int, player: int) -> int:
        if self.winner:
            return self.winner
        self.board[(row, col)] = player

        for pair in range(4):
            d, d_opp = pair * 2, pair * 2 + 1
            dx, dy = self._half_dirs[d]

            # a = backward run (neighbour in –d direction, extending further back)
            pr, pc = row - dx, col - dy
            a = self.run.get((pr, pc, d_opp), 0) if self.board.get((pr, pc)) == player else 0

            # b = forward run (neighbour in +d direction, extending further forward)
            sr, sc = row + dx, col + dy
            b = self.run.get((sr, sc, d), 0) if self.board.get((sr, sc)) == player else 0

            total = a + 1 + b

            # Update far endpoints of the merged run
            self.run[(row - dx * a, col - dy * a, d)] = total
            self.run[(row + dx * b, col + dy * b, d_opp)] = total

            if total >= self.k:
                self.winner = player
                return player

        return 0


def run_all_tests(cls):
    """Run the full test suite against the given TicTacToe class."""
    tag = cls.__name__

    # Test 1: Column win (main example)
    g = cls(4, 6, 4)
    assert g.move(0, 2, 1) == 0
    assert g.move(0, 3, 2) == 0
    assert g.move(1, 2, 1) == 0
    assert g.move(1, 3, 2) == 0
    assert g.move(2, 2, 1) == 0
    assert g.move(2, 3, 2) == 0
    assert g.move(3, 2, 1) == 1  # player 1 wins column 2

    # Test 2: Player 2 diagonal win
    g = cls(4, 4, 3)
    assert g.move(0, 1, 1) == 0
    assert g.move(0, 0, 2) == 0
    assert g.move(1, 0, 1) == 0
    assert g.move(1, 1, 2) == 0
    assert g.move(3, 3, 1) == 0
    assert g.move(2, 2, 2) == 2

    # Test 3: Horizontal win
    g = cls(3, 3, 3)
    assert g.move(0, 0, 1) == 0
    assert g.move(1, 0, 2) == 0
    assert g.move(0, 1, 1) == 0
    assert g.move(1, 1, 2) == 0
    assert g.move(0, 2, 1) == 1

    # Test 4: No winner until the last move
    g = cls(4, 4, 4)
    assert g.move(0, 0, 1) == 0
    assert g.move(1, 0, 2) == 0
    assert g.move(0, 1, 1) == 0
    assert g.move(1, 1, 2) == 0
    assert g.move(0, 2, 1) == 0
    assert g.move(1, 2, 2) == 0
    assert g.move(0, 3, 1) == 1

    # Test 5: k=1
    g = cls(3, 3, 1)
    assert g.move(1, 1, 1) == 1

    # Test 6: Anti-diagonal win
    g = cls(5, 5, 3)
    assert g.move(0, 0, 1) == 0
    assert g.move(0, 2, 2) == 0
    assert g.move(3, 3, 1) == 0
    assert g.move(1, 1, 2) == 0
    assert g.move(4, 4, 1) == 0
    assert g.move(2, 0, 2) == 2

    # Test 7: Large board
    g = cls(10000, 10000, 3)
    assert g.move(5000, 5000, 1) == 0
    assert g.move(0, 0, 2) == 0
    assert g.move(5001, 5000, 1) == 0
    assert g.move(0, 1, 2) == 0
    assert g.move(5002, 5000, 1) == 1

    # Test 8: Out-of-order placement (bridges a gap)
    g = cls(5, 5, 4)
    assert g.move(0, 0, 1) == 0
    assert g.move(1, 0, 2) == 0
    assert g.move(0, 3, 1) == 0
    assert g.move(1, 1, 2) == 0
    assert g.move(0, 1, 1) == 0
    assert g.move(1, 2, 2) == 0
    assert g.move(0, 2, 1) == 1  # bridges cols 0-1 and 3 → 4-in-a-row

    print(f"  {tag}: all tests passed")


if __name__ == "__main__":
    run_all_tests(TicTacToe)
    run_all_tests(TicTacToeO1)
    print("All tests passed!")

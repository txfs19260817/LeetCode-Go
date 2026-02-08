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


if __name__ == "__main__":
    # Test 1: Column win (main example)
    game = TicTacToe(4, 6, 4)
    assert game.move(0, 2, 1) == 0
    assert game.move(0, 3, 2) == 0
    assert game.move(1, 2, 1) == 0
    assert game.move(1, 3, 2) == 0
    assert game.move(2, 2, 1) == 0
    assert game.move(2, 3, 2) == 0
    assert game.move(3, 2, 1) == 1  # player 1 wins column 2

    # Test 2: Player 2 diagonal win
    game2 = TicTacToe(4, 4, 3)
    assert game2.move(0, 1, 1) == 0
    assert game2.move(0, 0, 2) == 0
    assert game2.move(1, 0, 1) == 0
    assert game2.move(1, 1, 2) == 0
    assert game2.move(3, 3, 1) == 0
    assert game2.move(2, 2, 2) == 2  # player 2 wins diagonal

    # Test 3: Horizontal win
    game3 = TicTacToe(3, 3, 3)
    assert game3.move(0, 0, 1) == 0
    assert game3.move(1, 0, 2) == 0
    assert game3.move(0, 1, 1) == 0
    assert game3.move(1, 1, 2) == 0
    assert game3.move(0, 2, 1) == 1  # player 1 wins row 0

    # Test 4: No winner until the last move
    game4 = TicTacToe(4, 4, 4)
    assert game4.move(0, 0, 1) == 0
    assert game4.move(1, 0, 2) == 0
    assert game4.move(0, 1, 1) == 0
    assert game4.move(1, 1, 2) == 0
    assert game4.move(0, 2, 1) == 0
    assert game4.move(1, 2, 2) == 0
    assert game4.move(0, 3, 1) == 1  # player 1 wins with 4 in row 0

    # Test 5: k=1 → first move always wins
    game5 = TicTacToe(3, 3, 1)
    assert game5.move(1, 1, 1) == 1  # immediately wins

    # Test 6: Anti-diagonal win
    game6 = TicTacToe(5, 5, 3)
    assert game6.move(0, 0, 1) == 0
    assert game6.move(0, 2, 2) == 0
    assert game6.move(3, 3, 1) == 0
    assert game6.move(1, 1, 2) == 0
    assert game6.move(4, 4, 1) == 0
    assert game6.move(2, 0, 2) == 2  # player 2 wins anti-diagonal

    # Test 7: Large board
    game7 = TicTacToe(10000, 10000, 3)
    assert game7.move(5000, 5000, 1) == 0
    assert game7.move(0, 0, 2) == 0
    assert game7.move(5001, 5000, 1) == 0
    assert game7.move(0, 1, 2) == 0
    assert game7.move(5002, 5000, 1) == 1  # player 1 wins vertical

    print("All tests passed!")

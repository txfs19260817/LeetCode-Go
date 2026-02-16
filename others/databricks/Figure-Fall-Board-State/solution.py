class Solution:
    """
    Time: O(R * C + F)
    ◦ R*C for scanning the board by column/row
    ◦ F for clearing and writing figure cells (F = number of 'F' cells)
    ◦ Since F <= R*C, overall is O(R*C).
    Space: O(F) extra
    ◦ due to positions list storing all figure coordinates.
    """

    def dropFigure(self, board: list[list[str]]) -> list[list[str]]:
        rows, cols = len(board), len(board[0])
        # Global lower bound on how far the whole figure can drop.
        min_move = rows
        # Save original figure cells so we can move them after min_move is known.
        positions: list[tuple[int, int]] = []

        for c in range(cols):
            # Track the lowest 'F' seen so far in this column.
            last_f = -1
            for r in range(rows):
                if board[r][c] == "F":
                    last_f = r
                    positions.append((r, c))
                elif board[r][c] == "#" and last_f != -1:
                    # Gap between the nearest upper F and obstacle in this column.
                    min_move = min(min_move, r - last_f - 1)

            # If no obstacle below, bottom boundary constrains movement.
            min_move = min(min_move, rows - last_f - 1)

        # Move in two phases to avoid overwrite when figure cells share columns.
        for r, c in positions:
            board[r][c] = "."
        for r, c in positions:
            board[r + min_move][c] = "F"

        return board


if __name__ == "__main__":
    sol = Solution()

    b1 = [
        list("...."),
        list(".F.."),
        list(".F.."),
        list("...."),
    ]
    e1 = [
        list("...."),
        list("...."),
        list(".F.."),
        list(".F.."),
    ]
    assert sol.dropFigure(b1) == e1

    b2 = [
        list("..F.."),
        list(".FFF."),
        list("....."),
        list("..#.."),
        list("....."),
    ]
    e2 = [
        list("....."),
        list("..F.."),
        list(".FFF."),
        list("..#.."),
        list("....."),
    ]
    assert sol.dropFigure(b2) == e2

    b3 = [list(".F."), list(".#."), list("...")]
    e3 = [list(".F."), list(".#."), list("...")]
    assert sol.dropFigure(b3) == e3

    b4 = [list(".#."), list("...")]
    e4 = [list(".#."), list("...")]
    assert sol.dropFigure(b4) == e4

    # Single-cell figure drops to bottom.
    b5 = [list("F.."), list("..."), list("...")]
    e5 = [list("..."), list("..."), list("F..")]
    assert sol.dropFigure(b5) == e5

    # Figure already touching bottom cannot move.
    b6 = [list("..."), list(".F."), list(".F.")]
    e6 = [list("..."), list(".F."), list(".F.")]
    assert sol.dropFigure(b6) == e6

    # Different columns have different free space; global move uses the minimum.
    b7 = [
        list(".FF..."),
        list(".FF..."),
        list("......"),
        list("...#.."),
        list("......"),
        list("......"),
    ]
    e7 = [
        list("......"),
        list("......"),
        list("......"),
        list("...#.."),
        list(".FF..."),
        list(".FF..."),
    ]
    assert sol.dropFigure(b7) == e7

    # Obstacle directly under one column blocks movement entirely.
    b8 = [
        list(".FF."),
        list(".#.."),
        list("...."),
    ]
    e8 = [
        list(".FF."),
        list(".#.."),
        list("...."),
    ]
    assert sol.dropFigure(b8) == e8

    print("All tests passed!")

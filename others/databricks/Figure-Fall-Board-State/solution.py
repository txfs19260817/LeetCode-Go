from __future__ import annotations

from typing import List


class Solution:
    def dropFigure(self, board: List[List[str]]) -> List[List[str]]:
        if not board or not board[0]:
            return board

        rows, cols = len(board), len(board[0])
        min_move = rows
        has_figure = False
        positions: list[tuple[int, int]] = []

        for c in range(cols):
            last_f = -1
            for r in range(rows):
                if board[r][c] == "F":
                    has_figure = True
                    last_f = r
                    positions.append((r, c))
                elif board[r][c] == "#" and last_f != -1:
                    min_move = min(min_move, r - last_f - 1)
            if last_f != -1:
                min_move = min(min_move, rows - last_f - 1)

        if not has_figure or min_move <= 0:
            return board

        for r, c in positions:
            board[r][c] = "."
        for r, c in positions:
            board[r + min_move][c] = "F"

        return board


if __name__ == "__main__":
    sol = Solution()

    b1 = [list("...."), list(".F.."), list(".F.."), list("....")]
    e1 = [list("...."), list("...."), list(".F.."), list(".F..")] 
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

    print("All tests passed!")
package k

import "math"

/*
We have a two-dimensional board game involving snakes.
The board has two types of squares on it:
+'s represent impassable squares where snakes cannot go, and 0's represent squares through which snakes can move.
Snakes can only enter on the edges of the board, and each snake can move in only one direction.
We'd like to find the places where a snake can pass through the entire board, moving in a straight line.

Here is an example board:
    col-->       0  1  2  3  4  5  6
    +---------------------
    row     0 |  +  +  +  0  +  0  0
    |       1 |  0  0  +  0  0  0  0
    |       2 |  0  0  0  0  +  0  0
    v       3 |  +  +  +  0  0  +  0
            4 |  0  0  0  0  0  0  0
Write a function that takes a rectangular board with only +'s and 0's, and returns two collections:
* one containing all of the row numbers whose row is completely passable by snakes, and
* the other containing all of the column numbers where the column is completely passable
    Sample Inputs:
    board1 = [['+', '+', '+', '0', '+', '0', '0'],
    ['0', '0', '+', '0', '0', '0', '0'],
    ['0', '0', '0', '0', '+', '0', '0'],
    ['+', '+', '+', '0', '0', '+', '0'],
    ['0', '0', '0', '0', '0', '0', '0']]
    board2 = [['+', '+', '+', '0', '+', '0', '0'],
    ['0', '0', '0', '0', '0', '+', '0'],
    ['0', '0', '+', '0', '0', '0', '0'],
    ['0', '0', '0', '0', '+', '0', '0'],
    ['+', '+', '+', '0', '0', '0', '+']]
    board3 = [['+', '+', '+', '0', '+', '0', '0'],
    ['0', '0', '0', '0', '0', '0', '0'],
    ['0', '0', '+', '+', '0', '+', '0'],
    ['0', '0', '0', '0', '+', '0', '0'],
    ['+', '+', '+', '0', '0', '0', '+']]
    board4 = [['+']]
    board5 = [['0']]
    All test cases:
    findPassableLanes(board1) => Rows: [4], Columns: [3, 6]
    findPassableLanes(board2) => Rows: [], Columns: [3]
    findPassableLanes(board3) => Rows: [1], Columns: []
    findPassableLanes(board4) => Rows: [], Columns: []
    findPassableLanes(board5) => Rows: [0], Columns: [0]
    Complexity Analysis:
    r: number of rows in the board
    c: number of columns in the board
*/

func findPassableLanes(board [][]byte) (rows []int, cols []int) {
	m, n := len(board), len(board[0])
	for i := range m {
		valid := true
		for j := range n {
			if board[i][j] == '+' {
				valid = false
				break
			}
		}
		if valid {
			rows = append(rows, i)
		}
	}
	for j := range n {
		valid := true
		for i := range m {
			if board[i][j] == '+' {
				valid = false
				break
			}
		}
		if valid {
			cols = append(cols, j)
		}
	}
	return
}

/*
We have a two-dimensional board game involving snakes. The board has two types of squares on it: +’s represent impassable squares where snakes cannot go, and 0’s represent squares through which snakes can move.
Snakes may move in any of four directions - up, down, left, or right - one square at a time,
but they will never return to a square that they’ve already visited.

If a snake enters the board on an edge square, we want to catch it at a different exit square on the board’s edge.
The snake is familiar with the board and will take the route to the nearest reachable exit,
in terms of the number of squares it has to move through to get there.

Write a function that takes a rectangular board with only +’s and 0’s,
along with a starting point on the edge of the board (given row first, then column),
and returns the coordinates of the nearest exit to which it can travel.
If multiple exits are equally close, give the one with the lowest numerical value for the row.
If there is still a tie, give the one of those with the lowest numerical value for the column.
If there is no answer, output -1 -1
The board will be non-empty and rectangular.
All values in the board will be either + or 0.
All coordinates (input and output) are zero-based.
All start positions will be 0, and be on the edge of the board. For example, (0,0) would be the top left corner of any size input.

 board1 = [['+', '+', '+', '+', '+', '+', '+', '0', '0'],
 ['+', '+', '0', '0', '0', '0', '0', '+', '+'],
 ['0', '0', '0', '0', '0', '+', '+', '0', '+'],
 ['+', '+', '0', '+', '+', '+', '+', '0', '0'],
 ['+', '+', '0', '0', '0', '0', '0', '0', '+'],
 ['+', '+', '0', '+', '+', '0', '+', '0', '+']]
 start1_1 = (2, 0) # Expected output = (5, 2)
 start1_2 = (0, 7) # Expected output = (0, 8)
 start1_3 = (5, 2) # Expected output = (2, 0) or (5, 5)
 start1_4 = (5, 5) # Expected output = (5, 7)
 board2 = [['+', '+', '+', '+', '+', '+', '+'],
 ['0', '0', '0', '0', '+', '0', '+'],
 ['+', '0', '+', '0', '+', '0', '0'],
 ['+', '0', '0', '0', '+', '+', '+'],
 ['+', '+', '+', '+', '+', '+', '+']]
 start2_1 = (1, 0) # Expected output = null (or a special value
 representing no possible exit)
 start2_2 = (2, 6) # Expected output = null
 board3 = [['+', '0', '+', '0', '+',],
 ['0', '0', '+', '0', '0',],
 ['+', '0', '+', '0', '+',],
 ['0', '0', '+', '0', '0',],
 ['+', '0', '+', '0', '+']]
 start3_1 = (0, 1) # Expected output = (1, 0)
 start3_2 = (4, 1) # Expected output = (3, 0)
 start3_3 = (0, 3) # Expected output = (1, 4)
 start3_4 = (4, 3) # Expected output = (3, 4)
 board4 = [['+', '0', '+', '0', '+',],
 ['0', '0', '0', '0', '0',],
 ['+', '+', '+', '+', '+',],
 ['0', '0', '0', '0', '0',],
 ['+', '0', '+', '0', '+']]
 start4_1 = (1, 0) # Expected output = (0, 1)
 start4_2 = (1, 4) # Expected output = (0, 3)
 start4_3 = (3, 0) # Expected output = (4, 1)
 start4_4 = (3, 4) # Expected output = (4, 3)
 board5 = [['+', '0', '0', '0', '+',],
 ['+', '0', '+', '0', '+',],
 ['+', '0', '0', '0', '+',],
 ['+', '0', '+', '0', '+']]
 start5_1 = (0, 1) # Expected output = (0, 2)
 start5_2 = (3, 1) # Expected output = (0, 1)

 All test cases:
 findExit(board1, start1_1) => (5, 2)
 findExit(board1, start1_2) => (0, 8)
 findExit(board1, start1_3) => (2, 0) or (5, 5)
 findExit(board1, start1_4) => (5, 7)
 findExit(board2, start2_1) => null (or a special value representing no possible exit)
 findExit(board2, start2_2) => null
 findExit(board3, start3_1) => (1, 0)
 findExit(board3, start3_2) => (3, 0)
 findExit(board3, start3_3) => (1, 4)
 findExit(board3, start3_4) => (3, 4)
 findExit(board4, start4_1) => (0, 1)
 findExit(board4, start4_2) => (0, 3)
 findExit(board4, start4_3) => (4, 1)
 findExit(board4, start4_4) => (4, 3)
 findExit(board5, start5_1) => (0, 2)
 findExit(board5, start5_2) => (0, 1)
*/

func findExit(board [][]byte, start [2]int) (end [2]int) {
	end = [2]int{-1, -1}
	if board[start[0]][start[1]] == '+' {
		return end
	}
	dirs := [4][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	m, n := len(board), len(board[0])
	isEdge := func(i, j int) bool { return i == 0 || i == m-1 || j == 0 || j == n-1 }
	bestStep := math.MaxInt
	var dfs func(i, j, step int)
	dfs = func(i, j, step int) {
		if isEdge(i, j) && [2]int{i, j} != start && bestStep > step {
			bestStep = step
			end = [2]int{i, j}
			return
		}

		board[i][j] = '+'
		for _, dir := range dirs {
			ni, nj := i+dir[0], j+dir[1]
			if ni >= 0 && nj >= 0 && ni < m && nj < n && board[ni][nj] == '0' {
				dfs(ni, nj, step+1)
			}
		}
		board[i][j] = '0'
	}
	dfs(start[0], start[1], 0)
	return
}

/*
Snakes make nests of open spaces, but they avoid connecting their nest with other snakes' nests.
Each nest (consisting of one or more connected 0's) houses one snake only,
though a single nest may have one or many entrances on the edges of the board, or possibly none at all.
Each 0 for a nest that is on an edge counts as its own entrance, even if two or more of them are next to one another.
Here is an example board:
    col-->        0  1  2  3  4  5  6  7  8
               +---------------------------
    row      0 |  +  +  +  +  +  +  +  0  0
     |       1 |  +  +  0  0  0  0  0  +  +
     |       2 |  0  0  0  0  0  +  +  0  +
     v       3 |  +  +  0  +  +  +  +  0  0
             4 |  +  +  0  0  0  0  0  0  +
             5 |  +  +  0  +  +  0  +  0  +
Given a board, return a collection that, for each nest, lists the nest's number of entrances (open spaces on the edge of the board) that that nest includes.
Sample inputs:
board1 = [['+', '+', '+', '+', '+', '+', '+', '0', '0'],
          ['+', '+', '0', '0', '0', '0', '0', '+', '+'],
          ['0', '0', '0', '0', '0', '+', '+', '0', '+'],
          ['+', '+', '0', '+', '+', '+', '+', '0', '0'],
          ['+', '+', '0', '0', '0', '0', '0', '0', '+'],
          ['+', '+', '0', '+', '+', '0', '+', '0', '+']]
board2 = [['+', '+', '+', '+', '+', '+'],
          ['0', '0', '0', '+', '0', '+'],
          ['+', '0', '+', '0', '0', '0'],
          ['+', '+', '+', '+', '+', '+']]
board3 = [['+', '0', '+', '+', '+', '0', '+', '0', '0'],
          ['+', '0', '+', '0', '0', '0', '0', '+', '+'],
          ['0', '0', '0', '0', '0', '+', '+', '0', '+'],
          ['+', '+', '+', '+', '+', '+', '+', '0', '0'],
          ['+', '+', '0', '0', '0', '0', '0', '0', '+'],
          ['+', '+', '0', '+', '+', '0', '+', '0', '+']]
board4 = [['+', '+', '+'],
          ['+', '0', '+'],
          ['+', '+', '+']]
board5 = [['+']]
board6 = [['0', '0'],
          ['0', '0']]
All test cases:
getNestEntranceCount(board1) => [2, 5]
getNestEntranceCount(board2) => [1, 1]
getNestEntranceCount(board3) => [2, 4, 3]
getNestEntranceCount(board4) => [0]
getNestEntranceCount(board5) => []
getNestEntranceCount(board6) => [4]
Complexity Analysis:
r: number of rows in the board
c: number of columns in the board
*/

func getNestEntranceCount(board [][]byte) (ans []int) {
	dirs := [4][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	m, n := len(board), len(board[0])
	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}
	isEdge := func(i, j int) bool { return i == 0 || i == m-1 || j == 0 || j == n-1 }

	for i := range m {
		for j := range n {
			if board[i][j] != '0' || visited[i][j] {
				continue
			}
			eCnt := 0
			visited[i][j] = true
			if isEdge(i, j) {
				eCnt++
			}
			q := [][2]int{{i, j}}
			for len(q) > 0 {
				cur := q[0]
				q = q[1:]

				for _, dir := range dirs {
					ni, nj := cur[0]+dir[0], cur[1]+dir[1]
					if ni >= 0 && nj >= 0 && ni < m && nj < n && board[ni][nj] == '0' && !visited[ni][nj] {
						visited[ni][nj] = true
						if isEdge(ni, nj) {
							eCnt++
						}
						q = append(q, [2]int{ni, nj})
					}
				}
			}

			ans = append(ans, eCnt)
		}
	}
	return
}

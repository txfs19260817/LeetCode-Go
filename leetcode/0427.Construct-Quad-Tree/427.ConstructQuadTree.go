package leetcode

type Node struct {
	Val         bool
	IsLeaf      bool
	TopLeft     *Node
	TopRight    *Node
	BottomLeft  *Node
	BottomRight *Node
}

func construct(grid [][]int) *Node {
	n := len(grid)

	// 2D prefix sum array to calculate the sum of the submatrix in O(1) time with O(n^2) space
	pre := make([][]int, n+1)
	for i := range pre {
		pre[i] = make([]int, n+1)
	}
	for i, row := range grid {
		for j, v := range row {
			pre[i+1][j+1] = pre[i][j+1] + pre[i+1][j] - pre[i][j] + v
		}
	}

	var dfs func(r0, c0, r1, c1 int) *Node
	dfs = func(r0, c0, r1, c1 int) *Node {
		sum := pre[r1][c1] - pre[r0][c1] - pre[r1][c0] + pre[r0][c0]
		if sum == 0 { // all 0
			return &Node{Val: false, IsLeaf: true}
		}
		if sum == (r1-r0)*(c1-c0) { // all 1
			return &Node{Val: true, IsLeaf: true}
		}
		rMid, cMid := (r0+r1)/2, (c0+c1)/2

		//      c0          cMid        c1
		// r0   +-----------+-----------+
		//      |  TopLeft  |  TopRight |
		// rMid +-----------+-----------+
		//      | BottomLeft|BottomRight|
		// r1   +-----------+-----------+

		// you can assign the val to True or False when isLeaf is False, and both are accepted in the answer
		return &Node{
			TopLeft:     dfs(r0, c0, rMid, cMid),
			TopRight:    dfs(r0, cMid, rMid, c1),
			BottomLeft:  dfs(rMid, c0, r1, cMid),
			BottomRight: dfs(rMid, cMid, r1, c1),
		}
	}

	return dfs(0, 0, n, n)
}

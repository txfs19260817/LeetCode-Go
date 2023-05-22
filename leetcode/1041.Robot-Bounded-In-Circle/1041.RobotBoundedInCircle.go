package leetcode

func isRobotBounded(instructions string) bool {
	dir, d, cur := [4][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}, 0, [2]int{0, 0}
	for _, i := range instructions {
		switch i {
		case 'G':
			cur[0] += dir[d][0]
			cur[1] += dir[d][1]
		case 'L':
			d++
			if d > 3 {
				d = 0
			}
		case 'R':
			d--
			if d < 0 {
				d = 3
			}
		}
	}
	return cur == [2]int{0, 0} || d != 0 //The robot stays in the circle iff (looking at the final vector) it changes direction (ie. doesn't stay pointing north), or it moves 0.
}

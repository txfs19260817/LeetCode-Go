package leetcode

func maxPoints(points [][]int) int {
	if len(points) <= 2 {
		return len(points)
	}
	var ans int
	for i, p := range points {
		if ans >= len(points)-i || ans > len(points)/2 {
			break
		}
		cnt := map[int]int{} // slope: cnt
		for _, q := range points[i+1:] {
			x, y := p[0]-q[0], p[1]-q[1]
			if x == 0 {
				y = 1
			} else if y == 0 {
				x = 1
			} else {
				if y < 0 { // only allow x<0
					x, y = -x, -y
				}
				g := gcd(abs(x), abs(y))
				x, y = x/g, y/g

			}
			cnt[x+20001*y]++
		}
		for _, c := range cnt {
			if count := c + 1; ans < count {
				ans = count
			}
		}
	}
	return ans
}

func abs(a int) int {
	if a < 0 {
		return -1 * a
	}
	return a
}

func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}

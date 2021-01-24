package _011_Container_With_Most_Water

func maxArea(height []int) int {
	ans := 0
	for l, r := 0, len(height)-1; l < r; {
		ans = max(ans, (r-l)*min(height[l], height[r]))
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

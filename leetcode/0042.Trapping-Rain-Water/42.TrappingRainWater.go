package leetcode

func trap(height []int) int {
	var res, mLeftHeight, mRightHeight int
	for l, r := 0, len(height)-1; l <= r; {
		if height[l] <= height[r] {
			if mLeftHeight < height[l] {
				mLeftHeight = height[l]
			} else {
				res += mLeftHeight - height[l]
			}
			l++
		} else {
			if mRightHeight < height[r] {
				mRightHeight = height[r]
			} else {
				res += mRightHeight - height[r]
			}
			r--
		}
	}
	return res
}

package leetcode

import "math"

func constructRectangle(area int) []int {
	ans := []int{area, 1}
	for i := int(math.Sqrt(float64(area))) + 1; i > 0; i-- {
		if area%i == 0 {
			ans[0], ans[1] = area/i, i
			if ans[0] < ans[1] {
				ans[0], ans[1] = ans[1], ans[0]
			}
			break
		}
	}
	return ans
}

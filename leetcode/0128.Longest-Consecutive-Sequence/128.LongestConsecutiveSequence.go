package leetcode

func longestConsecutive(nums []int) int {
	ans, m := 0, map[int]int{} // num:length
	for _, num := range nums {
		if m[num] == 0 {
			left, right := m[num-1], m[num+1]
			sum := 1 + left + right
			if sum > ans {
				ans = sum
			}
			m[num], m[num-left], m[num+right] = sum, sum, sum
		}
	}
	return ans
}

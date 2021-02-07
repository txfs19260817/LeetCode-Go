package _695_Maximum_Erasure_Value

func maximumUniqueSubarray(nums []int) int {
	l, r, freq, ans := 0, 0, map[int]int{}, 0
	for l < len(nums) {
		if r < len(nums) && freq[nums[r]] == 0 {
			freq[nums[r]]++
			r++
		} else {
			freq[nums[l]]--
			l++
		}
		if s := arraySum(nums, l, r); s > ans {
			ans = s
		}
	}
	return ans
}

func arraySum(nums []int, l, r int) int {
	var s int
	for i := l; i < r; i++ {
		s += nums[i]
	}
	return s
}

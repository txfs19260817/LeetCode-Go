package leetcode

func countSubarrays(nums []int, k int64) int64 {
	var ans, sum int64
	for l, r := 0, 0; r < len(nums); r++ {
		sum += int64(nums[r])
		// move the left side of the window, decreasing sum, to limit score <= k.
		for sum*int64(r-l+1) >= k {
			sum -= int64(nums[l])
			l++
		}
		ans += int64(r - l + 1) // tracking the sum of the subarray in the window
	}
	return ans
}

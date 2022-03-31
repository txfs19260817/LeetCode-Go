package _410_Split_Array_Largest_Sum

func splitArray(nums []int, m int) int {
	l, r := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > l {
			l = nums[i]
		}
		r += nums[i]
	}
	for l < r {
		mid := l + (r-l)/2
		if check(nums, m, mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return l
}

func check(nums []int, m, target int) bool {
	sum, cnt := 0, 1
	for _, num := range nums {
		if sum+num > target {
			sum = num
			cnt++
		} else {
			sum += num
		}
	}
	return cnt <= m
}

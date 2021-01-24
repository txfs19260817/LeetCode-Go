package _220_Contains_Duplicate_III

import "sort"

// https://bitbrave.github.io/2020/08/22/LeetCode(220.%20Contains%20Duplicate%20III)%E9%A2%98%E8%A7%A3/
func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	if len(nums) < 2 {
		return false
	}
	buckets := map[int]int{} // num/(t+1):num
	getBucket := func(num, t int) int {
		if num == 0 || t == 0 {
			return num
		}
		if num > 0 {
			return num / (t + 1)
		}
		return num/t - 1
	}
	for i, num := range nums {
		if i > k {
			delete(buckets, getBucket(nums[i-k-1], t))
		}
		key := getBucket(num, t)
		if _, ok := buckets[key]; ok {
			return true
		}
		if val, ok := buckets[key-1]; ok && abs(val-num) <= t {
			return true
		}
		if val, ok := buckets[key+1]; ok && abs(val-num) <= t {
			return true
		}
		buckets[key] = num
	}
	return false
}

func containsNearbyAlmostDuplicate1(nums []int, k int, t int) bool {
	if len(nums) < 2 {
		return false
	}
	copyNums := append([]int{}, nums...)
	sortInterval := k + 1
	if k+1 > len(nums) {
		sortInterval = len(nums)
	}
	sort.Ints(nums[:sortInterval])
	for i := 1; i <= k && i < len(nums); i++ {
		if abs(nums[i]-nums[i-1]) <= t {
			return true
		}
	}
	nums = copyNums
	for i := k + 1; i < len(nums); i++ {
		for j := i - 1; j > i-k-1; j-- {
			if abs(nums[i]-nums[j]) <= t {
				return true
			}
		}
	}
	return false
}

func abs(a int) int {
	if a < 0 {
		return -1 * a
	}
	return a
}

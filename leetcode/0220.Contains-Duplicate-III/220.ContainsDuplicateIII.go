package leetcode

func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	buckets := map[int]int{} // element:bucket, bucket size = t+1
	getID := func(x int) int {
		if x >= 0 {
			return x / (t + 1)
		}
		return (x+1)/(t+1) - 1
	}
	for i, x := range nums {
		id := getID(x)
		if _, ok := buckets[id]; ok {
			return true
		}
		if y, ok := buckets[id-1]; ok && abs(y-x) <= t {
			return true
		}
		if y, ok := buckets[id+1]; ok && abs(y-x) <= t {
			return true
		}
		buckets[id] = x
		if i >= k {
			delete(buckets, getID(nums[i-k]))
		}
	}
	return false
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

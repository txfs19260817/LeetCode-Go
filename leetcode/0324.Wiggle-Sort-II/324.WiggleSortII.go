package leetcode

func wiggleSort(nums []int) {
	if len(nums) < 2 {
		return
	}
	qs(nums, 0, len(nums)-1)
	res := make([]int, len(nums))
	mid := len(nums) / 2
	for i, j := 0, mid; j < len(nums); i, j = i+2, j+1 {
		res[i] = nums[j]
	}
	for i, j := 1, 0; j < mid; i, j = i+2, j+1 {
		res[i] = nums[j]
	}
	for i, r := range res {
		nums[i] = r
	}
}

func qs(nums []int, l, r int) {
	if l < r {
		pivot := partition(nums, l, r)
		qs(nums, l, pivot-1)
		qs(nums, pivot+1, r)
	}
}

func partition(nums []int, l, r int) int {
	pivotElem := nums[r]
	i := l
	for j := l; j < r; j++ {
		if nums[j] > pivotElem {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
	}
	nums[i], nums[r] = nums[r], nums[i]
	return i
}

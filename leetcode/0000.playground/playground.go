package _000_playground

/*
A playground to code benefiting from auto-completion within IDE.
*/
func findKthLargest(nums []int, k int) int {
	return qs(nums, 0, len(nums)-1, k)
}

func qs(nums []int, l, r, k int) int {
	if l == r {
		return nums[l]
	}
	pivot := partition(nums, l, r)
	if pivot == k {
		return nums[k]
	}
	if pivot > k {
		return qs(nums, l, pivot-1, k)
	}
	return qs(nums, pivot+1, r, k)
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

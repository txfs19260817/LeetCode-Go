package _031_Next_Permutation

func nextPermutation(nums []int) {
	i, j, k := len(nums)-2, len(nums)-1, len(nums)-1
	for i >= 0 && nums[i] >= nums[j] { // find the first ascending pair in inverted order
		i--
		j--
	}
	if i >= 0 { // if nums is not sorted in descending order
		for nums[i] >= nums[k] { // find the first k that nums[k] > nums[i] in inverted order
			k--
		}
		nums[i], nums[k] = nums[k], nums[i] // swap them
	}
	for i, j := j, len(nums)-1; i < j; i, j = i+1, j-1 { // reverse the remain descending part
		nums[i], nums[j] = nums[j], nums[i]
	}
}

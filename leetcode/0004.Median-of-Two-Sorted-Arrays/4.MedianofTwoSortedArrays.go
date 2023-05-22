package leetcode

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	totalLen := len(nums1) + len(nums2)
	if totalLen%2 == 1 {
		return float64(getKthElement(nums1, nums2, totalLen/2+1))
	}
	return float64(getKthElement(nums1, nums2, totalLen/2)+getKthElement(nums1, nums2, totalLen/2+1)) / 2
}

func getKthElement(nums1 []int, nums2 []int, k int) int { // k is a number/1-indexed
	for idx1, idx2 := 0, 0; ; {
		if idx1 == len(nums1) {
			return nums2[idx2+k-1]
		}
		if idx2 == len(nums2) {
			return nums1[idx1+k-1]
		}
		if k == 1 {
			return min(nums1[idx1], nums2[idx2])
		}
		newIdx1, newIdx2 := min(idx1+k/2, len(nums1))-1, min(idx2+k/2, len(nums2))-1
		pivot1, pivot2 := nums1[newIdx1], nums2[newIdx2]
		if pivot1 <= pivot2 {
			k -= newIdx1 + 1 - idx1
			idx1 = newIdx1 + 1
		} else {
			k -= newIdx2 + 1 - idx2
			idx2 = newIdx2 + 1
		}
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

package _088_Merge_Sorted_Array

func merge(nums1 []int, m int, nums2 []int, n int) {
	i, j, k := m-1, n-1, m+n-1
	for ; i >= 0 && j >= 0; k-- {
		if nums1[i] > nums2[j] {
			nums1[k] = nums1[i]
			i--
		} else {
			nums1[k] = nums2[j]
			j--
		}
	}
	for ; j >= 0; k-- {
		nums1[k] = nums2[j]
		j--
	}
}

func merge1(nums1 []int, m int, nums2 []int, n int) {
	for i, j := 0, 0; j < n; i, j, m = i+1, j+1, m+1 {
		for nums1[i] <= nums2[j] && i < m {
			i++
		}
		for k := m; k > i; k-- {
			nums1[k] = nums1[k-1]
		}
		nums1[i] = nums2[j]
	}
}

func mergeForTest(nums1 []int, m int, nums2 []int, n int) []int {
	i, j, k := m-1, n-1, m+n-1
	for ; i >= 0 && j >= 0; k-- {
		if nums1[i] > nums2[j] {
			nums1[k] = nums1[i]
			i--
		} else {
			nums1[k] = nums2[j]
			j--
		}
	}
	for ; j >= 0; k-- {
		nums1[k] = nums2[j]
		j--
	}
	return nums1
}

func mergeForTest1(nums1 []int, m int, nums2 []int, n int) []int {
	for i, j := 0, 0; j < n; i, j, m = i+1, j+1, m+1 {
		for nums1[i] <= nums2[j] && i < m {
			i++
		}
		for k := m; k > i; k-- {
			nums1[k] = nums1[k-1]
		}
		nums1[i] = nums2[j]
	}
	return nums1
}

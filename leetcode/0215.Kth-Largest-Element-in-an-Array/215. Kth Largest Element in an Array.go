package _215_Kth_Largest_Element_in_an_Array

func findKthLargest(nums []int, k int) int {
	if len(nums) == 0 {
		return 0
	}
	return topK(nums, 0, len(nums)-1, len(nums)-k)
}

func topK(data []int, left, right, k int) int {
	if left == right {
		return data[left]
	}
	pivot := partition(data, left, right)
	if pivot == k {
		return data[pivot]
	}
	if pivot > k {
		return topK(data, left, pivot-1, k)
	}
	return topK(data, pivot+1, right, k)
}

func partition(data []int, left, right int) int {
	pivotElem := data[right]
	i := left - 1
	for j := left; j < right; j++ {
		if data[j] < pivotElem {
			i++
			data[i], data[j] = data[j], data[i]
		}
	}
	i++
	data[i], data[right] = data[right], data[i]
	return i
}

package leetcode

func maximumGap(nums []int) int {
	if len(nums) < 2 {
		return 0
	}
	nums = radixSort(nums)
	res := 0
	for i := 1; i < len(nums); i++ {
		res = max(res, nums[i]-nums[i-1])
	}
	return res
}

func radixSort(data []int) []int {
	// find the largest number to obtain the max units.
	maxNum := data[0]
	for i := 1; i < len(data); i++ {
		maxNum = max(maxNum, data[i])
	}

	for exp := 1; maxNum/exp > 0; exp *= 10 {
		count := make([]int, 10)
		// count by radix
		for _, d := range data {
			count[(d/exp)%10]++
		}

		// convert counts to positions
		for i := 1; i < 10; i++ {
			count[i] += count[i-1]
		}

		//place elements
		temp := make([]int, len(data))
		for i := len(data) - 1; i >= 0; i-- {
			count[(data[i]/exp)%10]--
			temp[count[(data[i]/exp)%10]] = data[i]
		}

		// copy back
		for i, t := range temp {
			data[i] = t
		}
	}
	return data
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

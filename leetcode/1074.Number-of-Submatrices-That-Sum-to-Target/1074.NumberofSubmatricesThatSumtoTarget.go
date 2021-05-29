package _074_Number_of_Submatrices_That_Sum_to_Target

func numSubmatrixSumTarget(matrix [][]int, target int) int {
	var ans int
	for i := range matrix { // upper bound
		colSum := make([]int, len(matrix[0])) // columns sum
		for _, row := range matrix[i:] {      // lower bound
			for j, v := range row {
				colSum[j] += v
			}
			ans += subarraySum(colSum, target)
		}
	}
	return ans
}

func subarraySum(nums []int, k int) int {
	ans, preSum, m := 0, 0, map[int]int{0: 1}
	for _, num := range nums {
		preSum += num
		ans += m[preSum-k]
		m[preSum]++
	}
	return ans
}

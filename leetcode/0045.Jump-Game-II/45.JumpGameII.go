package leetcode

func jump(nums []int) int {
	var ans, pivot, maxFar int
	for i := 0; i < len(nums)-1 && pivot < len(nums)-1; i++ {
		if i+nums[i] > maxFar { // the farthest index can be reached from i
			maxFar = i + nums[i]
		}
		if i == pivot { // reach the right boundary of last jump range
			pivot = maxFar // start point of the next round of jump
			ans++          // enter the `ans`th step
		}
	}
	return ans
}

func jump1(nums []int) int {
	var ans int
	for lastIdx := len(nums) - 1; lastIdx > 0; {
		for i := 0; i < lastIdx; i++ {
			if nums[i] >= lastIdx-i {
				lastIdx = i
				ans++
				break
			}
		}
	}
	return ans
}

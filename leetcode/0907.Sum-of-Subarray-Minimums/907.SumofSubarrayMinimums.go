package leetcode

type pair struct {
	v, i int
}

// https://leetcode.com/problems/sum-of-subarray-minimums/discuss/178876/stack-solution-with-very-detailed-explanation-step-by-step/195709
func sumSubarrayMins(arr []int) int {
	var preLessStack, nextLessStack []pair
	ans, mod, left, right := 0, 1000000007, make([]int, len(arr)), make([]int, len(arr))
	for i := 0; i < len(arr); i++ {
		for len(preLessStack) > 0 && preLessStack[len(preLessStack)-1].v >= arr[i] {
			preLessStack = preLessStack[:len(preLessStack)-1]
		}
		if len(preLessStack) > 0 {
			left[i] = i - preLessStack[len(preLessStack)-1].i
		} else {
			left[i] = i + 1
		}
		preLessStack = append(preLessStack, pair{v: arr[i], i: i})
	}
	for i := len(arr) - 1; i >= 0; i-- {
		for len(nextLessStack) > 0 && nextLessStack[len(nextLessStack)-1].v > arr[i] {
			nextLessStack = nextLessStack[:len(nextLessStack)-1]
		}
		if len(nextLessStack) > 0 {
			right[i] = nextLessStack[len(nextLessStack)-1].i - i
		} else {
			right[i] = len(arr) - i
		}
		nextLessStack = append(nextLessStack, pair{v: arr[i], i: i})
	}
	for i := 0; i < len(arr); i++ {
		ans = (ans + arr[i]*left[i]*right[i]) % mod
	}
	return ans
}

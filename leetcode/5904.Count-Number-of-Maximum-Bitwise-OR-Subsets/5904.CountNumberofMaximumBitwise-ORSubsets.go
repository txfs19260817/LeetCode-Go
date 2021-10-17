package _904_Count_Number_of_Maximum_Bitwise_OR_Subsets

func countMaxOrSubsets(nums []int) int {
	getOrs := func(arr []int) (ors int) {
		for _, num := range arr {
			ors |= num
		}
		return
	}
	ans, maxOrs := 0, getOrs(nums)
	var recursion func(path []int, idx int)
	recursion = func(path []int, idx int) {
		if idx == len(nums) {
			if getOrs(path) == maxOrs {
				ans++
			}
			return
		}
		path = append(path, nums[idx])
		recursion(path, idx+1)
		path = path[:len(path)-1]
		recursion(path, idx+1)
	}
	recursion(make([]int, 0, len(nums)), 0)
	return ans
}

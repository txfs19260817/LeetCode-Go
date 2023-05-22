package leetcode

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
		path = append(path, nums[idx]) // pick
		recursion(path, idx+1)
		path = path[:len(path)-1] // skip
		recursion(path, idx+1)
	}
	recursion(make([]int, 0, len(nums)), 0)
	return ans
}

func countMaxOrSubsets2(nums []int) int {
	var hi, cnt int
	for i := 1; i < (1 << len(nums)); i++ { // combinations
		var tmp int
		for j := 0; j < len(nums); j++ {
			if (i & (1 << j)) != 0 { // 1: pick, 0: skip
				tmp |= nums[j]
			}
		}
		if tmp > hi {
			hi, cnt = tmp, 1
		} else if tmp == hi {
			cnt++
		}
	}
	return cnt
}

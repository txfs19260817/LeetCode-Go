package leetcode

import "slices"

func combinationSum3(k int, n int) [][]int {
	var ans [][]int
	path := make([]int, 0, k)
	var dfs func(int, int)
	dfs = func(i int, leftSum int) {
		d := k - len(path)                            // d nums to choose
		if leftSum < 0 || leftSum > d*(i+(i-d+1))/2 { // d * (first + last) / 2
			return
		}
		if d == 0 {
			ans = append(ans, slices.Clone(path))
			return
		}
		for j := i; j >= d; j-- {
			path = append(path, j)
			dfs(j-1, leftSum-j)
			path = path[:len(path)-1]
		}
	}
	dfs(9, n)
	return ans
}

func combinationSum32(k int, n int) [][]int {
	var ans [][]int
	path := make([]int, 0, k)
	var dfs func(int, int)
	dfs = func(i int, leftSum int) {
		d := k - len(path)
		if leftSum < 0 || leftSum > d*(i+(i-d+1))/2 {
			return
		}
		if d == 0 {
			ans = append(ans, slices.Clone(path))
			return
		}
		if i > d {
			dfs(i-1, leftSum)
		}
		path = append(path, i)
		dfs(i-1, leftSum-i)
		path = path[:len(path)-1]
	}
	dfs(9, n)
	return ans
}

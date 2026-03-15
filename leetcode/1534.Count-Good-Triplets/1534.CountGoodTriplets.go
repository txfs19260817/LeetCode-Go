package leetcode

import "sort"

func countGoodTriplets(arr []int, a int, b int, c int) int {
	var ans int
	// 对下标数组 idx 按照 arr[i] 的值从小到大排序。
	idx := make([]int, len(arr))
	for i := range idx {
		idx[i] = i
	}
	sort.Slice(idx, func(i, j int) bool { return arr[idx[i]] < arr[idx[j]] })

	for _, j := range idx {
		// 然后遍历 idx 的元素，作为下标 j。（枚举中间）
		// i < j && abs(arr[i]-y) <= a -> left[]
		// k > j && abs(arr[k]-y) <= b -> right[]
		y := arr[j]
		var left, right []int
		for _, i := range idx {
			if i < j && abs(arr[i]-y) <= a {
				left = append(left, arr[i])
			}
		}
		for _, k := range idx {
			if j < k && abs(arr[k]-y) <= b {
				right = append(right, arr[k])
			}
		}

		// 遍历 left 中的元素 x，计算 right 中有多少个元素 z 满足 ∣x−z∣≤c，即在 [x−c,x+c] 中的 z 的个数。
		// 等价于 right 中的 ≤x+c 的元素个数，减去 right 中的 <x−c 的元素个数。
		var k1, k2 int
		for _, x := range left {
			for k1 < len(right) && right[k1] < x-c {
				k1++
			}
			for k2 < len(right) && right[k2] <= x+c {
				k2++
			}
			ans += k2 - k1
		}
	}
	return ans
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

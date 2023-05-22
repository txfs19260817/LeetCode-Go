package leetcode

type pair struct {
	num, idx int
}

func minCostII(costs [][]int) int {
	ans, cur, pre := 2147483647, make([]int, len(costs[0])), costs[len(costs)-1]
	for i := len(costs) - 2; i >= 0; i-- {
		copy(cur, costs[i])
		first, second := minTwo(pre)
		for j := 0; j < len(cur); j++ {
			if j != first.idx {
				cur[j] += first.num
			} else {
				cur[j] += second.num
			}
		}
		copy(pre, cur)
	}
	for _, n := range pre {
		if n < ans {
			ans = n
		}
	}
	return ans
}

func minTwo(nums []int) (pair, pair) {
	first, second := pair{2147483647, -1}, pair{2147483647, -1}
	for i, n := range nums {
		if n < first.num {
			second.num, second.idx = first.num, first.idx
			first.num, first.idx = n, i
		} else if n < second.num {
			second.num, second.idx = n, i
		}
	}
	return first, second
}

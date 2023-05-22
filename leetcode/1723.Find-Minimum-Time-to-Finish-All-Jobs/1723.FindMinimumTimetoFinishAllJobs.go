package leetcode

func minimumTimeRequired(jobs []int, k int) int {
	ans := 1 << 31
	dfs(&ans, 0, jobs, make([]int, k))
	return ans
}

func dfs(ans *int, idx int, jobs, workload []int) {
	if idx == len(jobs) {
		if m := maxInSlice(workload); m < *ans {
			*ans = m
		}
		return
	}
	for i := 0; i < len(workload); i++ {
		if workload[i]+jobs[idx] > *ans { // prune since it cannot be the optimal choice
			continue
		}
		workload[i] += jobs[idx]
		dfs(ans, idx+1, jobs, workload)
		workload[i] -= jobs[idx]
		if workload[i] == 0 { // we only focus on Combination rather than Permutation
			break
		}
	}
}

func maxInSlice(a []int) (m int) {
	for _, n := range a {
		if n > m {
			m = n
		}
	}
	return
}

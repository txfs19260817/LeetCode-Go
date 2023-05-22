package leetcode

import "sort"

type job struct {
	diff, pro int
}

func maxProfitAssignment(difficulty []int, profit []int, worker []int) int {
	jobs, idx, ans := make([]job, len(difficulty)), 0, 0
	for i := 0; i < len(difficulty); i++ {
		jobs[i] = job{difficulty[i], profit[i]}
	}
	sort.Slice(jobs, func(i, j int) bool { return jobs[i].diff < jobs[j].diff })
	sort.Ints(worker)
	for i := 1; i < len(jobs); i++ {
		jobs[i].pro = max(jobs[i].pro, jobs[i-1].pro)
	}
	for _, w := range worker {
		for idx < len(jobs) && w >= jobs[idx].diff {
			idx++
		}
		if idx > 0 {
			ans += jobs[idx-1].pro
		}
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

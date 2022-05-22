package _071_Minimum_Rounds_to_Complete_All_Tasks

func minimumRounds(tasks []int) int {
	ans, t2cnt := 0, map[int]int{}
	for _, task := range tasks {
		t2cnt[task]++
	}
	for _, cnt := range t2cnt {
		v := getMinimum23Pair(cnt)
		if v == -1 {
			return -1
		}
		ans += v
	}
	return ans
}

func getMinimum23Pair(n int) int {
	times := 1 << 30
	for i := 0; i*3 <= n; i++ { // 3
		for j := 0; 3*i+2*j <= n; j++ { // 2
			if 3*i+2*j == n {
				times = min(times, i+j)
			}
		}
	}
	if times == 1<<30 {
		return -1
	}
	return times
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

package leetcode

func numPairsDivisibleBy60(time []int) int {
	ans, min2cnt := 0, [60]int{}
	for _, t := range time {
		modt := t % 60
		ans += min2cnt[(60-modt)%60]
		min2cnt[modt]++
	}
	return ans
}

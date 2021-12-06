package sn

func profitTargets(stocksProfit []int, target int) (ans int) {
	set, foundSamePair := map[int]bool{}, false
	for _, p := range stocksProfit {
		if set[p] || p > target {
			if !foundSamePair && p*2 == target {
				foundSamePair = true
				ans++
			}
			continue
		}
		if set[target-p] {
			ans++
		}
		set[p] = true
	}
	return
}

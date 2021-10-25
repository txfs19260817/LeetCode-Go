package _638_Shopping_Offers

func shoppingOffers(price []int, special [][]int, needs []int) int {
	var cost int
	for i, p := range price {
		cost += p * needs[i]
	}

	sliceEqual := func(a, b []int) bool {
		if len(a) != len(b) {
			return false
		}
		for i := 0; i < len(a); i++ {
			if a[i] != b[i] {
				return false
			}
		}
		return true
	}

	var dfs func(bracket []int, curCost, idx int)
	dfs = func(bracket []int, curCost, idx int) {
		if sliceEqual(bracket, needs) {
			if cost > curCost {
				cost = curCost
			}
			return
		}
		for i := idx; i < len(special); i++ {
			var removeFlag bool
			for j := 0; j < len(special[i])-1; j++ {
				if special[i][j]+bracket[j] > needs[j] {
					removeFlag = true
					break
				}
			}
			if removeFlag {
				continue
			}
			for j := 0; j < len(special[i])-1; j++ {
				bracket[j] += special[i][j]
			}
			dfs(bracket, curCost+special[i][len(special[i])-1], idx) // idx for we can pick an offer several times
			for j := 0; j < len(special[i])-1; j++ {
				bracket[j] -= special[i][j]
			}
		}
		for i, p := range price {
			curCost += p * (needs[i] - bracket[i])
		}
		if cost > curCost {
			cost = curCost
		}
	}
	dfs(make([]int, len(price)), 0, 0)
	return cost
}

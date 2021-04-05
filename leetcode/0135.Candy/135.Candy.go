package _135_Candy

func candy(ratings []int) int {
	ans, pre, decNum := 1, 1, 0
	for i := 1; i < len(ratings); i++ {
		if ratings[i] >= ratings[i-1] {
			if decNum > 0 {
				ans += decNum * (decNum + 1) / 2
				if pre <= decNum {
					ans += decNum - pre + 1
				}
				pre, decNum = 1, 0
			}
			if ratings[i] == ratings[i-1] {
				pre = 1
			} else {
				pre++
			}
			ans += pre
		} else {
			decNum++
		}
	}
	if decNum > 0 {
		ans += decNum * (decNum + 1) / 2
		if pre <= decNum {
			ans += decNum - pre + 1
		}
	}
	return ans
}

func candy1(ratings []int) int {
	candies := make([]int, len(ratings))
	for i := range candies {
		candies[i] = 1
	}

	for i := 1; i < len(ratings); i++ {
		if ratings[i] > ratings[i-1] {
			candies[i] = candies[i-1] + 1
		}
	}
	ans := candies[len(ratings)-1]
	for i := len(ratings) - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			candies[i] = max(candies[i], candies[i+1]+1)
		}
		ans += candies[i]
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

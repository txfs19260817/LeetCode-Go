package _904_Fruit_Into_Baskets

func totalFruit(tree []int) int {
	ans, baskets := 0, map[int]int{} // type:num
	for l, r := 0, 0; l < len(tree); {
		if r < len(tree) && getBasketLen(baskets, tree[r]) <= 2 {
			baskets[tree[r]]++
			r++
		} else {
			baskets[tree[l]]--
			if baskets[tree[l]] == 0 {
				delete(baskets, tree[l])
			}
			l++
		}
		if r-l > ans {
			ans = r - l
		}
	}
	return ans
}

func getBasketLen(b map[int]int, next int) int {
	if _, ok := b[next]; ok {
		return len(b)
	}
	return len(b) + 1
}

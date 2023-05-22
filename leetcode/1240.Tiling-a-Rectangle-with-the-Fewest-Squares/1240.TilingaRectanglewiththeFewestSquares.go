package leetcode

// https://leetcode-cn.com/problems/tiling-a-rectangle-with-the-fewest-squares/solution/1240-pu-ci-zhuan-dong-tai-gui-hua-by-acw_weian/
func tilingRectangle(n int, m int) int {
	cache := map[[2]int]int{}
	get := func(a, b int) (int, bool) {
		if v, ok := cache[[2]int{a, b}]; ok {
			return v, true
		}
		if v, ok := cache[[2]int{b, a}]; ok {
			return v, true
		}
		return 0, false
	}
	set := func(a, b, v int) {
		cache[[2]int{a, b}], cache[[2]int{b, a}] = v, v
	}
	var dfs func(int, int) int
	dfs = func(n, m int) int {
		if n == m {
			return 1
		}
		if n == 1 {
			return m
		}
		if m == 1 {
			return n
		}
		ans := n * m

		for i := 1; i < n/2+1; i++ {
			var tmp int
			if v, ok := get(i, m); ok {
				tmp += v
			} else {
				v = dfs(i, m)
				set(i, m, v)
				tmp += v
			}
			if v, ok := get(n-i, m); ok {
				tmp += v
			} else {
				v = dfs(n-i, m)
				set(n-i, m, v)
				tmp += v
			}
			ans = min(ans, tmp)
		}
		for j := 1; j < m/2+1; j++ {
			var tmp int
			if v, ok := get(n, j); ok {
				tmp += v
			} else {
				v = dfs(n, j)
				set(n, j, v)
				tmp += v
			}
			if v, ok := get(n, m-j); ok {
				tmp += v
			} else {
				v = dfs(n, m-j)
				set(n, m-j, v)
				tmp += v
			}
			ans = min(ans, tmp)
		}
		for a := 1; a < min(n, m); a++ { // sq len
			for i := 1; i < n-a; i++ {
				for j := 1; j < m-a; j++ {
					tmp := 1
					if v, ok := get(a+i, j); ok {
						tmp += v
					} else {
						v = dfs(a+i, j)
						set(a+i, j, v)
						tmp += v
					}
					if v, ok := get(i, m-j); ok {
						tmp += v
					} else {
						v = dfs(i, m-j)
						set(i, m-j, v)
						tmp += v
					}
					if v, ok := get(n-i, m-a-j); ok {
						tmp += v
					} else {
						v = dfs(n-i, m-a-j)
						set(n-i, m-a-j, v)
						tmp += v
					}
					if v, ok := get(n-a-i, a+j); ok {
						tmp += v
					} else {
						v = dfs(n-a-i, a+j)
						set(n-a-i, a+j, v)
						tmp += v
					}
					ans = min(ans, tmp)
				}
			}
		}
		return ans
	}
	return dfs(n, m)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

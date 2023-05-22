package leetcode

func clumsy(N int) int {
	var ans int
	var flag bool
	for i := N; i >= 1; i -= 4 {
		switch {
		case i >= 4:
			e := -i * (i - 1) / (i - 2)
			if !flag {
				e = -e
				flag = true
			}
			ans += e + i - 3
		case i == 3:
			ans -= 6
		case i == 2:
			ans -= 2
		case i == 1:
			ans -= 1
		}
		if !flag {
			ans = -ans
			flag = true
		}
	}
	return ans
}

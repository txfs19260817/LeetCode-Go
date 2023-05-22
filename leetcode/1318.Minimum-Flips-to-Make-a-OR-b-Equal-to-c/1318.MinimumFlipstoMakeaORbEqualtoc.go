package leetcode

func minFlips(a int, b int, c int) int {
	var ans int
	for i := 0; i < 31; i++ {
		if (c>>i)&1 == 1 && (a>>i)&1 == 0 && (b>>i)&1 == 0 {
			ans++
		} else if (c>>i)&1 == 0 {
			if (a>>i)&1 == 1 {
				ans++
			}
			if (b>>i)&1 == 1 {
				ans++
			}
		}
	}
	return ans
}

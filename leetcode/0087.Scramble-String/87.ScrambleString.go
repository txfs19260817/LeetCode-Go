package _087_Scramble_String

func isScramble(s1 string, s2 string) bool {
	dp := make([][][]int8, len(s1))
	for i := range dp {
		dp[i] = make([][]int8, len(s2))
		for j := range dp[i] {
			dp[i][j] = make([]int8, len(s1)+1)
			for k := range dp[i][j] {
				dp[i][j][k] = -1
			}
		}
	}
	var dfs func(int, int, int) int8
	dfs = func(i1, i2, length int) (res int8) {
		d := &dp[i1][i2][length]
		if *d != -1 {
			return *d
		}
		defer func() { *d = res }()
		x, y := s1[i1:i1+length], s2[i2:i2+length]
		if x == y {
			return 1
		}
		var freq [26]int8
		for i, c := range x {
			freq[c-'a']++
			freq[y[i]-'a']--
		}
		for _, f := range freq {
			if f != 0 {
				return 0
			}
		}
		for i := 1; i < length; i++ {
			if dfs(i1, i2, i) == 1 && dfs(i1+i, i2+i, length-i) == 1 {
				return 1
			}
			if dfs(i1, i2+length-i, i) == 1 && dfs(i1+i, i2, length-i) == 1 {
				return 1
			}
		}
		return 0
	}
	return dfs(0, 0, len(s1)) == 1
}

package _784_Letter_Case_Permutation

func letterCasePermutation(S string) []string {
	var ans []string
	dfs(&ans, 0, []byte{}, S)
	return ans
}

func dfs(ans *[]string, index int, path []byte, S string) {
	if index == len(S) {
		*ans = append(*ans, string(path))
		return
	}
	for i := index; i < len(S) && len(path) < len(S); i++ {
		lower, upper, isAlpha := getLowerAndUpper(S[i])
		if !isAlpha {
			path = append(path, S[i])
			dfs(ans, index+1, path, S)
			break
		}
		path = append(path, lower)
		dfs(ans, index+1, path, S)
		path = path[:len(path)-1]
		path = append(path, upper)
		dfs(ans, index+1, path, S)
	}
}

func getLowerAndUpper(c byte) (byte, byte, bool) {
	if c >= 65 && c <= 90 {
		return c + 32, c, true
	}
	if c >= 97 && c <= 122 {
		return c, c - 32, true
	}
	return c, c, false
}

func letterCasePermutation1(S string) []string {
	bytes := make([][]byte, 1)
	for i := 0; i < len(S); i++ {
		lower, upper, isAlpha := getLowerAndUpper(S[i])
		if !isAlpha {
			for j := 0; j < len(bytes); j++ {
				bytes[j] = append(bytes[j], S[i])
			}
			continue
		}
		for j, n := 0, len(bytes); j < n; j++ {
			temp := append([]byte{}, bytes[j]...)
			bytes = append(bytes, temp)
			bytes[j] = append(bytes[j], lower)
			bytes[n+j] = append(bytes[n+j], upper)
		}
	}
	ans := make([]string, 0, len(bytes))
	for _, b := range bytes {
		ans = append(ans, string(b))
	}
	return ans
}

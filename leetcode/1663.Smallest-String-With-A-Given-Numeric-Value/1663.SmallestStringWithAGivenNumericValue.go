package _663_Smallest_String_With_A_Given_Numeric_Value

func getSmallestString(n int, k int) string {
	ans := make([]byte, n)
	for i := n - 1; i >= 0; i-- {
		ans[i] = byte(max(min(int('z'), k-i+'a'-1), 'a')) // the highest possible character (a or remaining k or z)
		k -= int(ans[i] - 'a' + 1)
	}
	return string(ans)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

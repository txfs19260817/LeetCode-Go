package leetcode

func removeDigit(number string, digit byte) string {
	var ans string
	bs := []byte(number)
	for i, b := range bs {
		if b == digit {
			candidate := string(append([]byte{}, bs[:i]...)) + string(append([]byte{}, bs[i+1:]...))
			if len(ans) == 0 || ans < candidate {
				ans = candidate
			}
		}
	}
	return ans
}

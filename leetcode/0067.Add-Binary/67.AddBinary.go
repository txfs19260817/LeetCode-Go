package leetcode

func addBinary(a string, b string) string {
	var ans []byte
	for i, j, c := len(a)-1, len(b)-1, 0; i >= 0 || j >= 0 || c != 0; i, j = i-1, j-1 {
		res := c
		if i >= 0 {
			res += int(a[i] - '0')
		}
		if j >= 0 {
			res += int(b[j] - '0')
		}
		switch res {
		case 0, 1:
			ans = append([]byte{byte(res + '0')}, ans...)
			c = 0
		case 2, 3:
			ans = append([]byte{byte(res - 2 + '0')}, ans...)
			c = 1
		}
	}
	return string(ans)
}

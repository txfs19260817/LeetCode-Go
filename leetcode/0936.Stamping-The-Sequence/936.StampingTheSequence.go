package leetcode

func movesToStamp(stamp string, target string) []int {
	ans, tSlice := make([]int, 0, 10*len(target)), []byte(target)
	stamping := func(start int) (stamped bool) {
		for i := 0; i < len(stamp); i++ {
			if tSlice[start+i] == '?' {
				continue
			}
			if tSlice[start+i] != stamp[i] {
				return false
			}
			stamped = true
		}
		if stamped {
			for i := 0; i < len(stamp); i++ {
				tSlice[start+i] = '?'
			}
			ans = append(ans, start)
		}
		return stamped
	}
	for {
		var isChanged bool
		for i := 0; i < len(target)-len(stamp)+1; i++ {
			isChanged = isChanged || stamping(i)
		}
		if !isChanged {
			break
		}
	}
	for _, c := range tSlice {
		if c != '?' {
			return nil
		}
	}
	for l, r := 0, len(ans)-1; l < r; l, r = l+1, r-1 {
		ans[l], ans[r] = ans[r], ans[l]
	}
	return ans
}

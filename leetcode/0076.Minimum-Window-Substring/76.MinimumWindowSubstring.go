package leetcode

func minWindow(s string, t string) string {
	ms, mt := [128]int{}, [128]int{}
	for _, c := range t {
		mt[c]++
	}

	contain := func() bool {
		for _, c := range t {
			if ms[c] < mt[c] {
				return false
			}
		}
		return true
	}
	l, ansL, ansR := 0, -1, len(s)
	for r := range s {
		ms[s[r]]++
		for contain() {
			if r-l < ansR-ansL {
				ansL, ansR = l, r
			}
			ms[s[l]]--
			l++
		}
	}
	if ansL < 0 {
		return ""
	}
	return s[ansL : ansR+1]
}

func minWindow2(s, t string) string {
	cnt := [128]int{}
	less := 0
	for _, c := range t {
		if cnt[c] == 0 {
			less++ // 有 less 种字母的出现次数 < t 中的字母出现次数
		}
		cnt[c]++
	}

	ansLeft, ansRight := -1, len(s)
	left := 0
	for right, c := range s { // 移动子串右端点
		cnt[c]-- // 右端点字母移入子串
		if cnt[c] == 0 {
			// 原来窗口内 c 的出现次数比 t 的少，现在一样多
			less--
		}
		for less == 0 { // 涵盖：所有字母的出现次数都是 >=
			if right-left < ansRight-ansLeft { // 找到更短的子串
				ansLeft, ansRight = left, right // 记录此时的左右端点
			}
			x := s[left] // 左端点字母
			if cnt[x] == 0 {
				// x 移出窗口之前，检查出现次数，
				// 如果窗口内 x 的出现次数和 t 一样，
				// 那么 x 移出窗口后，窗口内 x 的出现次数比 t 的少
				less++
			}
			cnt[x]++ // 左端点字母移出子串
			left++
		}
	}
	if ansLeft < 0 {
		return ""
	}
	return s[ansLeft : ansRight+1]
}

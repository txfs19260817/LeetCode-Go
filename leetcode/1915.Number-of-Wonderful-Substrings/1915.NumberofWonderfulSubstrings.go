package _915_Number_of_Wonderful_Substrings

// https://leetcode-cn.com/problems/number-of-wonderful-substrings/solution/qian-zhui-he-chang-jian-ji-qiao-by-endle-t57t/
func wonderfulSubstrings(word string) int64 {
	cnt := [1024]int{1} // 初始前缀和为 0，需将其计入出现次数
	ans, pre := 0, 0
	for _, ch := range word {
		pre ^= 1 << (ch - 'a')          // 计算当前前缀和
		ans += cnt[pre]                 // 所有字母均出现偶数次
		for i := 1; i < 1024; i <<= 1 { // 枚举其中一个字母出现奇数次
			ans += cnt[pre^i] // 反转第 i 个字母的出现次数的奇偶性
		}
		cnt[pre]++ // 更新前缀和出现次数
	}
	return int64(ans)
}

// TLE
func wonderfulSubstrings2(word string) int64 {
	ans, prefix := int64(0), make([]int, len(word)+1)
	valid := func(n int) bool {
		for count := 0; n > 0; {
			n &= n - 1
			count++
			if count > 1 {
				return false
			}
		}
		return true
	}
	for i, c := range word {
		prefix[i+1] = (1 << int(c-'a')) ^ prefix[i]
	}
	for i := 0; i < len(prefix); i++ {
		for j := i + 1; j < len(prefix); j++ {
			if valid(prefix[j] ^ prefix[i]) {
				ans++
			}
		}
	}
	return ans
}

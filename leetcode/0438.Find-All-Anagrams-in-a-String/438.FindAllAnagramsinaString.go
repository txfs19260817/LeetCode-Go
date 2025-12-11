package leetcode

func findAnagrams(s string, p string) []int {
	k := len(p)
	ans, ms, mp := []int{}, [26]int{}, [26]int{}
	for _, c := range p {
		mp[c-'a']++
	}
	for i, c := range s {
		ms[c-'a']++

		if i < k-1 {
			continue
		}

		if mp == ms {
			ans = append(ans, i-k+1)
		}

		ms[s[i-k+1]-'a']--
	}
	return ans
}

package leetcode

func replaceNonCoprimes(nums []int) []int {
	var ans []int
	for _, num := range nums {
		ans = append(ans, num)
		for len(ans) > 1 && GCD(ans[len(ans)-1], ans[len(ans)-2]) > 1 {
			a, b := ans[len(ans)-1], ans[len(ans)-2]
			g := GCD(a, b)
			ans = ans[:len(ans)-2]
			ans = append(ans, a*b/g)
		}
	}

	return ans
}

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

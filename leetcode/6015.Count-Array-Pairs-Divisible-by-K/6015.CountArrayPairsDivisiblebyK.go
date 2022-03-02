package _015_Count_Array_Pairs_Divisible_by_K

func coutPairs(nums []int, k int) int64 {
	ans, cnt := int64(0), make([]int, len(nums))
	for i, num := range nums {
		if num%k == 0 {
			ans += int64(i)
			cnt[0]++
			continue
		}
		div := gcd(k, nums[i])
		for d := 0; d <= div; d++ {
			ans += int64(cnt[k/div*d])
		}
		cnt[div]++
	}
	return ans
}

// greatest common divisor (GCD) via Euclidean algorithm
func gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

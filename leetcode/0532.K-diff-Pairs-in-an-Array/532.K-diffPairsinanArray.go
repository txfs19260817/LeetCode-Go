package _532_K_diff_Pairs_in_an_Array

func findPairs(nums []int, k int) int {
	ans, m := 0, map[int]int{} //num:freq
	for _, num := range nums {
		m[num]++
	}
	for key := range m {
		if k == 0 && m[key] > 1 {
			ans++
		}
		if k > 0 && m[key+k] > 0 {
			ans++
		}
	}
	return ans
}

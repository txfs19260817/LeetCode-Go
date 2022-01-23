package _150_Find_All_Lonely_Numbers_in_the_Array

func findLonely(nums []int) []int {
	var ans []int
	count := map[int]int{}
	for _, num := range nums {
		count[num]++
	}
	for num, cnt := range count {
		if cnt == 1 && count[num+1] == 0 && count[num-1] == 0 {
			ans = append(ans, num)
		}
	}
	return ans
}

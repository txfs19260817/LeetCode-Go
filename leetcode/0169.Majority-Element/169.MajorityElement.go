package leetcode

func majorityElement(nums []int) int {
	var major, count int
	for _, num := range nums {
		if count == 0 {
			major = num
		}
		if num == major {
			count++
		} else {
			count--
		}
	}
	return major
}

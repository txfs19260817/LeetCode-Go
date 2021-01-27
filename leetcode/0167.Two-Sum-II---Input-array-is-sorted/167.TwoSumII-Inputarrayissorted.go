package _167_Two_Sum_II___Input_array_is_sorted

func twoSum(numbers []int, target int) []int {
	for l, r := 0, len(numbers)-1; l < r; {
		if numbers[l]+numbers[r] == target {
			return []int{l + 1, r + 1}
		}
		if numbers[l]+numbers[r] > target {
			r--
		} else {
			l++
		}
	}
	return []int{1, len(numbers)}
}

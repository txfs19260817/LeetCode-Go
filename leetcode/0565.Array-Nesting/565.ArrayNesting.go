package _565_Array_Nesting

func arrayNesting(nums []int) int { // Space O(n)
	ans, visited := 0, make([]bool, len(nums))
	for i, next := range nums {
		if visited[i] {
			continue
		}
		var l int
		for j := i; !visited[j]; {
			visited[j] = true
			j = next
			next = nums[next]
			l++
		}
		if l > ans {
			ans = l
		}
	}
	return ans
}

func arrayNesting2(nums []int) int { // Space O(1)
	var ans int
	const MAX = 1<<31 - 1
	for i, next := range nums {
		if next == MAX {
			continue
		}
		var l int
		for j := i; next != MAX; {
			nums[j] = MAX
			j = next
			next = nums[next]
			l++
		}
		if l > ans {
			ans = l
		}
	}
	return ans
}

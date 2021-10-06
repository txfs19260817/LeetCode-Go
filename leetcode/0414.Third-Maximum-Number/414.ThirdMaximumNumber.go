package _414_Third_Maximum_Number

func thirdMax(nums []int) int {
	first, second, third, updated := -1<<32, -1<<32, -1<<32, false
	for _, num := range nums {
		if num == first || num == second || num == third {
			continue
		}
		oldThird := third
		if num > first {
			third = second
			second = first
			first = num
		} else if num > second {
			third = second
			second = num
		} else if num > third {
			third = num
		}
		if third > oldThird {
			updated = true
		}
	}
	if updated {
		return third
	}
	return first
}

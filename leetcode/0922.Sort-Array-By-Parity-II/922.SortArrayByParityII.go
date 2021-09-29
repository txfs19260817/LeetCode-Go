package _922_Sort_Array_By_Parity_II

func sortArrayByParityII(nums []int) []int {
	for i, j := 0, 1; i < len(nums) && j < len(nums); {
		for i < len(nums) && nums[i]%2 == 0 {
			i += 2
		}
		for j < len(nums) && nums[j]%2 == 1 {
			j += 2
		}
		if i < len(nums) && j < len(nums) {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	return nums
}

func sortArrayByParityII2(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		if nums[i]%2 == 1 && i%2 == 0 {
			for j := i + 1; j < len(nums); j++ {
				if nums[j]%2 == 0 && j%2 == 1 {
					nums[i], nums[j] = nums[j], nums[i]
					break
				}
			}
		} else if nums[i]%2 == 0 && i%2 == 1 {
			for j := i + 1; j < len(nums); j++ {
				if nums[j]%2 == 1 && j%2 == 0 {
					nums[i], nums[j] = nums[j], nums[i]
					break
				}
			}
		}
	}
	return nums
}

func sortArrayByParityII3(A []int) []int {
	odds, evens := make([]int, 0, len(A)/2), make([]int, 0, len(A)/2)
	for _, a := range A {
		if a%2 == 0 {
			evens = append(evens, a)
		} else {
			odds = append(odds, a)
		}
	}
	odds = append(odds, evens...)
	for i, num := range odds {
		A[mappedIndex(i, len(A))] = num
	}
	return A
}

func mappedIndex(i, n int) int {
	return (1 + 2*i) % (n | 1)
}

package _229_Majority_Element_II

func majorityElement(nums []int) []int {
	var major1, vote1, cnt1, major2, vote2, cnt2 int
	for _, num := range nums {
		if vote1 > 0 && major1 == num {
			vote1++
		} else if vote2 > 0 && major2 == num {
			vote2++
		} else if vote1 == 0 {
			major1, vote1 = num, 1
		} else if vote2 == 0 {
			major2, vote2 = num, 1
		} else {
			vote1--
			vote2--
		}
	}
	for _, num := range nums {
		if num == major1 {
			cnt1++
		} else if num == major2 {
			cnt2++
		}
	}
	ans := make([]int, 0, 2)
	if cnt1 > len(nums)/3 {
		ans = append(ans, major1)
	}
	if cnt2 > len(nums)/3 {
		ans = append(ans, major2)
	}
	return ans
}

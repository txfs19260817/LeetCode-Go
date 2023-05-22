package leetcode

func countElements(nums []int) int {
	ma, mi, cntMa, cntMi := nums[0], nums[0], 1, 1
	for _, num := range nums[1:] {
		if num > ma {
			ma, cntMa = num, 1
		} else if num == ma {
			cntMa++
		}
		if num < mi {
			mi, cntMi = num, 1
		} else if num == mi {
			cntMi++
		}
	}
	if mi == ma {
		return 0
	}
	return len(nums) - cntMi - cntMa
}

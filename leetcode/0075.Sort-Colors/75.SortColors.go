package _075_Sort_Colors

func sortColors(nums []int) {
	if len(nums) < 2 {
		return
	}
	r0, w1, b2 := 0, 0, 0
	for _, x := range nums {
		switch x {
		case 0:
			nums[b2] = 2
			b2++
			nums[w1] = 1
			w1++
			nums[r0] = 0
			r0++
		case 1:
			nums[b2] = 2
			b2++
			nums[w1] = 1
			w1++
		case 2:
			b2++
		}
	}
}

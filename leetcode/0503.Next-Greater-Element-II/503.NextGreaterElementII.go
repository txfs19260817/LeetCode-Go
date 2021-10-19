package _503_Next_Greater_Element_II

func nextGreaterElements(nums []int) []int {
	type ele struct {
		num, idx int
	}
	ans, monoStack := make([]int, len(nums)), make([]ele, 0, len(nums))
	for i := range ans { // in case all elements are the same
		ans[i] = -1
	}
	for idx, num := range nums {
		for len(monoStack) > 0 && monoStack[len(monoStack)-1].num < num {
			ans[monoStack[len(monoStack)-1].idx] = num
			monoStack = monoStack[:len(monoStack)-1]
		}
		monoStack = append(monoStack, ele{num, idx})
	}
	for idx, num := range nums { // 2nd iteration
		for len(monoStack) > 0 && monoStack[len(monoStack)-1].num < num {
			ans[monoStack[len(monoStack)-1].idx] = num
			monoStack = monoStack[:len(monoStack)-1]
		}
		if monoStack[len(monoStack)-1].num == num && monoStack[len(monoStack)-1].idx == idx {
			ans[monoStack[len(monoStack)-1].idx] = -1
			monoStack = monoStack[:len(monoStack)-1]
		}
		if len(monoStack) == 0 { // ans gets filled fully
			break
		}
		monoStack = append(monoStack, ele{num, idx})
	}
	return ans
}

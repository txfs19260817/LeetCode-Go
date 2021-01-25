package _027_Remove_Element

func removeElement(nums []int, val int) int {
	i := 0
	for j := 0; j < len(nums); j++ {
		if nums[j] != val {
			nums[i] = nums[j]
			i++
		}
	}
	return i
}

/*
[3,2,2,3]
3

[0,1,2,2,3,0,4,2]
2

[2]
3
*/

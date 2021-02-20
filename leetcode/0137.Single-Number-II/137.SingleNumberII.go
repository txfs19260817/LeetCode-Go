package _137_Single_Number_II

func singleNumber(nums []int) int {
	var X, Y int
	for _, n := range nums {
		Y = Y ^ n & ^X
		X = X ^ n & ^Y
	}
	return Y
}

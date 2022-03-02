package _004_Count_Operations_to_Obtain_Zero

func countOperations(num1 int, num2 int) int {
	var ans int
	for ; num1 != 0 && num2 != 0; ans++ {
		if num1 > num2 {
			num1 -= num2
		} else {
			num2 -= num1
		}
	}
	return ans
}

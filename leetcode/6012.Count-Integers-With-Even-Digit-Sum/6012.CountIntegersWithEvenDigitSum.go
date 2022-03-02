package _012_Count_Integers_With_Even_Digit_Sum

func countEven(num int) int {
	var ans int
	for i := 1; i <= num; i++ {
		var sum int
		for j := i; j > 0; j /= 10 {
			sum += j % 10
		}
		if sum%2 == 0 {
			ans++
		}
	}
	return ans
}

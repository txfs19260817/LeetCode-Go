package _588_Sum_of_All_Odd_Length_Subarrays

func sumOddLengthSubarrays(arr []int) int {
	var ans int
	for i := 1; i < len(arr); i++ {
		arr[i] += arr[i-1]
	}
	for i := len(arr) - 1; i >= 0; i-- {
		if i%2 == 0 {
			ans += arr[i]
		}
		for j := i - 1; j >= 0; j -= 2 {
			ans += arr[i] - arr[j]
		}
	}
	return ans
}

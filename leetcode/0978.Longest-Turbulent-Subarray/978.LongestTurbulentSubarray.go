package _978_Longest_Turbulent_Subarray

func maxTurbulenceSize(arr []int) int {
	ans, start, intCmp := 1, 0, func(i, j int) int {
		if i > j {
			return 1
		} else if i < j {
			return -1
		}
		return 0
	}
	for i := 1; i < len(arr); i++ {
		if c := intCmp(arr[i-1], arr[i]); c == 0 {
			start = i
		} else if i == len(arr)-1 || c*intCmp(arr[i], arr[i+1]) != -1 {
			if l := i - start + 1; ans < l {
				ans = l
			}
			start = i
		}
	}
	return ans
}

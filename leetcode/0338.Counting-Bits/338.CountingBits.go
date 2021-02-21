package _338_Counting_Bits

func countBits(num int) []int {
	res := make([]int, num+1)
	for i := 1; i < len(res); i++ {
		res[i] = res[i/2] + (i % 2)
	}
	return res
}

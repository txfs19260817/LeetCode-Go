package _412_Fizz_Buzz

import "strconv"

func fizzBuzz(n int) []string {
	ans := make([]string, n)
	for i := 1; i <= n; i++ {
		switch {
		case i%3 == 0 && i%5 == 0:
			ans[i-1] = "FizzBuzz"
		case i%3 == 0:
			ans[i-1] = "Fizz"
		case i%5 == 0:
			ans[i-1] = "Buzz"
		default:
			ans[i-1] = strconv.Itoa(i)
		}
	}
	return ans
}

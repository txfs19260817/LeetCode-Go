package _109_Corporate_Flight_Bookings

func corpFlightBookings(bookings [][]int, n int) []int {
	ans := make([]int, n)
	for _, book := range bookings {
		ans[book[0]-1] += book[2]
		if book[1] < n {
			ans[book[1]] -= book[2]
		}
	}
	for i := 1; i < n; i++ {
		ans[i] += ans[i-1]
	}
	return ans
}

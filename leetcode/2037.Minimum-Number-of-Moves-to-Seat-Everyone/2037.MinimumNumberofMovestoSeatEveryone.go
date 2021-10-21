package _037_Minimum_Number_of_Moves_to_Seat_Everyone

import "sort"

func minMovesToSeat(seats []int, students []int) int {
	var ans int
	sort.Ints(seats)
	sort.Ints(students)
	for i := range seats {
		ans += abs(seats[i] - students[i])
	}
	return ans
}

func abs(a int) int {
	if a < 0 {
		return -1 * a
	}
	return a
}

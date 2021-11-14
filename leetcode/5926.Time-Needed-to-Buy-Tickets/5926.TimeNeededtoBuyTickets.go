package _926_Time_Needed_to_Buy_Tickets

func timeRequiredToBuy(tickets []int, k int) int {
	var ans int
	for i, ticket := range tickets {
		if ticket < tickets[k] {
			ans += ticket
			continue
		}
		if i <= k {
			ans += tickets[k]
		} else {
			ans += tickets[k] - 1
		}
	}
	return ans
}

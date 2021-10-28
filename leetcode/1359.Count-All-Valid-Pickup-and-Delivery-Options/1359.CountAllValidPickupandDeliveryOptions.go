package _359_Count_All_Valid_Pickup_and_Delivery_Options

// https://leetcode-cn.com/problems/count-all-valid-pickup-and-delivery-options/solution/you-xiao-de-kuai-di-xu-lie-shu-mu-by-leetcode-solu/
func countOrders(n int) int {
	ans := 1
	for i := 2; i <= n; i++ {
		ans *= ((2*i - 1) % 1000000007) * i % 1000000007
		ans %= 1000000007
	}
	return ans
}

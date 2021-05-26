package _810_Chalkboard_XOR_Game

func xorGame(nums []int) bool {
	if len(nums)%2 == 0 {
		return true
	}
	var xor int
	for _, num := range nums {
		xor ^= num
	}
	return xor == 0
}

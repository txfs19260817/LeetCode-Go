package _055_Jump_Game

func canJump(nums []int) bool {
	var end, maxFar int
	for i := 0; i < len(nums); i++ {
		if i > end {
			return false
		}
		if maxFar < i+nums[i] {
			maxFar = i + nums[i]
		}
		if i == end {
			end = maxFar
		}
	}
	return true
}

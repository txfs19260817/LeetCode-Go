package leetcode

func slowestKey(releaseTimes []int, keysPressed string) byte {
	ans, long := keysPressed[0], releaseTimes[0]
	for i := 1; i < len(releaseTimes); i++ {
		if t := releaseTimes[i] - releaseTimes[i-1]; t > long || t == long && keysPressed[i] > ans {
			ans, long = keysPressed[i], t
		}
	}
	return ans
}

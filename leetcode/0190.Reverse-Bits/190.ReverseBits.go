package leetcode

func reverseBits(num uint32) uint32 {
	var ans uint32
	for i := 0; i < 32; i++ {
		lastBit := num & 1     // get the last bit from num
		num >>= 1              // remove the last bit
		ans = ans<<1 + lastBit // append the bit
	}
	return ans
}

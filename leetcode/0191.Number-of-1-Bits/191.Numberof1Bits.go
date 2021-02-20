package _191_Number_of_1_Bits

func hammingWeight(num uint32) int {
	var count int
	for i := 0; i < 32; i++ {
		count += int(num & 1)
		num >>= 1
	}
	return count
}

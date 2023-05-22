package leetcode

func xorOperation(n int, start int) int {
	// the lowest bit is 1 if and only if n and start are both odd numbers
	lowestBit := n & start & 1

	// start ^ (start+2) ^ (start+4) ^ ... ^(start + 2*(n-1))
	// let s = start / 2
	// <=> (s ^ (s+1) ^ (s+2) ^ ... ^ (s+n-1)) * 2 + lowestBit
	s := start >> 1

	// ^(N ... M) = ^(1 ... N-1) ^ ^(1 ... M)
	temp := computeXOR(s-1) ^ computeXOR(s+n-1)
	return temp<<1 | lowestBit
}

func computeXOR(x int) int {
	// The pattern of XOR of the first n numbers
	//      n   binary ^(0...N)   return
	//      1    0001    0001        1
	//      2    0010    0011       n+1
	//      3    0011    0000        0
	//      4    0100    0100        n
	//      5    0101    0001        1
	//      6    0110    0111       n+1
	//      ...  ...     ...        ...
	switch x % 4 {
	case 0:
		return x
	case 1:
		return 1
	case 2:
		return x + 1
	default: // 3
		return 0
	}
}

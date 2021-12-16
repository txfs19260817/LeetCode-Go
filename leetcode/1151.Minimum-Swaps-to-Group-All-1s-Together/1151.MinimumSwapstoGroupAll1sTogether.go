package _151_Minimum_Swaps_to_Group_All_1s_Together

func minSwaps(data []int) int {
	var totalOnes, maxLocalOnes int
	for _, v := range data {
		if v == 1 {
			totalOnes++
		}
	}
	for l, r, localOnes := 0, 0, 0; r < len(data); { // sliding window with size `ones`
		localOnes += data[r]
		r++
		if l < r-totalOnes {
			localOnes -= data[l]
			l++
		}
		if localOnes > maxLocalOnes {
			maxLocalOnes = localOnes
		}
	}
	return totalOnes - maxLocalOnes
}

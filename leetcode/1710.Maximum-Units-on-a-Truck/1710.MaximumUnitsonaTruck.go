package leetcode

import "sort"

func maximumUnits(boxTypes [][]int, truckSize int) int {
	var ans int
	sort.Slice(boxTypes, func(i, j int) bool { return boxTypes[i][1] > boxTypes[j][1] })
	for _, box := range boxTypes {
		if truckSize > box[0] {
			truckSize -= box[0]
			ans += box[1] * box[0]
		} else {
			ans += box[1] * truckSize
			break
		}
	}
	return ans
}

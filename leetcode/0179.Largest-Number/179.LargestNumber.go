package leetcode

import (
	"strconv"
	"strings"
)

func largestNumber(nums []int) string {
	strSlice := make([]string, len(nums))
	for i, num := range nums {
		strSlice[i] = strconv.Itoa(num)
	}
	//sort.Slice(strSlice, func(i, j int) bool {return strSlice[i]+strSlice[j] > strSlice[j]+strSlice[i]})
	quickSort(strSlice, 0, len(strSlice)-1)
	res := strings.Join(strSlice, "")
	if res[0] == '0' {
		return "0"
	}
	return res
}

func quickSort(data []string, left, right int) {
	if left < right {
		pivot := partition(data, left, right)
		quickSort(data, left, pivot-1)
		quickSort(data, pivot+1, right)
	}
}

func partition(data []string, left, right int) int {
	pivotElem := data[right]
	i := left - 1
	for j := left; j < right; j++ {
		if data[j]+pivotElem > pivotElem+data[j] {
			i++
			data[i], data[j] = data[j], data[i]
		}
	}
	i++
	data[i], data[right] = data[right], data[i]
	return i
}

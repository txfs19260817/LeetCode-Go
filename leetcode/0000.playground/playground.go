package _000_playground

/*
A playground to code benefiting from auto-completion within IDE.
*/
func pancakeSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	var res []int
	for right := len(arr); right > 0; right-- {
		idx := find(arr, right)
		if idx+1 != right {
			reverse(arr, 0, idx)
			reverse(arr, 0, right-1)
			res = append(res, idx+1, right)
		}
	}
	return res
}

func reverse(arr []int, l, r int) {
	for ; l < r; l, r = l+1, r-1 {
		arr[l], arr[r] = arr[r], arr[l]
	}
}

func find(arr []int, t int) int {
	for i, a := range arr {
		if a == t {
			return i
		}
	}
	return -1
}

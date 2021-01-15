package _969_Pancake_Sorting

func pancakeSort(arr []int) []int {
	if len(arr) == 0 {
		return arr
	}
	res := make([]int, 0, len(arr))
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

func find(arr []int, t int) (i int) {
	for i, n := range arr {
		if n == t {
			return i
		}
	}
	return -1
}

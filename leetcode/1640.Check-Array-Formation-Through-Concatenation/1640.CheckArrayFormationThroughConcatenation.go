package _640_Check_Array_Formation_Through_Concatenation

func canFormArray(arr []int, pieces [][]int) bool {
	connected, dict := make([]int, 0, len(arr)), map[int][]int{} // pieces[i][0]:pieces[i]
	for _, piece := range pieces {
		dict[piece[0]] = piece
	}
	for _, a := range arr {
		connected = append(connected, dict[a]...)
	}
	if len(arr) != len(connected) {
		return false
	}
	for i := 0; i < len(arr); i++ {
		if arr[i] != connected[i] {
			return false
		}
	}
	return true
}

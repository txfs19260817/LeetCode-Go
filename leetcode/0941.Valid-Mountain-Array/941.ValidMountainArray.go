package _941_Valid_Mountain_Array

func validMountainArray(arr []int) bool {
	var down bool
	for i := 1; i < len(arr); i++ {
		if arr[i-1] > arr[i] {
			down = true
			continue
		}
		if !down && arr[i-1] >= arr[i] || down && arr[i-1] <= arr[i] {
			return false
		}
	}
	return !(len(arr) < 2 || !down || down && arr[0] > arr[1])
}

func validMountainArray2(arr []int) bool {
	var i int
	for i+1 < len(arr) && arr[i] < arr[i+1] {
		i++
	}
	if i == 0 || i == len(arr)-1 {
		return false
	}
	for i+1 < len(arr) && arr[i] > arr[i+1] {
		i++
	}
	return i == len(arr)-1
}

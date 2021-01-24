package sorts

func BubbleSort(data []int) []int {
	for i := 0; i < len(data); i++ {
		for j := 1; j < len(data)-i; j++ {
			if data[j-1] > data[j] {
				data[j-1], data[j] = data[j], data[j-1]
			}
		}
	}
	return data
}

func SelectionSort(data []int) []int {
	for i := 0; i < len(data); i++ {
		maxIdx := 0
		for j := 1; j < len(data)-i; j++ {
			if data[maxIdx] < data[j] {
				maxIdx = j
			}
		}
		data[maxIdx], data[len(data)-i-1] = data[len(data)-i-1], data[maxIdx]
	}
	return data
}

func InsertionSort(data []int) []int {
	for i := 1; i < len(data); i++ {
		if data[i-1] > data[i] {
			j := i - 1
			temp := data[i]
			for j >= 0 && data[j] > temp {
				data[j+1] = data[j]
				j--
			}
			data[j+1] = temp
		}
	}
	return data
}

func ShellSort(data []int) []int {
	h := 1
	for h < len(data)/3 {
		h = h*3 + 1
	}
	for ; h > 0; h /= 3 {
		for i := h; i < len(data); i++ {
			j := i - h
			temp := data[i]
			for j >= 0 && data[j] > temp {
				data[j+h] = data[j]
				j -= h
			}
			data[j+h] = temp
		}
	}
	return data
}

func QuickSort(data []int) []int {
	quickSort(data, 0, len(data)-1)
	return data
}

func quickSort(data []int, left, right int) {
	if left < right {
		pivot := partition2(data, left, right)
		quickSort(data, left, pivot-1)
		quickSort(data, pivot+1, right)
	}
}

func partition(data []int, left, right int) int {
	pivotElem := data[left]
	for left < right {
		for left < right && data[right] >= pivotElem {
			right--
		}
		if left < right {
			data[left] = data[right]
		}
		for left < right && data[left] <= pivotElem {
			left++
		}
		if left < right {
			data[right] = data[left]
		}
	}
	data[left] = pivotElem
	return left
}

func partition2(data []int, left, right int) int {
	pivotElem := data[right]
	i := left
	for j := left; j < right; j++ {
		if data[j] < pivotElem {
			data[i], data[j] = data[j], data[i]
			i++
		}
	}
	data[i], data[right] = data[right], data[i]
	return i
}

func MergeSort(data []int) []int {
	if len(data) < 2 {
		return data
	}
	mid := len(data) / 2
	return merge(MergeSort(data[:mid]), MergeSort(data[mid:]))
}

func merge(left, right []int) []int {
	res := make([]int, 0, len(left)+len(right))
	for len(left) > 0 && len(right) > 0 {
		if left[0] < right[0] {
			res = append(res, left[0])
			left = left[1:]
		} else {
			res = append(res, right[0])
			right = right[1:]
		}
	}
	if len(left) > 0 {
		res = append(res, left...)
	}
	if len(right) > 0 {
		res = append(res, right...)
	}
	return res
}

func HeapSort(data []int) []int {
	for i := len(data)/2 - 1; i >= 0; i-- {
		adjustHeap(data, i, len(data))
	}
	for i := len(data) - 1; i > 0; i-- {
		data[i], data[0] = data[0], data[i]
		adjustHeap(data, 0, i)
	}
	return data
}

func adjustHeap(data []int, i, length int) {
	temp := data[i]
	for k := i*2 + 1; k < length; k = k*2 + 1 {
		if k+1 < length && data[k+1] > data[k] {
			k++
		}
		if data[k] > temp {
			data[i] = data[k]
			i = k
		} else {
			break
		}
	}
	data[i] = temp
}

func RadixSort(data []int) []int {
	maxNum := data[0]
	for _, num := range data {
		if num > maxNum {
			maxNum = num
		}
	}
	exp := 1
	for ; maxNum/exp > 0; exp *= 10 {
		counts := make([]int, 10)

		for _, num := range data {
			counts[(num/exp)%10]++
		}

		for i := 1; i < 10; i++ {
			counts[i] += counts[i-1]
		}

		temp := make([]int, len(data))
		for i := len(data) - 1; i >= 0; i-- {
			counts[(data[i]/exp)%10]--
			temp[counts[(data[i]/exp)%10]] = data[i]
		}

		for i, t := range temp {
			data[i] = t
		}
	}
	return data
}

package utils

import (
	"fmt"
	"reflect"
	"testing"
)

var ints = []int{74, 59, 238, -784, 9845, 959, 905, 0, 0, 42, 7586, -5467984, 7586}
var sorted = []int{-5467984, -784, 0, 0, 42, 59, 74, 238, 905, 959, 7586, 7586, 9845}

func copyInts(a []int) []int {
	b := make([]int, len(a))
	for i, num := range a {
		b[i] = num
	}
	return b
}

func TestSort(t *testing.T) {
	var arr []int
	arr = copyInts(ints)
	if !reflect.DeepEqual(sorted, BubbleSort(arr)) {
		panic(fmt.Errorf("BubbleSort not equal: %v", arr))
	}
	arr = copyInts(ints)
	if !reflect.DeepEqual(sorted, SelectionSort(arr)) {
		panic(fmt.Errorf("SelectionSort not equal: %v", arr))
	}
	arr = copyInts(ints)
	if !reflect.DeepEqual(sorted, InsertionSort(arr)) {
		panic(fmt.Errorf("InsertionSort not equal: %v", arr))
	}
	arr = copyInts(ints)
	if !reflect.DeepEqual(sorted, ShellSort(arr)) {
		panic(fmt.Errorf("ShellSort not equal: %v", arr))
	}
	arr = copyInts(ints)
	if !reflect.DeepEqual(sorted, QuickSort(arr)) {
		panic(fmt.Errorf("QuickSort not equal: %v", arr))
	}
	arr = copyInts(ints)
	if !reflect.DeepEqual(sorted, MergeSort(arr)) {
		panic(fmt.Errorf("MergeSort not equal: %v", arr))
	}
	arr = copyInts(ints)
	if !reflect.DeepEqual(sorted, HeapSort(arr)) {
		panic(fmt.Errorf("HeapSort not equal: %v", arr))
	}
	//arr = copyInts(ints)
	//if !reflect.DeepEqual(sorted, RadixSort(arr)) {
	//	panic(fmt.Errorf("RadixSort not equal: %v", arr))
	//}
}

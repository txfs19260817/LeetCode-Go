package _922_Sort_Array_By_Parity_II

func sortArrayByParityII(A []int) []int {
	odds, evens := make([]int, 0, len(A)/2), make([]int, 0, len(A)/2)
	for _, a := range A {
		if a%2 == 0 {
			evens = append(evens, a)
		} else {
			odds = append(odds, a)
		}
	}
	odds = append(odds, evens...)
	for i, num := range odds {
		A[mappedIndex(i, len(A))] = num
	}
	return A
}

func mappedIndex(i, n int) int {
	return (1 + 2*i) % (n | 1)
}

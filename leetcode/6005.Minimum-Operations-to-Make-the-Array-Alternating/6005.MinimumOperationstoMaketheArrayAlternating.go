package _005_Minimum_Operations_to_Make_the_Array_Alternating

func minimumOperations(nums []int) int {
	cntOdd, cntEven := map[int]int{}, map[int]int{}
	for i, num := range nums {
		if i%2 == 1 {
			cntOdd[num]++
		} else {
			cntEven[num]++
		}
	}
	kOdd, maxOdd := findMaxCntFromMap(cntOdd)
	kEven, maxEven := findMaxCntFromMap(cntEven)
	if kOdd != kEven {
		return len(nums) - maxEven - maxOdd
	}
	delete(cntOdd, kOdd)
	delete(cntEven, kEven)
	_, secondOdd := findMaxCntFromMap(cntOdd)
	_, secondEven := findMaxCntFromMap(cntEven)
	if maxOdd+secondEven > maxEven+secondOdd {
		return len(nums) - maxOdd - secondEven
	}
	return len(nums) - maxEven - secondOdd
}

func findMaxCntFromMap(m map[int]int) (key, c int) {
	for k, v := range m {
		if v > c {
			key, c = k, v
		}
	}
	return
}

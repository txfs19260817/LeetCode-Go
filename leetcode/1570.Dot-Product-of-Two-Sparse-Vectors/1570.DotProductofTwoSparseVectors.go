package leetcode

type SparseVector struct {
	i2v map[int]int
}

func Constructor(nums []int) SparseVector {
	m := map[int]int{}
	for i, num := range nums {
		if num != 0 {
			m[i] = num
		}
	}
	return SparseVector{i2v: m}
}

// Return the dotProduct of two sparse vectors
func (this *SparseVector) dotProduct(vec SparseVector) int {
	var ans int
	if len(vec.i2v) < len(this.i2v) { // traverse the "shorter" vec
		for i, num := range vec.i2v {
			ans += num * this.i2v[i]
		}
	} else {
		for i, num := range this.i2v {
			ans += num * vec.i2v[i]
		}
	}
	return ans
}

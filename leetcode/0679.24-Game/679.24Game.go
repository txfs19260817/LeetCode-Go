package _679_24_Game

const (
	ADD = iota
	SUBTRACT
	MULTIPLY
	DIVIDE
	eps    = 1e-6
	target = 24
)

func judgePoint24(nums []int) bool {
	f := make([]float64, len(nums))
	for i, num := range nums {
		f[i] = float64(num)
	}
	return dfs(f)
}

func dfs(nums []float64) bool {
	if len(nums) == 0 {
		return false
	}
	if len(nums) == 1 {
		return abs(nums[0]-target) < eps
	}
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if i == j {
				continue
			}
			list := make([]float64, 0, len(nums))
			for k, num := range nums {
				if k != i && k != j {
					list = append(list, num)
				}
			}
			for k := 0; k < 4; k++ {
				switch k {
				case ADD:
					list = append(list, nums[i]+nums[j])
				case SUBTRACT:
					list = append(list, nums[i]-nums[j])
				case MULTIPLY:
					list = append(list, nums[i]*nums[j])
				case DIVIDE:
					if nums[j] < eps {
						continue
					}
					list = append(list, nums[i]/nums[j])
				}
				if dfs(list) {
					return true
				}
				list = list[:len(list)-1]
			}
		}
	}
	return false
}

func abs(a float64) float64 {
	if a < 0 {
		return -a
	}
	return a
}

func judgePoint241(nums []int) bool {
	return ternaryOp(float64(nums[0])+float64(nums[1]), float64(nums[2]), float64(nums[3])) ||
		ternaryOp(float64(nums[0])-float64(nums[1]), float64(nums[2]), float64(nums[3])) ||
		ternaryOp(float64(nums[0])*float64(nums[1]), float64(nums[2]), float64(nums[3])) ||
		ternaryOp(float64(nums[0])/float64(nums[1]), float64(nums[2]), float64(nums[3])) ||
		ternaryOp(float64(nums[1])-float64(nums[0]), float64(nums[2]), float64(nums[3])) ||
		ternaryOp(float64(nums[1])/float64(nums[0]), float64(nums[2]), float64(nums[3])) ||

		ternaryOp(float64(nums[0])+float64(nums[2]), float64(nums[1]), float64(nums[3])) ||
		ternaryOp(float64(nums[0])-float64(nums[2]), float64(nums[1]), float64(nums[3])) ||
		ternaryOp(float64(nums[0])*float64(nums[2]), float64(nums[1]), float64(nums[3])) ||
		ternaryOp(float64(nums[0])/float64(nums[2]), float64(nums[1]), float64(nums[3])) ||
		ternaryOp(float64(nums[2])-float64(nums[0]), float64(nums[1]), float64(nums[3])) ||
		ternaryOp(float64(nums[2])/float64(nums[0]), float64(nums[1]), float64(nums[3])) ||

		ternaryOp(float64(nums[3])+float64(nums[0]), float64(nums[1]), float64(nums[2])) ||
		ternaryOp(float64(nums[3])-float64(nums[0]), float64(nums[1]), float64(nums[2])) ||
		ternaryOp(float64(nums[3])*float64(nums[0]), float64(nums[1]), float64(nums[2])) ||
		ternaryOp(float64(nums[3])/float64(nums[0]), float64(nums[1]), float64(nums[2])) ||
		ternaryOp(float64(nums[0])-float64(nums[3]), float64(nums[1]), float64(nums[2])) ||
		ternaryOp(float64(nums[0])/float64(nums[3]), float64(nums[1]), float64(nums[2])) ||

		ternaryOp(float64(nums[1])+float64(nums[2]), float64(nums[3]), float64(nums[0])) ||
		ternaryOp(float64(nums[1])-float64(nums[2]), float64(nums[3]), float64(nums[0])) ||
		ternaryOp(float64(nums[1])*float64(nums[2]), float64(nums[3]), float64(nums[0])) ||
		ternaryOp(float64(nums[1])/float64(nums[2]), float64(nums[3]), float64(nums[0])) ||
		ternaryOp(float64(nums[2])-float64(nums[1]), float64(nums[3]), float64(nums[0])) ||
		ternaryOp(float64(nums[2])/float64(nums[1]), float64(nums[3]), float64(nums[0])) ||

		ternaryOp(float64(nums[2])+float64(nums[3]), float64(nums[0]), float64(nums[1])) ||
		ternaryOp(float64(nums[2])-float64(nums[3]), float64(nums[0]), float64(nums[1])) ||
		ternaryOp(float64(nums[2])*float64(nums[3]), float64(nums[0]), float64(nums[1])) ||
		ternaryOp(float64(nums[2])/float64(nums[3]), float64(nums[0]), float64(nums[1])) ||
		ternaryOp(float64(nums[3])-float64(nums[2]), float64(nums[0]), float64(nums[1])) ||
		ternaryOp(float64(nums[3])/float64(nums[2]), float64(nums[0]), float64(nums[1])) ||

		ternaryOp(float64(nums[1])+float64(nums[3]), float64(nums[0]), float64(nums[2])) ||
		ternaryOp(float64(nums[1])-float64(nums[3]), float64(nums[0]), float64(nums[2])) ||
		ternaryOp(float64(nums[1])*float64(nums[3]), float64(nums[0]), float64(nums[2])) ||
		ternaryOp(float64(nums[1])/float64(nums[3]), float64(nums[0]), float64(nums[2])) ||
		ternaryOp(float64(nums[3])-float64(nums[1]), float64(nums[0]), float64(nums[2])) ||
		ternaryOp(float64(nums[3])/float64(nums[1]), float64(nums[0]), float64(nums[2]))
}

func ternaryOp(a, b, c float64) bool {
	return binaryOp(a+b, c) ||
		binaryOp(a-b, c) ||
		binaryOp(a*b, c) ||
		binaryOp(a/b, c) ||
		binaryOp(b-a, c) ||
		binaryOp(b/a, c) ||

		binaryOp(b+c, a) ||
		binaryOp(b-c, a) ||
		binaryOp(b*c, a) ||
		binaryOp(b/c, a) ||
		binaryOp(c-b, a) ||
		binaryOp(c/b, a) ||

		binaryOp(a+c, b) ||
		binaryOp(a-c, b) ||
		binaryOp(a*c, b) ||
		binaryOp(a/c, b) ||
		binaryOp(c-a, b) ||
		binaryOp(c/a, b)
}

func binaryOp(a, b float64) bool {
	return a+b < 24+1e-6 && a+b > 24-1e-6 ||
		a-b < 24+1e-6 && a-b > 24-1e-6 ||
		a*b < 24+1e-6 && a*b > 24-1e-6 ||
		a/b < 24+1e-6 && a/b > 24-1e-6 ||
		b-a < 24+1e-6 && b-a > 24-1e-6 ||
		b/a < 24+1e-6 && b/a > 24-1e-6
}

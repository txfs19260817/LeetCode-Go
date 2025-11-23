package leetcode

import (
	"testing"
)

func TestSolution(t *testing.T) {
	// 小工具：给定矩阵，返回一个求名人的函数
	mkSolver := func(matrix [][]bool) func(int) int {
		rel := Relation{matrix: matrix}
		return solution(rel.knows)
	}

	// 辅助函数：构造有名人的矩阵
	buildMatrixWithCelebrity := func(n, celeb int) [][]bool {
		m := make([][]bool, n)
		for i := 0; i < n; i++ {
			m[i] = make([]bool, n)
		}
		for i := 0; i < n; i++ {
			if i == celeb {
				// 名人不认识任何人：整行保持 false
				continue
			}
			for j := 0; j < n; j++ {
				if j == celeb {
					// 所有人都认识名人
					m[i][j] = true
				} else {
					m[i][j] = false
				}
			}
		}
		return m
	}

	// 辅助函数：构造「一定不存在名人」的矩阵（全 false）
	buildMatrixNoCelebrity := func(n int) [][]bool {
		m := make([][]bool, n)
		for i := 0; i < n; i++ {
			m[i] = make([]bool, n)
			for j := 0; j < n; j++ {
				m[i][j] = false
			}
		}
		return m
	}

	// === 边界用例 ===
	t.Run("two people celebrity is 0", func(t *testing.T) {
		// 0 是名人：1 认识 0，0 不认识任何人
		matrix := [][]bool{
			{false, false}, // 0 -> nobody
			{true, false},  // 1 -> 0
		}
		findCelebrity := mkSolver(matrix)
		got := findCelebrity(len(matrix))
		if got != 0 {
			t.Fatalf("got %d, want 0", got)
		}
	})

	t.Run("two people no celebrity", func(t *testing.T) {
		// 互相认识，不存在名人
		matrix := [][]bool{
			{false, true}, // 0 -> 1
			{true, false}, // 1 -> 0
		}
		findCelebrity := mkSolver(matrix)
		got := findCelebrity(len(matrix))
		if got != -1 {
			t.Fatalf("got %d, want -1", got)
		}
	})

	t.Run("everyone knows everyone no celebrity", func(t *testing.T) {
		// 3 个人，每个人都认识其他所有人：没有人满足“名人不认识别人”
		matrix := [][]bool{
			{false, true, true},
			{true, false, true},
			{true, true, false},
		}
		findCelebrity := mkSolver(matrix)
		got := findCelebrity(len(matrix))
		if got != -1 {
			t.Fatalf("got %d, want -1", got)
		}
	})

	t.Run("three people celebrity is 1", func(t *testing.T) {
		// 复用你之前在 main 里那个例子
		matrix := [][]bool{
			{false, true, false},  // 0 -> 1
			{false, false, false}, // 1 -> nobody
			{false, true, false},  // 2 -> 1
		}
		findCelebrity := mkSolver(matrix)
		got := findCelebrity(len(matrix))
		if got != 1 {
			t.Fatalf("got %d, want 1", got)
		}
	})

	t.Run("celebrity is last index", func(t *testing.T) {
		// 4 个人，3 是名人
		matrix := [][]bool{
			{false, false, false, true},  // 0 -> 3
			{false, false, false, true},  // 1 -> 3
			{false, false, false, true},  // 2 -> 3
			{false, false, false, false}, // 3 -> nobody
		}
		findCelebrity := mkSolver(matrix)
		got := findCelebrity(len(matrix))
		if got != 3 {
			t.Fatalf("got %d, want 3", got)
		}
	})

	// === 大输入用例 ===

	t.Run("large n=1000 with celebrity in middle", func(t *testing.T) {
		n := 1000
		celeb := 567
		matrix := buildMatrixWithCelebrity(n, celeb)

		findCelebrity := mkSolver(matrix)
		got := findCelebrity(n)
		if got != celeb {
			t.Fatalf("got %d, want %d", got, celeb)
		}
	})

	t.Run("large n=2000 with no celebrity", func(t *testing.T) {
		n := 2000
		matrix := buildMatrixNoCelebrity(n)

		findCelebrity := mkSolver(matrix)
		got := findCelebrity(n)
		if got != -1 {
			t.Fatalf("got %d, want -1", got)
		}
	})
}

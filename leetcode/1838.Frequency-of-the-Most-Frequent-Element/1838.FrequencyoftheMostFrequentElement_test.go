package leetcode

import (
	"math/rand"
	"slices"
	"testing"
)

func Test_maxFrequency(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example_1",
			args: args{nums: []int{1, 2, 4}, k: 5},
			want: 3,
		},
		{
			name: "example_2",
			args: args{nums: []int{1, 4, 8, 13}, k: 5},
			want: 2,
		},
		{
			name: "example_3",
			args: args{nums: []int{3, 9, 6}, k: 2},
			want: 1,
		},
		{
			name: "single_element",
			args: args{nums: []int{7}, k: 0},
			want: 1,
		},
		{
			name: "already_all_equal_k0",
			args: args{nums: []int{5, 5, 5, 5}, k: 0},
			want: 4,
		},
		{
			name: "k0_with_duplicates",
			args: args{nums: []int{1, 2, 2, 3, 3, 3}, k: 0},
			want: 3,
		},
		{
			name: "large_k_makes_all_equal",
			args: args{nums: []int{1, 2, 3, 4}, k: 100},
			want: 4,
		},
		{
			name: "choose_best_target_not_max_value",
			args: args{nums: []int{1, 2, 2, 2, 100}, k: 3},
			want: 4, // make 1->2
		},
		{
			name: "large_values",
			args: args{nums: []int{999999998, 999999999, 1000000000}, k: 3},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nums1 := slices.Clone(tt.args.nums)
			if got := maxFrequency(nums1, tt.args.k); got != tt.want {
				t.Errorf("maxFrequency() = %v, want %v", got, tt.want)
			}
			nums2 := slices.Clone(tt.args.nums)
			if got := maxFrequency2(nums2, tt.args.k); got != tt.want {
				t.Errorf("maxFrequency2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maxFrequency_randomized_against_bruteforce(t *testing.T) {
	r := rand.New(rand.NewSource(1))
	for tc := 0; tc < 300; tc++ {
		n := 1 + r.Intn(8)
		nums := make([]int, n)
		for i := range nums {
			nums[i] = r.Intn(11) // 0..10
		}
		k := r.Intn(21) // 0..20

		want := bruteMaxFrequency(nums, k)

		if got := maxFrequency(slices.Clone(nums), k); got != want {
			t.Fatalf("maxFrequency mismatch: nums=%v k=%d got=%d want=%d", nums, k, got, want)
		}
		if got := maxFrequency2(slices.Clone(nums), k); got != want {
			t.Fatalf("maxFrequency2 mismatch: nums=%v k=%d got=%d want=%d", nums, k, got, want)
		}
	}
}

func Benchmark_maxFrequency(b *testing.B) {
	benchMaxFrequencyImpl(b, maxFrequency)
}

func Benchmark_maxFrequency2(b *testing.B) {
	benchMaxFrequencyImpl(b, maxFrequency2)
}

func benchMaxFrequencyImpl(b *testing.B, fn func([]int, int) int) {
	cases := []struct {
		name string
		n    int
		k    int
	}{
		{name: "n=32_k=1e2", n: 32, k: 100},
		{name: "n=1e3_k=1e6", n: 1_000, k: 1_000_000},
		{name: "n=2e4_k=1e9", n: 20_000, k: 1_000_000_000},
	}

	r := rand.New(rand.NewSource(2))
	for _, c := range cases {
		base := make([]int, c.n)
		for i := range base {
			base[i] = r.Intn(1_000_000_000)
		}

		b.Run(c.name, func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				nums := slices.Clone(base)
				_ = fn(nums, c.k)
			}
		})
	}
}

func bruteMaxFrequency(nums []int, k int) int {
	// O(n^2 log n) brute (small n): try each value as target (after sorting).
	a := slices.Clone(nums)
	// local sort to avoid depending on production helpers
	for i := 0; i < len(a); i++ {
		for j := i + 1; j < len(a); j++ {
			if a[j] < a[i] {
				a[i], a[j] = a[j], a[i]
			}
		}
	}

	best := 1
	for r := 0; r < len(a); r++ {
		need := 0
		for l := r; l >= 0; l-- {
			need += a[r] - a[l]
			if need > k {
				break
			}
			if freq := r - l + 1; freq > best {
				best = freq
			}
		}
	}
	return best
}

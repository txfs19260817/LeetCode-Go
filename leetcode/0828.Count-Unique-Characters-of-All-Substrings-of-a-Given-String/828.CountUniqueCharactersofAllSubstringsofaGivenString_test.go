package _828_Count_Unique_Characters_of_All_Substrings_of_a_Given_String

import "testing"

func Test_uniqueLetterString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: `s = "LEETCODE"`,
			args: args{"LEETCODE"},
			want: 92,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := uniqueLetterString(tt.args.s); got != tt.want {
				t.Errorf("uniqueLetterString() = %v, want %v", got, tt.want)
			}
			if got := uniqueLetterString2(tt.args.s); got != tt.want {
				t.Errorf("uniqueLetterString2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_uniqueLetterString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		uniqueLetterString("LEETCODELEETCODELEETCODELEETCODELEETCODELEETCODELEETCODELEETCODELEETCODELEETCODELEETCOD")
	}
}

func Benchmark_uniqueLetterString2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		uniqueLetterString2("LEETCODELEETCODELEETCODELEETCODELEETCODELEETCODELEETCODELEETCODELEETCODELEETCODELEETCOD")
	}
}

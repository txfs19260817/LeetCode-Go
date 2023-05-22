package leetcode

import "testing"

func Test_secondHighest(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "dfa12321afd",
			args: args{"dfa12321afd"},
			want: 2,
		},
		{
			name: "abc1111",
			args: args{"abc1111"},
			want: -1,
		},
		{
			name: "ck077",
			args: args{"ck077"},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := secondHighest(tt.args.s); got != tt.want {
				t.Errorf("secondHighest() = %v, want %v", got, tt.want)
			}
		})
	}
}

package leetcode

import "testing"

func Test_numSpecialEquivGroups(t *testing.T) {
	type args struct {
		words []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: `["abc","acb","bac","bca","cab","cba"]`,
			args: args{[]string{"abc", "acb", "bac", "bca", "cab", "cba"}},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numSpecialEquivGroups(tt.args.words); got != tt.want {
				t.Errorf("numSpecialEquivGroups() = %v, want %v", got, tt.want)
			}
		})
	}
}

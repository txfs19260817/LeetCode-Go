package _255_Verify_Preorder_Sequence_in_Binary_Search_Tree

import "testing"

func Test_verifyPreorder(t *testing.T) {
	type args struct {
		preorder []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "[5,2,6,1,3]",
			args: args{[]int{5, 2, 6, 1, 3}},
			want: false,
		},
		{
			name: "[5,2,1,3,6]",
			args: args{[]int{5, 2, 1, 3, 6}},
			want: true,
		},
		{
			name: "[1,3,2]",
			args: args{[]int{1, 3, 2}},
			want: true,
		},
		{
			name: "[2,3,1]",
			args: args{[]int{2, 3, 1}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := verifyPreorder(tt.args.preorder); got != tt.want {
				t.Errorf("verifyPreorder() = %v, want %v", got, tt.want)
			}
		})
	}
}

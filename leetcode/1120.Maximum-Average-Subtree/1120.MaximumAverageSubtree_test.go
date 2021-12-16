package _120_Maximum_Average_Subtree

import "testing"

func Test_maximumAverageSubtree(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "root = [5,6,1]",
			args: args{&TreeNode{5, &TreeNode{Val: 6}, &TreeNode{Val: 1}}},
			want: 6,
		},
		{
			name: "root = [0,null,1]",
			args: args{&TreeNode{0, nil, &TreeNode{Val: 1}}},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumAverageSubtree(tt.args.root); got != tt.want {
				t.Errorf("maximumAverageSubtree() = %v, want %v", got, tt.want)
			}
		})
	}
}

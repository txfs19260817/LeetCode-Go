package _993_Cousins_in_Binary_Tree

import "testing"

func Test_isCousins(t *testing.T) {
	type args struct {
		root *TreeNode
		x    int
		y    int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "root = [1,2,3,4], x = 4, y = 3",
			args: args{&TreeNode{
				Val:   1,
				Left:  &TreeNode{Val: 2, Left: &TreeNode{Val: 4}},
				Right: &TreeNode{Val: 3},
			}, 4, 3},
			want: false,
		},
		{
			name: "root = [1,2,3,null,4,null,5], x = 5, y = 4",
			args: args{&TreeNode{
				Val:   1,
				Left:  &TreeNode{Val: 2, Right: &TreeNode{Val: 4}},
				Right: &TreeNode{Val: 3, Right: &TreeNode{Val: 5}},
			}, 5, 4},
			want: true,
		},
		{
			name: "root = [1,2,3,null,4], x = 2, y = 3",
			args: args{&TreeNode{
				Val:   1,
				Left:  &TreeNode{Val: 2, Right: &TreeNode{Val: 4}},
				Right: &TreeNode{Val: 3},
			}, 2, 3},
			want: false,
		},
		{
			name: "root = [1,2,3,null,4], x = 3, y = 2",
			args: args{&TreeNode{
				Val:   1,
				Left:  &TreeNode{Val: 2, Right: &TreeNode{Val: 4}},
				Right: &TreeNode{Val: 3},
			}, 3, 2},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isCousins(tt.args.root, tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("isCousins() = %v, want %v", got, tt.want)
			}
		})
	}
}

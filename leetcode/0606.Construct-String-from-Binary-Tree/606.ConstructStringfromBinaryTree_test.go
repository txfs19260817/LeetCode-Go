package _606_Construct_String_from_Binary_Tree

import "testing"

func Test_tree2str(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "root = [1,2,3,4]",
			args: args{&TreeNode{
				Val:   1,
				Left:  &TreeNode{Val: 2, Left: &TreeNode{Val: 4}},
				Right: &TreeNode{Val: 3},
			}},
			want: "1(2(4))(3)",
		},
		{
			name: "root = [1,2,3,null,4]",
			args: args{&TreeNode{
				Val:   1,
				Left:  &TreeNode{Val: 2, Right: &TreeNode{Val: 4}},
				Right: &TreeNode{Val: 3},
			}},
			want: "1(2()(4))(3)",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tree2str(tt.args.root); got != tt.want {
				t.Errorf("tree2str() = %v, want %v", got, tt.want)
			}
		})
	}
}

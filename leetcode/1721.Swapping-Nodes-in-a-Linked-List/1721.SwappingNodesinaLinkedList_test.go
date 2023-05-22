package leetcode

import (
	"reflect"
	"testing"
)

func Test_swapNodes(t *testing.T) {
	type args struct {
		head *ListNode
		k    int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "head = [7,9,6,6,7,8,3,0,9,5], k = 5",
			args: args{buildLinkedListFromInts([]int{7, 9, 6, 6, 7, 8, 3, 0, 9, 5}), 5},
			want: buildLinkedListFromInts([]int{7, 9, 6, 6, 8, 7, 3, 0, 9, 5}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := swapNodes(tt.args.head, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("swapNodes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func buildLinkedListFromInts(nums []int) *ListNode {
	head := &ListNode{}
	p := head
	for _, num := range nums {
		node := &ListNode{num, nil}
		p.Next = node
		p = p.Next
	}
	return head.Next
}

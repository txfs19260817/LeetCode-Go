package leetcode

import (
	"reflect"
	"testing"
)

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

func Test_mergeNodes(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "head = [0,3,1,0,4,5,2,0]",
			args: args{buildLinkedListFromInts([]int{0, 3, 1, 0, 4, 5, 2, 0})},
			want: &ListNode{
				Val:  4,
				Next: &ListNode{Val: 11},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeNodes(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeNodes() = %v, want %v", got, tt.want)
			}
		})
	}
}

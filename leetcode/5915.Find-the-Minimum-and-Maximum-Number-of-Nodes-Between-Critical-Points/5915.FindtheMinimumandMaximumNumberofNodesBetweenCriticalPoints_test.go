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

func Test_nodesBetweenCriticalPoints(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "head = [3,1]",
			args: args{buildLinkedListFromInts([]int{3, 1})},
			want: []int{-1, -1},
		},
		{
			name: "head = [2,2,1,3]",
			args: args{buildLinkedListFromInts([]int{2, 2, 1, 3})},
			want: []int{-1, -1},
		},
		{
			name: "head = [5,3,1,2,5,1,2]",
			args: args{buildLinkedListFromInts([]int{5, 3, 1, 2, 5, 1, 2})},
			want: []int{1, 3},
		},
		{
			name: "head = [1,3,2,2,3,2,2,2,7]",
			args: args{buildLinkedListFromInts([]int{1, 3, 2, 2, 3, 2, 2, 2, 7})},
			want: []int{3, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nodesBetweenCriticalPoints(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("nodesBetweenCriticalPoints() = %v, want %v", got, tt.want)
			}
		})
	}
}

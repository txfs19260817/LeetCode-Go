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

func Test_rotateRight(t *testing.T) {
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
			name: "head = [1,2,3,4,5], k = 2",
			args: args{buildLinkedListFromInts([]int{1, 2, 3, 4, 5}), 2},
			want: buildLinkedListFromInts([]int{4, 5, 1, 2, 3}),
		},
		{
			name: "head = [0,1,2], k = 4",
			args: args{buildLinkedListFromInts([]int{0, 1, 2}), 4},
			want: buildLinkedListFromInts([]int{2, 0, 1}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rotateRight(tt.args.head, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("rotateRight() = %v, want %v", got, tt.want)
			}
		})
	}
}

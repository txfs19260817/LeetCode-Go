package leetcode

import (
	"reflect"
	"testing"
)

func Test_reverseEvenLengthGroups(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "head = [5,2,6,3,9,1,7,3,8,4]",
			args: args{buildLinkedListFromInts([]int{5, 2, 6, 3, 9, 1, 7, 3, 8, 4})},
			want: buildLinkedListFromInts([]int{5, 6, 2, 3, 9, 1, 4, 8, 3, 7}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseEvenLengthGroups(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseEvenLengthGroups() = %v, want %v", got, tt.want)
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

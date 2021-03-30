package _445_Add_Two_Numbers_II_

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

func Test_addTwoNumbers(t *testing.T) {
	type args struct {
		l1 *ListNode
		l2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "(7 -> 2 -> 4 -> 3) + (5 -> 6 -> 4)",
			args: args{buildLinkedListFromInts([]int{7, 2, 4, 3}), buildLinkedListFromInts([]int{5, 6, 4})},
			want: buildLinkedListFromInts([]int{7, 8, 0, 7}),
		},
		{
			name: "(7 -> 2 -> 4 -> 3) + (1 -> 5 -> 6 -> 4)",
			args: args{buildLinkedListFromInts([]int{7, 2, 4, 3}), buildLinkedListFromInts([]int{1, 5, 6, 4})},
			want: buildLinkedListFromInts([]int{8, 8, 0, 7}),
		},
		{
			name: "(7 -> 2 -> 4 -> 3) + (4 -> 5 -> 6 -> 4)",
			args: args{buildLinkedListFromInts([]int{7, 2, 4, 3}), buildLinkedListFromInts([]int{4, 5, 6, 4})},
			want: buildLinkedListFromInts([]int{1, 1, 8, 0, 7}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addTwoNumbers(tt.args.l1, tt.args.l2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addTwoNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}

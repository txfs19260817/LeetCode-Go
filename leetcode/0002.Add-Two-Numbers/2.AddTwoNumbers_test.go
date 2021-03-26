package _002_Add_Two_Numbers

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
			name: "l1 = [2,4,3], l2 = [5,6,4]",
			args: args{buildLinkedListFromInts([]int{2, 4, 3}), buildLinkedListFromInts([]int{5, 6, 4})},
			want: buildLinkedListFromInts([]int{7, 0, 8}),
		},
		{
			name: "l1 = [0], l2 = [0]",
			args: args{buildLinkedListFromInts([]int{0}), buildLinkedListFromInts([]int{0})},
			want: buildLinkedListFromInts([]int{0}),
		},
		{
			name: "l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]",
			args: args{buildLinkedListFromInts([]int{9, 9, 9, 9, 9, 9, 9}), buildLinkedListFromInts([]int{9, 9, 9, 9})},
			want: buildLinkedListFromInts([]int{8, 9, 9, 9, 0, 0, 0, 1}),
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

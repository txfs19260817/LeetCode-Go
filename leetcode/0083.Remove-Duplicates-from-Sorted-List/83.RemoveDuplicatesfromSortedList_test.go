package _083_Remove_Duplicates_from_Sorted_List

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

func Test_deleteDuplicates(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "head = [1,1,2]",
			args: args{buildLinkedListFromInts([]int{1, 1, 2})},
			want: buildLinkedListFromInts([]int{1, 2}),
		},
		{
			name: "head = [1,1,2,3,3]",
			args: args{buildLinkedListFromInts([]int{1, 1, 2, 3, 3})},
			want: buildLinkedListFromInts([]int{1, 2, 3}),
		},
		{
			name: "head = [1,1]",
			args: args{buildLinkedListFromInts([]int{1, 1})},
			want: buildLinkedListFromInts([]int{1}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := deleteDuplicates(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("deleteDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}
}

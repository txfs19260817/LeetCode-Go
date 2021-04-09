package _021_Merge_Two_Sorted_Lists

import (
	"reflect"
	"testing"
)

var benchCase = [][]int{
	{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
	{0, 0, 1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 6, 6, 7, 7, 8, 8, 9, 9, 10},
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

func Test_mergeTwoLists(t *testing.T) {
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
			name: "l1 = [1,2,4], l2 = [1,3,4]",
			args: args{buildLinkedListFromInts([]int{1, 2, 4}), buildLinkedListFromInts([]int{1, 3, 4})},
			want: buildLinkedListFromInts([]int{1, 1, 2, 3, 4, 4}),
		},
		{
			name: "l1 = [], l2 = []",
			args: args{nil, nil},
			want: nil,
		},
		{
			name: "l1 = [], l2 = [0]",
			args: args{nil, &ListNode{0, nil}},
			want: &ListNode{0, nil},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeTwoLists(tt.args.l1, tt.args.l2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeTwoLists() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mergeTwoLists1(t *testing.T) {
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
			name: "l1 = [1,2,4], l2 = [1,3,4]",
			args: args{buildLinkedListFromInts([]int{1, 2, 4}), buildLinkedListFromInts([]int{1, 3, 4})},
			want: buildLinkedListFromInts([]int{1, 1, 2, 3, 4, 4}),
		},
		{
			name: "l1 = [], l2 = []",
			args: args{nil, nil},
			want: nil,
		},
		{
			name: "l1 = [], l2 = [0]",
			args: args{nil, &ListNode{0, nil}},
			want: &ListNode{0, nil},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeTwoLists1(tt.args.l1, tt.args.l2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeTwoLists1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_mergeTwoLists(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mergeTwoLists(buildLinkedListFromInts(benchCase[0]), buildLinkedListFromInts(benchCase[1]))
	}
}

func Benchmark_mergeTwoLists1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mergeTwoLists1(buildLinkedListFromInts(benchCase[0]), buildLinkedListFromInts(benchCase[1]))
	}
}

package leetcode

import (
	"reflect"
	"testing"
)

type firstUniqueCase struct {
	name   string
	ops    []string
	args   [][]int
	expect []interface{}
}

func runFirstUniqueCase(t *testing.T, tc firstUniqueCase) {
	t.Helper()

	var fu FirstUnique
	actual := make([]interface{}, 0, len(tc.ops))

	for i, op := range tc.ops {
		switch op {
		case "FirstUnique":
			fu = Constructor(tc.args[i])
			actual = append(actual, nil)
		case "showFirstUnique":
			actual = append(actual, fu.ShowFirstUnique())
		case "add":
			fu.Add(tc.args[i][0])
			actual = append(actual, nil)
		default:
			t.Fatalf("unsupported op %q", op)
		}
	}

	if !reflect.DeepEqual(actual, tc.expect) {
		t.Fatalf("case %q got %v, want %v", tc.name, actual, tc.expect)
	}
}

func TestFirstUniqueExamples(t *testing.T) {
	tests := []firstUniqueCase{
		{
			name: "example-1",
			ops: []string{
				"FirstUnique", "showFirstUnique", "add", "showFirstUnique",
				"add", "showFirstUnique", "add", "showFirstUnique",
			},
			args: [][]int{
				{2, 3, 5}, {}, {5}, {}, {2}, {}, {3}, {},
			},
			expect: []interface{}{
				nil, 2, nil, 2, nil, 3, nil, -1,
			},
		},
		{
			name: "example-2",
			ops: []string{
				"FirstUnique", "showFirstUnique", "add", "add", "add", "add", "add", "showFirstUnique",
			},
			args: [][]int{
				{7, 7, 7, 7, 7, 7}, {}, {7}, {3}, {3}, {7}, {17}, {},
			},
			expect: []interface{}{
				nil, -1, nil, nil, nil, nil, nil, 17,
			},
		},
		{
			name: "example-3",
			ops: []string{
				"FirstUnique", "showFirstUnique", "add", "showFirstUnique",
			},
			args: [][]int{
				{809}, {}, {809}, {},
			},
			expect: []interface{}{
				nil, 809, nil, -1,
			},
		},
		{
			name: "example-4",
			ops: []string{
				"FirstUnique", "add", "add", "showFirstUnique",
			},
			args: [][]int{
				{1}, {1}, {1}, {},
			},
			expect: []interface{}{
				nil, nil, nil, -1,
			},
		},
	}

	for _, tc := range tests {
		runFirstUniqueCase(t, tc)
	}
}

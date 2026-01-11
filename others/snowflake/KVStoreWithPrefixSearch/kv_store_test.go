package snowflake

import (
	"reflect"
	"sort"
	"testing"
)

func TestKVStoreOperations(t *testing.T) {
	tests := []struct {
		name     string
		commands []string
		args     []any
		expected []any
	}{
		{
			name:     "Test Case 1",
			commands: []string{"KVStore", "set", "set", "get", "update", "prefixSearch", "deleteKey", "get", "prefixSearch"},
			args: []any{
				[]string{},
				[]any{"aaple", 3},
				[]any{"aap", 2},
				[]string{"aap"},
				[]any{"aap", 5},
				[]string{"aap"},
				[]string{"aap"},
				[]string{"aap"},
				[]string{"aaple"},
			},
			expected: []any{nil, nil, nil, 2, nil, []int{5, 3}, nil, -1, []int{3}},
		},
		{
			name:     "Test Case 2",
			commands: []string{"KVStore", "set", "set", "set", "prefixSearch", "update", "prefixSearch", "deleteKey", "set", "prefixSearch", "get"},
			args: []any{
				[]string{},
				[]any{"foo", 1},
				[]any{"bar", 2},
				[]any{"foobar", 3},
				[]string{"foo"},
				[]any{"foo", 4},
				[]string{"foo"},
				[]string{"bar"},
				[]any{"baz", 5},
				[]string{"ba"},
				[]string{"bar"},
			},
			expected: []any{nil, nil, nil, nil, []int{1, 3}, nil, []int{4, 3}, nil, nil, []int{5}, -1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var obj KVStore
			for i, cmd := range tt.commands {
				arg := tt.args[i]
				want := tt.expected[i]

				switch cmd {
				case "KVStore":
					obj = NewKVStore()
				case "set":
					args := arg.([]any)
					obj.Set(args[0].(string), args[1].(int))
				case "get":
					got := obj.Get(arg.([]string)[0])
					if want != nil && got != want.(int) {
						t.Errorf("step %d: Get(%s) = %d, want %d", i, arg.([]string)[0], got, want)
					}
				case "update":
					args := arg.([]any)
					obj.Update(args[0].(string), args[1].(int))
				case "deleteKey":
					obj.DeleteKey(arg.([]string)[0])
				case "prefixSearch":
					got := obj.PrefixSearch(arg.([]string)[0])
					wantSlice := want.([]int)
					sort.Ints(got)
					sort.Ints(wantSlice)
					if !reflect.DeepEqual(got, wantSlice) {
						t.Errorf("step %d: PrefixSearch(%s) = %v, want %v", i, arg.([]string)[0], got, wantSlice)
					}
				}
			}
		})
	}
}

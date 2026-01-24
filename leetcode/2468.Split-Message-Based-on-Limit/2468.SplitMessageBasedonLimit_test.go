package leetcode

import (
	"reflect"
	"testing"
)

func Test_splitMessage(t *testing.T) {
	type args struct {
		message string
		limit   int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "Example 1",
			args: args{
				message: "this is really a very awesome message",
				limit:   9,
			},
			want: []string{"thi<1/14>", "s i<2/14>", "s r<3/14>", "eal<4/14>", "ly <5/14>", "a v<6/14>", "ery<7/14>", " aw<8/14>", "eso<9/14>", "me<10/14>", " m<11/14>", "es<12/14>", "sa<13/14>", "ge<14/14>"},
		},
		{
			name: "Example 2",
			args: args{
				message: "short message",
				limit:   15,
			},
			want: []string{"short mess<1/2>", "age<2/2>"},
		},
		{
			name: "Example 3",
			args: args{
				message: "baaaababab aabaaba",
				limit:   7,
			},
			want: []string{"ba<1/9>", "aa<2/9>", "ab<3/9>", "ab<4/9>", "ab<5/9>", " a<6/9>", "ab<7/9>", "aa<8/9>", "ba<9/9>"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := splitMessage(tt.args.message, tt.args.limit); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("splitMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}

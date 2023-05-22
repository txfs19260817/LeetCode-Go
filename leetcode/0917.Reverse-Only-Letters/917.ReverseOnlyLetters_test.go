package leetcode

import "testing"

func Test_reverseOnlyLetters(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "ab-cd",
			args: args{"ab-cd"},
			want: "dc-ba",
		},
		{
			name: "a-bC-dEf-ghIj",
			args: args{"a-bC-dEf-ghIj"},
			want: "j-Ih-gfE-dCba",
		},
		{
			name: "Test1ng-Leet=code-Q!",
			args: args{"Test1ng-Leet=code-Q!"},
			want: "Qedo1ct-eeLg=ntse-T!",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseOnlyLetters(tt.args.s); got != tt.want {
				t.Errorf("reverseOnlyLetters() = %v, want %v", got, tt.want)
			}
		})
	}
}

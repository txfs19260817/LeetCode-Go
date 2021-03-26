package _150_Evaluate_Reverse_Polish_Notation

import "testing"

func Test_evalRPN(t *testing.T) {
	type args struct {
		tokens []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: `"2","1","+","3","*"`,
			args: args{[]string{"2", "1", "+", "3", "*"}},
			want: 9,
		},
		{
			name: `"4","13","5","/","+"`,
			args: args{[]string{"4", "13", "5", "/", "+"}},
			want: 6,
		},
		{
			name: `"10","6","9","3","+","-11","*","/","*","17","+","5","+"`,
			args: args{[]string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}},
			want: 22,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := evalRPN(tt.args.tokens); got != tt.want {
				t.Errorf("evalRPN() = %v, want %v", got, tt.want)
			}
		})
	}
}

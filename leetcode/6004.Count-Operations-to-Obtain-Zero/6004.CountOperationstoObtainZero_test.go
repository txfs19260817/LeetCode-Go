package _004_Count_Operations_to_Obtain_Zero

import "testing"

func Test_countOperations(t *testing.T) {
	type args struct {
		num1 int
		num2 int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "2,3",
			args: args{2, 3},
			want: 3,
		},
		{
			name: "10,10",
			args: args{10, 10},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countOperations(tt.args.num1, tt.args.num2); got != tt.want {
				t.Errorf("countOperations() = %v, want %v", got, tt.want)
			}
		})
	}
}

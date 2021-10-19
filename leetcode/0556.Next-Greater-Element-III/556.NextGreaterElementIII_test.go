package _556_Next_Greater_Element_III

import "testing"

func Test_nextGreaterElement(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "12",
			args: args{12},
			want: 21,
		},
		{
			name: "1321",
			args: args{1321},
			want: 2113,
		},
		{
			name: "21",
			args: args{21},
			want: -1,
		},
		{
			name: "2147483486",
			args: args{2147483486},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nextGreaterElement(tt.args.n); got != tt.want {
				t.Errorf("nextGreaterElement() = %v, want %v", got, tt.want)
			}
		})
	}
}

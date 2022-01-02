package _970_Maximum_Employees_to_Be_Invited_to_a_Meeting

import "testing"

func Test_maximumInvitations(t *testing.T) {
	type args struct {
		favorite []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "favorite = [2,2,1,2]",
			args: args{[]int{2, 2, 1, 2}},
			want: 3,
		},
		{
			name: "favorite = [1,2,0]",
			args: args{[]int{1, 2, 0}},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumInvitations(tt.args.favorite); got != tt.want {
				t.Errorf("maximumInvitations() = %v, want %v", got, tt.want)
			}
		})
	}
}

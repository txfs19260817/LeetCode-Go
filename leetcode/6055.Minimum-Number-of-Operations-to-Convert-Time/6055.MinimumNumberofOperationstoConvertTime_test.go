package _055_Minimum_Number_of_Operations_to_Convert_Time

import "testing"

func Test_convertTime(t *testing.T) {
	type args struct {
		current string
		correct string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: `current = "02:30", correct = "04:35"`,
			args: args{"02:30", "04:35"},
			want: 3,
		},
		{
			name: `current = "11:00", correct = "11:01"`,
			args: args{"11:00", "11:01"},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertTime(tt.args.current, tt.args.correct); got != tt.want {
				t.Errorf("convertTime() = %v, want %v", got, tt.want)
			}
		})
	}
}

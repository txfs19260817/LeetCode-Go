package _252_Meeting_Rooms

import "testing"

func Test_canAttendMeetings(t *testing.T) {
	type args struct {
		intervals [][]int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "intervals = [[0,30],[5,10],[15,20]]",
			args: args{[][]int{{0, 30}, {5, 10}, {15, 20}}},
			want: false,
		},
		{
			name: "intervals = [[7,10],[2,4]]",
			args: args{[][]int{{7, 10}, {2, 4}}},
			want: true,
		},
		{
			name: "intervals = [[1,5],[8,9]]",
			args: args{[][]int{{1, 5}, {8, 9}}},
			want: true,
		},
		{
			name: "intervals = []",
			args: args{nil},
			want: true,
		},
		{
			name: "intervals = [[465,497],[386,462],[354,380],[134,189],[199,282],[18,104],[499,562],[4,14],[111,129],[292,345]]",
			args: args{[][]int{{465, 497}, {386, 462}, {354, 380}, {134, 189}, {199, 282}, {18, 104}, {499, 562}, {4, 14}, {111, 129}, {292, 345}}},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canAttendMeetings(tt.args.intervals); got != tt.want {
				t.Errorf("canAttendMeetings() = %v, want %v", got, tt.want)
			}
			if got := canAttendMeetings2(tt.args.intervals); got != tt.want {
				t.Errorf("canAttendMeetings2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_canAttendMeetings(b *testing.B) {
	for i := 0; i < b.N; i++ {
		canAttendMeetings([][]int{{465, 497}, {386, 462}, {354, 380}, {134, 189}, {199, 282}, {18, 104}, {499, 562}, {4, 14}, {111, 129}, {292, 345}})
	}
}

func Benchmark_canAttendMeetings2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		canAttendMeetings2([][]int{{465, 497}, {386, 462}, {354, 380}, {134, 189}, {199, 282}, {18, 104}, {499, 562}, {4, 14}, {111, 129}, {292, 345}})
	}
}

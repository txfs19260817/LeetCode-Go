package _841_Keys_and_Rooms

import "testing"

func Test_canVisitAllRooms(t *testing.T) {
	type args struct {
		rooms [][]int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "rooms = [[1],[2],[3],[]]",
			args: args{[][]int{{1}, {2}, {3}, {}}},
			want: true,
		},
		{
			name: "rooms = [[1,3],[3,0,1],[2],[0]]",
			args: args{[][]int{{1, 3}, {3, 0, 1}, {2}, {0}}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canVisitAllRooms(tt.args.rooms); got != tt.want {
				t.Errorf("canVisitAllRooms() = %v, want %v", got, tt.want)
			}
		})
	}
}

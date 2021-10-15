package _690_Employee_Importance

import "testing"

func Test_getImportance(t *testing.T) {
	type args struct {
		employees []*Employee
		id        int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "[[1,2,[2]], [2,3,[]]], 2",
			args: args{[]*Employee{{1, 2, []int{2}}, {2, 3, nil}}, 2},
			want: 3,
		},
		{
			name: "[[1,2,[2]], [2,3,[]]], 3",
			args: args{[]*Employee{{1, 2, []int{2}}, {2, 3, nil}}, 3},
			want: 0,
		},
		{
			name: "[[1,5,[2,3]],[2,3,[]],[3,3,[]]],1",
			args: args{[]*Employee{{1, 5, []int{2, 3}}, {2, 3, nil}, {3, 3, nil}}, 1},
			want: 11,
		},
		{
			name: "[[1,15,[2]], [2,10,[3]], [3,5,[]]], 1",
			args: args{[]*Employee{{1, 15, []int{2}}, {2, 10, []int{3}}, {3, 5, nil}}, 1},
			want: 30,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getImportance(tt.args.employees, tt.args.id); got != tt.want {
				t.Errorf("getImportance() = %v, want %v", got, tt.want)
			}
		})
	}
}

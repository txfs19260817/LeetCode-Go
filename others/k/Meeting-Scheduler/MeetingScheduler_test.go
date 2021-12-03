package k

import (
	"reflect"
	"testing"
)

var p1_meetings = [][]int{{1230, 1300}, {845, 900}, {1300, 1500}}

var p2_meetings = [][]int{{0, 844}, {930, 1200}, {1515, 1546}, {1600, 2400}}

var p3_meetings = [][]int{{845, 915}, {1515, 1545}, {1235, 1245}}

var p4_meetings = [][]int{{1, 5}, {844, 900}, {1515, 1600}}

func Test_spareTime(t *testing.T) {
	type args struct {
		meetings [][][]int
	}
	tests := []struct {
		name    string
		args    args
		wantAns [][]int
	}{
		{
			name:    "123",
			args:    args{[][][]int{p1_meetings, p2_meetings, p3_meetings}},
			wantAns: [][]int{{844, 845}, {915, 930}, {1200, 1230}, {1500, 1515}, {1546, 1600}},
		},
		{
			name:    "13",
			args:    args{[][][]int{p1_meetings, p3_meetings}},
			wantAns: [][]int{{0, 845}, {915, 1230}, {1500, 1515}, {1545, 2400}},
		},
		{
			name:    "24",
			args:    args{[][][]int{p2_meetings, p4_meetings}},
			wantAns: [][]int{{900, 930}, {1200, 1515}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := spareTime(tt.args.meetings); !reflect.DeepEqual(gotAns, tt.wantAns) {
				t.Errorf("spareTime() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}

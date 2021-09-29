package _721_Accounts_Merge

import (
	"reflect"
	"testing"
)

func Test_accountsMerge(t *testing.T) {
	type args struct {
		accounts [][]string
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{
			name: "Fern",
			args: args{[][]string{{"Fern","Fern8@m.co","Fern9@m.co"},{"Fern","Fern7@m.co","Fern8@m.co"},{"Fern","Fern4@m.co","Fern5@m.co"},{"Fern","Fern10@m.co","Fern11@m.co"},{"Fern","Fern9@m.co","Fern10@m.co"},{"Fern","Fern6@m.co","Fern7@m.co"},{"Fern","Fern1@m.co","Fern2@m.co"},{"Fern","Fern11@m.co","Fern12@m.co"},{"Fern","Fern3@m.co","Fern4@m.co"},{"Fern","Fern2@m.co","Fern3@m.co"},{"Fern","Fern5@m.co","Fern6@m.co"},{"Fern","Fern0@m.co","Fern1@m.co"}}},
			want: [][]string{{"Fern","Fern0@m.co","Fern10@m.co","Fern11@m.co","Fern12@m.co","Fern1@m.co","Fern2@m.co","Fern3@m.co","Fern4@m.co","Fern5@m.co","Fern6@m.co","Fern7@m.co","Fern8@m.co","Fern9@m.co"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := accountsMerge(tt.args.accounts); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("accountsMerge() = %v, want %v", got, tt.want)
			}
		})
	}
}

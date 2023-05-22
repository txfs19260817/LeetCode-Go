package leetcode

import "testing"

func Test_maxProfit(t *testing.T) {
	type args struct {
		inventory []int
		orders    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "inventory = [3,5], orders = 6",
			args: args{[]int{3, 5}, 6},
			want: 19,
		},
		{
			name: "inventory = [497978859,167261111,483575207,591815159], orders = 836556809",
			args: args{[]int{497978859, 167261111, 483575207, 591815159}, 836556809},
			want: 373219333,
		},
		{
			name: "inventory = [215304442, ...], orders = 736018741",
			args: args{[]int{215304442, 215520137, 789340189, 671515111, 162866516, 34775245, 440362748, 807594688, 771805560, 457696035, 933697922, 74781394, 55203629, 478443718, 61254303, 764025706, 813169464, 319568349, 209224238, 938193904, 463246620, 894931258, 294930990, 974756402, 50124566, 216580040, 728203093, 991462942, 210175195, 694836099, 375598500, 182268382, 787941805, 129316712, 351790475, 78646550, 361370990, 554717209, 752316728, 447978710, 146899252, 583712441, 529216382, 961753636, 758773559}, 736018741},
			want: 458631340,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfit(tt.args.inventory, tt.args.orders); got != tt.want {
				t.Errorf("maxProfit() = %v, want %v", got, tt.want)
			}
		})
	}
}

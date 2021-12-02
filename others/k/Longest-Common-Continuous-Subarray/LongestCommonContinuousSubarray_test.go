package k

import (
	"reflect"
	"testing"
)

func Test_longestCommonContinuousHistory(t *testing.T) {
	type args struct {
		history1 []string
		history2 []string
	}
	tests := []struct {
		name    string
		args    args
		wantAns []string
	}{
		{
			name:    "1",
			args:    args{[]string{"3234.html", "xys.html", "7hsaa.html"}, []string{"3234.html", "sdhsfjdsh.html", "xys.html", "7hsaa.html"}},
			wantAns: []string{"xys.html", "7hsaa.html"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := longestCommonContinuousHistory(tt.args.history1, tt.args.history2); !reflect.DeepEqual(gotAns, tt.wantAns) {
				t.Errorf("longestCommonContinuousHistory() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}

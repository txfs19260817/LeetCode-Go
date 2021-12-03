package k

import (
	"reflect"
	"testing"
)

func Test_reflowAndJustify(t *testing.T) {
	type args struct {
		lines  []string
		maxLen int
	}
	tests := []struct {
		name    string
		args    args
		wantAns []string
	}{
		{
			name:    "1",
			args:    args{[]string{"The day began as still as the", "night abruptly lighted with", "brilliant flame"}, 24},
			wantAns: []string{"The--day--began-as-still", "as--the--night--abruptly", "lighted--with--brilliant", "flame"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := reflowAndJustify(tt.args.lines, tt.args.maxLen); !reflect.DeepEqual(gotAns, tt.wantAns) {
				t.Errorf("reflowAndJustify() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}

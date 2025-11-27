package k

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_mismatches(t *testing.T) {
	type args struct {
		records [][]string
	}
	tests := []struct {
		name    string
		args    args
		wantAns map[string][]string
	}{
		{
			name: "Records 1 (Complex)",
			args: args{
				records: [][]string{
					{"Paul", "enter"},
					{"Pauline", "exit"},
					{"Paul", "enter"},
					{"Paul", "exit"},
					{"Martha", "exit"},
					{"Joe", "enter"},
					{"Martha", "enter"},
					{"Steve", "enter"},
					{"Martha", "exit"},
					{"Jennifer", "enter"},
					{"Joe", "enter"},
					{"Curtis", "exit"},
					{"Curtis", "enter"},
					{"Joe", "exit"},
					{"Martha", "enter"},
					{"Martha", "exit"},
					{"Jennifer", "exit"},
					{"Joe", "enter"},
					{"Joe", "enter"},
					{"Martha", "exit"},
					{"Joe", "exit"},
					{"Joe", "exit"},
				},
			},
			wantAns: map[string][]string{
				"enter": {"Steve", "Curtis", "Paul", "Joe"},
				"exit":  {"Martha", "Pauline", "Curtis", "Joe"},
			},
		},
		{
			name: "Records 2 (Empty)",
			args: args{
				records: [][]string{
					{"Paul", "enter"},
					{"Paul", "exit"},
				},
			},
			wantAns: map[string][]string{
				"enter": {},
				"exit":  {},
			},
		},
		{
			name: "Records 3 (Double Enter/Exit)",
			args: args{
				records: [][]string{
					{"Paul", "enter"},
					{"Paul", "enter"},
					{"Paul", "exit"},
					{"Paul", "exit"},
				},
			},
			wantAns: map[string][]string{
				"enter": {"Paul"},
				"exit":  {"Paul"},
			},
		},
		{
			name: "Records 4 (Mixed)",
			args: args{
				records: [][]string{
					{"Raj", "enter"},
					{"Paul", "enter"},
					{"Paul", "exit"},
					{"Paul", "exit"},
					{"Paul", "enter"},
					{"Raj", "enter"},
				},
			},
			wantAns: map[string][]string{
				"enter": {"Raj", "Paul"},
				"exit":  {"Paul"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotAns := mismatches(tt.args.records)
			assert.ElementsMatch(t, tt.wantAns["enter"], gotAns["enter"])
			assert.ElementsMatch(t, tt.wantAns["exit"], gotAns["exit"])
		})
	}
}

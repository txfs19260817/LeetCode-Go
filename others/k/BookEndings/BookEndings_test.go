package k

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findMostReadPage(t *testing.T) {
	type args struct {
		endings []int
		choices [][]int
	}
	tests := []struct {
		name         string
		args         args
		mostReadPage int
		reads        int
	}{
		{
			name: "example1",
			args: args{
				endings: []int{5, 10},
				choices: [][]int{{3, 7, 9}, {9, 10, 8}},
			},
			mostReadPage: 9,
			reads:        6,
		},
		{
			name: "singleEndingStart",
			args: args{
				endings: []int{1},
				choices: [][]int{},
			},
			mostReadPage: 1,
			reads:        1,
		},
		{
			name: "LoopCase",
			args: args{
				endings: []int{3},
				choices: [][]int{{2, 1, 3}},
			},
			mostReadPage: 1, // 1 or 2
			reads:        3,
		},
		{
			name: "InvalidStoryline",
			args: args{
				endings: []int{4},
				choices: [][]int{{2, 1, 1}},
			},
			mostReadPage: -1,
			reads:        0,
		},
		{
			name: "StraightLine",
			args: args{
				endings: []int{5},
				choices: [][]int{},
			},
			mostReadPage: 1, // any
			reads:        1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotMostReadPage, gotReads := findMostReadPage(tt.args.endings, tt.args.choices)
			if tt.name == "LoopCase" {
				assert.Equal(t, 3, gotReads)
				assert.Contains(t, []int{1, 2}, gotMostReadPage)
				return
			}
			if tt.name == "StraightLine" {
				assert.Equal(t, 1, gotReads)
				assert.GreaterOrEqual(t, gotMostReadPage, 1)
				assert.LessOrEqual(t, gotMostReadPage, 5)
				return
			}
			assert.Equal(t, tt.mostReadPage, gotMostReadPage, "expected most read page to be %d, got %d", tt.mostReadPage, gotMostReadPage)
			assert.Equal(t, tt.reads, gotReads, "expected reads to be %d, got %d", tt.reads, gotReads)
		})
	}
}

func Test_validMoves(t *testing.T) {
	type args struct {
		start []string
		end   []string
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{
			name: "example1",
			args: args{
				start: []string{"R", "_", "B", "B"},
				end:   []string{"B", "_", "B", "R"},
			},
			want: [][]string{
				{"R", "_", "B", "B"},
				{"_", "R", "B", "B"},
				{"B", "R", "_", "B"},
				{"B", "R", "B", "_"},
				{"B", "_", "B", "R"},
			},
		},
		{
			name: "example2",
			args: args{
				start: []string{"R", "R", "_"},
				end:   []string{"_", "R", "R"},
			},
			want: [][]string{
				{"R", "R", "_"},
				{"R", "_", "R"},
				{"_", "R", "R"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, validMoves(tt.args.start, tt.args.end), "validMoves(%v, %v)", tt.args.start, tt.args.end)
		})
	}
}

func Test_sightseeing(t *testing.T) {
	type args struct {
		trails      [][]string
		attractions []string
	}
	trails1 := [][]string{
		{"Beaver Dam", "Frozen Ocean"},
		{"Beaver Dam", "Frozen Ocean"},
		{"Parking Lot", "Beaver Dam"},
		{"Parking Lot", "Liberty Lake"},
		{"Beaver Dam", "Campsite"},
		{"Eel Weir", "Campsite"},
		{"Eel Weir", "Campsite"},
	}
	trails2 := [][]string{
		{"Mason's Cabin", "Liberty Lake"},
		{"Parking Lot", "Mill Falls"},
		{"Mason's Cabin", "Jeremy's Bay"},
		{"Eel Weir", "Hardwood Forest"},
		{"Outdoor Theater", "Campsite"},
		{"Jeremy's Bay", "Horseshoe Falls"},
		{"Mason's Cabin", "Parking Lot"},
		{"Mason's Cabin", "Liberty Lake"},
		{"Mill Falls", "Horseshoe Falls"},
		{"Mill Falls", "Eel Weir"},
		{"Hardwood Forest", "Campsite"},
		{"Eel Weir", "Outdoor Theater"},
		{"Liberty Lake", "Mason's Cabin"},
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "trails1_1",
			args: args{
				trails:      trails1,
				attractions: []string{"Frozen Ocean"},
			},
			want: true,
		},
		{
			name: "trails1_2",
			args: args{
				trails:      trails1,
				attractions: []string{"Liberty Lake", "Beaver Dam"},
			},
			want: false,
		},
		{
			name: "trails1_3",
			args: args{
				trails:      trails1,
				attractions: []string{"Eel Weir"},
			},
			want: true,
		},
		{
			name: "trails2_1",
			args: args{
				trails:      trails2,
				attractions: []string{"Jeremy's Bay", "Mason's Cabin", "Outdoor Theater"},
			},
			want: true,
		},
		{
			name: "trails2_2",
			args: args{
				trails:      trails2,
				attractions: []string{"Outdoor Theater", "Eel Weir", "Hardwood Forest"},
			},
			want: false,
		},
		{
			name: "trails2_3",
			args: args{
				trails:      trails2,
				attractions: []string{"Liberty Lake"},
			},
			want: true,
		},
		{
			name: "trails2_4",
			args: args{
				trails:      trails2,
				attractions: []string{"Horseshoe Falls", "Eel Weir"},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, sightseeing(tt.args.trails, tt.args.attractions), "sightseeing(%v, %v)", tt.args.trails, tt.args.attractions)
		})
	}
}

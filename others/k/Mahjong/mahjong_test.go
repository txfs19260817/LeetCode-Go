package k

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_isCompleteHand(t *testing.T) {
	tests := []struct {
		name  string
		tiles string
		want  bool
	}{
		{"tiles1", "11133555", true},
		{"tiles2", "111333555", false},
		{"tiles3", "00000111", true},
		{"tiles4", "13233121", true},
		{"tiles5", "11223344555", false},
		{"tiles6", "99999999", true},
		{"tiles7", "999999999", false},
		{"tiles8", "9", false},
		{"tiles9", "99", true},
		{"tiles10", "000022", false},
		{"tiles11", "888889", false},
		{"tiles12", "889", false},
		{"tiles13", "88888844", true},
		{"tiles14", "77777777777777", true},
		{"tiles15", "1111111", false},
		{"tiles16", "1111122222", false},
		{"tiles17", "2244", false},
		{"tiles18", "22222", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, isCompleteHand(tt.tiles))
		})
	}
}

func Test_advanced(t *testing.T) {
	tests := []struct {
		name string
		hand string
		want bool
	}{
		{"hand1", "11123", true},
		{"hand2", "12131", true},
		{"hand3", "11123455", true},
		{"hand4", "11122334", true},
		{"hand5", "11234", true},
		{"hand6", "123456", false},
		{"hand7", "11133355577", true},
		{"hand8", "11223344556677", true},
		{"hand9", "12233444556677", true},
		{"hand10", "1123456789", false},
		{"hand11", "00123457", false},
		{"hand12", "0012345", false},
		{"hand13", "11890", false},
		{"hand14", "99", true},
		{"hand15", "11122344", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, advanced(tt.hand))
		})
	}
}

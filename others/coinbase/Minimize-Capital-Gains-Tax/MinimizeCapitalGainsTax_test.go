package minimizecapitalgainstax

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateMinimalTax(t *testing.T) {
	solver := &Solution{}
	cases := []struct {
		name         string
		transactions [][]string
		want         float64
	}{
		{
			name: "Example 1",
			transactions: [][]string{
				{"1", "buy", "100", "20"},
				{"2", "buy", "50", "30"},
				{"3", "sell", "80", "25"},
				{"4", "sell", "60", "35"},
			},
			want: 105.0,
		},
		{
			name: "Example 2",
			transactions: [][]string{
				{"1", "buy", "20", "50"},
				{"2", "sell", "10", "60"},
				{"3", "buy", "15", "55"},
				{"4", "sell", "10", "65"},
				{"5", "sell", "10", "70"},
			},
			want: 37.5,
		},
		{
			name: "Example 3",
			transactions: [][]string{
				{"1", "buy", "10", "10"},
				{"2", "buy", "20", "20"},
				{"3", "buy", "30", "105"},
				{"4", "sell", "10", "100"},
				{"5", "sell", "20", "120"},
				{"6", "sell", "30", "50"},
			},
			want: 130.0,
		},
		{
			name: "Additional 1",
			transactions: [][]string{
				{"1", "buy", "50", "25"},
				{"2", "buy", "10", "40"},
				{"3", "sell", "20", "6"},
				{"4", "sell", "60", "12"},
			},
			want: 0.0,
		},
		{
			name: "Additional 2",
			transactions: [][]string{
				{"1", "buy", "50", "10"},
				{"2", "sell", "20", "15"},
				{"3", "buy", "30", "12"},
				{"4", "sell", "40", "20"},
				{"5", "buy", "10", "25"},
				{"6", "buy", "30", "50"},
				{"7", "sell", "10", "5"},
			},
			want: 44.0,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := solver.CalculateMinimalTax(tc.transactions)
			assert.InDelta(t, tc.want, got, 1e-9)
		})
	}
}

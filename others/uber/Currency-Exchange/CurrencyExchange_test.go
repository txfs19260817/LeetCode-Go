package uber

import (
	"math"
	"testing"
)

func TestCurrencyConverter(t *testing.T) {
	tests := []struct {
		name    string
		fromArr []string
		toArr   []string
		rateArr []float64
		queries []struct {
			from     string
			to       string
			expected float64
		}
	}{
		{
			name:    "Example 1",
			fromArr: []string{"GBP", "USD", "USD", "USD", "CNY"},
			toArr:   []string{"JPY", "JPY", "GBP", "CAD", "EUR"},
			rateArr: []float64{155.0, 112.0, 0.9, 1.3, 0.14},
			queries: []struct {
				from     string
				to       string
				expected float64
			}{
				{"USD", "JPY", 139.5},
				{"JPY", "GBP", 0.0080357}, // README says "BGP" (typo) and 0.00803. Calculated: (1/112) * 0.9
				{"XYZ", "GBP", -1.0},
				{"CNY", "CAD", -1.0},
			},
		},
		{
			name:    "Cycle",
			fromArr: []string{"USD", "CAD", "GBP", "EUR", "AUD"},
			toArr:   []string{"CAD", "GBP", "EUR", "AUD", "USD"},
			rateArr: []float64{1.3, 0.8, 1.1, 1.5, 0.75},
			queries: []struct {
				from     string
				to       string
				expected float64
			}{
				{"USD", "EUR", 1.144},
				{"EUR", "USD", 1.125},
				{"CAD", "AUD", 1.32},
				{"USD", "USD", 1.0},
			},
		},
		{
			name:    "Crypto",
			fromArr: []string{"BTC", "ETH", "BTC", "LTC"},
			toArr:   []string{"ETH", "XRP", "LTC", "ETH"},
			rateArr: []float64{15, 0.05, 190, 0.2},
			queries: []struct {
				from     string
				to       string
				expected float64
			}{
				{"XRP", "LTC", 253.33333},
				{"BTC", "ETH", 38.0},
				{"ETH", "XRP", 0.05},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := Constructor(test.fromArr, test.toArr, test.rateArr)
			for _, q := range test.queries {
				result := c.GetBestRate(q.from, q.to)

				// Check for exact match (e.g. -1.0 or 1.0)
				if result == q.expected {
					continue
				}

				// Check for float equality with epsilon
				if math.Abs(result-q.expected) > 1e-5 {
					t.Errorf("GetBestRate(%s, %s) = %v; want %v", q.from, q.to, result, q.expected)
				}
			}
		})
	}
}

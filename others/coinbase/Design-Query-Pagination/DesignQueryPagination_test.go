package coinbase

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func sampleRecords() [][]string {
	return [][]string{
		{"1", "id-1", "user-1", "USD", "5"},
		{"2", "id-2", "user-2", "USD", "10"},
		{"3", "id-3", "user-1", "CAD", "20"},
		{"4", "id-4", "user-1", "CAD", "10"},
		{"5", "id-5", "user-1", "AUD", "30"},
		{"6", "id-6", "user-1", "JPY", "100"},
	}
}

func TestQuerySystem(t *testing.T) {
	tests := []struct {
		name string
		run  func(t *testing.T)
	}{
		{
			name: "example pagination",
			run: func(t *testing.T) {
				system := NewQuerySystem(sampleRecords())
				system.SetPageSize(2)
				system.SetTimeRange(1, 5)
				system.SetUserId("user-1")

				assert.Equal(t, [][]string{
					{"1", "id-1", "user-1", "USD", "5"},
					{"3", "id-3", "user-1", "CAD", "20"},
				}, system.NextPage())
				assert.Equal(t, [][]string{
					{"4", "id-4", "user-1", "CAD", "10"},
					{"5", "id-5", "user-1", "AUD", "30"},
				}, system.NextPage())
				assert.Equal(t, [][]string{}, system.NextPage())
			},
		},
		{
			name: "no page size returns all then empty",
			run: func(t *testing.T) {
				system := NewQuerySystem(sampleRecords())
				system.SetCurrency("USD")
				assert.Equal(t, [][]string{
					{"1", "id-1", "user-1", "USD", "5"},
					{"2", "id-2", "user-2", "USD", "10"},
				}, system.NextPage())
				assert.Equal(t, [][]string{}, system.NextPage())
			},
		},
		{
			name: "filter change resets cursor",
			run: func(t *testing.T) {
				system := NewQuerySystem(sampleRecords())
				system.SetPageSize(1)
				system.SetUserID("user-1")
				assert.Equal(t, [][]string{{"1", "id-1", "user-1", "USD", "5"}}, system.NextPage())

				system.SetAmountRange(10, 100)
				assert.Equal(t, [][]string{{"3", "id-3", "user-1", "CAD", "20"}}, system.NextPage())
			},
		},
		{
			name: "records are sorted by timestamp",
			run: func(t *testing.T) {
				records := [][]string{
					{"10", "id-10", "u1", "USD", "1"},
					{"2", "id-2", "u1", "USD", "1"},
					{"7", "id-7", "u1", "USD", "1"},
				}
				system := NewQueryPaginationSystem(records)
				assert.Equal(t, [][]string{
					{"2", "id-2", "u1", "USD", "1"},
					{"7", "id-7", "u1", "USD", "1"},
					{"10", "id-10", "u1", "USD", "1"},
				}, system.NextPage())
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, tc.run)
	}
}

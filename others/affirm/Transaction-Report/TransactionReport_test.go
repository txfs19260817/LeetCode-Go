package affirm

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestBuildLast24hReport(t *testing.T) {
	now := time.Date(2025, 12, 2, 0, 0, 0, 0, time.UTC)

	source1 := []Transaction{
		{Amount: 20, MerchantID: "M1", Status: StatusCompleted, UserID: "U1", StartTime: time.Date(2025, 12, 1, 1, 0, 0, 0, time.UTC)},
		{Amount: 5, MerchantID: "M1", Status: StatusCompleted, UserID: "U2", StartTime: time.Date(2025, 12, 1, 2, 0, 0, 0, time.UTC)},
		{Amount: 7, MerchantID: "M2", Status: StatusCompleted, UserID: "U3", StartTime: time.Date(2025, 12, 1, 23, 0, 0, 0, time.UTC)},
	}

	source2 := []Transaction{
		{Amount: 9, MerchantID: "M1", Status: StatusCompleted, UserID: "U4", StartTime: time.Date(2025, 12, 1, 1, 30, 0, 0, time.UTC)},
		{Amount: 99, MerchantID: "M1", Status: StatusPending, UserID: "U5", StartTime: time.Date(2025, 12, 1, 3, 0, 0, 0, time.UTC)},
		{Amount: 10, MerchantID: "M2", Status: StatusCompleted, UserID: "U6", StartTime: time.Date(2025, 12, 1, 0, 0, 0, 0, time.UTC)}, // diff 24, excluded
		{Amount: 11, MerchantID: "M3", Status: StatusCompleted, UserID: "U7", StartTime: time.Date(2025, 12, 2, 1, 0, 0, 0, time.UTC)}, // future
	}

	all := append(source1, source2...)

	reporter := NewTransactionReporter()
	result := reporter.BuildLast24hReport(all, now)

	expected := map[string]map[int]int64{
		"M1": {
			23: 20, // 12-01 01:00 -> diff 23
			22: 5,  // 12-01 02:00 -> diff 22
		},
		"M2": {
			1: 7, // 12-01 23:00 -> diff 1
		},
	}
	// M1 extra at 12-01 01:30 diff 22 -> should aggregate with diff 22 bucket
	expected["M1"][22] += 9

	assert.Equal(t, expected, result)
}

func TestBuildLast24hReportEmpty(t *testing.T) {
	reporter := NewTransactionReporter()
	result := reporter.BuildLast24hReport(nil, time.Now())
	assert.Empty(t, result)
}

package affirm

import (
	"math"
	"time"
)

type Status int

const (
	StatusUnknown Status = iota
	StatusCompleted
	StatusPending
	StatusFailed
)

type Transaction struct {
	Amount     int64
	MerchantID string
	Status     Status
	UserID     string
	StartTime  time.Time
}

// HourDiff returns the integer hour difference between now and start time.
// In the interview setting, this helper is provided to you.
func HourDiff(now, start time.Time) int {
	return int(math.Floor(now.Sub(start).Hours()))
}

type TransactionReporter struct{}

func NewTransactionReporter() *TransactionReporter {
	return &TransactionReporter{}
}

// BuildLast24hReport returns: merchantId -> (hourDiff -> sumAmount).
func (tr *TransactionReporter) BuildLast24hReport(
	transactions []Transaction,
	now time.Time,
) map[string]map[int]int64 {
	report := make(map[string]map[int]int64)

	for _, tx := range transactions {
		if tx.Status != StatusCompleted {
			continue
		}

		diff := HourDiff(now, tx.StartTime)
		if diff < 0 || diff >= 24 {
			continue
		}

		inner, ok := report[tx.MerchantID]
		if !ok {
			inner = make(map[int]int64)
			report[tx.MerchantID] = inner
		}
		inner[diff] += tx.Amount
	}

	return report
}

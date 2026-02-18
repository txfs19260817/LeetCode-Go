package coinbase

import (
	"sort"
	"strconv"
)

const (
	recordLen      = 5
	timestampIndex = 0
	userIDIndex    = 2
	currencyIndex  = 3
	amountIndex    = 4
)

type numberRange struct {
	start   int
	end     int
	enabled bool
}

type transactionRecord struct {
	raw       []string
	timestamp int
	amount    int
	userID    string
	currency  string
}

type QuerySystem struct {
	records      []transactionRecord
	pageSize     int
	nextIndex    int
	timeRange    numberRange
	amountRange  numberRange
	userIDFilter *string
	currencyOnly *string
}

type QueryPaginationSystem = QuerySystem

func NewQuerySystem(records [][]string) *QuerySystem {
	parsed := make([]transactionRecord, 0, len(records))
	for _, row := range records {
		if len(row) < recordLen {
			continue
		}
		timestamp, err := strconv.Atoi(row[timestampIndex])
		if err != nil {
			continue
		}
		amount, err := strconv.Atoi(row[amountIndex])
		if err != nil {
			continue
		}
		rawCopy := append([]string(nil), row...)
		parsed = append(parsed, transactionRecord{
			raw:       rawCopy,
			timestamp: timestamp,
			amount:    amount,
			userID:    row[userIDIndex],
			currency:  row[currencyIndex],
		})
	}

	sort.SliceStable(parsed, func(i, j int) bool {
		return parsed[i].timestamp < parsed[j].timestamp
	})

	return &QuerySystem{records: parsed}
}

func NewQueryPaginationSystem(records [][]string) *QuerySystem {
	return NewQuerySystem(records)
}

func (qs *QuerySystem) SetPageSize(size int) {
	qs.pageSize = size
	qs.resetPagination()
}

func (qs *QuerySystem) SetTimeRange(start int, end int) {
	if start > end {
		start, end = end, start
	}
	qs.timeRange = numberRange{start: start, end: end, enabled: true}
	qs.resetPagination()
}

func (qs *QuerySystem) SetAmountRange(start int, end int) {
	if start > end {
		start, end = end, start
	}
	qs.amountRange = numberRange{start: start, end: end, enabled: true}
	qs.resetPagination()
}

func (qs *QuerySystem) SetUserID(id string) {
	value := id
	qs.userIDFilter = &value
	qs.resetPagination()
}

func (qs *QuerySystem) SetUserId(id string) {
	qs.SetUserID(id)
}

func (qs *QuerySystem) SetCurrency(currency string) {
	value := currency
	qs.currencyOnly = &value
	qs.resetPagination()
}

func (qs *QuerySystem) NextPage() [][]string {
	limit := len(qs.records)
	if qs.pageSize > 0 {
		limit = qs.pageSize
	}

	result := make([][]string, 0, limit)
	for qs.nextIndex < len(qs.records) && len(result) < limit {
		record := qs.records[qs.nextIndex]
		qs.nextIndex++
		if !qs.matches(record) {
			continue
		}
		result = append(result, append([]string(nil), record.raw...))
	}
	return result
}

func (qs *QuerySystem) resetPagination() {
	qs.nextIndex = 0
}

func (qs *QuerySystem) matches(record transactionRecord) bool {
	if qs.timeRange.enabled {
		if record.timestamp < qs.timeRange.start || record.timestamp > qs.timeRange.end {
			return false
		}
	}
	if qs.amountRange.enabled {
		if record.amount < qs.amountRange.start || record.amount > qs.amountRange.end {
			return false
		}
	}
	if qs.userIDFilter != nil && record.userID != *qs.userIDFilter {
		return false
	}
	if qs.currencyOnly != nil && record.currency != *qs.currencyOnly {
		return false
	}
	return true
}

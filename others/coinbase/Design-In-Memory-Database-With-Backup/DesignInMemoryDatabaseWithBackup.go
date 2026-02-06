package designinmemorydatabasewithbackup

import (
	"sort"
	"strings"
)

type valueWithTTL struct {
	value        string
	setTimestamp int
	ttl          int
	hasTTL       bool
}

func (v valueWithTTL) isAliveAt(timestamp int) bool {
	if !v.hasTTL {
		return true
	}
	return timestamp >= v.setTimestamp && timestamp < v.setTimestamp+v.ttl
}

type backupSnapshot struct {
	timestamp int
	records   map[string]map[string]valueWithTTL
}

type InMemoryDB struct {
	records map[string]map[string]valueWithTTL
	backups []backupSnapshot
}

func Constructor() *InMemoryDB {
	return &InMemoryDB{records: make(map[string]map[string]valueWithTTL)}
}

func (db *InMemoryDB) SetData(key string, field string, value string) {
	record := db.ensureRecord(key)
	record[field] = valueWithTTL{value: value}
}

func (db *InMemoryDB) GetData(key string, field string) string {
	if db.records == nil {
		return ""
	}

	record, exists := db.records[key]
	if !exists {
		return ""
	}
	entry, ok := record[field]
	if !ok {
		return ""
	}
	return entry.value
}

func (db *InMemoryDB) DeleteData(key string, field string) bool {
	if db.records == nil {
		return false
	}

	record, exists := db.records[key]
	if !exists {
		return false
	}
	if _, ok := record[field]; !ok {
		return false
	}
	delete(record, field)
	if len(record) == 0 {
		delete(db.records, key)
	}
	return true
}

func (db *InMemoryDB) ScanData(key string) []string {
	if db.records == nil {
		return []string{}
	}

	record, exists := db.records[key]
	if !exists {
		return []string{}
	}
	fields := make([]string, 0, len(record))
	for field := range record {
		fields = append(fields, field)
	}
	sort.Strings(fields)
	result := make([]string, 0, len(fields))
	for _, field := range fields {
		result = append(result, field+"("+record[field].value+")")
	}
	return result
}

func (db *InMemoryDB) ScanDataByPrefix(key string, prefix string) []string {
	if db.records == nil {
		return []string{}
	}

	record, exists := db.records[key]
	if !exists {
		return []string{}
	}
	fields := make([]string, 0, len(record))
	for field := range record {
		if strings.HasPrefix(field, prefix) {
			fields = append(fields, field)
		}
	}
	sort.Strings(fields)
	result := make([]string, 0, len(fields))
	for _, field := range fields {
		result = append(result, field+"("+record[field].value+")")
	}
	return result
}

func (db *InMemoryDB) SetDataAt(key string, field string, value string, timestamp int) {
	record := db.ensureRecord(key)
	record[field] = valueWithTTL{value: value, setTimestamp: timestamp}
}

func (db *InMemoryDB) SetDataAtWithTtl(key string, field string, value string, timestamp int, ttl int) {
	record := db.ensureRecord(key)
	record[field] = valueWithTTL{value: value, setTimestamp: timestamp, ttl: ttl, hasTTL: true}
}

func (db *InMemoryDB) DeleteDataAt(key string, field string, timestamp int) bool {
	if db.records == nil {
		return false
	}

	record, exists := db.records[key]
	if !exists {
		return false
	}
	entry, ok := record[field]
	if !ok || !entry.isAliveAt(timestamp) {
		return false
	}
	delete(record, field)
	if len(record) == 0 {
		delete(db.records, key)
	}
	return true
}

func (db *InMemoryDB) GetDataAt(key string, field string, timestamp int) string {
	if db.records == nil {
		return ""
	}

	record, exists := db.records[key]
	if !exists {
		return ""
	}
	entry, ok := record[field]
	if ok && entry.isAliveAt(timestamp) {
		return entry.value
	}
	return ""
}

func (db *InMemoryDB) ScanDataAt(key string, timestamp int) []string {
	if db.records == nil {
		return []string{}
	}

	record, exists := db.records[key]
	if !exists {
		return []string{}
	}
	fields := make([]string, 0, len(record))
	for field, entry := range record {
		if entry.isAliveAt(timestamp) {
			fields = append(fields, field)
		}
	}
	sort.Strings(fields)
	result := make([]string, 0, len(fields))
	for _, field := range fields {
		result = append(result, field+"("+record[field].value+")")
	}
	return result
}

func (db *InMemoryDB) ScanDataByPrefixAt(key string, prefix string, timestamp int) []string {
	if db.records == nil {
		return []string{}
	}

	record, exists := db.records[key]
	if !exists {
		return []string{}
	}
	fields := make([]string, 0, len(record))
	for field, entry := range record {
		if strings.HasPrefix(field, prefix) && entry.isAliveAt(timestamp) {
			fields = append(fields, field)
		}
	}
	sort.Strings(fields)
	result := make([]string, 0, len(fields))
	for _, field := range fields {
		result = append(result, field+"("+record[field].value+")")
	}
	return result
}

func (db *InMemoryDB) Backup(timestamp int) int {
	snapshot := make(map[string]map[string]valueWithTTL)
	count := 0
	for key, record := range db.records {
		newRecord := make(map[string]valueWithTTL)
		for field, entry := range record {
			if !entry.isAliveAt(timestamp) {
				continue
			}
			if entry.hasTTL {
				remaining := entry.setTimestamp + entry.ttl - timestamp
				if remaining <= 0 {
					continue
				}
				newRecord[field] = valueWithTTL{
					value:        entry.value,
					setTimestamp: timestamp,
					ttl:          remaining,
					hasTTL:       true,
				}
			} else {
				newRecord[field] = valueWithTTL{
					value:        entry.value,
					setTimestamp: timestamp,
				}
			}
		}
		if len(newRecord) > 0 {
			snapshot[key] = newRecord
			count++
		}
	}
	db.backups = append(db.backups, backupSnapshot{
		timestamp: timestamp,
		records:   snapshot,
	})
	return count
}

func (db *InMemoryDB) Restore(timestamp int, timestampToRestore int) {
	var chosen *backupSnapshot
	for i := len(db.backups) - 1; i >= 0; i-- {
		if db.backups[i].timestamp <= timestampToRestore {
			chosen = &db.backups[i]
			break
		}
	}
	if chosen == nil {
		return
	}
	db.records = make(map[string]map[string]valueWithTTL, len(chosen.records))
	for key, record := range chosen.records {
		newRecord := make(map[string]valueWithTTL, len(record))
		for field, entry := range record {
			newEntry := valueWithTTL{
				value:        entry.value,
				setTimestamp: timestamp,
				ttl:          entry.ttl,
				hasTTL:       entry.hasTTL,
			}
			newRecord[field] = newEntry
		}
		db.records[key] = newRecord
	}
}

func (db *InMemoryDB) ensureRecord(key string) map[string]valueWithTTL {
	if db.records == nil {
		db.records = make(map[string]map[string]valueWithTTL)
	}
	record, exists := db.records[key]
	if !exists {
		record = make(map[string]valueWithTTL)
		db.records[key] = record
	}
	return record
}

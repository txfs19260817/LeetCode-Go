package coinbase

import (
	"sort"
	"strconv"
	"strings"
)

type workSession struct {
	start        int
	end          int
	compensation int
}

type worker struct {
	position              string
	compensation          int
	inOffice              bool
	lastEntry             int
	totalTime             int
	workTimes             map[string]int
	sessions              []workSession
	hasPromotion          bool
	promotionStart        int
	promotionPosition     string
	promotionCompensation int
}

type interval struct {
	start int
	end   int
}

type OfficeManager struct {
	workers    map[string]*worker
	doublePaid []interval
}

func Constructor() *OfficeManager {
	return &OfficeManager{workers: make(map[string]*worker)}
}

func (om *OfficeManager) AddWorker(workerId string, position string, compensation int) bool {
	if om.workers == nil {
		om.workers = make(map[string]*worker)
	}
	if _, exists := om.workers[workerId]; exists {
		return false
	}

	om.workers[workerId] = &worker{
		position:     position,
		compensation: compensation,
		workTimes:    make(map[string]int),
	}
	return true
}

func (om *OfficeManager) RegisterWorker(workerId string, timestamp int) string {
	if om.workers == nil {
		return "invalid_request"
	}

	w, exists := om.workers[workerId]
	if !exists {
		return "invalid_request"
	}

	if !w.inOffice {
		if w.hasPromotion && timestamp >= w.promotionStart {
			w.position = w.promotionPosition
			w.compensation = w.promotionCompensation
			w.hasPromotion = false
		}
		w.inOffice = true
		w.lastEntry = timestamp
		return "registered"
	}

	w.inOffice = false
	duration := timestamp - w.lastEntry
	w.totalTime += duration
	if w.workTimes == nil {
		w.workTimes = make(map[string]int)
	}
	w.workTimes[w.position] += duration
	w.sessions = append(w.sessions, workSession{
		start:        w.lastEntry,
		end:          timestamp,
		compensation: w.compensation,
	})
	return "registered"
}

func (om *OfficeManager) Get(workerId string) int {
	if om.workers == nil {
		return -1
	}

	w, exists := om.workers[workerId]
	if !exists {
		return -1
	}
	return w.totalTime
}

func (om *OfficeManager) TopNWorkers(n int, position string) string {
	if n <= 0 {
		return ""
	}
	if om.workers == nil {
		return ""
	}

	type entry struct {
		id   string
		time int
	}
	var entries []entry
	for id, w := range om.workers {
		if w.position != position {
			continue
		}
		time := w.workTimes[position]
		entries = append(entries, entry{id: id, time: time})
	}
	if len(entries) == 0 {
		return ""
	}

	sort.Slice(entries, func(i, j int) bool {
		if entries[i].time == entries[j].time {
			return entries[i].id < entries[j].id
		}
		return entries[i].time > entries[j].time
	})

	if n > len(entries) {
		n = len(entries)
	}

	var b strings.Builder
	for i := 0; i < n; i++ {
		if i > 0 {
			b.WriteString(", ")
		}
		b.WriteString(entries[i].id)
		b.WriteByte('(')
		b.WriteString(strconv.Itoa(entries[i].time))
		b.WriteByte(')')
	}
	return b.String()
}

func (om *OfficeManager) Promote(workerId string, newPosition string, newCompensation string, startTimestamp int) string {
	if om.workers == nil {
		return "invalid_request"
	}
	w, exists := om.workers[workerId]
	if !exists {
		return "invalid_request"
	}
	if w.hasPromotion {
		return "invalid_request"
	}

	comp, err := strconv.Atoi(newCompensation)
	if err != nil {
		return "invalid_request"
	}

	w.hasPromotion = true
	w.promotionStart = startTimestamp
	w.promotionPosition = newPosition
	w.promotionCompensation = comp
	return "success"
}

func (om *OfficeManager) CalcSalary(workerId string, startTimestamp int, endTimestamp int) int {
	if om.workers == nil {
		return -1
	}
	w, exists := om.workers[workerId]
	if !exists {
		return -1
	}

	total := 0
	for _, session := range w.sessions {
		sessionStart := maxInt(session.start, startTimestamp)
		sessionEnd := minInt(session.end, endTimestamp)
		if sessionStart >= sessionEnd {
			continue
		}
		duration := sessionEnd - sessionStart
		total += duration * session.compensation
		extra := om.doublePaidOverlap(sessionStart, sessionEnd)
		total += extra * session.compensation
	}
	return total
}

func (om *OfficeManager) SetDoublePaid(startTimestamp int, endTimestamp int) {
	if startTimestamp >= endTimestamp {
		return
	}
	om.doublePaid = append(om.doublePaid, interval{start: startTimestamp, end: endTimestamp})
	sort.Slice(om.doublePaid, func(i, j int) bool {
		return om.doublePaid[i].start < om.doublePaid[j].start
	})

	merged := make([]interval, 0, len(om.doublePaid))
	for _, current := range om.doublePaid {
		if len(merged) == 0 {
			merged = append(merged, current)
			continue
		}
		last := &merged[len(merged)-1]
		if current.start <= last.end {
			if current.end > last.end {
				last.end = current.end
			}
			continue
		}
		merged = append(merged, current)
	}
	om.doublePaid = merged
}

func (om *OfficeManager) doublePaidOverlap(start int, end int) int {
	if start >= end || len(om.doublePaid) == 0 {
		return 0
	}
	total := 0
	for _, period := range om.doublePaid {
		if period.end <= start {
			continue
		}
		if period.start >= end {
			break
		}
		overlapStart := maxInt(start, period.start)
		overlapEnd := minInt(end, period.end)
		if overlapStart < overlapEnd {
			total += overlapEnd - overlapStart
		}
	}
	return total
}

func maxInt(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func minInt(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

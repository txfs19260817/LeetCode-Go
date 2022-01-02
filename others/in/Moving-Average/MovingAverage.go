package in

import "time"

const window = 5 * time.Minute

type event struct {
	t   time.Time
	val int
}

type movingAverage struct {
	sum   int
	queue []event
}

func NewMovingAverage() *movingAverage {
	return &movingAverage{}
}

// Record adds a record
func (m movingAverage) Record(val int) {
	m.queue = append(m.queue, event{m.getNow(), val})
	m.sum += val
	m.removeExpiredData()
}

// GetAvg return the avg within the time window
func (m movingAverage) GetAvg() float64 {
	m.removeExpiredData()
	if len(m.queue) == 0 {
		return .0
	}
	return float64(m.sum) / float64(len(m.queue))
}

func (m *movingAverage) removeExpiredData() {
	for len(m.queue) > 0 && m.getNow().Sub(m.queue[0].t) > window {
		top := m.queue[0]
		m.sum -= top.val
	}
}

// 这个是每次记录读进来的时间的,这个不用自己写,就是直接返回当前系统时间
func (m movingAverage) getNow() time.Time {
	return time.Now()
}

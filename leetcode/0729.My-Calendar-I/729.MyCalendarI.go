package leetcode

import "sort"

type MyCalendar struct {
	start, end []int
}

func Constructor() MyCalendar {
	return MyCalendar{}
}

func (c *MyCalendar) Book(startTime int, endTime int) bool {
	i := sort.SearchInts(c.start, endTime)
	if i > 0 && c.end[i-1] > startTime {
		return false
	}
	c.start = append(c.start, 0)
	c.end = append(c.end, 0)
	copy(c.start[i+1:], c.start[i:])
	copy(c.end[i+1:], c.end[i:])
	c.start[i] = startTime
	c.end[i] = endTime
	return true
}

// Delete deletes a booking (follow-up requirement)
func (c *MyCalendar) Delete(startTime int, endTime int) bool {
	i := sort.SearchInts(c.start, startTime)
	if i < len(c.start) && c.start[i] == startTime && c.end[i] == endTime {
		c.start = append(c.start[:i], c.start[i+1:]...)
		c.end = append(c.end[:i], c.end[i+1:]...)
		return true
	}
	return false
}

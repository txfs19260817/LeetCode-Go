package uber

import "sort"

type Room struct {
	roomName     string
	starts, ends []int
}

func (r *Room) book(start int, end int) bool {
	idx := sort.SearchInts(r.starts, end)
	if idx > 0 && r.ends[idx-1] > start {
		return false
	}

	r.starts = append(r.starts, 0)
	r.ends = append(r.ends, 0)
	copy(r.starts[idx+1:], r.starts[idx:])
	copy(r.ends[idx+1:], r.ends[idx:])
	r.starts[idx] = start
	r.ends[idx] = end
	return true
}

type MeetingScheduler struct {
	rooms []*Room
}

func Constructor(roomList []string) MeetingScheduler {
	sort.Strings(roomList)
	rooms := make([]*Room, len(roomList))
	for i := range rooms {
		rooms[i] = &Room{roomName: roomList[i]}
	}
	return MeetingScheduler{rooms}
}

func (ms *MeetingScheduler) Schedule(start int, end int) string {
	for _, r := range ms.rooms {
		if r.book(start, end) {
			return r.roomName
		}
	}
	return ""
}

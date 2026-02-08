package uber

// spotType
const (
	SpotMotorcycle = iota
	SpotCompact
	SpotLarge
)

// vehicleType
const (
	Motorcycle = "motorcycle"
	Car        = "car"
	Van        = "van"
	Bus        = "bus"
)

type ParkingLot struct {
	// spots2cnt holds available counts by spot type index.
	spots2cnt [3]int
	// parked maps vehicle id -> (spotType, count) for easy removal.
	parked map[string][2]int
	// rules defines preferred allocations in priority order per vehicle type.
	rules map[string][][2]int
}

// NewParkingLot builds a lot with fixed spot counts and default allocation rules.
func NewParkingLot(motorcycle, compact, large int) *ParkingLot {
	return &ParkingLot{
		spots2cnt: [3]int{motorcycle, compact, large},
		parked:    map[string][2]int{},
		rules: map[string][][2]int{
			Motorcycle: {{SpotMotorcycle, 1}, {SpotCompact, 1}, {SpotLarge, 1}},
			Car:        {{SpotCompact, 1}, {SpotLarge, 1}},
			Van:        {{SpotLarge, 1}, {SpotCompact, 3}},
			Bus:        {{SpotLarge, 1}, {SpotCompact, 5}},
		},
	}
}

// Park tries to allocate a spot (or compact count) based on vehicle rules.
func (p *ParkingLot) Park(id, vehicleType string) bool {
	if _, exists := p.parked[id]; exists {
		return false
	}
	options := p.rules[vehicleType]
	for _, opt := range options {
		spot, count := opt[0], opt[1]
		if p.spots2cnt[spot] >= count {
			p.spots2cnt[spot] -= count
			p.parked[id] = opt
			return true
		}
	}
	return false
}

// Remove frees the spots previously used by the vehicle id.
func (p *ParkingLot) Remove(id string) bool {
	alloc, ok := p.parked[id] // (spotType, count)
	if !ok {
		return false
	}
	p.spots2cnt[alloc[0]] += alloc[1]
	delete(p.parked, id)
	return true
}

// Available returns remaining count for a given spot type.
func (p *ParkingLot) Available(spotType int) int {
	if spotType < 0 || spotType > SpotLarge {
		return 0
	}
	return p.spots2cnt[spotType]
}

package uber

const (
	SpotMotorcycle = iota
	SpotCompact
	SpotLarge
)

type ParkingLot struct {
	// spots holds available counts by spot type index.
	spots [3]int
	// parked maps vehicle id -> (spotType, count) for easy removal.
	parked map[string][2]int
	// rules defines preferred allocations in priority order per vehicle type.
	rules map[string][][2]int
}

// NewParkingLot builds a lot with fixed spot counts and default allocation rules.
func NewParkingLot(motorcycle, compact, large int) *ParkingLot {
	return &ParkingLot{
		spots:  [3]int{motorcycle, compact, large},
		parked: map[string][2]int{},
		rules: map[string][][2]int{
			"motorcycle": {{SpotMotorcycle, 1}, {SpotCompact, 1}, {SpotLarge, 1}},
			"car":        {{SpotCompact, 1}, {SpotLarge, 1}},
			"van":        {{SpotLarge, 1}, {SpotCompact, 3}},
			"bus":        {{SpotLarge, 1}, {SpotCompact, 5}},
		},
	}
}

// Park tries to allocate a spot (or compact count) based on vehicle rules.
func (p *ParkingLot) Park(id, vehicleType string) bool {
	if id == "" {
		return false
	}
	if _, exists := p.parked[id]; exists {
		return false
	}
	options, ok := p.rules[vehicleType]
	if !ok {
		return false
	}
	for _, opt := range options {
		spot, count := opt[0], opt[1]
		if p.spots[spot] >= count {
			p.spots[spot] -= count
			p.parked[id] = opt
			return true
		}
	}
	return false
}

// Remove frees the spots previously used by the vehicle id.
func (p *ParkingLot) Remove(id string) bool {
	alloc, ok := p.parked[id]
	if !ok {
		return false
	}
	p.spots[alloc[0]] += alloc[1]
	delete(p.parked, id)
	return true
}

// Available returns remaining count for a given spot type.
func (p *ParkingLot) Available(spotType int) int {
	if spotType < 0 || spotType > SpotLarge {
		return 0
	}
	return p.spots[spotType]
}

package k

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

// LogEntry represents a single log line
type LogEntry struct {
	Timestamp    float64
	LicensePlate string
	BoothType    string
	Location     int
	Direction    string
}

// NewLogEntry parses a log line string into a LogEntry struct
func NewLogEntry(logLine string) (LogEntry, error) {
	tokens := strings.Split(logLine, " ")
	if len(tokens) < 4 {
		return LogEntry{}, fmt.Errorf("invalid log line format")
	}

	// 1. Fix Timestamp parsing (Java used float, Go needs ParseFloat)
	ts, err := strconv.ParseFloat(tokens[0], 64)
	if err != nil {
		return LogEntry{}, err
	}

	license := tokens[1]

	// 2. Fix Location parsing
	// Token[2] is like "210E". Substring until last char is location.
	locStr := tokens[2][:len(tokens[2])-1]
	loc, err := strconv.Atoi(locStr)
	if err != nil {
		return LogEntry{}, err
	}

	// 3. Parse Direction
	dirChar := tokens[2][len(tokens[2])-1:]
	var direction string
	if dirChar == "E" {
		direction = "EAST"
	} else if dirChar == "W" {
		direction = "WEST"
	} else {
		return LogEntry{}, fmt.Errorf("invalid direction")
	}

	boothType := tokens[3]

	return LogEntry{
		Timestamp:    ts,
		LicensePlate: license,
		BoothType:    boothType,
		Location:     loc,
		Direction:    direction,
	}, nil
}

type LogFile struct {
	Entries []LogEntry
}

func NewLogFile(r io.Reader) (*LogFile, error) {
	scanner := bufio.NewScanner(r)
	var entries []LogEntry
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		entry, err := NewLogEntry(line)
		if err != nil {
			return nil, err
		}
		entries = append(entries, entry)
	}
	return &LogFile{Entries: entries}, nil
}

// Task 2-1: Count Journeys
// A complete journey starts with ENTRY and ends with EXIT.
// We assume logs contain complete journeys.
// So we can just count the number of "EXIT" entries, or count pairs.
// Since "You may assume that the log only contains complete journeys",
// counting "EXIT" events is sufficient and accurate.
func (lf *LogFile) CountJourneys() int {
	count := 0
	for _, entry := range lf.Entries {
		if entry.BoothType == "EXIT" {
			count++
		}
	}
	return count
}

// Task 3-1: Catch Speeders
func (lf *LogFile) CatchSpeeders() (speeders []string) {
	// Group entries by license plate to reconstruct journeys
	// Since logs are time-ordered (implied), we can iterate through.
	// However, entries for different cars might be interleaved.
	// We need to track the active journey for each car.

	type JourneyState struct {
		LastEntry LogEntry
		// Track high speed segments count for the "120km/h in two segments" rule
		SegmentsOver120 int
		IsSpeeding      bool // To ensure we only mark once per journey
	}

	activeJourneys := make(map[string]*JourneyState)

	for _, entry := range lf.Entries {
		plate := entry.LicensePlate

		if entry.BoothType == "ENTRY" {
			// Start of a new journey
			activeJourneys[plate] = &JourneyState{
				LastEntry: entry,
			}
		} else {
			// MAINROAD or EXIT
			state, exists := activeJourneys[plate]
			if !exists {
				// Should not happen based on problem statement (complete journeys)
				continue
			}

			// Calculate speed for this segment
			// Segment distance is absolute difference in location
			// Locations are in km. Timestamps in seconds.
			distKm := float64(abs(entry.Location - state.LastEntry.Location))
			timeSec := entry.Timestamp - state.LastEntry.Timestamp

			// Avoid divide by zero
			if timeSec > 0 {
				speedKmh := (distKm * 3600.0) / timeSec

				// Check rules
				// 1. > 130 km/h in any single 10km segment
				// Note: Problem says "any individual 10km segment".
				// The toll booths are placed every 10km. So consecutive entries are exactly 10km apart
				// unless a car skips a sensor (unlikely based on problem description of booths).
				// Assuming segments are always 10km.

				isSegmentOver130 := speedKmh >= 130.0
				isSegmentOver120 := speedKmh >= 120.0

				if !state.IsSpeeding {
					if isSegmentOver130 {
						speeders = append(speeders, plate)
						state.IsSpeeding = true
					} else if isSegmentOver120 {
						state.SegmentsOver120++
						if state.SegmentsOver120 >= 2 {
							speeders = append(speeders, plate)
							state.IsSpeeding = true
						}
					}
				}
			}

			// Update last entry for next segment
			state.LastEntry = entry

			// If EXIT, journey ends. Remove from active map.
			if entry.BoothType == "EXIT" {
				delete(activeJourneys, plate)
			}
		}
	}

	return speeders
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

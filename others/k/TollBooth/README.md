# Toll Booth Log Analysis

## Problem Description

We are writing software to analyze logs for toll booths on a highway. This highway is a divided highway with limited access; the only way on to or off of the highway is through a toll booth.

There are three types of toll booths:

- **ENTRY** (E in the diagram) toll booths, where a car goes through a booth as it enters the highway.
- **EXIT** (X in the diagram) toll booths, where a car goes through a booth as it exits the highway.
- **MAINROAD** (M in the diagram), which have sensors that record a license plate as a car drives through at full speed.

```
           Entry Booth                    Exit Booth
                |                          |
                E                          X
               /                            \
---<------------<---------M---------<-----------<---------<----
                                                 (West-bound side)
===============================================================
                                                 (East-bound side)
------>--------->---------M--------->--------->--------->------
               \                            /
                X                          E
                |                          |
            Exit Booth                Entry Booth
```

## Tasks

### Task 1: Bug Fix

The provided `LogEntry` parsing code has a bug. We need to fix it so that tests pass.

### Task 2: Count Journeys

We are interested in how many people are using the highway. We would like to count how many complete journeys are taken in the log file.

A complete journey consists of:

- A driver entering the highway through an **ENTRY** toll booth.
- The driver passing through some number of **MAINROAD** toll booths (possibly 0).
- The driver exiting the highway through an **EXIT** toll booth.

You may assume that the log only contains complete journeys, and there are no missing entries.

**Goal:** Write a function `countJourneys()` that returns how many complete journeys there are in the given LogFile.

### Task 3: Catch Speeders

We would like to identify journeys where a driver does either of the following:

- Drive an average of **130 km/h or greater** in any individual 10km segment of tollway.
- Drive an average of **120 km/h or greater** in any two 10km segments of tollway.

**Example:**

```
1000.000 TST002 270W ENTRY
1275.000 TST002 260W EXIT
```

Distance: 10 km. Time: 275 seconds.
Speed = (10 km \* 3600 sec/hr) / 275 sec = 130.91 km/hr.

**Rules:**

- A license plate may have multiple journeys in one file. If they speed in both, both should be counted.
- We do not mark speeding if they are not on the highway (i.e. between EXIT and ENTRY).
- Speeding is only marked once per journey.

**Goal:** Write a function `catchSpeeders` that returns a collection of license plates that drove at unsafe speeds. If the same license plate speeds in two different journeys, it should appear twice.

## Solution Code (Go)

Since the original request provided Java code but we are in a Go project, here is the Go implementation of the solution.

### `tollbooth.go`

```go
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
func (lf *LogFile) CatchSpeeders() []string {
	speeders := []string{}

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
```

### `tollbooth_test.go`

```go
package k

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewLogEntry(t *testing.T) {
	line := "44776.619 KTB918 310E MAINROAD"
	entry, err := NewLogEntry(line)
	assert.NoError(t, err)
	assert.InDelta(t, 44776.619, entry.Timestamp, 0.0001)
	assert.Equal(t, "KTB918", entry.LicensePlate)
	assert.Equal(t, 310, entry.Location)
	assert.Equal(t, "EAST", entry.Direction)
	assert.Equal(t, "MAINROAD", entry.BoothType)
}

func TestCountJourneys(t *testing.T) {
	logData := `90750.191 JOX304 250E ENTRY
91081.684 JOX304 260E MAINROAD
91483.251 JOX304 270E MAINROAD
91874.493 JOX304 280E EXIT`

	lf, err := NewLogFile(strings.NewReader(logData))
	assert.NoError(t, err)
	assert.Equal(t, 1, lf.CountJourneys())
}

func TestCatchSpeeders(t *testing.T) {
	// Case 1: 130km/h rule
	// 10km in 275s = 130.9 km/h -> Speeding
	logData1 := `1000.000 TST002 270W ENTRY
1275.000 TST002 260W EXIT`
	lf1, _ := NewLogFile(strings.NewReader(logData1))
	speeders1 := lf1.CatchSpeeders()
	assert.Contains(t, speeders1, "TST002")

	// Case 2: 120km/h rule (2 segments)
	// Segment 1: 10km in 290s = 124 km/h
	// Segment 2: 10km in 290s = 124 km/h
	// Total: 2 segments >= 120 -> Speeding
	logData2 := `1000.000 TST003 270W ENTRY
1290.000 TST003 260W MAINROAD
1580.000 TST003 250W EXIT`
	lf2, _ := NewLogFile(strings.NewReader(logData2))
	speeders2 := lf2.CatchSpeeders()
	assert.Contains(t, speeders2, "TST003")

	// Case 3: Normal driving
	// 10km in 400s = 90 km/h
	logData3 := `1000.000 OK001 270W ENTRY
1400.000 OK001 260W EXIT`
	lf3, _ := NewLogFile(strings.NewReader(logData3))
	speeders3 := lf3.CatchSpeeders()
	assert.Empty(t, speeders3)
}
```
